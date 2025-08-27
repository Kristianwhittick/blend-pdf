// Copyright 2025 Kristian Whittick
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"os/signal"
	"path/filepath"
	"strings"
	"syscall"
	"time"
)

func main() {
	// Initialize loggers if needed
	if DEBUG {
		initLoggers()
	}
	
	// Set up signal handling for graceful shutdown
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		fmt.Printf("\n%sShutting down gracefully...%s\n", YELLOW, NC)
		cleanup()
		os.Exit(0)
	}()

	// Setup lock file to prevent multiple instances
	if err := setupLock(); err != nil {
		if strings.Contains(err.Error(), "exit code 6") {
			os.Exit(6)
		}
		printError(err.Error())
		os.Exit(1)
	}

	// Parse command line arguments
	folder, err := parseArgs()
	if err != nil {
		printError(err.Error())
		cleanupLock()
		os.Exit(1)
	}

	// Setup directories
	if err := setupDirectories(folder); err != nil {
		printError(err.Error())
		cleanupLock()
		os.Exit(1)
	}

	// Main program loop
	for CONTINUE {
		processMenu()
	}

	cleanup()
}

func cleanup() {
	showStatistics()
	cleanupLock()
}

func showHelp() {
	fmt.Printf("BlendPDF v%s - A tool for merging PDF files\n\n", VERSION)
	fmt.Printf("Usage: %s [options] [folder]\n\n", filepath.Base(os.Args[0]))
	fmt.Printf("Command line options:\n")
	fmt.Printf("  -h, --help     Show this help information and exit\n")
	fmt.Printf("  -v, --version  Show version information and exit\n")
	fmt.Printf("  -V, --verbose  Enable verbose mode (show all program output)\n")
	fmt.Printf("  -D, --debug    Enable debug mode (includes verbose + structured logging)\n")
	fmt.Printf("  [folder]       Specify folder to watch (default: current directory)\n\n")
	fmt.Printf("Examples:\n")
	fmt.Printf("  %s -h                # Show help\n", filepath.Base(os.Args[0]))
	fmt.Printf("  %s -v                # Show version\n", filepath.Base(os.Args[0]))
	fmt.Printf("  %s -V                # Run in verbose mode\n", filepath.Base(os.Args[0]))
	fmt.Printf("  %s -D                # Run in debug mode\n", filepath.Base(os.Args[0]))
	fmt.Printf("  %s /path/to/pdfs     # Watch specific folder\n", filepath.Base(os.Args[0]))
	fmt.Printf("  %s -V /path/to/pdfs  # Verbose mode with specific folder\n", filepath.Base(os.Args[0]))
	fmt.Printf("  %s                   # Watch current directory\n\n", filepath.Base(os.Args[0]))
	fmt.Printf("Interactive options:\n")
	fmt.Printf("  S - Move a single PDF file to the output directory\n")
	fmt.Printf("  M - Merge two PDF files (first file + reversed second file)\n")
	fmt.Printf("  H - Show this help information\n")
	fmt.Printf("  V - Toggle verbose mode\n")
	fmt.Printf("  D - Toggle debug mode\n")
	fmt.Printf("  Q - Quit the program\n\n")
}

func processMenu() {
	fmt.Println()
	displayFileCounts()
	showFilePreview()
	
	fmt.Printf("Options: %s[S]%single, %s[M]%serge, %s[H]%selp, %s[V]%serbose, %s[D]%sebug, %s[Q]%suit\n", 
		YELLOW, NC, YELLOW, NC, YELLOW, NC, YELLOW, NC, YELLOW, NC, YELLOW, NC)
	
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter choice (S/M/H/V/D/Q): ")
	
	// Create a channel to receive the input
	inputChan := make(chan string, 1)
	errorChan := make(chan error, 1)
	
	// Start a goroutine to read input
	go func() {
		input, err := reader.ReadString('\n')
		if err != nil {
			errorChan <- err
		} else {
			inputChan <- input
		}
	}()
	
	// Wait for input or timeout (300 seconds = 5 minutes)
	select {
	case input := <-inputChan:
		input = strings.TrimSpace(strings.ToUpper(input))
		
		switch input {
		case "S":
			processSingleFileWithValidation()
		case "M":
			processMergeFilesWithValidation()
		case "H":
			showHelp()
		case "V":
			VERBOSE = !VERBOSE
			if VERBOSE {
				printSuccess("Verbose mode enabled")
			} else {
				printInfo("Verbose mode disabled")
			}
		case "D":
			DEBUG = !DEBUG
			if DEBUG {
				VERBOSE = true // Debug implies verbose
				if debugLogger == nil {
					initLoggers()
				}
				printSuccess("Debug mode enabled (includes verbose)")
			} else {
				printInfo("Debug mode disabled")
			}
		case "Q":
			fmt.Printf("%sExiting program...%s\n", YELLOW, NC)
			CONTINUE = false
		default:
			printWarning("Invalid choice. Please enter S, M, H, V, D, or Q.")
		}
		
	case err := <-errorChan:
		// Handle EOF gracefully (happens when input is piped)
		if err == io.EOF {
			fmt.Printf("\n%sEnd of input reached. Exiting...%s\n", YELLOW, NC)
			CONTINUE = false
			return
		}
		printError(fmt.Sprintf("Error reading input: %v", err))
		
	case <-time.After(300 * time.Second): // 5 minutes timeout
		fmt.Printf("\n%sTimeout reached (5 minutes). Exiting...%s\n", YELLOW, NC)
		CONTINUE = false
		cleanup()
		os.Exit(7) // Exit code 7 for timeout
	}
}

// Enhanced single file processing with validation
func processSingleFileWithValidation() {
	startTime := time.Now()
	
	files, err := findPDFFiles()
	if err != nil {
		printError(fmt.Sprintf("Error finding PDF files: %v", err))
		return
	}

	if len(files) == 0 {
		printWarning("No PDF files found in " + FOLDER)
		return
	}

	file := files[0]
	filename := filepath.Base(file)
	
	if DEBUG {
		printDebug(fmt.Sprintf("Processing single file: %s", filename))
	}
	
	// Validate PDF before processing
	if err := validatePDFFile(file); err != nil {
		printError(fmt.Sprintf("'%s' is not a valid PDF file: %v", filename, err))
		// Move invalid file to error directory
		destFile := filepath.Join(ERROR_DIR, filename)
		if moveErr := moveFileWithRecovery(file, destFile); moveErr != nil {
			printError(fmt.Sprintf("Failed to move invalid file to error directory: %v", moveErr))
		} else {
			ERROR_COUNT++
			printError("Invalid PDF moved to error folder")
		}
		logOperation("SINGLE_FILE_INVALID", filename, "", "FAILED")
		return
	}
	
	if VERBOSE {
		filesize := getHumanReadableSize(file)
		fmt.Printf("Processing: %s%s%s (%s)\n", YELLOW, filename, NC, filesize)
	}

	// Get file size for performance monitoring
	var fileSize int64
	if info, err := os.Stat(file); err == nil {
		fileSize = info.Size()
	}

	// Move file to output directory
	destFile := filepath.Join(OUTPUT, filename)
	if err := moveFileWithRecovery(file, destFile); err != nil {
		printError(fmt.Sprintf("Failed to move %s: %v", filename, err))
		logOperation("SINGLE_FILE_MOVE", filename, "", "FAILED")
		return
	}

	duration := time.Since(startTime)
	COUNTER++
	printSuccess(fmt.Sprintf("File moved. (%d)", COUNTER))
	
	logOperation("SINGLE_FILE_MOVE", filename, "", "SUCCESS")
	logPerformance("SINGLE_FILE_MOVE", duration, fileSize)
}

// Enhanced merge processing with validation
func processMergeFilesWithValidation() {
	startTime := time.Now()
	
	files, err := findPDFFiles()
	if err != nil {
		printError(fmt.Sprintf("Error finding PDF files: %v", err))
		return
	}

	if len(files) < 2 {
		printWarning(fmt.Sprintf("Did not find two PDF files in %s", FOLDER))
		return
	}

	file1 := files[0]
	file2 := files[1]
	
	if DEBUG {
		printDebug(fmt.Sprintf("Processing merge: %s + %s", filepath.Base(file1), filepath.Base(file2)))
	}

	// Validate both PDFs before processing
	if err := validatePDFFile(file1); err != nil {
		printError(fmt.Sprintf("First PDF '%s' is invalid: %v", filepath.Base(file1), err))
		moveInvalidFiles(file1, file2)
		logOperation("MERGE_INVALID", filepath.Base(file1), filepath.Base(file2), "FAILED")
		return
	}
	
	if err := validatePDFFile(file2); err != nil {
		printError(fmt.Sprintf("Second PDF '%s' is invalid: %v", filepath.Base(file2), err))
		moveInvalidFiles(file1, file2)
		logOperation("MERGE_INVALID", filepath.Base(file1), filepath.Base(file2), "FAILED")
		return
	}

	fmt.Printf("Merging: %s%s%s %s%s%s -> %s%s%s\n", 
		BLUE, filepath.Base(file1), NC, 
		BLUE, filepath.Base(file2), NC,
		GREEN, filepath.Base(file1)+"-"+filepath.Base(file2), NC)

	// Get combined file size for performance monitoring
	var totalSize int64
	if info1, err := os.Stat(file1); err == nil {
		totalSize += info1.Size()
	}
	if info2, err := os.Stat(file2); err == nil {
		totalSize += info2.Size()
	}

	if VERBOSE {
		size1 := getHumanReadableSize(file1)
		size2 := getHumanReadableSize(file2)
		fmt.Printf("File 1 size: %s\n", size1)
		fmt.Printf("File 2 size: %s\n", size2)
	}

	// Create output filename (combine both names with hyphen)
	name1 := strings.TrimSuffix(filepath.Base(file1), filepath.Ext(file1))
	name2 := strings.TrimSuffix(filepath.Base(file2), filepath.Ext(file2))
	outputFile := filepath.Join(OUTPUT, name1+"-"+name2+".pdf")

	// Process and merge the files with smart page reversal logic
	processAndMerge(outputFile, file1, file2, 0) // pages parameter not used in new implementation
	
	duration := time.Since(startTime)
	logOperation("MERGE", filepath.Base(file1), filepath.Base(file2), "COMPLETED")
	logPerformance("MERGE", duration, totalSize)
}

// Move invalid files to error directory
func moveInvalidFiles(files ...string) {
	for _, file := range files {
		if file == "" {
			continue
		}
		filename := filepath.Base(file)
		destFile := filepath.Join(ERROR_DIR, filename)
		if err := moveFileWithRecovery(file, destFile); err != nil {
			printError(fmt.Sprintf("Failed to move invalid file %s: %v", filename, err))
		}
	}
	ERROR_COUNT++
	printError("Invalid PDF files moved to error folder")
}

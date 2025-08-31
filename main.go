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

	"github.com/kris/blendpdfgo/ui"
)

func main() {
	initializeApplication()
	setupSignalHandling()
	
	if err := setupLockFile(); err != nil {
		handleLockFileError(err)
	}

	folder, err := parseCommandLineArgs()
	if err != nil {
		handleStartupError(err)
	}

	if err := setupApplicationDirectories(folder); err != nil {
		handleStartupError(err)
	}

	// Try to run TUI, fallback to original interface if needed
	if err := runTUI(); err != nil {
		// Fallback to original interface
		runMainLoop()
	}
	
	cleanup()
}

// Run the new TUI interface
func runTUI() error {
	// Create file operations bridge
	bridge := ui.NewFileOpsBridge(FOLDER, ARCHIVE, OUTPUT, ERROR_DIR)
	
	// Set function pointers to existing operations
	bridge.SetFunctions(
		func(dir string) ([]string, error) { return findPDFFiles() }, // Wrapper for existing function
		countPDFFiles,
		getHumanReadableSize,
		func() error { processSingleFileOperation(); return nil }, // Wrapper for void function
		func() error { processMergeOperation(); return nil },       // Wrapper for void function
	)
	
	// Create and run enhanced menu (no complex TUI)
	menu := ui.NewEnhancedMenu(FOLDER, ARCHIVE, OUTPUT, ERROR_DIR, VERSION, bridge)
	return menu.Run()
}

// Initialize application components
func initializeApplication() {
	if DEBUG {
		initLoggers()
	}
}

// Setup signal handling for graceful shutdown
func setupSignalHandling() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		fmt.Printf("\n%sShutting down gracefully...%s\n", YELLOW, NC)
		cleanup()
		os.Exit(0)
	}()
}

// Handle lock file setup errors
func setupLockFile() error {
	return setupLock()
}

// Handle lock file errors with appropriate exit codes
func handleLockFileError(err error) {
	if strings.Contains(err.Error(), "exit code 6") {
		os.Exit(6)
	}
	printError(err.Error())
	os.Exit(1)
}

// Parse command line arguments with error handling
func parseCommandLineArgs() (string, error) {
	return parseArgs()
}

// Setup application directories
func setupApplicationDirectories(folder string) error {
	return setupDirectories(folder)
}

// Handle startup errors
func handleStartupError(err error) {
	printError(err.Error())
	cleanupLock()
	os.Exit(1)
}

// Run the main application loop
func runMainLoop() {
	for CONTINUE {
		processUserMenu()
	}
}

// Process user menu interaction
func processUserMenu() {
	displayApplicationStatus()
	choice, err := getUserChoice()
	
	if err != nil {
		handleInputError(err)
		return
	}
	
	executeUserChoice(choice)
}

// Display current application status
func displayApplicationStatus() {
	fmt.Println()
	displayFileCounts()
	showFilePreview()
	displayMenuOptions()
}

// Display menu options
func displayMenuOptions() {
	fmt.Printf("Options: %s[S]%single, %s[M]%serge, %s[H]%selp, %s[V]%serbose, %s[D]%sebug, %s[Q]%suit\n", 
		YELLOW, NC, YELLOW, NC, YELLOW, NC, YELLOW, NC, YELLOW, NC, YELLOW, NC)
}

// Get user choice
func getUserChoice() (string, error) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter choice (S/M/H/V/D/Q): ")
	
	input, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}
	
	return strings.TrimSpace(strings.ToUpper(input)), nil
}

// Handle input errors
func handleInputError(err error) {
	if err == io.EOF {
		fmt.Printf("\n%sEnd of input reached. Exiting...%s\n", YELLOW, NC)
		CONTINUE = false
		return
	}
	printError(fmt.Sprintf("Error reading input: %v", err))
}

// Execute user menu choice
func executeUserChoice(choice string) {
	switch choice {
	case "S":
		processSingleFileOperation()
	case "M":
		processMergeOperation()
	case "H":
		showApplicationHelp()
	case "V":
		toggleVerboseMode()
	case "D":
		toggleDebugMode()
	case "Q":
		exitApplication()
	default:
		printWarning("Invalid choice. Please enter S, M, H, V, D, or Q.")
	}
}

// Process single file operation
func processSingleFileOperation() {
	processSingleFileWithValidation()
}

// Process merge operation
func processMergeOperation() {
	processMergeFilesWithValidation()
}

// Show application help
func showApplicationHelp() {
	showHelp()
}

// Toggle verbose mode
func toggleVerboseMode() {
	VERBOSE = !VERBOSE
	if VERBOSE {
		printSuccess("Verbose mode enabled")
	} else {
		printInfo("Verbose mode disabled")
	}
}

// Toggle debug mode
func toggleDebugMode() {
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
}

// Exit application
func exitApplication() {
	fmt.Printf("%sExiting program...%s\n", YELLOW, NC)
	CONTINUE = false
}

// Show comprehensive help information
func showHelp() {
	fmt.Printf("BlendPDF v%s - A tool for merging PDF files\n\n", VERSION)
	fmt.Printf("Usage: %s [options] [folder]\n\n", filepath.Base(os.Args[0]))
	
	showCommandLineOptions()
	showUsageExamples()
	showInteractiveOptions()
}

// Show command line options
func showCommandLineOptions() {
	fmt.Printf("Command line options:\n")
	fmt.Printf("  -h, --help     Show this help information and exit\n")
	fmt.Printf("  -v, --version  Show version information and exit\n")
	fmt.Printf("  -V, --verbose  Enable verbose mode (show all program output)\n")
	fmt.Printf("  -D, --debug    Enable debug mode (includes verbose + structured logging)\n")
	fmt.Printf("  [folder]       Specify folder to watch (default: current directory)\n\n")
}

// Show usage examples
func showUsageExamples() {
	baseName := filepath.Base(os.Args[0])
	fmt.Printf("Examples:\n")
	fmt.Printf("  %s -h                # Show help\n", baseName)
	fmt.Printf("  %s -v                # Show version\n", baseName)
	fmt.Printf("  %s -V                # Run in verbose mode\n", baseName)
	fmt.Printf("  %s -D                # Run in debug mode\n", baseName)
	fmt.Printf("  %s /path/to/pdfs     # Watch specific folder\n", baseName)
	fmt.Printf("  %s -V /path/to/pdfs  # Verbose mode with specific folder\n", baseName)
	fmt.Printf("  %s                   # Watch current directory\n\n", baseName)
}

// Show interactive options
func showInteractiveOptions() {
	fmt.Printf("Interactive options:\n")
	fmt.Printf("  S - Move a single PDF file to the output directory\n")
	fmt.Printf("  M - Merge two PDF files (first file + reversed second file)\n")
	fmt.Printf("  H - Show this help information\n")
	fmt.Printf("  V - Toggle verbose mode\n")
	fmt.Printf("  D - Toggle debug mode\n")
	fmt.Printf("  Q - Quit the program\n\n")
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
	
	logDebugOperation("Processing single file", filename)
	
	if err := validateAndProcessSingleFile(file, filename, startTime); err != nil {
		handleSingleFileError(file, filename, err)
	}
}

// Validate and process single file
func validateAndProcessSingleFile(file, filename string, startTime time.Time) error {
	if err := validatePDFFile(file); err != nil {
		return fmt.Errorf("validation failed: %v", err)
	}
	
	if VERBOSE {
		filesize := getHumanReadableSize(file)
		fmt.Printf("Processing: %s%s%s (%s)\n", YELLOW, filename, NC, filesize)
	}

	fileSize := getFileSize(file)
	destFile := filepath.Join(OUTPUT, filename)
	
	if err := moveFileWithRecovery(file, destFile); err != nil {
		return fmt.Errorf("move failed: %v", err)
	}

	recordSuccessfulOperation(startTime, filename, fileSize)
	return nil
}

// Handle single file processing errors
func handleSingleFileError(file, filename string, err error) {
	printError(fmt.Sprintf("'%s' processing failed: %v", filename, err))
	
	destFile := filepath.Join(ERROR_DIR, filename)
	if moveErr := moveFileWithRecovery(file, destFile); moveErr != nil {
		printError(fmt.Sprintf("Failed to move invalid file to error directory: %v", moveErr))
	} else {
		ERROR_COUNT++
		printError("Invalid PDF moved to error folder")
	}
	logOperation("SINGLE_FILE_INVALID", filename, "", "FAILED")
}

// Record successful operation
func recordSuccessfulOperation(startTime time.Time, filename string, fileSize int64) {
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

	file1, file2 := files[0], files[1]
	logDebugOperation("Processing merge", fmt.Sprintf("%s + %s", filepath.Base(file1), filepath.Base(file2)))

	if err := validateAndProcessMerge(file1, file2, startTime); err != nil {
		handleMergeError(file1, file2, err)
	}
}

// Validate and process merge operation
func validateAndProcessMerge(file1, file2 string, startTime time.Time) error {
	if err := validateBothPDFs(file1, file2); err != nil {
		return err
	}

	displayMergeInfo(file1, file2)
	totalSize := getFileSize(file1) + getFileSize(file2)
	outputFile := createMergeOutputPath(file1, file2)

	processAndMerge(outputFile, file1, file2, 0)
	
	duration := time.Since(startTime)
	logOperation("MERGE", filepath.Base(file1), filepath.Base(file2), "COMPLETED")
	logPerformance("MERGE", duration, totalSize)
	
	return nil
}

// Validate both PDFs for merge
func validateBothPDFs(file1, file2 string) error {
	if err := validatePDFFile(file1); err != nil {
		return fmt.Errorf("first PDF '%s' is invalid: %v", filepath.Base(file1), err)
	}
	
	if err := validatePDFFile(file2); err != nil {
		return fmt.Errorf("second PDF '%s' is invalid: %v", filepath.Base(file2), err)
	}
	
	return nil
}

// Display merge information
func displayMergeInfo(file1, file2 string) {
	name1 := strings.TrimSuffix(filepath.Base(file1), filepath.Ext(file1))
	name2 := strings.TrimSuffix(filepath.Base(file2), filepath.Ext(file2))
	fmt.Printf("Merging: %s%s%s %s%s%s -> %s%s%s\n", 
		BLUE, filepath.Base(file1), NC, 
		BLUE, filepath.Base(file2), NC,
		GREEN, name1+"-"+name2+".pdf", NC)

	if VERBOSE {
		size1 := getHumanReadableSize(file1)
		size2 := getHumanReadableSize(file2)
		fmt.Printf("File 1 size: %s\n", size1)
		fmt.Printf("File 2 size: %s\n", size2)
	}
}

// Create merge output file path
func createMergeOutputPath(file1, file2 string) string {
	name1 := strings.TrimSuffix(filepath.Base(file1), filepath.Ext(file1))
	name2 := strings.TrimSuffix(filepath.Base(file2), filepath.Ext(file2))
	return filepath.Join(OUTPUT, name1+"-"+name2+".pdf")
}

// Handle merge processing errors
func handleMergeError(file1, file2 string, err error) {
	printError(fmt.Sprintf("Merge processing failed: %v", err))
	moveInvalidFiles(file1, file2)
	logOperation("MERGE_INVALID", filepath.Base(file1), filepath.Base(file2), "FAILED")
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

// Get file size safely
func getFileSize(filepath string) int64 {
	if info, err := os.Stat(filepath); err == nil {
		return info.Size()
	}
	return 0
}

// Log debug operation if debug mode is enabled
func logDebugOperation(operation, details string) {
	if DEBUG {
		printDebug(fmt.Sprintf("%s: %s", operation, details))
	}
}

// Cleanup resources and show statistics
func cleanup() {
	// Statistics now handled by enhanced menu UI
	cleanupLock()
}

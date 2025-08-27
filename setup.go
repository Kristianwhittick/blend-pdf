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
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

// Setup lock file to prevent multiple instances
func setupLock() error {
	LOCKFILE = filepath.Join(os.TempDir(), "blendpdfgo.lock")

	// Check if lock file exists
	if _, err := os.Stat(LOCKFILE); err == nil {
		printError("Another instance is already running")
		printInfo(fmt.Sprintf("Lock file exists: %s", LOCKFILE))
		printInfo("If you're sure no other instance is running, remove the lock file manually")
		return fmt.Errorf("already running (exit code 6)")
	}

	// Create lock file with process ID
	file, err := os.Create(LOCKFILE)
	if err != nil {
		return fmt.Errorf("failed to create lock file: %v", err)
	}
	
	// Write process ID to lock file
	_, err = file.WriteString(strconv.Itoa(os.Getpid()))
	if err != nil {
		file.Close()
		os.Remove(LOCKFILE)
		return fmt.Errorf("failed to write to lock file: %v", err)
	}
	
	file.Close()
	
	if VERBOSE {
		printInfo(fmt.Sprintf("Created lock file: %s", LOCKFILE))
	}

	return nil
}

// Clean up lock file and resources
func cleanupLock() {
	if LOCKFILE != "" {
		if err := os.Remove(LOCKFILE); err != nil && VERBOSE {
			printWarning(fmt.Sprintf("Failed to remove lock file: %v", err))
		} else if VERBOSE {
			printInfo("Removed lock file")
		}
	}
}

// Parse command line arguments with enhanced error handling
func parseArgs() (string, error) {
	args := os.Args[1:]
	folder := ""
	
	for i, arg := range args {
		switch arg {
		case "-h", "--help":
			showHelp()
			os.Exit(0)
		case "-v", "--version":
			fmt.Printf("BlendPDFGo v%s\n", VERSION)
			os.Exit(0)
		case "-V", "--verbose":
			VERBOSE = true
			printSuccess("Verbose mode enabled")
		default:
			// Check if it's a flag we don't recognize
			if strings.HasPrefix(arg, "-") {
				return "", fmt.Errorf("unknown flag: %s", arg)
			}
			// Assume it's a folder path
			if folder == "" {
				folder = arg
			} else {
				return "", fmt.Errorf("multiple folder paths specified")
			}
		}
		
		// Handle combined flags like -V /path/to/folder
		if arg == "-V" || arg == "--verbose" {
			if i+1 < len(args) && !strings.HasPrefix(args[i+1], "-") {
				folder = args[i+1]
				break
			}
		}
	}
	
	// Use current directory if no folder specified
	if folder == "" {
		var err error
		folder, err = os.Getwd()
		if err != nil {
			return "", fmt.Errorf("failed to get current directory: %v", err)
		}
	}
	
	// Resolve absolute path
	absFolder, err := filepath.Abs(folder)
	if err != nil {
		return "", fmt.Errorf("failed to resolve absolute path for '%s': %v", folder, err)
	}
	
	return absFolder, nil
}

// Enhanced directory setup with comprehensive error handling
func setupDirectories(folder string) error {
	// Validate folder exists and is accessible
	info, err := os.Stat(folder)
	if os.IsNotExist(err) {
		return fmt.Errorf("directory '%s' does not exist", folder)
	}
	if err != nil {
		return fmt.Errorf("cannot access directory '%s': %v", folder, err)
	}
	if !info.IsDir() {
		return fmt.Errorf("'%s' is not a directory", folder)
	}

	// Check if directory is writable
	testFile := filepath.Join(folder, ".blendpdfgo_test")
	file, err := os.Create(testFile)
	if err != nil {
		return fmt.Errorf("directory '%s' is not writable: %v", folder, err)
	}
	file.Close()
	os.Remove(testFile)

	FOLDER = folder
	ARCHIVE = filepath.Join(folder, "archive")
	OUTPUT = filepath.Join(folder, "output")
	ERROR_DIR = filepath.Join(folder, "error")

	// Create required directories with error handling
	dirs := []string{ARCHIVE, OUTPUT, ERROR_DIR}
	for _, dir := range dirs {
		if err := os.MkdirAll(dir, 0755); err != nil {
			return fmt.Errorf("failed to create directory '%s': %v", dir, err)
		}
		
		// Verify directory was created and is writable
		if info, err := os.Stat(dir); err != nil {
			return fmt.Errorf("failed to verify directory '%s': %v", dir, err)
		} else if !info.IsDir() {
			return fmt.Errorf("'%s' exists but is not a directory", dir)
		}
	}

	// Display setup information
	fmt.Printf("Watching folder: %s%s%s\n", BLUE, FOLDER, NC)
	fmt.Printf("Archive  folder: %s%s%s\n", BLUE, ARCHIVE, NC)
	fmt.Printf("Output   folder: %s%s%s\n", BLUE, OUTPUT, NC)
	fmt.Printf("Error    folder: %s%s%s\n", BLUE, ERROR_DIR, NC)
	fmt.Println()

	return nil
}

// Validate PDF file with enhanced error reporting
func validatePDFFile(file string) error {
	// Check if file exists
	info, err := os.Stat(file)
	if os.IsNotExist(err) {
		return fmt.Errorf("file does not exist")
	}
	if err != nil {
		return fmt.Errorf("cannot access file: %v", err)
	}
	
	// Check if it's a regular file
	if !info.Mode().IsRegular() {
		return fmt.Errorf("not a regular file")
	}
	
	// Check file size
	if info.Size() == 0 {
		return fmt.Errorf("file is empty")
	}
	
	// Check file extension
	if !strings.HasSuffix(strings.ToLower(file), ".pdf") {
		return fmt.Errorf("file does not have .pdf extension")
	}
	
	// Validate PDF structure using pdfcpu
	if !validatePDF(file) {
		return fmt.Errorf("invalid PDF structure")
	}
	
	return nil
}

// Enhanced file operation with error recovery
func moveFileWithRecovery(src, dst string) error {
	// Ensure destination directory exists
	dstDir := filepath.Dir(dst)
	if err := os.MkdirAll(dstDir, 0755); err != nil {
		return fmt.Errorf("failed to create destination directory: %v", err)
	}
	
	// Check if destination file already exists
	if _, err := os.Stat(dst); err == nil {
		// Generate unique filename
		base := strings.TrimSuffix(filepath.Base(dst), filepath.Ext(dst))
		ext := filepath.Ext(dst)
		counter := 1
		
		for {
			newDst := filepath.Join(dstDir, fmt.Sprintf("%s_%d%s", base, counter, ext))
			if _, err := os.Stat(newDst); os.IsNotExist(err) {
				dst = newDst
				break
			}
			counter++
			if counter > 1000 {
				return fmt.Errorf("too many duplicate files")
			}
		}
		
		if VERBOSE {
			printWarning(fmt.Sprintf("Destination exists, using: %s", filepath.Base(dst)))
		}
	}
	
	// Perform the move
	err := os.Rename(src, dst)
	if err != nil {
		// Try copy and delete as fallback
		if copyErr := copyFile(src, dst); copyErr != nil {
			return fmt.Errorf("move failed: %v, copy fallback failed: %v", err, copyErr)
		}
		
		if deleteErr := os.Remove(src); deleteErr != nil {
			printWarning(fmt.Sprintf("Original file not deleted: %v", deleteErr))
		}
	}
	
	return nil
}

// Copy file as fallback for move operations
func copyFile(src, dst string) error {
	sourceFile, err := os.Open(src)
	if err != nil {
		return err
	}
	defer sourceFile.Close()
	
	destFile, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer destFile.Close()
	
	_, err = destFile.ReadFrom(sourceFile)
	return err
}

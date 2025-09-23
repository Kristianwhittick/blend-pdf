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
	"log"
	"os"
	"path/filepath"
	"sort"
	"time"
)

// Directory and file management functions

// Setup application directories
func setupDirectories(folder string) error {
	FOLDER = folder
	ARCHIVE = filepath.Join(folder, "archive")
	OUTPUT = filepath.Join(folder, "output")
	ERROR_DIR = filepath.Join(folder, "error")

	if err := createRequiredDirectories(); err != nil {
		return err
	}

	displayDirectoryPaths()
	return nil
}

// Create required directories if they don't exist
func createRequiredDirectories() error {
	// Always create archive and error directories
	dirs := []string{ARCHIVE, ERROR_DIR}

	// Only create default output directory if not using multi-output folders
	if CONFIG == nil || len(CONFIG.OutputFolders) == 0 {
		dirs = append(dirs, OUTPUT)
	} else {
		// Create multi-output folders at startup
		dirs = append(dirs, CONFIG.OutputFolders...)
	}

	for _, dir := range dirs {
		if err := os.MkdirAll(dir, 0750); err != nil {
			return fmt.Errorf("failed to create directory %s: %v", dir, err)
		}
	}
	return nil
}

// Display directory paths
func displayDirectoryPaths() {
	fmt.Printf("Watching folder: %s%s%s\n", BLUE, FOLDER, NC)
	fmt.Printf("Archive  folder: %s%s%s\n", BLUE, ARCHIVE, NC)
	fmt.Printf("Output   folder: %s%s%s\n", BLUE, OUTPUT, NC)
	fmt.Printf("Error    folder: %s%s%s\n", BLUE, ERROR_DIR, NC)
	fmt.Println()
}

// Find PDF files in the main directory
func findPDFFiles() ([]string, error) {
	pattern := filepath.Join(FOLDER, "*.pdf")
	files, err := filepath.Glob(pattern)
	if err != nil {
		return nil, err
	}

	sort.Strings(files)
	return files, nil
}

// File counting and display functions

// Get file counts for all directories
func getFileCounts() (int, int, int, int) {
	return countPDFFiles(FOLDER), countPDFFiles(ARCHIVE), countPDFFiles(OUTPUT), countPDFFiles(ERROR_DIR)
}

// Count PDF files in a specific directory
func countPDFFiles(dir string) int {
	files, err := filepath.Glob(filepath.Join(dir, "*.pdf"))
	if err != nil {
		return 0
	}
	return len(files)
}

// Display file counts before menu
func displayFileCounts() {
	main, archive, output, errorCount := getFileCounts()
	fmt.Printf("Files: Main(%d) Archive(%d) Output(%d) Error(%d)\n", main, archive, output, errorCount)
}

// Show file preview in verbose mode
func showFilePreview() {
	if !VERBOSE {
		return
	}

	files, err := findPDFFiles()
	if err != nil || len(files) == 0 {
		return
	}

	displayFileList(files)
}

// Display list of files with size information
func displayFileList(files []string) {
	fmt.Printf("%sAvailable PDF files:%s\n", BLUE, NC)

	maxDisplay := 5
	displayedCount := displayFiles(files, maxDisplay)

	if len(files) > maxDisplay {
		remaining := len(files) - displayedCount
		fmt.Printf("  ... and %d more file(s)\n", remaining)
	}
	fmt.Println()
}

// Display individual files up to maximum count
func displayFiles(files []string, maxDisplay int) int {
	displayed := 0
	for _, file := range files {
		if displayed >= maxDisplay {
			break
		}

		filename := filepath.Base(file)
		filesize := getHumanReadableSize(file)
		fmt.Printf("  %s%s%s (%s)\n", YELLOW, filename, NC, filesize)
		displayed++
	}
	return displayed
}

// File size utilities

// Get human readable file size
func getHumanReadableSize(filepath string) string {
	info, err := os.Stat(filepath)
	if err != nil {
		return "unknown"
	}

	return formatFileSize(info.Size())
}

// Format file size in human readable format
func formatFileSize(size int64) string {
	const (
		KB = 1024
		MB = KB * 1024
		GB = MB * 1024
	)

	switch {
	case size < KB:
		return fmt.Sprintf("%dB", size)
	case size < MB:
		return fmt.Sprintf("%.1fK", float64(size)/KB)
	case size < GB:
		return fmt.Sprintf("%.1fM", float64(size)/MB)
	default:
		return fmt.Sprintf("%.1fG", float64(size)/GB)
	}
}

// File movement and processing functions

// Move processed files to destination directory with enhanced error handling
func moveProcessedFiles(destination, message string, files ...string) {
	moveResults := processMoveOperations(destination, files)
	handleMoveResults(moveResults, destination, message)
}

// Process individual move operations
func processMoveOperations(destination string, files []string) []moveResult {
	var results []moveResult

	for _, file := range files {
		if file == "" {
			continue
		}

		result := moveResult{
			filename: filepath.Base(file),
			success:  false,
		}

		destFile := filepath.Join(destination, result.filename)
		if err := moveFileWithRecovery(file, destFile); err != nil {
			result.error = err
		} else {
			result.success = true
		}

		results = append(results, result)
	}

	return results
}

// Handle results of move operations
func handleMoveResults(results []moveResult, destination, message string) {
	allMoved := true

	for _, result := range results {
		if !result.success {
			printError(fmt.Sprintf("Failed to move %s: %v", result.filename, result.error))
			allMoved = false
		} else if VERBOSE {
			printInfo(fmt.Sprintf("Moved %s to %s", result.filename, filepath.Base(destination)))
		}
	}

	updateCountersBasedOnResults(allMoved, destination, message)
}

// Update counters based on move operation results
func updateCountersBasedOnResults(allMoved bool, destination, message string) {
	if allMoved {
		if destination == ARCHIVE {
			COUNTER++
			printSuccess(fmt.Sprintf("%s (%d)", message, COUNTER))
		} else {
			ERROR_COUNT++
			printError(message)
		}
	} else {
		ERROR_COUNT++
		printError("Some files could not be moved")
	}
}

// Statistics and reporting functions

// Show session statistics on exit - REMOVED: Now handled by UI
// func showStatistics() - Moved to enhanced menu UI

// Display operation counts
func displayOperationCounts() {
	fmt.Printf("Successful operations: %s%d%s\n", GREEN, COUNTER, NC)
	fmt.Printf("Errors encountered: %s%d%s\n", RED, ERROR_COUNT, NC)
}

// Display elapsed time in appropriate format
func displayElapsedTime(elapsed time.Duration) {
	minutes := int(elapsed.Minutes())
	seconds := int(elapsed.Seconds()) % 60

	if minutes > 0 {
		fmt.Printf("Time elapsed: %dm %ds\n", minutes, seconds)
	} else {
		fmt.Printf("Time elapsed: %ds\n", seconds)
	}
}

// Output and logging functions

// Colored output helper functions
func printSuccess(message string) {
	fmt.Printf("%sSuccess:%s %s\n", GREEN, NC, message)
	logIfDebugEnabled("SUCCESS", message, infoLogger)
}

func printError(message string) {
	fmt.Printf("%sError:%s %s\n", RED, NC, message)
	ERROR_COUNT++
	logIfDebugEnabled("ERROR", message, errorLogger)
}

func printWarning(message string) {
	fmt.Printf("%sWarning:%s %s\n", YELLOW, NC, message)
	logIfDebugEnabled("WARNING", message, warnLogger)
}

func printInfo(message string) {
	fmt.Printf("%sInfo:%s %s\n", BLUE, NC, message)
	logIfDebugEnabled("INFO", message, infoLogger)
}

func printDebug(message string) {
	if DEBUG {
		fmt.Printf("[DEBUG] %s\n", message)
		logIfDebugEnabled("DEBUG", message, debugLogger)
	}
}

// Structured logging functions
func logOperation(operation, file1, file2, result string) {
	if !DEBUG {
		return
	}

	if file2 != "" {
		infoLogger.Printf("OPERATION: %s | Files: %s, %s | Result: %s", operation, file1, file2, result)
	} else {
		infoLogger.Printf("OPERATION: %s | File: %s | Result: %s", operation, file1, result)
	}
}

func logPerformance(operation string, duration time.Duration, fileSize int64) {
	if DEBUG && duration.Seconds() > 0 {
		speed := float64(fileSize) / (1024 * 1024) / duration.Seconds()
		infoLogger.Printf("PERFORMANCE: %s | Duration: %v | Size: %d bytes | Speed: %.2f MB/s",
			operation, duration, fileSize, speed)
	}
}

// Helper types and functions

// moveResult represents the result of a file move operation
type moveResult struct {
	filename string
	success  bool
	error    error
}

// Log message if debug mode is enabled
func logIfDebugEnabled(level, message string, logger interface{}) {
	if DEBUG && logger != nil {
		switch l := logger.(type) {
		case *log.Logger:
			l.Printf("%s: %s", level, message)
		}
	}
}

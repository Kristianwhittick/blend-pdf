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
	"sort"
	"strings"
	"time"
)

// Get file counts for display
func getFileCounts() (int, int, int, int) {
	mainCount := countPDFFiles(FOLDER)
	archiveCount := countPDFFiles(ARCHIVE)
	outputCount := countPDFFiles(OUTPUT)
	errorCount := countPDFFiles(ERROR_DIR)
	return mainCount, archiveCount, outputCount, errorCount
}

// Count PDF files in a directory
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
	
	files, err := filepath.Glob(filepath.Join(FOLDER, "*.pdf"))
	if err != nil || len(files) == 0 {
		return
	}
	
	sort.Strings(files)
	fmt.Printf("%sAvailable PDF files:%s\n", BLUE, NC)
	
	maxDisplay := 5
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
	
	if len(files) > maxDisplay {
		remaining := len(files) - maxDisplay
		fmt.Printf("  ... and %d more file(s)\n", remaining)
	}
	fmt.Println()
}

// Get human readable file size
func getHumanReadableSize(filepath string) string {
	info, err := os.Stat(filepath)
	if err != nil {
		return "unknown"
	}
	
	size := info.Size()
	if size < 1024 {
		return fmt.Sprintf("%dB", size)
	} else if size < 1024*1024 {
		return fmt.Sprintf("%.1fK", float64(size)/1024)
	} else if size < 1024*1024*1024 {
		return fmt.Sprintf("%.1fM", float64(size)/(1024*1024))
	} else {
		return fmt.Sprintf("%.1fG", float64(size)/(1024*1024*1024))
	}
}

// Show session statistics on exit
func showStatistics() {
	elapsed := time.Since(START_TIME)
	minutes := int(elapsed.Minutes())
	seconds := int(elapsed.Seconds()) % 60
	
	fmt.Printf("\n%sSession Statistics:%s\n", BLUE, NC)
	fmt.Printf("Successful operations: %s%d%s\n", GREEN, COUNTER, NC)
	fmt.Printf("Errors encountered: %s%d%s\n", RED, ERROR_COUNT, NC)
	
	if minutes > 0 {
		fmt.Printf("Time elapsed: %dm %ds\n", minutes, seconds)
	} else {
		fmt.Printf("Time elapsed: %ds\n", seconds)
	}
}

// Colored output helper functions
func printSuccess(message string) {
	fmt.Printf("%sSuccess:%s %s\n", GREEN, NC, message)
}

func printError(message string) {
	fmt.Printf("%sError:%s %s\n", RED, NC, message)
	ERROR_COUNT++
}

func printWarning(message string) {
	fmt.Printf("%sWarning:%s %s\n", YELLOW, NC, message)
}

func printInfo(message string) {
	fmt.Printf("%sInfo:%s %s\n", BLUE, NC, message)
}

// Setup directories
func setupDirectories(folder string) error {
	FOLDER = folder
	ARCHIVE = filepath.Join(folder, "archive")
	OUTPUT = filepath.Join(folder, "output")
	ERROR_DIR = filepath.Join(folder, "error")

	// Create directories if they don't exist
	dirs := []string{ARCHIVE, OUTPUT, ERROR_DIR}
	for _, dir := range dirs {
		if err := os.MkdirAll(dir, 0755); err != nil {
			return fmt.Errorf("failed to create directory %s: %v", dir, err)
		}
	}

	fmt.Printf("Watching folder: %s%s%s\n", BLUE, FOLDER, NC)
	fmt.Printf("Archive  folder: %s%s%s\n", BLUE, ARCHIVE, NC)
	fmt.Printf("Output   folder: %s%s%s\n", BLUE, OUTPUT, NC)
	fmt.Printf("Error    folder: %s%s%s\n", BLUE, ERROR_DIR, NC)
	fmt.Println()

	return nil
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

// Move processed files to destination directory with enhanced error handling
func moveProcessedFiles(destination, message string, files ...string) {
	allMoved := true
	
	for _, file := range files {
		if file == "" {
			continue
		}
		
		filename := filepath.Base(file)
		destFile := filepath.Join(destination, filename)
		
		if err := moveFileWithRecovery(file, destFile); err != nil {
			printError(fmt.Sprintf("Failed to move %s: %v", filename, err))
			allMoved = false
		} else if VERBOSE {
			printInfo(fmt.Sprintf("Moved %s to %s", filename, filepath.Base(destination)))
		}
	}
	
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

// Process single file
func processSingleFile() {
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
	
	if VERBOSE {
		filesize := getHumanReadableSize(file)
		fmt.Printf("Processing: %s%s%s (%s)\n", YELLOW, filename, NC, filesize)
	}

	// Move file to output directory
	destFile := filepath.Join(OUTPUT, filename)
	if err := os.Rename(file, destFile); err != nil {
		printError(fmt.Sprintf("Failed to move %s: %v", filename, err))
		return
	}

	COUNTER++
	printSuccess(fmt.Sprintf("File moved. (%d)", COUNTER))
}

// Process merge files
func processMergeFiles() {
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

	fmt.Printf("Merging: %s%s%s %s%s%s -> %s%s%s\n", 
		BLUE, filepath.Base(file1), NC, 
		BLUE, filepath.Base(file2), NC,
		GREEN, filepath.Base(file1)+"-"+filepath.Base(file2), NC)

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
}

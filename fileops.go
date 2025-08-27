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
)

func moveSingleFile() {
	// Find PDF files with both upper and lower case extensions
	lowerFiles, err1 := filepath.Glob(filepath.Join(FOLDER, "*.pdf"))
	upperFiles, err2 := filepath.Glob(filepath.Join(FOLDER, "*.PDF"))

	// Combine the results
	files := append(lowerFiles, upperFiles...)

	if (err1 != nil && err2 != nil) || len(files) == 0 {
		fmt.Printf("%sWarning:%s No PDF files found in %s\n", YELLOW, NC, FOLDER)
		return
	}

	sort.Strings(files)
	file := files[0]

	if validatePDF(file) {
		fmt.Printf("Moving %s%s%s to output...\n", BLUE, filepath.Base(file), NC)
		dest := filepath.Join(OUTPUT, filepath.Base(file))
		err := os.Rename(file, dest)
		if err != nil {
			fmt.Printf("%sError:%s Failed to move file: %v\n", RED, NC, err)
			return
		}
		fmt.Printf("%sSuccess:%s File moved successfully.\n", GREEN, NC)
	}
}

func moveProcessedFiles(destination, message, file1, file2 string, file3 string) {
	var err error

	// Move files
	err = os.Rename(file1, filepath.Join(destination, filepath.Base(file1)))
	if err == nil {
		err = os.Rename(file2, filepath.Join(destination, filepath.Base(file2)))
		if err == nil && file3 != "" {
			err = os.Rename(file3, filepath.Join(destination, filepath.Base(file3)))
		}
	}

	if err != nil {
		fmt.Printf("%sError:%s Failed to move files: %v\n", RED, NC, err)
		return
	}

	if destination == ARCHIVE {
		fmt.Printf("%sSuccess:%s %s\n", GREEN, NC, message)
	} else {
		fmt.Printf("%sError:%s %s\n", RED, NC, message)
	}
}

func mergeFiles() {
	// Find PDF files with both upper and lower case extensions
	lowerFiles, err1 := filepath.Glob(filepath.Join(FOLDER, "*.pdf"))
	upperFiles, err2 := filepath.Glob(filepath.Join(FOLDER, "*.PDF"))

	// Combine the results
	files := append(lowerFiles, upperFiles...)

	if (err1 != nil && err2 != nil) || len(files) < 2 {
		fmt.Printf("%sWarning:%s Did not find two PDF files in %s\n", YELLOW, NC, FOLDER)
		return
	}

	sort.Strings(files)
	file1 := files[0]
	file2 := files[1]

	// Validate both files
	if !validatePDF(file1) || !validatePDF(file2) {
		return
	}

	// Get page counts for both files
	pages1 := getPageCount(file1)
	pages2 := getPageCount(file2)
	
	if pages1 == -1 || pages2 == -1 {
		fmt.Printf("%sError:%s Could not determine page counts\n", RED, NC)
		moveProcessedFiles(ERROR_DIR, "Page count error. Moving files to error folder...", file1, file2, "")
		return
	}

	// Check if page counts match (exact match required)
	if pages1 != pages2 {
		fmt.Printf("%sError:%s Page count mismatch - %s has %d pages, %s has %d pages\n", 
			RED, NC, filepath.Base(file1), pages1, filepath.Base(file2), pages2)
		moveProcessedFiles(ERROR_DIR, "Page count mismatch. Moving files to error folder...", file1, file2, "")
		return
	}

	// Create output filename (combine both names without "_merged")
	name1 := strings.TrimSuffix(filepath.Base(file1), filepath.Ext(file1))
	name2 := strings.TrimSuffix(filepath.Base(file2), filepath.Ext(file2))
	outputFile := filepath.Join(OUTPUT, name1+"_"+name2+".pdf")

	fmt.Printf("Found files:  %s%s%s (%d pages)   %s%s%s (%d pages)\n", 
		BLUE, filepath.Base(file1), NC, pages1, BLUE, filepath.Base(file2), NC, pages2)
	fmt.Printf("Output file:  %s%s%s\n", GREEN, filepath.Base(outputFile), NC)
	fmt.Printf("Merging with interleaved pattern (Doc1_Page1, Doc2_Page%d, Doc1_Page2, Doc2_Page%d, ...)\n", pages2, pages2-1)

	// Process and merge the files with interleaved pattern
	processAndMerge(outputFile, file1, file2, pages1)
}
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
	"strings"

	"github.com/pdfcpu/pdfcpu/pkg/api"
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/model"
)

// Validate PDF file structure
func validatePDF(file string) bool {
	if _, err := os.Stat(file); os.IsNotExist(err) {
		printError(fmt.Sprintf("File '%s' does not exist", file))
		return false
	}

	// Use pdfcpu API to validate the PDF
	conf := model.NewDefaultConfiguration()
	conf.ValidationMode = model.ValidationRelaxed

	validateErr := api.ValidateFile(file, conf)
	if validateErr != nil {
		printError(fmt.Sprintf("'%s' is not a valid PDF file: %v", file, validateErr))
		return false
	}

	return true
}

// Get page count using pdfcpu API
func getPageCount(file string) (int, error) {
	pageCount, err := api.PageCountFile(file)
	if err != nil {
		return -1, fmt.Errorf("could not determine page count for '%s': %v", file, err)
	}

	return pageCount, nil
}

// Create reversed copy of PDF for multi-page files
func createReversedPDF(inputFile, outputFile string, pageCount int) error {
	if pageCount <= 1 {
		return fmt.Errorf("cannot reverse single-page PDF")
	}

	conf := model.NewDefaultConfiguration()
	
	// Extract individual pages in reverse order
	var tempFiles []string
	for i := pageCount; i >= 1; i-- {
		tempFile := fmt.Sprintf("temp_reverse_%d.pdf", i)
		pageSelection, err := api.ParsePageSelection(fmt.Sprintf("%d", i))
		if err != nil {
			return fmt.Errorf("failed to parse page selection for page %d: %v", i, err)
		}
		
		err = api.TrimFile(inputFile, tempFile, pageSelection, conf)
		if err != nil {
			return fmt.Errorf("failed to extract page %d: %v", i, err)
		}
		
		tempFiles = append(tempFiles, tempFile)
	}
	
	if VERBOSE {
		fmt.Printf("rev = %s%d", BLUE, pageCount)
		for i := pageCount - 1; i >= 1; i-- {
			fmt.Printf(",%d", i)
		}
		fmt.Printf("%s\n", NC)
	}

	// Merge the extracted pages to create reversed PDF
	err := api.MergeCreateFile(tempFiles, outputFile, false, conf)
	if err != nil {
		// Clean up temp files on error
		for _, tempFile := range tempFiles {
			os.Remove(tempFile)
		}
		return fmt.Errorf("failed to merge reversed pages: %v", err)
	}
	
	// Clean up temporary files
	for _, tempFile := range tempFiles {
		if removeErr := os.Remove(tempFile); removeErr != nil && VERBOSE {
			fmt.Printf("Warning: Failed to remove temp file %s: %v\n", tempFile, removeErr)
		}
	}

	return nil
}

// Smart merge: direct merge for single-page, reversed merge for multi-page
func smartMerge(file1, file2, outputFile string, pages1, pages2 int) error {
	conf := model.NewDefaultConfiguration()
	
	if pages2 == 1 {
		// Single-page second file: direct merge (no reversal)
		if VERBOSE {
			printInfo("Single-page second file detected - merging directly without reversal")
		}
		
		inputFiles := []string{file1, file2}
		return api.MergeCreateFile(inputFiles, outputFile, false, conf)
		
	} else {
		// Multi-page second file: create reversed copy, then merge
		if VERBOSE {
			printInfo(fmt.Sprintf("Multi-page second file detected (%d pages) - creating reversed copy", pages2))
		}
		
		// Create temporary reversed file
		reversedFile := strings.TrimSuffix(file2, filepath.Ext(file2)) + "-reverse.pdf"
		
		err := createReversedPDF(file2, reversedFile, pages2)
		if err != nil {
			return fmt.Errorf("failed to create reversed PDF: %v", err)
		}
		
		// Ensure cleanup of temporary file
		defer func() {
			if err := os.Remove(reversedFile); err != nil && VERBOSE {
				printWarning(fmt.Sprintf("Failed to clean up temporary file %s: %v", reversedFile, err))
			}
		}()
		
		// Create interleaved merge
		return createInterleavedMerge(file1, reversedFile, outputFile, pages1)
	}
}

// Create interleaved merge pattern (Doc1_Page1, Doc2_Page3, Doc1_Page2, Doc2_Page2, etc.)
func createInterleavedMerge(file1, file2, outputFile string, pageCount int) error {
	conf := model.NewDefaultConfiguration()
	
	// Create temporary files for individual pages
	tempFiles := make([]string, 0, pageCount*2)
	
	// Extract pages in interleaved pattern
	for i := 1; i <= pageCount; i++ {
		// Extract page from first document
		tempFile1 := fmt.Sprintf("temp_A_%d.pdf", i)
		pageSelection1, err := api.ParsePageSelection(fmt.Sprintf("%d", i))
		if err != nil {
			return fmt.Errorf("failed to parse page selection for file1 page %d: %v", i, err)
		}
		
		err = api.TrimFile(file1, tempFile1, pageSelection1, conf)
		if err != nil {
			return fmt.Errorf("failed to extract page %d from file1: %v", i, err)
		}
		
		// Extract corresponding page from second document (reversed file)
		// For interleaved pattern: Doc1_Page1 + Doc2_Page3, Doc1_Page2 + Doc2_Page2, Doc1_Page3 + Doc2_Page1
		// Since file2 is already reversed (3,2,1), we extract page i to get the correct interleaved page
		tempFile2 := fmt.Sprintf("temp_B_%d.pdf", i)
		pageSelection2, err := api.ParsePageSelection(fmt.Sprintf("%d", i))
		if err != nil {
			return fmt.Errorf("failed to parse page selection for file2 page %d: %v", i, err)
		}
		
		err = api.TrimFile(file2, tempFile2, pageSelection2, conf)
		if err != nil {
			return fmt.Errorf("failed to extract page %d from file2: %v", i, err)
		}
		
		// Add both pages to merge list (interleaved)
		tempFiles = append(tempFiles, tempFile1, tempFile2)
	}
	
	// Merge all temporary files using zip mode for proper interleaving
	err := api.MergeCreateFile(tempFiles, outputFile, false, conf)
	
	// Clean up temporary files
	for _, tempFile := range tempFiles {
		if err := os.Remove(tempFile); err != nil && VERBOSE {
			printWarning(fmt.Sprintf("Failed to clean up temporary file %s: %v", tempFile, err))
		}
	}
	
	return err
}

// Enhanced PDF validation before processing
func validatePDFsForMerge(file1, file2 string) (int, int, error) {
	// Validate both files exist and are valid PDFs
	if !validatePDF(file1) {
		return 0, 0, fmt.Errorf("first PDF validation failed")
	}
	
	if !validatePDF(file2) {
		return 0, 0, fmt.Errorf("second PDF validation failed")
	}
	
	// Get page counts
	pages1, err1 := getPageCount(file1)
	pages2, err2 := getPageCount(file2)
	
	if err1 != nil {
		return 0, 0, fmt.Errorf("failed to get page count for first file: %v", err1)
	}
	
	if err2 != nil {
		return 0, 0, fmt.Errorf("failed to get page count for second file: %v", err2)
	}
	
	// Check for exact page count match
	if pages1 != pages2 {
		return pages1, pages2, fmt.Errorf("page count mismatch - %s has %d pages, %s has %d pages", 
			filepath.Base(file1), pages1, filepath.Base(file2), pages2)
	}
	
	if VERBOSE {
		fmt.Printf("pages = %s%d%s\n", BLUE, pages1, NC)
	}
	
	return pages1, pages2, nil
}

// Process and merge files with smart page reversal
func processAndMerge(outputFile, file1, file2 string, pages int) {
	// Validate PDFs before processing
	pages1, pages2, err := validatePDFsForMerge(file1, file2)
	if err != nil {
		printError(err.Error())
		moveProcessedFiles(ERROR_DIR, "PDF validation failed. Moving files to error folder...", file1, file2)
		return
	}
	
	// Use smart merge logic
	err = smartMerge(file1, file2, outputFile, pages1, pages2)
	if err != nil {
		printError(fmt.Sprintf("Failed to merge PDFs: %v", err))
		moveProcessedFiles(ERROR_DIR, "Merge failed. Moving files to error folder...", file1, file2)
	} else {
		moveProcessedFiles(ARCHIVE, "Files merged and moved.", file1, file2)
	}
}

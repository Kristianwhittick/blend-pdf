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
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/pdfcpu/pdfcpu/pkg/api"
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/model"
)

// PDF validation functions

// Validate PDF file structure using pdfcpu
func validatePDF(file string) bool {
	if !fileExists(file) {
		printError(fmt.Sprintf("File '%s' does not exist", file))
		return false
	}

	return validatePDFStructure(file)
}

// Check if file exists
func fileExists(file string) bool {
	_, err := os.Stat(file)
	return !os.IsNotExist(err)
}

// Validate PDF structure using pdfcpu API
func validatePDFStructure(file string) bool {
	conf := createValidationConfig()

	if err := api.ValidateFile(file, conf); err != nil {
		printError(fmt.Sprintf("'%s' is not a valid PDF file: %v", file, err))
		return false
	}

	return true
}

// Create validation configuration
func createValidationConfig() *model.Configuration {
	conf := model.NewDefaultConfiguration()
	conf.ValidationMode = model.ValidationRelaxed
	return conf
}

// PDF page operations

// Get page count using pdfcpu API
func getPageCount(file string) (int, error) {
	pageCount, err := api.PageCountFile(file)
	if err != nil {
		return -1, fmt.Errorf("could not determine page count for '%s': %v", file, err)
	}
	return pageCount, nil
}

// Enhanced PDF validation for merge operations
func validatePDFsForMerge(file1, file2 string) (int, int, error) {
	if err := validateBothPDFFiles(file1, file2); err != nil {
		return 0, 0, err
	}

	pages1, pages2, err := getPageCountsForBothFiles(file1, file2)
	if err != nil {
		return 0, 0, err
	}

	if err := validatePageCountMatch(file1, file2, pages1, pages2); err != nil {
		return pages1, pages2, err
	}

	displayPageCount(pages1)
	return pages1, pages2, nil
}

// Validate both PDF files
func validateBothPDFFiles(file1, file2 string) error {
	if !validatePDF(file1) {
		return fmt.Errorf("first PDF validation failed")
	}

	if !validatePDF(file2) {
		return fmt.Errorf("second PDF validation failed")
	}

	return nil
}

// Get page counts for both files
func getPageCountsForBothFiles(file1, file2 string) (int, int, error) {
	pages1, err1 := getPageCount(file1)
	if err1 != nil {
		return 0, 0, fmt.Errorf("failed to get page count for first file: %v", err1)
	}

	pages2, err2 := getPageCount(file2)
	if err2 != nil {
		return 0, 0, fmt.Errorf("failed to get page count for second file: %v", err2)
	}

	return pages1, pages2, nil
}

// Validate page count match between files
func validatePageCountMatch(file1, file2 string, pages1, pages2 int) error {
	if pages1 != pages2 {
		return fmt.Errorf("page count mismatch - %s has %d pages, %s has %d pages",
			filepath.Base(file1), pages1, filepath.Base(file2), pages2)
	}
	return nil
}

// Display page count in verbose mode
func displayPageCount(pages int) {
	if VERBOSE {
		fmt.Printf("pages = %s%d%s\n", BLUE, pages, NC)
	}
}

// PDF reversal operations

// Display reversal information in verbose mode
func displayReversalInfo(pageCount int) {
	if VERBOSE {
		fmt.Printf("rev = %s%d", BLUE, pageCount)
		for i := pageCount - 1; i >= 1; i-- {
			fmt.Printf(",%d", i)
		}
		fmt.Printf("%s\n", NC)
	}
}

// Clean up temporary files
func cleanupTempFiles(tempFiles []string) {
	for _, tempFile := range tempFiles {
		if err := os.Remove(tempFile); err != nil && VERBOSE {
			printWarning(fmt.Sprintf("Failed to remove temp file %s: %v", tempFile, err))
		}
	}
}

// PDF merging operations

// Smart merge: direct merge for single-page, reversed merge for multi-page
func smartMerge(file1, file2, outputFile string, pages1, pages2 int) error {
	if pages2 == 1 {
		return performDirectMerge(file1, file2, outputFile)
	}

	return performReversedMerge(file1, file2, outputFile, pages1, pages2)
}

// Perform direct merge for single-page second file
func performDirectMerge(file1, file2, outputFile string) error {
	if VERBOSE {
		printInfo("Single-page second file detected - merging directly without reversal")
	}

	conf := model.NewDefaultConfiguration()
	inputFiles := []string{file1, file2}
	return api.MergeCreateFile(inputFiles, outputFile, false, conf)
}

// Perform reversed merge for multi-page second file using in-memory processing
func performReversedMerge(file1, file2, outputFile string, pages1, pages2 int) error {
	if VERBOSE {
		printInfo(fmt.Sprintf("Multi-page second file detected (%d pages) - processing in memory", pages2))
	}

	return createInterleavedMerge(file1, file2, outputFile, pages1)
}

// Create filename for reversed PDF
func createReversedFileName(file string) string {
	return strings.TrimSuffix(file, filepath.Ext(file)) + "-reverse.pdf"
}

// Remove reversed file with error handling
func removeReversedFile(reversedFile string) {
	if err := os.Remove(reversedFile); err != nil && VERBOSE {
		printWarning(fmt.Sprintf("Failed to clean up temporary file %s: %v", reversedFile, err))
	}
}

// Create interleaved merge pattern using pure in-memory stream-based approach
func createInterleavedMerge(file1, file2, outputFile string, pageCount int) error {
	conf := model.NewDefaultConfiguration()

	// Load both PDFs into memory
	bytes1, err := os.ReadFile(file1)
	if err != nil {
		return fmt.Errorf("failed to read file1 into memory: %v", err)
	}

	bytes2, err := os.ReadFile(file2)
	if err != nil {
		return fmt.Errorf("failed to read file2 into memory: %v", err)
	}

	// Create reverse page selection for second document (3,2,1 for 3-page doc)
	reversePages := ""
	for i := pageCount; i >= 1; i-- {
		if reversePages != "" {
			reversePages += ","
		}
		reversePages += fmt.Sprintf("%d", i)
	}

	// Reverse second document in memory using stream-based Trim
	reader2 := bytes.NewReader(bytes2)
	var reversedBuffer bytes.Buffer
	reverseSelection, err := api.ParsePageSelection(reversePages)
	if err != nil {
		return fmt.Errorf("failed to parse reverse page selection: %v", err)
	}

	err = api.Trim(reader2, &reversedBuffer, reverseSelection, conf)
	if err != nil {
		return fmt.Errorf("failed to reverse document in memory: %v", err)
	}

	// Zip merge for perfect interleaving using stream-based MergeCreateZip
	reader1 := bytes.NewReader(bytes1)
	reversedReader := bytes.NewReader(reversedBuffer.Bytes())
	var finalBuffer bytes.Buffer

	err = api.MergeCreateZip(reader1, reversedReader, &finalBuffer, conf)
	if err != nil {
		return fmt.Errorf("failed to create interleaved merge in memory: %v", err)
	}

	// Write final result to output file
	err = os.WriteFile(outputFile, finalBuffer.Bytes(), 0644)
	if err != nil {
		return fmt.Errorf("failed to write output file: %v", err)
	}

	return nil
}

// Main processing function

// Process and merge files with smart page reversal
// Process and merge to temporary file (for multi-output support)
func processAndMergeToTemp(outputFile, file1, file2 string, pages int) {
	pages1, pages2, err := validatePDFsForMerge(file1, file2)
	if err != nil {
		handleMergeValidationError(file1, file2, err)
		return
	}

	if err := smartMerge(file1, file2, outputFile, pages1, pages2); err != nil {
		handleMergeExecutionError(file1, file2, err)
	}
	// Note: Don't move source files here - that's handled by the caller
}

// Process and merge PDFs with file movement
func processAndMerge(outputFile, file1, file2 string, pages int) {
	pages1, pages2, err := validatePDFsForMerge(file1, file2)
	if err != nil {
		handleMergeValidationError(file1, file2, err)
		return
	}

	if err := smartMerge(file1, file2, outputFile, pages1, pages2); err != nil {
		handleMergeExecutionError(file1, file2, err)
	} else {
		handleMergeSuccess(file1, file2)
	}
}

// Handle merge validation errors
func handleMergeValidationError(file1, file2 string, err error) {
	printError(err.Error())
	moveProcessedFiles(ERROR_DIR, "PDF validation failed. Moving files to error folder...", file1, file2)
}

// Handle merge execution errors
func handleMergeExecutionError(file1, file2 string, err error) {
	printError(fmt.Sprintf("Failed to merge PDFs: %v", err))
	moveProcessedFiles(ERROR_DIR, "Merge failed. Moving files to error folder...", file1, file2)
}

// Handle successful merge
func handleMergeSuccess(file1, file2 string) {
	moveProcessedFiles(ARCHIVE, "Files merged and moved.", file1, file2)
}

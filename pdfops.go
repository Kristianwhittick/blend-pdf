package main

import (
	"fmt"
	"os"

	"github.com/pdfcpu/pdfcpu/pkg/api"
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/model"
)

func validatePDF(file string) bool {
	if _, err := os.Stat(file); os.IsNotExist(err) {
		fmt.Printf("%sError:%s File '%s' does not exist.\n", RED, NC, file)
		return false
	}

	// Use pdfcpu API to validate the PDF
	conf := model.NewDefaultConfiguration()
	conf.ValidationMode = model.ValidationRelaxed

	validateErr := api.ValidateFile(file, conf)
	if validateErr != nil {
		fmt.Printf("%sError:%s '%s' is not a valid PDF file: %v\n", RED, NC, file, validateErr)
		return false
	}

	return true
}

func getPageCount(file string) int {
	pageCount, err := api.PageCountFile(file)
	if err != nil {
		fmt.Printf("%sError:%s Could not determine page count for '%s': %v\n", RED, NC, file, err)
		return -1
	}

	return pageCount
}



func createInterleavedMerge(file1, file2, outputFile string, pageCount int) error {
	conf := model.NewDefaultConfiguration()
	
	// Create temporary files for individual pages
	tempFiles := make([]string, 0, pageCount*2)
	
	// Extract pages from first document (in order)
	for i := 1; i <= pageCount; i++ {
		tempFile1 := fmt.Sprintf("temp_A_%d.pdf", i)
		pageSelection, err := api.ParsePageSelection(fmt.Sprintf("%d", i))
		if err != nil {
			return fmt.Errorf("failed to parse page selection for file1 page %d: %v", i, err)
		}
		
		err = api.TrimFile(file1, tempFile1, pageSelection, conf)
		if err != nil {
			return fmt.Errorf("failed to extract page %d from file1: %v", i, err)
		}
		
		// Extract corresponding page from second document (in reverse order)
		reversePage := pageCount - i + 1
		tempFile2 := fmt.Sprintf("temp_B_%d.pdf", i)
		pageSelection2, err := api.ParsePageSelection(fmt.Sprintf("%d", reversePage))
		if err != nil {
			return fmt.Errorf("failed to parse page selection for file2 page %d: %v", reversePage, err)
		}
		
		err = api.TrimFile(file2, tempFile2, pageSelection2, conf)
		if err != nil {
			return fmt.Errorf("failed to extract page %d from file2: %v", reversePage, err)
		}
		
		// Add both pages to merge list (interleaved)
		tempFiles = append(tempFiles, tempFile1, tempFile2)
	}
	
	// Merge all temporary files
	err := api.MergeCreateFile(tempFiles, outputFile, false, conf)
	
	// Clean up temporary files
	for _, tempFile := range tempFiles {
		os.Remove(tempFile)
	}
	
	return err
}

func processAndMerge(outputFile, file1, file2 string, pages int) {
	err := createInterleavedMerge(file1, file2, outputFile, pages)
	if err != nil {
		fmt.Printf("%sError:%s Failed to create interleaved merge: %v\n", RED, NC, err)
		moveProcessedFiles(ERROR_DIR, "Merge failed. Moving files to error folder...", file1, file2, "")
	} else {
		moveProcessedFiles(ARCHIVE, "Files merged and moved to archive successfully.", file1, file2, "")
	}
}

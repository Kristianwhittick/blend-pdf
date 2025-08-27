package main

import (
	"fmt"
	"log"
	"os"

	"github.com/pdfcpu/pdfcpu/pkg/api"
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/model"
)

func main() {
	fmt.Println("=== Test 08: Complete Interleaved Pattern ===")
	fmt.Println("Expected: Doc1_Page1, Doc2_Page3, Doc1_Page2, Doc2_Page2, Doc1_Page3, Doc2_Page1")
	fmt.Println("Which is: A1, B3, A2, B2, A3, B1")
	
	// Create default configuration
	conf := model.NewDefaultConfiguration()
	
	// Get page counts
	pageCount, err := api.PageCountFile("Doc_A.pdf")
	if err != nil {
		log.Fatalf("Failed to get page count: %v", err)
	}
	
	fmt.Printf("Both documents have %d pages\n", pageCount)
	
	// Create temporary files for individual pages
	tempFiles := make([]string, 0, pageCount*2)
	
	// Extract pages from first document (in order) and second document (in reverse order)
	for i := 1; i <= pageCount; i++ {
		// Extract page i from Doc_A
		tempFileA := fmt.Sprintf("temp_A_%d.pdf", i)
		pageSelection, err := api.ParsePageSelection(fmt.Sprintf("%d", i))
		if err != nil {
			log.Fatalf("Failed to parse page selection for A%d: %v", i, err)
		}
		
		err = api.TrimFile("Doc_A.pdf", tempFileA, pageSelection, conf)
		if err != nil {
			log.Fatalf("Failed to extract A%d: %v", i, err)
		}
		
		// Extract corresponding page from Doc_B (in reverse order)
		reversePage := pageCount - i + 1
		tempFileB := fmt.Sprintf("temp_B_%d.pdf", i)
		pageSelection, err = api.ParsePageSelection(fmt.Sprintf("%d", reversePage))
		if err != nil {
			log.Fatalf("Failed to parse page selection for B%d: %v", reversePage, err)
		}
		
		err = api.TrimFile("Doc_B.pdf", tempFileB, pageSelection, conf)
		if err != nil {
			log.Fatalf("Failed to extract B%d: %v", reversePage, err)
		}
		
		// Add both pages to merge list (interleaved)
		tempFiles = append(tempFiles, tempFileA, tempFileB)
		
		fmt.Printf("Extracted: A%d and B%d (from page %d of Doc_B)\n", i, pageCount-i+1, reversePage)
	}
	
	// Merge all temporary files
	fmt.Println("Merging all pages in interleaved pattern...")
	err = api.MergeCreateFile(tempFiles, "output/test08_interleaved.pdf", false, conf)
	if err != nil {
		log.Fatalf("Failed to merge files: %v", err)
	}
	
	// Clean up temporary files
	for _, tempFile := range tempFiles {
		os.Remove(tempFile)
	}
	
	fmt.Println("Successfully created output/test08_interleaved.pdf")
	
	// Verify the result
	resultPageCount, err := api.PageCountFile("output/test08_interleaved.pdf")
	if err != nil {
		log.Fatalf("Failed to get page count of result: %v", err)
	}
	
	fmt.Printf("Result file has %d page(s)\n", resultPageCount)
	
	expectedPages := pageCount * 2
	if resultPageCount == expectedPages {
		fmt.Printf("✅ Test 08 PASSED - Interleaved pattern created successfully!\n")
	} else {
		fmt.Printf("❌ Test 08 FAILED - Expected %d pages, got %d\n", expectedPages, resultPageCount)
	}
}

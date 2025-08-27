package main

import (
	"fmt"
	"log"

	"github.com/pdfcpu/pdfcpu/pkg/api"
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/model"
)

func main() {
	fmt.Println("=== Test 06: Simple Merge Two Files ===")
	
	// Create default configuration
	conf := model.NewDefaultConfiguration()
	
	// Simple merge: Doc_A.pdf + Doc_B.pdf
	fmt.Println("Merging Doc_A.pdf + Doc_B.pdf...")
	
	inputFiles := []string{"Doc_A.pdf", "Doc_B.pdf"}
	
	err := api.MergeCreateFile(inputFiles, "output/test06_simple_merge.pdf", false, conf)
	if err != nil {
		log.Fatalf("Failed to merge files: %v", err)
	}
	
	fmt.Println("Successfully merged to output/test06_simple_merge.pdf")
	
	// Verify the result
	pageCount, err := api.PageCountFile("output/test06_simple_merge.pdf")
	if err != nil {
		log.Fatalf("Failed to get page count of result: %v", err)
	}
	
	fmt.Printf("Result file has %d page(s)\n", pageCount)
	
	if pageCount == 6 {
		fmt.Println("✅ Test 06 PASSED - Files merged successfully!")
	} else {
		fmt.Printf("❌ Test 06 FAILED - Expected 6 pages, got %d\n", pageCount)
	}
}

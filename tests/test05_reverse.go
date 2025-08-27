package main

import (
	"fmt"
	"log"

	"github.com/pdfcpu/pdfcpu/pkg/api"
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/model"
)

func main() {
	fmt.Println("=== Test 05: Extract Pages in Reverse Order ===")
	
	// Create default configuration
	conf := model.NewDefaultConfiguration()
	
	// Extract pages 3,2,1 from Doc_B.pdf (reverse order)
	fmt.Println("Extracting pages 3,2,1 from Doc_B.pdf...")
	
	pageSelection, err := api.ParsePageSelection("3,2,1")
	if err != nil {
		log.Fatalf("Failed to parse page selection: %v", err)
	}
	
	err = api.TrimFile("Doc_B.pdf", "output/test05_reverse.pdf", pageSelection, conf)
	if err != nil {
		log.Fatalf("Failed to extract pages: %v", err)
	}
	
	fmt.Println("Successfully extracted pages 3,2,1 to output/test05_reverse.pdf")
	
	// Verify the result
	pageCount, err := api.PageCountFile("output/test05_reverse.pdf")
	if err != nil {
		log.Fatalf("Failed to get page count of result: %v", err)
	}
	
	fmt.Printf("Result file has %d page(s)\n", pageCount)
	
	if pageCount == 3 {
		fmt.Println("✅ Test 05 PASSED - Pages extracted in reverse order!")
	} else {
		fmt.Printf("❌ Test 05 FAILED - Expected 3 pages, got %d\n", pageCount)
	}
}

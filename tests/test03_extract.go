package main

import (
	"fmt"
	"log"

	"github.com/pdfcpu/pdfcpu/pkg/api"
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/model"
)

func main() {
	fmt.Println("=== Test 03: Extract Single Page ===")
	
	// Create default configuration
	conf := model.NewDefaultConfiguration()
	
	// Extract page 1 from Doc_A.pdf
	fmt.Println("Extracting page 1 from Doc_A.pdf...")
	
	pageSelection, err := api.ParsePageSelection("1")
	if err != nil {
		log.Fatalf("Failed to parse page selection: %v", err)
	}
	
	err = api.TrimFile("Doc_A.pdf", "output/test03_single_page.pdf", pageSelection, conf)
	if err != nil {
		log.Fatalf("Failed to extract page: %v", err)
	}
	
	fmt.Println("Successfully extracted page 1 to output/test03_single_page.pdf")
	
	// Verify the result
	pageCount, err := api.PageCountFile("output/test03_single_page.pdf")
	if err != nil {
		log.Fatalf("Failed to get page count of result: %v", err)
	}
	
	fmt.Printf("Result file has %d page(s)\n", pageCount)
	
	if pageCount == 1 {
		fmt.Println("✅ Test 03 PASSED - Single page extracted successfully!")
	} else {
		fmt.Printf("❌ Test 03 FAILED - Expected 1 page, got %d\n", pageCount)
	}
}

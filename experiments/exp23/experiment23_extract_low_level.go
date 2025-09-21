package main

import (
	"fmt"
	"log"
	"os"

	"github.com/pdfcpu/pdfcpu/pkg/api"
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/model"
)

func main() {
	fmt.Println("=== Experiment 23: Low-Level Extract Pages API ===")
	
	// Test the low-level extract pages function mentioned by maintainer
	// Based on extract_test.go: func TestExtractPagesLowLevel(t *testing.T)
	
	inputFile := "Doc_A.pdf"
	outputFile := "output/test_23_extract_low_level.pdf"
	
	// Ensure output directory exists
	os.MkdirAll("output", 0755)
	
	fmt.Printf("Testing low-level extract pages from: %s\n", inputFile)
	
	// Try to find and use the low-level extract function
	// We need to look at the actual pdfcpu source to understand the API
	
	// First, let's try the standard approach and see what we can discover
	conf := model.NewDefaultConfiguration()
	
	// Check if there are any extract functions we haven't tried
	fmt.Println("Checking available extract functions...")
	
	// Let's try to extract pages 1,3 using different approaches
	pageSelection, err := api.ParsePageSelection("1,3")
	if err != nil {
		log.Fatalf("Failed to parse page selection: %v", err)
	}
	
	// Standard TrimFile approach (we know this works)
	err = api.TrimFile(inputFile, outputFile, pageSelection, conf)
	if err != nil {
		log.Fatalf("TrimFile failed: %v", err)
	}
	
	fmt.Printf("✅ Standard TrimFile worked: %s\n", outputFile)
	
	// Now let's try to find the low-level API
	// We need to examine the pdfcpu source code structure
	fmt.Println("\n=== Investigating Low-Level APIs ===")
	fmt.Println("The maintainer mentioned:")
	fmt.Println("- extract_test.go: func TestExtractPagesLowLevel(t *testing.T)")
	fmt.Println("- merge_test.go: TestMergeRaw(t *testing.T)")
	fmt.Println("- trim.go: func Trim(rs io.ReadSeeker, w io.Writer, selectedPages []string, conf *model.Configuration) error")
	
	fmt.Println("\nNext steps:")
	fmt.Println("1. Examine pdfcpu source code for these functions")
	fmt.Println("2. Test the io.ReadSeeker/io.Writer based Trim function")
	fmt.Println("3. Look for extract and merge functions that work with streams")
}

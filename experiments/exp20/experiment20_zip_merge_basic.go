package main

import (
	"fmt"
	"log"
	"os"

	"github.com/pdfcpu/pdfcpu/pkg/api"
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/model"
)

func main() {
	fmt.Println("=== Experiment 20: Basic Zip Merge Test ===")
	
	// Test files (paths from project root)
	file1 := "archive/Doc_A.pdf" // A1, A2, A3
	file2 := "archive/Doc_B.pdf" // M, 9, f
	output := "output/test_20_zip_basic.pdf"
	
	// Verify input files exist
	if _, err := os.Stat(file1); os.IsNotExist(err) {
		log.Fatalf("Input file 1 not found: %s", file1)
	}
	if _, err := os.Stat(file2); os.IsNotExist(err) {
		log.Fatalf("Input file 2 not found: %s", file2)
	}
	
	// Get page counts
	count1, err := api.PageCountFile(file1)
	if err != nil {
		log.Fatalf("Failed to get page count for file1: %v", err)
	}
	
	count2, err := api.PageCountFile(file2)
	if err != nil {
		log.Fatalf("Failed to get page count for file2: %v", err)
	}
	
	fmt.Printf("File 1 (%s): %d pages\n", file1, count1)
	fmt.Printf("File 2 (%s): %d pages\n", file2, count2)
	
	// Test basic zip merge (no reversal)
	conf := model.NewDefaultConfiguration()
	
	fmt.Println("\nTesting MergeCreateZipFile...")
	err = api.MergeCreateZipFile(file1, file2, output, conf)
	if err != nil {
		log.Fatalf("MergeCreateZipFile failed: %v", err)
	}
	
	// Verify output
	outputCount, err := api.PageCountFile(output)
	if err != nil {
		log.Fatalf("Failed to get output page count: %v", err)
	}
	
	fmt.Printf("✓ Zip merge successful!\n")
	fmt.Printf("Output file: %s\n", output)
	fmt.Printf("Output pages: %d (expected: %d)\n", outputCount, count1+count2)
	
	if outputCount == count1+count2 {
		fmt.Println("✓ Page count matches expected total")
	} else {
		fmt.Println("✗ Page count mismatch!")
	}
	
	fmt.Println("\n=== Expected Pattern (no reversal) ===")
	fmt.Println("A1, M, A2, 9, A3, f")
	fmt.Println("\nUse pdftotext to verify:")
	fmt.Printf("pdftotext %s -\n", output)
}

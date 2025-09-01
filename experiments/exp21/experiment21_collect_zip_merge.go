package main

import (
	"fmt"
	"log"
	"os"

	"github.com/pdfcpu/pdfcpu/pkg/api"
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/model"
)

func main() {
	fmt.Println("=== Experiment 21: CollectFile + Zip Merge Test ===")

	// Test files (paths from project root)
	file1 := "archive/Doc_A.pdf" // A1, A2, A3
	file2 := "archive/Doc_B.pdf" // M, 9, f
	reversedFile := "output/test_21_reversed.pdf"
	output := "output/test_21_collect_zip.pdf"

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

	conf := model.NewDefaultConfiguration()

	// Step 1: Reverse file2 using CollectFile
	fmt.Println("\nStep 1: Reversing file2 with CollectFile...")
	pageSelection, err := api.ParsePageSelection("3,2,1")
	if err != nil {
		log.Fatalf("Failed to parse page selection: %v", err)
	}

	err = api.CollectFile(file2, reversedFile, pageSelection, conf)
	if err != nil {
		log.Fatalf("CollectFile failed: %v", err)
	}

	// Verify reversed file
	reversedCount, err := api.PageCountFile(reversedFile)
	if err != nil {
		log.Fatalf("Failed to get reversed file page count: %v", err)
	}

	fmt.Printf("✓ Reversed file created: %s\n", reversedFile)
	fmt.Printf("Reversed file pages: %d (expected: %d)\n", reversedCount, count2)

	// Step 2: Zip merge original file1 with reversed file2
	fmt.Println("\nStep 2: Zip merging file1 + reversed file2...")
	err = api.MergeCreateZipFile(file1, reversedFile, output, conf)
	if err != nil {
		log.Fatalf("MergeCreateZipFile failed: %v", err)
	}

	// Verify final output
	outputCount, err := api.PageCountFile(output)
	if err != nil {
		log.Fatalf("Failed to get output page count: %v", err)
	}

	fmt.Printf("✓ Final zip merge successful!\n")
	fmt.Printf("Output file: %s\n", output)
	fmt.Printf("Output pages: %d (expected: %d)\n", outputCount, count1+count2)

	if outputCount == count1+count2 {
		fmt.Println("✓ Page count matches expected total")
	} else {
		fmt.Println("✗ Page count mismatch!")
	}

	fmt.Println("\n=== Expected Pattern (with reversal) ===")
	fmt.Println("Original Doc_B: M, 9, f")
	fmt.Println("Reversed Doc_B: f, 9, M")
	fmt.Println("Zip merge result: A1, f, A2, 9, A3, M")
	fmt.Println("\nUse pdftotext to verify:")
	fmt.Printf("pdftotext %s -\n", output)
	fmt.Printf("pdftotext %s -\n", reversedFile)

	// Cleanup option
	fmt.Println("\nCleanup commands:")
	fmt.Printf("rm %s\n", reversedFile)
}

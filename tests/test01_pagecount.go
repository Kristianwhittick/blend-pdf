package main

import (
	"fmt"
	"log"

	"github.com/pdfcpu/pdfcpu/pkg/api"
)

func main() {
	fmt.Println("=== Test 01: Page Count API ===")
	
	// Test Doc_A.pdf
	fmt.Println("Testing Doc_A.pdf...")
	countA, err := api.PageCountFile("Doc_A.pdf")
	if err != nil {
		log.Fatalf("Error getting page count for Doc_A.pdf: %v", err)
	}
	fmt.Printf("Doc_A.pdf has %d pages\n", countA)
	
	// Test Doc_B.pdf
	fmt.Println("Testing Doc_B.pdf...")
	countB, err := api.PageCountFile("Doc_B.pdf")
	if err != nil {
		log.Fatalf("Error getting page count for Doc_B.pdf: %v", err)
	}
	fmt.Printf("Doc_B.pdf has %d pages\n", countB)
	
	// Summary
	fmt.Printf("\nSummary:\n")
	fmt.Printf("- Doc_A.pdf: %d pages\n", countA)
	fmt.Printf("- Doc_B.pdf: %d pages\n", countB)
	fmt.Printf("- Pages match: %t\n", countA == countB)
	
	fmt.Println("Test 01 completed successfully!")
}

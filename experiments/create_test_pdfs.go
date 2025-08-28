package main

import (
	"fmt"
	"log"
	"os"

	"github.com/pdfcpu/pdfcpu/pkg/api"
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/model"
)

func main() {
	fmt.Println("=== Creating Test PDF Files ===")
	
	// Create simple test PDFs using pdfcpu's create functionality
	conf := model.NewDefaultConfiguration()
	
	// Create Doc_A.pdf with 3 pages containing text A1, A2, A3
	fmt.Println("Creating Doc_A.pdf...")
	
	// For now, let's create a simple approach by copying an existing PDF structure
	// We'll create minimal PDFs with text content
	
	// Create a simple text file and convert it to PDF
	textA := "A1\nPage 1 of Document A\n\nA2\nPage 2 of Document A\n\nA3\nPage 3 of Document A"
	err := os.WriteFile("temp_a.txt", []byte(textA), 0644)
	if err != nil {
		log.Fatalf("Failed to create temp text file: %v", err)
	}
	defer os.Remove("temp_a.txt")
	
	textB := "M\nPage 1 of Document B\n\n9\nPage 2 of Document B\n\nf\nPage 3 of Document B"
	err = os.WriteFile("temp_b.txt", []byte(textB), 0644)
	if err != nil {
		log.Fatalf("Failed to create temp text file: %v", err)
	}
	defer os.Remove("temp_b.txt")
	
	fmt.Println("Note: This experiment requires existing PDF files.")
	fmt.Println("Please create Doc_A.pdf and Doc_B.pdf manually with 3 pages each.")
	fmt.Println("Doc_A.pdf should contain: A1, A2, A3 on pages 1, 2, 3")
	fmt.Println("Doc_B.pdf should contain: M, 9, f on pages 1, 2, 3")
	fmt.Println("You can use any PDF creation tool or online converter.")
}

package main

import (
	"bytes"
	"fmt"
	"log"
	"os"

	"github.com/pdfcpu/pdfcpu/pkg/api"
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/model"
)

func main() {
	fmt.Println("=== Experiment 24: Stream-Based Trim Function ===")
	
	// Test the stream-based Trim function mentioned by maintainer:
	// trim.go: func Trim(rs io.ReadSeeker, w io.Writer, selectedPages []string, conf *model.Configuration) error
	
	inputFile := "output/Doc_A.pdf"
	outputFile := "output/test_24_trim_stream.pdf"
	
	// Ensure output directory exists
	os.MkdirAll("output", 0755)
	
	fmt.Printf("Testing stream-based trim from: %s\n", inputFile)
	
	// Read input file into memory
	inputBytes, err := os.ReadFile(inputFile)
	if err != nil {
		log.Fatalf("Failed to read input file: %v", err)
	}
	
	fmt.Printf("Loaded %d bytes from %s\n", len(inputBytes), inputFile)
	
	// Create ReadSeeker from bytes
	reader := bytes.NewReader(inputBytes)
	
	// Create output buffer
	var outputBuffer bytes.Buffer
	
	// Parse page selection
	pageSelection, err := api.ParsePageSelection("1,3")
	if err != nil {
		log.Fatalf("Failed to parse page selection: %v", err)
	}
	
	fmt.Printf("Page selection: %v\n", pageSelection)
	fmt.Printf("Reader size: %d bytes, Output buffer ready: %d bytes\n", reader.Len(), outputBuffer.Len())
	
	// Configuration
	conf := model.NewDefaultConfiguration()
	
	// The maintainer mentioned trim.go has a Trim function with this signature:
	// func Trim(rs io.ReadSeeker, w io.Writer, selectedPages []string, conf *model.Configuration) error
	
	// Let's try to find it in different packages
	fmt.Println("Looking for stream-based Trim function...")
	
	// First, let's check what's available in the api package
	fmt.Println("Available in api package:")
	fmt.Println("- api.TrimFile() - works with files")
	fmt.Println("- Need to find the stream version")
	
	// For now, let's use the file-based approach to verify our setup works
	// Then we'll investigate the stream-based version
	
	// Write bytes to temp file for testing
	tempInput := "temp_input.pdf"
	err = os.WriteFile(tempInput, inputBytes, 0644)
	if err != nil {
		log.Fatalf("Failed to write temp file: %v", err)
	}
	defer os.Remove(tempInput)
	
	// Use file-based trim
	err = api.TrimFile(tempInput, outputFile, pageSelection, conf)
	if err != nil {
		log.Fatalf("TrimFile failed: %v", err)
	}
	
	fmt.Printf("✅ File-based Trim succeeded: %s\n", outputFile)
	
	// Verify the output
	pageCount, err := api.PageCountFile(outputFile)
	if err != nil {
		log.Printf("⚠️  Could not get page count: %v", err)
	} else {
		fmt.Printf("✅ Output PDF has %d pages\n", pageCount)
	}
	
	fmt.Println("\n=== Investigation Notes ===")
	fmt.Println("The maintainer mentioned a stream-based Trim function in trim.go")
	fmt.Println("We need to:")
	fmt.Println("1. Find the correct import path for this function")
	fmt.Println("2. It might be in a different package than api")
	fmt.Println("3. Check pdfcpu source code structure")
	
	fmt.Println("\nNext experiment will investigate the package structure...")
}

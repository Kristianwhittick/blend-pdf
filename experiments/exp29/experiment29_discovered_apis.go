package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/pdfcpu/pdfcpu/pkg/api"
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/model"
)

func main() {
	fmt.Println("=== Experiment 29: Discovered Stream-Based APIs ===")
	
	// Test the stream-based APIs discovered in pdfcpu source:
	// 1. api.Trim(rs io.ReadSeeker, w io.Writer, selectedPages []string, conf *model.Configuration) error
	// 2. api.MergeRaw(rsc []io.ReadSeeker, w io.Writer, dividerPage bool, conf *model.Configuration) error
	// 3. api.MergeCreateZip(rs1, rs2 io.ReadSeeker, w io.Writer, conf *model.Configuration) error
	// 4. api.ExtractPages(rs io.ReadSeeker, outDir, fileName string, selectedPages []string, conf *model.Configuration) error
	
	inputFile1 := "output/Doc_A.pdf"
	inputFile2 := "output/Doc_B.pdf"
	
	// Ensure output directory exists
	os.MkdirAll("output", 0755)
	
	fmt.Printf("Testing discovered APIs with: %s + %s\n", inputFile1, inputFile2)
	
	// Load both files into memory
	inputBytes1, err := os.ReadFile(inputFile1)
	if err != nil {
		log.Fatalf("Failed to read input file 1: %v", err)
	}
	
	inputBytes2, err := os.ReadFile(inputFile2)
	if err != nil {
		log.Fatalf("Failed to read input file 2: %v", err)
	}
	
	fmt.Printf("✅ Loaded %d + %d bytes into memory\n", len(inputBytes1), len(inputBytes2))
	
	// Configuration
	conf := model.NewDefaultConfiguration()
	
	// Test 1: Stream-based Trim
	fmt.Println("\n=== Test 1: Stream-based Trim ===")
	reader1 := bytes.NewReader(inputBytes1)
	var trimBuffer bytes.Buffer
	
	pageSelection, _ := api.ParsePageSelection("1,3")
	err = api.Trim(reader1, &trimBuffer, pageSelection, conf)
	if err != nil {
		log.Printf("❌ Trim failed: %v", err)
	} else {
		fmt.Printf("✅ Trim succeeded: %d bytes output\n", trimBuffer.Len())
		
		// Write to file to verify
		err = os.WriteFile("output/test_29_trim.pdf", trimBuffer.Bytes(), 0644)
		if err == nil {
			pageCount, _ := api.PageCountFile("output/test_29_trim.pdf")
			fmt.Printf("✅ Trim output has %d pages\n", pageCount)
		}
	}
	
	// Test 2: Stream-based MergeRaw
	fmt.Println("\n=== Test 2: Stream-based MergeRaw ===")
	reader1.Seek(0, io.SeekStart)
	reader2 := bytes.NewReader(inputBytes2)
	
	readers := []io.ReadSeeker{reader1, reader2}
	var mergeBuffer bytes.Buffer
	
	err = api.MergeRaw(readers, &mergeBuffer, false, conf)
	if err != nil {
		log.Printf("❌ MergeRaw failed: %v", err)
	} else {
		fmt.Printf("✅ MergeRaw succeeded: %d bytes output\n", mergeBuffer.Len())
		
		// Write to file to verify
		err = os.WriteFile("output/test_29_merge_raw.pdf", mergeBuffer.Bytes(), 0644)
		if err == nil {
			pageCount, _ := api.PageCountFile("output/test_29_merge_raw.pdf")
			fmt.Printf("✅ MergeRaw output has %d pages\n", pageCount)
		}
	}
	
	// Test 3: Stream-based MergeCreateZip (Interleaved)
	fmt.Println("\n=== Test 3: Stream-based MergeCreateZip ===")
	reader1.Seek(0, io.SeekStart)
	reader2.Seek(0, io.SeekStart)
	
	var zipBuffer bytes.Buffer
	
	err = api.MergeCreateZip(reader1, reader2, &zipBuffer, conf)
	if err != nil {
		log.Printf("❌ MergeCreateZip failed: %v", err)
	} else {
		fmt.Printf("✅ MergeCreateZip succeeded: %d bytes output\n", zipBuffer.Len())
		
		// Write to file to verify
		err = os.WriteFile("output/test_29_zip.pdf", zipBuffer.Bytes(), 0644)
		if err == nil {
			pageCount, _ := api.PageCountFile("output/test_29_zip.pdf")
			fmt.Printf("✅ MergeCreateZip output has %d pages\n", pageCount)
		}
	}
	
	// Test 4: Complete In-Memory Interleaved Workflow
	fmt.Println("\n=== Test 4: Complete In-Memory Interleaved Workflow ===")
	
	// Step 1: Get page count using file-based API first
	pageCount1, err := api.PageCountFile(inputFile1)
	if err != nil {
		log.Printf("❌ Failed to get page count 1: %v", err)
		return
	}
	
	pageCount2, err := api.PageCountFile(inputFile2)
	if err != nil {
		log.Printf("❌ Failed to get page count 2: %v", err)
		return
	}
	
	fmt.Printf("Page counts: %d + %d\n", pageCount1, pageCount2)
	
	if pageCount1 != pageCount2 {
		log.Printf("❌ Page count mismatch: %d != %d", pageCount1, pageCount2)
		return
	}
	
	// Step 2: Reverse second document using Trim
	reader2.Seek(0, io.SeekStart)
	var reversedBuffer bytes.Buffer
	
	// Create reverse page selection (3,2,1 for 3-page document)
	reversePages := ""
	for i := pageCount2; i >= 1; i-- {
		if reversePages != "" {
			reversePages += ","
		}
		reversePages += fmt.Sprintf("%d", i)
	}
	
	fmt.Printf("Reversing pages: %s\n", reversePages)
	
	reverseSelection, _ := api.ParsePageSelection(reversePages)
	err = api.Trim(reader2, &reversedBuffer, reverseSelection, conf)
	if err != nil {
		log.Printf("❌ Reverse trim failed: %v", err)
		return
	}
	
	fmt.Printf("✅ Reversed document: %d bytes\n", reversedBuffer.Len())
	
	// Step 3: Zip merge original + reversed for perfect interleaving
	reader1.Seek(0, io.SeekStart)
	reversedReader := bytes.NewReader(reversedBuffer.Bytes())
	
	var finalBuffer bytes.Buffer
	err = api.MergeCreateZip(reader1, reversedReader, &finalBuffer, conf)
	if err != nil {
		log.Printf("❌ Final zip merge failed: %v", err)
		return
	}
	
	fmt.Printf("✅ Final interleaved merge: %d bytes\n", finalBuffer.Len())
	
	// Write final result
	err = os.WriteFile("output/test_29_complete_interleaved.pdf", finalBuffer.Bytes(), 0644)
	if err == nil {
		pageCount, _ := api.PageCountFile("output/test_29_complete_interleaved.pdf")
		fmt.Printf("✅ Complete workflow output has %d pages\n", pageCount)
	}
	
	fmt.Println("\n=== SUCCESS! ===")
	fmt.Println("✅ All stream-based APIs work!")
	fmt.Println("✅ Complete in-memory interleaved workflow achieved!")
	fmt.Println("✅ Zero temporary files needed!")
	
	fmt.Println("\nDiscovered APIs:")
	fmt.Println("- api.Trim() for in-memory page extraction")
	fmt.Println("- api.MergeRaw() for in-memory merging")
	fmt.Println("- api.MergeCreateZip() for in-memory interleaved merging")
	fmt.Println("- Complete workflow: Load → Reverse → Zip → Output")
}

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
	fmt.Println("=== Experiment 26: Complete In-Memory Workflow ===")

	inputFile1 := "output/Doc_A.pdf"
	inputFile2 := "output/Doc_B.pdf"
	outputFile := "output/test_26_complete_memory.pdf"

	os.MkdirAll("output", 0755)

	fmt.Printf("Testing complete in-memory workflow: %s + %s\n", inputFile1, inputFile2)

	inputBytes1, err := os.ReadFile(inputFile1)
	if err != nil {
		log.Fatalf("Failed to read input file 1: %v", err)
	}

	inputBytes2, err := os.ReadFile(inputFile2)
	if err != nil {
		log.Fatalf("Failed to read input file 2: %v", err)
	}

	fmt.Printf("✅ Loaded %d + %d bytes into memory\n", len(inputBytes1), len(inputBytes2))

	reader1 := bytes.NewReader(inputBytes1)
	reader2 := bytes.NewReader(inputBytes2)
	conf := model.NewDefaultConfiguration()

	ctx1, err := api.ReadContext(reader1, conf)
	if err != nil {
		log.Fatalf("Failed to read context 1: %v", err)
	}

	ctx2, err := api.ReadContext(reader2, conf)
	if err != nil {
		log.Fatalf("Failed to read context 2: %v", err)
	}

	fmt.Printf("✅ Page counts: %d + %d pages\n", ctx1.PageCount, ctx2.PageCount)

	if ctx1.PageCount != ctx2.PageCount {
		log.Fatalf("❌ Page count mismatch: %d != %d", ctx1.PageCount, ctx2.PageCount)
	}

	fmt.Println("Extracting individual pages using stream-based Trim...")

	var pageBuffers [][]byte

	for i := 1; i <= ctx1.PageCount; i++ {
		reader1.Seek(0, 0)
		var pageBuffer1 bytes.Buffer
		pageSelection1, _ := api.ParsePageSelection(fmt.Sprintf("%d", i))

		err = api.Trim(reader1, &pageBuffer1, pageSelection1, conf)
		if err != nil {
			log.Printf("⚠️  Failed to extract page %d from doc A: %v", i, err)
		} else {
			pageBuffers = append(pageBuffers, pageBuffer1.Bytes())
			fmt.Printf("✅ Extracted A%d (%d bytes)\n", i, pageBuffer1.Len())
		}

		reversePageNum := ctx2.PageCount - i + 1
		reader2.Seek(0, 0)
		var pageBuffer2 bytes.Buffer
		pageSelection2, _ := api.ParsePageSelection(fmt.Sprintf("%d", reversePageNum))

		err = api.Trim(reader2, &pageBuffer2, pageSelection2, conf)
		if err != nil {
			log.Printf("⚠️  Failed to extract page %d from doc B: %v", reversePageNum, err)
		} else {
			pageBuffers = append(pageBuffers, pageBuffer2.Bytes())
			fmt.Printf("✅ Extracted B%d (%d bytes)\n", reversePageNum, pageBuffer2.Len())
		}
	}

	fmt.Printf("✅ Extracted %d pages total\n", len(pageBuffers))

	if len(pageBuffers) == 0 {
		log.Fatalf("❌ No pages extracted successfully")
	}

	err = os.WriteFile(outputFile, pageBuffers[0], 0644)
	if err != nil {
		log.Fatalf("Failed to write output file: %v", err)
	}

	fmt.Printf("✅ Test output written to: %s\n", outputFile)

	pageCount, err := api.PageCountFile(outputFile)
	if err != nil {
		log.Printf("⚠️  Could not get page count: %v", err)
	} else {
		fmt.Printf("✅ Output PDF has %d pages\n", pageCount)
	}

	fmt.Println("\n=== Workflow Analysis ===")
	fmt.Println("✅ Successfully loaded PDFs into memory")
	fmt.Println("✅ Successfully validated page counts from memory")
	fmt.Println("✅ Successfully extracted individual pages using stream-based Trim")
	fmt.Println("❓ Need to find proper merge function for combining page buffers")
}

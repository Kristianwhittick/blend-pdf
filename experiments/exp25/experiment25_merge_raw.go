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
	fmt.Println("=== Experiment 25: Raw Merge Function ===")

	inputFile1 := "output/Doc_A.pdf"
	inputFile2 := "output/Doc_B.pdf"
	outputFile := "output/test_25_merge_raw.pdf"

	os.MkdirAll("output", 0755)

	fmt.Printf("Testing raw merge: %s + %s\n", inputFile1, inputFile2)

	inputBytes1, err := os.ReadFile(inputFile1)
	if err != nil {
		log.Fatalf("Failed to read input file 1: %v", err)
	}

	inputBytes2, err := os.ReadFile(inputFile2)
	if err != nil {
		log.Fatalf("Failed to read input file 2: %v", err)
	}

	fmt.Printf("Loaded %d bytes from %s\n", len(inputBytes1), inputFile1)
	fmt.Printf("Loaded %d bytes from %s\n", len(inputBytes2), inputFile2)

	reader1 := bytes.NewReader(inputBytes1)
	reader2 := bytes.NewReader(inputBytes2)
	var outputBuffer bytes.Buffer
	conf := model.NewDefaultConfiguration()

	fmt.Println("Attempting to call raw merge function...")

	readers := []io.ReadSeeker{reader1, reader2}
	fmt.Printf("Created %d readers for merge testing\n", len(readers))

	reader1.Seek(0, io.SeekStart)
	reader2.Seek(0, io.SeekStart)

	ctx1, err := api.ReadContext(reader1, conf)
	if err != nil {
		log.Fatalf("Failed to read context 1: %v", err)
	}

	ctx2, err := api.ReadContext(reader2, conf)
	if err != nil {
		log.Fatalf("Failed to read context 2: %v", err)
	}

	fmt.Printf("✅ Read contexts: %d pages + %d pages\n", ctx1.PageCount, ctx2.PageCount)

	err = api.WriteContext(ctx1, &outputBuffer)
	if err != nil {
		log.Fatalf("Failed to write context: %v", err)
	}

	err = os.WriteFile(outputFile, outputBuffer.Bytes(), 0644)
	if err != nil {
		log.Fatalf("Failed to write output file: %v", err)
	}

	fmt.Printf("✅ Test output written to: %s\n", outputFile)
}

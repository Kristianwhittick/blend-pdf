// Copyright 2025 Kristian Whittick
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/pdfcpu/pdfcpu/pkg/api"
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/model"
)

func main() {
	fmt.Println("=== Experiment 16: Final Memory Approach ===")
	fmt.Println("Demonstrates optimal in-memory PDF processing with minimal temp files")

	conf := model.NewDefaultConfiguration()

	// Step 1: Load PDF data into memory for validation
	fmt.Println("\n1. Loading PDF data into memory...")

	// Keep original PDF data in memory
	bytesA, err := ioutil.ReadFile("Doc_A.pdf")
	if err != nil {
		log.Fatalf("Error reading Doc_A.pdf: %v", err)
	}

	bytesB, err := ioutil.ReadFile("Doc_B.pdf")
	if err != nil {
		log.Fatalf("Error reading Doc_B.pdf: %v", err)
	}

	fmt.Printf("   ✅ Doc_A: %d bytes loaded\n", len(bytesA))
	fmt.Printf("   ✅ Doc_B: %d bytes loaded\n", len(bytesB))

	// Create contexts for validation (file-based is more reliable)
	ctxA, err := api.ReadContextFile("Doc_A.pdf")
	if err != nil {
		log.Fatalf("Error creating Doc_A context: %v", err)
	}

	ctxB, err := api.ReadContextFile("Doc_B.pdf")
	if err != nil {
		log.Fatalf("Error creating Doc_B context: %v", err)
	}

	fmt.Printf("   ✅ Doc_A context: %d pages\n", ctxA.PageCount)
	fmt.Printf("   ✅ Doc_B context: %d pages\n", ctxB.PageCount)

	// Step 2: Validate in memory
	fmt.Println("\n2. Validating in memory...")

	if ctxA.PageCount != ctxB.PageCount {
		log.Fatalf("❌ Page count mismatch: Doc_A has %d pages, Doc_B has %d pages",
			ctxA.PageCount, ctxB.PageCount)
	}

	err = api.ValidateContext(ctxA)
	if err != nil {
		log.Printf("⚠️  Doc_A validation warning: %v", err)
	} else {
		fmt.Println("   ✅ Doc_A context is valid")
	}

	err = api.ValidateContext(ctxB)
	if err != nil {
		log.Printf("⚠️  Doc_B validation warning: %v", err)
	} else {
		fmt.Println("   ✅ Doc_B context is valid")
	}

	fmt.Printf("   ✅ Both documents have %d pages - ready for interleaved merge\n", ctxA.PageCount)

	// Step 3: Process with minimal temporary files
	fmt.Println("\n3. Processing with minimal temporary files...")

	tempDir := "temp_final"
	os.MkdirAll(tempDir, 0755)
	defer func() {
		fmt.Println("   🧹 Cleaning up temporary files...")
		os.RemoveAll(tempDir)
	}()

	// Function to extract a single page and return its bytes
	extractPageBytes := func(ctx *model.Context, pageNum int, docLabel string) ([]byte, error) {
		// Write context to temp file (minimal I/O)
		tempCtxFile := fmt.Sprintf("%s/%s_ctx.pdf", tempDir, docLabel)
		err := api.WriteContextFile(ctx, tempCtxFile)
		if err != nil {
			return nil, fmt.Errorf("context write error: %v", err)
		}
		defer os.Remove(tempCtxFile)

		// Extract page
		pageFile := fmt.Sprintf("%s/%s_p%d.pdf", tempDir, docLabel, pageNum)
		pageSelection, err := api.ParsePageSelection(fmt.Sprintf("%d", pageNum))
		if err != nil {
			return nil, fmt.Errorf("page selection error: %v", err)
		}

		err = api.TrimFile(tempCtxFile, pageFile, pageSelection, conf)
		if err != nil {
			return nil, fmt.Errorf("page extraction error: %v", err)
		}
		defer os.Remove(pageFile)

		// Read page into memory
		pageBytes, err := ioutil.ReadFile(pageFile)
		if err != nil {
			return nil, fmt.Errorf("page read error: %v", err)
		}

		return pageBytes, nil
	}

	// Step 4: Create interleaved merge in memory
	fmt.Println("\n4. Creating interleaved merge...")

	var mergeSequence [][]byte
	var sequenceLabels []string

	for i := 1; i <= ctxA.PageCount; i++ {
		// Extract page i from Doc_A
		fmt.Printf("   Extracting A page %d...", i)
		pageABytes, err := extractPageBytes(ctxA, i, "docA")
		if err != nil {
			fmt.Printf(" ❌ %v\n", err)
			continue
		}
		fmt.Printf(" ✅ (%d bytes)\n", len(pageABytes))

		mergeSequence = append(mergeSequence, pageABytes)
		sequenceLabels = append(sequenceLabels, fmt.Sprintf("A%d", i))

		// Extract corresponding page from Doc_B (reverse order)
		bPageNum := ctxB.PageCount - i + 1
		fmt.Printf("   Extracting B page %d...", bPageNum)
		pageBBytes, err := extractPageBytes(ctxB, bPageNum, "docB")
		if err != nil {
			fmt.Printf(" ❌ %v\n", err)
			continue
		}
		fmt.Printf(" ✅ (%d bytes)\n", len(pageBBytes))

		mergeSequence = append(mergeSequence, pageBBytes)
		sequenceLabels = append(sequenceLabels, fmt.Sprintf("B%d", bPageNum))
	}

	fmt.Printf("   ✅ Extracted %d pages into memory\n", len(mergeSequence))
	fmt.Printf("   📄 Sequence: %v\n", sequenceLabels)

	// Step 5: Final merge from memory
	fmt.Println("\n5. Final merge from memory...")

	// Write all pages to temp files for final merge
	var finalMergeFiles []string
	for i, pageBytes := range mergeSequence {
		tempPageFile := fmt.Sprintf("%s/final_%02d.pdf", tempDir, i)
		err := ioutil.WriteFile(tempPageFile, pageBytes, 0644)
		if err != nil {
			log.Printf("Error writing temp page %d: %v", i, err)
			continue
		}
		finalMergeFiles = append(finalMergeFiles, tempPageFile)
	}

	// Perform final merge
	outputFile := "output/experiment16_final_interleaved.pdf"
	err = api.MergeCreateFile(finalMergeFiles, outputFile, false, conf)
	if err != nil {
		log.Printf("❌ Final merge error: %v", err)
	} else {
		fmt.Printf("   ✅ Successfully created: %s\n", outputFile)

		// Verify result
		resultCount, err := api.PageCountFile(outputFile)
		if err != nil {
			log.Printf("Error verifying result: %v", err)
		} else {
			fmt.Printf("   📊 Result: %d pages\n", resultCount)

			// Show memory usage summary
			totalOriginalBytes := len(bytesA) + len(bytesB)
			totalExtractedBytes := 0
			for _, pageBytes := range mergeSequence {
				totalExtractedBytes += len(pageBytes)
			}

			fmt.Printf("\n📈 Memory Usage Summary:\n")
			fmt.Printf("   Original PDFs: %d bytes\n", totalOriginalBytes)
			fmt.Printf("   Extracted pages: %d bytes\n", totalExtractedBytes)
			fmt.Printf("   Memory efficiency: %.1f%%\n",
				float64(totalExtractedBytes)/float64(totalOriginalBytes)*100)
		}
	}

	fmt.Println("\n🎉 Experiment 16 completed - Final memory approach demonstrated!")
}

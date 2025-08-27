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
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/pdfcpu/pdfcpu/pkg/api"
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/model"
)

func main() {
	fmt.Println("=== Test 14: Working Memory Approach ===")
	
	conf := model.NewDefaultConfiguration()
	
	// Approach: Use in-memory contexts for validation and processing,
	// but use temporary files for operations that require file paths
	
	fmt.Println("1. Loading PDFs into memory contexts...")
	
	// Read original files as bytes
	bytesA, err := ioutil.ReadFile("Doc_A.pdf")
	if err != nil {
		log.Fatalf("Error reading Doc_A.pdf: %v", err)
	}
	
	bytesB, err := ioutil.ReadFile("Doc_B.pdf")
	if err != nil {
		log.Fatalf("Error reading Doc_B.pdf: %v", err)
	}
	
	fmt.Printf("   Doc_A: %d bytes\n", len(bytesA))
	fmt.Printf("   Doc_B: %d bytes\n", len(bytesB))
	
	// Create contexts from bytes (for validation and info)
	readerA := bytes.NewReader(bytesA)
	ctxA, err := api.ReadContext(readerA, conf)
	if err != nil {
		log.Printf("   Warning: Could not create context from Doc_A bytes: %v", err)
		// Fall back to file-based context
		ctxA, err = api.ReadContextFile("Doc_A.pdf")
		if err != nil {
			log.Fatalf("Error loading Doc_A: %v", err)
		}
	}
	
	readerB := bytes.NewReader(bytesB)
	ctxB, err := api.ReadContext(readerB, conf)
	if err != nil {
		log.Printf("   Warning: Could not create context from Doc_B bytes: %v", err)
		// Fall back to file-based context
		ctxB, err = api.ReadContextFile("Doc_B.pdf")
		if err != nil {
			log.Fatalf("Error loading Doc_B: %v", err)
		}
	}
	
	fmt.Printf("   Doc_A context: %d pages\n", ctxA.PageCount)
	fmt.Printf("   Doc_B context: %d pages\n", ctxB.PageCount)
	
	// 2. Validate page counts match
	fmt.Println("2. Validating page counts...")
	if ctxA.PageCount != ctxB.PageCount {
		log.Fatalf("Page count mismatch: Doc_A has %d pages, Doc_B has %d pages", 
			ctxA.PageCount, ctxB.PageCount)
	}
	fmt.Printf("   ✅ Both documents have %d pages\n", ctxA.PageCount)
	
	// 3. Create interleaved merge using temporary files
	fmt.Println("3. Creating interleaved merge...")
	
	// Create temporary directory for intermediate files
	tempDir := "temp_merge"
	os.MkdirAll(tempDir, 0755)
	defer os.RemoveAll(tempDir)
	
	// Write contexts to temporary files for processing
	tempA := fmt.Sprintf("%s/doc_a.pdf", tempDir)
	tempB := fmt.Sprintf("%s/doc_b.pdf", tempDir)
	
	err = api.WriteContextFile(ctxA, tempA)
	if err != nil {
		log.Fatalf("Error writing temp Doc_A: %v", err)
	}
	
	err = api.WriteContextFile(ctxB, tempB)
	if err != nil {
		log.Fatalf("Error writing temp Doc_B: %v", err)
	}
	
	// Extract individual pages using TrimFile
	var pageFiles []string
	
	for i := 1; i <= ctxA.PageCount; i++ {
		// Extract page i from Doc_A
		pageAFile := fmt.Sprintf("%s/a_page_%d.pdf", tempDir, i)
		pageSelection, err := api.ParsePageSelection(fmt.Sprintf("%d", i))
		if err != nil {
			log.Printf("Error parsing page selection for A page %d: %v", i, err)
			continue
		}
		
		err = api.TrimFile(tempA, pageAFile, pageSelection, conf)
		if err != nil {
			log.Printf("Error extracting A page %d: %v", i, err)
			continue
		}
		
		// Extract corresponding page from Doc_B (in reverse order)
		bPageNum := ctxB.PageCount - i + 1
		pageBFile := fmt.Sprintf("%s/b_page_%d.pdf", tempDir, bPageNum)
		pageSelectionB, err := api.ParsePageSelection(fmt.Sprintf("%d", bPageNum))
		if err != nil {
			log.Printf("Error parsing page selection for B page %d: %v", bPageNum, err)
			continue
		}
		
		err = api.TrimFile(tempB, pageBFile, pageSelectionB, conf)
		if err != nil {
			log.Printf("Error extracting B page %d: %v", bPageNum, err)
			continue
		}
		
		// Add to merge list in interleaved order
		pageFiles = append(pageFiles, pageAFile, pageBFile)
		fmt.Printf("   Added A page %d and B page %d to merge list\n", i, bPageNum)
	}
	
	// Merge all pages
	if len(pageFiles) > 0 {
		fmt.Printf("4. Merging %d page files...\n", len(pageFiles))
		err = api.MergeCreateFile(pageFiles, "output/test14_interleaved.pdf", false, conf)
		if err != nil {
			log.Printf("Error merging pages: %v", err)
		} else {
			fmt.Println("   ✅ Successfully created interleaved merge!")
			
			// Verify result
			resultCount, err := api.PageCountFile("output/test14_interleaved.pdf")
			if err != nil {
				log.Printf("Error checking result: %v", err)
			} else {
				fmt.Printf("   Result has %d pages\n", resultCount)
				expectedPages := ctxA.PageCount + ctxB.PageCount
				if resultCount == expectedPages {
					fmt.Printf("   ✅ Page count correct: %d pages\n", resultCount)
				} else {
					fmt.Printf("   ❌ Page count mismatch: expected %d, got %d\n", expectedPages, resultCount)
				}
			}
		}
	}
	
	fmt.Println("Test 14 completed!")
}

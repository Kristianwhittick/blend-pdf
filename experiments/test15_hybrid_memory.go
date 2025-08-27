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
	fmt.Println("=== Test 15: Hybrid Memory Approach ===")
	
	conf := model.NewDefaultConfiguration()
	
	// Strategy: Keep PDF data in memory as bytes, use file-based contexts
	// for operations, but minimize disk I/O by using temporary files
	
	fmt.Println("1. Loading PDF data into memory...")
	
	// Read original files into memory
	bytesA, err := ioutil.ReadFile("Doc_A.pdf")
	if err != nil {
		log.Fatalf("Error reading Doc_A.pdf: %v", err)
	}
	
	bytesB, err := ioutil.ReadFile("Doc_B.pdf")
	if err != nil {
		log.Fatalf("Error reading Doc_B.pdf: %v", err)
	}
	
	fmt.Printf("   Doc_A: %d bytes in memory\n", len(bytesA))
	fmt.Printf("   Doc_B: %d bytes in memory\n", len(bytesB))
	
	// Create contexts using file-based approach (more reliable)
	ctxA, err := api.ReadContextFile("Doc_A.pdf")
	if err != nil {
		log.Fatalf("Error loading Doc_A context: %v", err)
	}
	
	ctxB, err := api.ReadContextFile("Doc_B.pdf")
	if err != nil {
		log.Fatalf("Error loading Doc_B context: %v", err)
	}
	
	fmt.Printf("   Doc_A context: %d pages\n", ctxA.PageCount)
	fmt.Printf("   Doc_B context: %d pages\n", ctxB.PageCount)
	
	// 2. Validate and process in memory
	fmt.Println("2. Processing in memory...")
	if ctxA.PageCount != ctxB.PageCount {
		log.Fatalf("Page count mismatch: Doc_A has %d pages, Doc_B has %d pages", 
			ctxA.PageCount, ctxB.PageCount)
	}
	
	// 3. Create interleaved merge with minimal temp files
	fmt.Println("3. Creating interleaved merge with minimal temp files...")
	
	// Create a single temporary directory
	tempDir := "temp_hybrid"
	os.MkdirAll(tempDir, 0755)
	defer os.RemoveAll(tempDir)
	
	// Function to create temp file from context and extract page
	extractPageToBytes := func(ctx *model.Context, pageNum int, label string) ([]byte, error) {
		// Write context to temp file
		tempFile := fmt.Sprintf("%s/%s_temp.pdf", tempDir, label)
		err := api.WriteContextFile(ctx, tempFile)
		if err != nil {
			return nil, fmt.Errorf("error writing temp file: %v", err)
		}
		defer os.Remove(tempFile)
		
		// Extract the page
		pageFile := fmt.Sprintf("%s/%s_page_%d.pdf", tempDir, label, pageNum)
		pageSelection, err := api.ParsePageSelection(fmt.Sprintf("%d", pageNum))
		if err != nil {
			return nil, fmt.Errorf("error parsing page selection: %v", err)
		}
		
		err = api.TrimFile(tempFile, pageFile, pageSelection, conf)
		if err != nil {
			return nil, fmt.Errorf("error extracting page: %v", err)
		}
		defer os.Remove(pageFile)
		
		// Read extracted page into memory
		pageBytes, err := ioutil.ReadFile(pageFile)
		if err != nil {
			return nil, fmt.Errorf("error reading extracted page: %v", err)
		}
		
		return pageBytes, nil
	}
	
	// Extract all pages into memory
	var pageDataList [][]byte
	var pageLabels []string
	
	for i := 1; i <= ctxA.PageCount; i++ {
		// Extract page i from Doc_A
		fmt.Printf("   Extracting A page %d...\n", i)
		pageABytes, err := extractPageToBytes(ctxA, i, "docA")
		if err != nil {
			log.Printf("Error extracting A page %d: %v", i, err)
			continue
		}
		pageDataList = append(pageDataList, pageABytes)
		pageLabels = append(pageLabels, fmt.Sprintf("A%d", i))
		
		// Extract corresponding page from Doc_B (reverse order)
		bPageNum := ctxB.PageCount - i + 1
		fmt.Printf("   Extracting B page %d...\n", bPageNum)
		pageBBytes, err := extractPageToBytes(ctxB, bPageNum, "docB")
		if err != nil {
			log.Printf("Error extracting B page %d: %v", bPageNum, err)
			continue
		}
		pageDataList = append(pageDataList, pageBBytes)
		pageLabels = append(pageLabels, fmt.Sprintf("B%d", bPageNum))
	}
	
	fmt.Printf("   Extracted %d pages into memory\n", len(pageDataList))
	
	// 4. Merge pages from memory
	fmt.Println("4. Merging pages from memory...")
	
	// Write all page data to temp files for merging
	var mergeFiles []string
	for i, pageData := range pageDataList {
		tempPageFile := fmt.Sprintf("%s/merge_page_%d.pdf", tempDir, i)
		err := ioutil.WriteFile(tempPageFile, pageData, 0644)
		if err != nil {
			log.Printf("Error writing temp page %d: %v", i, err)
			continue
		}
		mergeFiles = append(mergeFiles, tempPageFile)
	}
	
	// Perform the merge
	if len(mergeFiles) > 0 {
		err = api.MergeCreateFile(mergeFiles, "output/test15_hybrid_interleaved.pdf", false, conf)
		if err != nil {
			log.Printf("Error merging: %v", err)
		} else {
			fmt.Println("   ✅ Successfully created hybrid interleaved merge!")
			
			// Verify result
			resultCount, err := api.PageCountFile("output/test15_hybrid_interleaved.pdf")
			if err != nil {
				log.Printf("Error checking result: %v", err)
			} else {
				expectedPages := ctxA.PageCount + ctxB.PageCount
				fmt.Printf("   Result: %d pages (expected %d)\n", resultCount, expectedPages)
				
				if resultCount == expectedPages {
					fmt.Println("   ✅ Hybrid memory approach successful!")
					
					// Show the page order
					fmt.Println("   Page order:", pageLabels)
				}
			}
		}
	}
	
	fmt.Println("Test 15 completed!")
}

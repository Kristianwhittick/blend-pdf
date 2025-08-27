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
	"log"

	"github.com/pdfcpu/pdfcpu/pkg/api"
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/model"
)

func main() {
	fmt.Println("=== Test 11: Memory Context Merging ===")
	
	// Create default configuration
	conf := model.NewDefaultConfiguration()
	
	// Load both PDFs into memory contexts
	fmt.Println("Loading PDFs into memory contexts...")
	ctxA, err := api.ReadContextFile("Doc_A.pdf")
	if err != nil {
		log.Fatalf("Error loading Doc_A.pdf: %v", err)
	}
	fmt.Printf("Doc_A loaded: %d pages\n", ctxA.PageCount)
	
	ctxB, err := api.ReadContextFile("Doc_B.pdf")
	if err != nil {
		log.Fatalf("Error loading Doc_B.pdf: %v", err)
	}
	fmt.Printf("Doc_B loaded: %d pages\n", ctxB.PageCount)
	
	// Method 1: Try MergeContext function
	fmt.Println("\nTesting api.MergeContext...")
	contexts := []*model.Context{ctxA, ctxB}
	ctxMerged, err := api.MergeContext(contexts, conf)
	if err != nil {
		log.Printf("MergeContext error: %v", err)
	} else {
		fmt.Printf("MergeContext success! Merged context has %d pages\n", ctxMerged.PageCount)
		
		// Write merged result
		err = api.WriteContextFile(ctxMerged, "output/test11_merged_simple.pdf")
		if err != nil {
			log.Printf("Error writing merged file: %v", err)
		} else {
			fmt.Println("Successfully wrote merged file to output/test11_merged_simple.pdf")
		}
	}
	
	// Method 2: Try manual page-by-page merging
	fmt.Println("\nTesting manual page extraction and merging...")
	
	// Extract individual pages from Doc_A
	page1A, err := api.ExtractPages(ctxA, []string{"1"}, conf)
	if err != nil {
		log.Printf("Error extracting page 1 from Doc_A: %v", err)
		return
	}
	
	page2A, err := api.ExtractPages(ctxA, []string{"2"}, conf)
	if err != nil {
		log.Printf("Error extracting page 2 from Doc_A: %v", err)
		return
	}
	
	// Extract individual pages from Doc_B (in reverse order)
	page3B, err := api.ExtractPages(ctxB, []string{"3"}, conf)
	if err != nil {
		log.Printf("Error extracting page 3 from Doc_B: %v", err)
		return
	}
	
	page2B, err := api.ExtractPages(ctxB, []string{"2"}, conf)
	if err != nil {
		log.Printf("Error extracting page 2 from Doc_B: %v", err)
		return
	}
	
	fmt.Println("Individual pages extracted successfully")
	
	// Try to merge in interleaved pattern: A1, B3, A2, B2
	fmt.Println("Attempting interleaved merge...")
	interleavedContexts := []*model.Context{page1A, page3B, page2A, page2B}
	ctxInterleaved, err := api.MergeContext(interleavedContexts, conf)
	if err != nil {
		log.Printf("Interleaved merge error: %v", err)
	} else {
		fmt.Printf("Interleaved merge success! Result has %d pages\n", ctxInterleaved.PageCount)
		
		// Write interleaved result
		err = api.WriteContextFile(ctxInterleaved, "output/test11_interleaved.pdf")
		if err != nil {
			log.Printf("Error writing interleaved file: %v", err)
		} else {
			fmt.Println("Successfully wrote interleaved file to output/test11_interleaved.pdf")
		}
	}
	
	fmt.Println("Test 11 completed!")
}

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
	"os"

	"github.com/pdfcpu/pdfcpu/pkg/api"
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/model"
)

func main() {
	fmt.Println("=== Experiment 10: Memory Page Extraction (Simple) ===")
	
	// Create default configuration
	conf := model.NewDefaultConfiguration()
	
	// Load Doc_A.pdf into memory context
	fmt.Println("Loading Doc_A.pdf into memory context...")
	ctxA, err := api.ReadContextFile("Doc_A.pdf")
	if err != nil {
		log.Fatalf("Error loading Doc_A.pdf: %v", err)
	}
	fmt.Printf("Doc_A loaded: %d pages\n", ctxA.PageCount)
	
	// Try to write the context back to a file (this should work)
	fmt.Println("Testing WriteContextFile...")
	err = api.WriteContextFile(ctxA, "output/experiment10_context_copy.pdf")
	if err != nil {
		log.Printf("WriteContextFile error: %v", err)
	} else {
		fmt.Println("Successfully wrote context to output/experiment10_context_copy.pdf")
		
		// Verify the copy
		pageCount, err := api.PageCountFile("output/experiment10_context_copy.pdf")
		if err != nil {
			log.Printf("Error checking copied file: %v", err)
		} else {
			fmt.Printf("Copied file has %d pages\n", pageCount)
		}
	}
	
	// Try to use TrimFile with a context-based approach
	fmt.Println("\nTesting extraction using file operations on context...")
	
	// First, write context to a temp file
	tempFile := "temp_doc_a.pdf"
	err = api.WriteContextFile(ctxA, tempFile)
	if err != nil {
		log.Printf("Error writing temp file: %v", err)
		return
	}
	
	// Extract page 1 using TrimFile
	pageSelection, err := api.ParsePageSelection("1")
	if err != nil {
		log.Printf("Error parsing page selection: %v", err)
		return
	}
	
	err = api.TrimFile(tempFile, "output/experiment10_extracted_page1.pdf", pageSelection, conf)
	if err != nil {
		log.Printf("TrimFile error: %v", err)
	} else {
		fmt.Println("Successfully extracted page 1 to output/experiment10_extracted_page1.pdf")
		
		// Load the extracted page back into context
		ctxExtracted, err := api.ReadContextFile("output/experiment10_extracted_page1.pdf")
		if err != nil {
			log.Printf("Error loading extracted page: %v", err)
		} else {
			fmt.Printf("Extracted page context has %d pages\n", ctxExtracted.PageCount)
		}
	}
	
	// Clean up temp file
	os.Remove(tempFile)
	
	fmt.Println("Experiment 10 completed!")
}

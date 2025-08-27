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
	fmt.Println("=== Test 10: Memory Page Extraction ===")
	
	// Create default configuration
	conf := model.NewDefaultConfiguration()
	
	// Load Doc_A.pdf into memory context
	fmt.Println("Loading Doc_A.pdf into memory context...")
	ctxA, err := api.ReadContextFile("Doc_A.pdf")
	if err != nil {
		log.Fatalf("Error loading Doc_A.pdf: %v", err)
	}
	
	// Try to extract page 1 in memory
	fmt.Println("Attempting to extract page 1 from Doc_A...")
	
	// Method 1: Try ExtractPages function
	fmt.Println("Testing api.ExtractPages...")
	selectedPages := []string{"1"}
	ctxExtracted, err := api.ExtractPages(ctxA, selectedPages, conf)
	if err != nil {
		log.Printf("ExtractPages error: %v", err)
	} else {
		fmt.Printf("ExtractPages success! Extracted context has %d pages\n", ctxExtracted.PageCount)
		
		// Try to write extracted page to file
		err = api.WriteContextFile(ctxExtracted, "output/test10_extracted_page1.pdf")
		if err != nil {
			log.Printf("Error writing extracted page: %v", err)
		} else {
			fmt.Println("Successfully wrote extracted page to output/test10_extracted_page1.pdf")
		}
	}
	
	// Method 2: Try TrimContext function
	fmt.Println("\nTesting api.TrimContext...")
	ctxTrimmed, err := api.TrimContext(ctxA, selectedPages, conf)
	if err != nil {
		log.Printf("TrimContext error: %v", err)
	} else {
		fmt.Printf("TrimContext success! Trimmed context has %d pages\n", ctxTrimmed.PageCount)
		
		// Try to write trimmed page to file
		err = api.WriteContextFile(ctxTrimmed, "output/test10_trimmed_page1.pdf")
		if err != nil {
			log.Printf("Error writing trimmed page: %v", err)
		} else {
			fmt.Println("Successfully wrote trimmed page to output/test10_trimmed_page1.pdf")
		}
	}
	
	fmt.Println("Test 10 completed!")
}

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
	fmt.Println("=== Experiment 12: Memory Bytes Processing ===")
	
	// Create default configuration
	conf := model.NewDefaultConfiguration()
	
	// Read PDF files as byte arrays
	fmt.Println("Reading Doc_A.pdf as bytes...")
	bytesA, err := ioutil.ReadFile("Doc_A.pdf")
	if err != nil {
		log.Fatalf("Error reading Doc_A.pdf: %v", err)
	}
	fmt.Printf("Doc_A.pdf size: %d bytes\n", len(bytesA))
	
	fmt.Println("Reading Doc_B.pdf as bytes...")
	bytesB, err := ioutil.ReadFile("Doc_B.pdf")
	if err != nil {
		log.Fatalf("Error reading Doc_B.pdf: %v", err)
	}
	fmt.Printf("Doc_B.pdf size: %d bytes\n", len(bytesB))
	
	// Try to create contexts from byte arrays
	fmt.Println("\nTesting api.ReadContext from bytes...")
	
	// Method 1: ReadContext from bytes
	ctxA, err := api.ReadContext(bytesA, conf)
	if err != nil {
		log.Printf("Error creating context from Doc_A bytes: %v", err)
	} else {
		fmt.Printf("Doc_A context from bytes: %d pages\n", ctxA.PageCount)
	}
	
	ctxB, err := api.ReadContext(bytesB, conf)
	if err != nil {
		log.Printf("Error creating context from Doc_B bytes: %v", err)
	} else {
		fmt.Printf("Doc_B context from bytes: %d pages\n", ctxB.PageCount)
	}
	
	// Test writing context back to bytes
	if ctxA != nil {
		fmt.Println("\nTesting api.WriteContext to bytes...")
		resultBytes, err := api.WriteContext(ctxA, conf)
		if err != nil {
			log.Printf("Error writing context to bytes: %v", err)
		} else {
			fmt.Printf("Successfully converted context back to bytes: %d bytes\n", len(resultBytes))
			
			// Write bytes to file for verification
			err = ioutil.WriteFile("output/experiment12_from_bytes.pdf", resultBytes, 0644)
			if err != nil {
				log.Printf("Error writing bytes to file: %v", err)
			} else {
				fmt.Println("Successfully wrote bytes to output/experiment12_from_bytes.pdf")
			}
		}
	}
	
	// Test page extraction from byte-based context
	if ctxA != nil && ctxA.PageCount > 0 {
		fmt.Println("\nTesting page extraction from byte-based context...")
		page1Ctx, err := api.ExtractPages(ctxA, []string{"1"}, conf)
		if err != nil {
			log.Printf("Error extracting page from byte-based context: %v", err)
		} else {
			fmt.Printf("Extracted page 1: %d pages in result\n", page1Ctx.PageCount)
			
			// Convert extracted page back to bytes
			page1Bytes, err := api.WriteContext(page1Ctx, conf)
			if err != nil {
				log.Printf("Error converting extracted page to bytes: %v", err)
			} else {
				fmt.Printf("Extracted page as bytes: %d bytes\n", len(page1Bytes))
				
				// Write to file
				err = ioutil.WriteFile("output/experiment12_extracted_page1.pdf", page1Bytes, 0644)
				if err != nil {
					log.Printf("Error writing extracted page: %v", err)
				} else {
					fmt.Println("Successfully wrote extracted page to output/experiment12_extracted_page1.pdf")
				}
			}
		}
	}
	
	fmt.Println("Experiment 12 completed!")
}

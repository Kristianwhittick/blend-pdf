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
	fmt.Println("=== Experiment 10: Memory Page Extraction ===")

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

	// Method 1: Try TrimFile function (file-based extraction)
	fmt.Println("Testing api.TrimFile...")
	selectedPages, err := api.ParsePageSelection("1")
	if err != nil {
		log.Printf("ParsePageSelection error: %v", err)
		return
	}

	// Write context to temp file first
	tempFile := "temp_doc_a.pdf"
	err = api.WriteContextFile(ctxA, tempFile)
	if err != nil {
		log.Printf("Error writing temp file: %v", err)
		return
	}

	err = api.TrimFile(tempFile, "output/experiment10_extracted_page1.pdf", selectedPages, conf)
	if err != nil {
		log.Printf("TrimFile error: %v", err)
	} else {
		fmt.Println("Successfully extracted page 1 to output/experiment10_extracted_page1.pdf")
	}

	// Method 2: Try CollectFile function (order-preserving extraction)
	fmt.Println("\nTesting api.CollectFile...")
	err = api.CollectFile(tempFile, "output/experiment10_collected_page1.pdf", selectedPages, conf)
	if err != nil {
		log.Printf("CollectFile error: %v", err)
	} else {
		fmt.Println("Successfully collected page 1 to output/experiment10_collected_page1.pdf")
	}

	fmt.Println("Experiment 10 completed!")
}

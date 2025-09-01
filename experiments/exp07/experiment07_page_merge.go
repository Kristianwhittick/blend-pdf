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
	fmt.Println("=== Experiment 07: Merge Individual Pages ===")

	// Create default configuration
	conf := model.NewDefaultConfiguration()

	// Extract individual pages first
	fmt.Println("Extracting individual pages...")

	// Extract A1 (page 1 from Doc_A)
	pageSelection, _ := api.ParsePageSelection("1")
	err := api.TrimFile("Doc_A.pdf", "temp_A1.pdf", pageSelection, conf)
	if err != nil {
		log.Fatalf("Failed to extract A1: %v", err)
	}

	// Extract A2 (page 2 from Doc_A)
	pageSelection, _ = api.ParsePageSelection("2")
	err = api.TrimFile("Doc_A.pdf", "temp_A2.pdf", pageSelection, conf)
	if err != nil {
		log.Fatalf("Failed to extract A2: %v", err)
	}

	// Extract B3 (page 1 from Doc_B, which contains B3)
	pageSelection, _ = api.ParsePageSelection("1")
	err = api.TrimFile("Doc_B.pdf", "temp_B3.pdf", pageSelection, conf)
	if err != nil {
		log.Fatalf("Failed to extract B3: %v", err)
	}

	// Extract B2 (page 2 from Doc_B, which contains B2)
	pageSelection, _ = api.ParsePageSelection("2")
	err = api.TrimFile("Doc_B.pdf", "temp_B2.pdf", pageSelection, conf)
	if err != nil {
		log.Fatalf("Failed to extract B2: %v", err)
	}

	// Now merge in interleaved order: A1, B3, A2, B2
	fmt.Println("Merging in interleaved order: A1, B3, A2, B2...")

	inputFiles := []string{"temp_A1.pdf", "temp_B3.pdf", "temp_A2.pdf", "temp_B2.pdf"}

	err = api.MergeCreateFile(inputFiles, "output/experiment07_page_merge.pdf", false, conf)
	if err != nil {
		log.Fatalf("Failed to merge individual pages: %v", err)
	}

	// Clean up temporary files
	os.Remove("temp_A1.pdf")
	os.Remove("temp_A2.pdf")
	os.Remove("temp_B3.pdf")
	os.Remove("temp_B2.pdf")

	fmt.Println("Successfully merged individual pages to output/experiment07_page_merge.pdf")

	// Verify the result
	pageCount, err := api.PageCountFile("output/experiment07_page_merge.pdf")
	if err != nil {
		log.Fatalf("Failed to get page count of result: %v", err)
	}

	fmt.Printf("Result file has %d page(s)\n", pageCount)

	if pageCount == 4 {
		fmt.Println("✅ Test 07 PASSED - Individual pages merged successfully!")
	} else {
		fmt.Printf("❌ Test 07 FAILED - Expected 4 pages, got %d\n", pageCount)
	}
}

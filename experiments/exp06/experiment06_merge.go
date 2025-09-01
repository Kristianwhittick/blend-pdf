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
	fmt.Println("=== Experiment 06: Simple Merge Two Files ===")

	// Create default configuration
	conf := model.NewDefaultConfiguration()

	// Simple merge: Doc_A.pdf + Doc_B.pdf
	fmt.Println("Merging Doc_A.pdf + Doc_B.pdf...")

	inputFiles := []string{"Doc_A.pdf", "Doc_B.pdf"}

	err := api.MergeCreateFile(inputFiles, "output/experiment06_simple_merge.pdf", false, conf)
	if err != nil {
		log.Fatalf("Failed to merge files: %v", err)
	}

	fmt.Println("Successfully merged to output/experiment06_simple_merge.pdf")

	// Verify the result
	pageCount, err := api.PageCountFile("output/experiment06_simple_merge.pdf")
	if err != nil {
		log.Fatalf("Failed to get page count of result: %v", err)
	}

	fmt.Printf("Result file has %d page(s)\n", pageCount)

	if pageCount == 6 {
		fmt.Println("✅ Test 06 PASSED - Files merged successfully!")
	} else {
		fmt.Printf("❌ Test 06 FAILED - Expected 6 pages, got %d\n", pageCount)
	}
}

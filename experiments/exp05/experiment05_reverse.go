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
	fmt.Println("=== Experiment 05: Extract Pages in Reverse Order ===")
	
	// Create default configuration
	conf := model.NewDefaultConfiguration()
	
	// Extract pages 3,2,1 from Doc_B.pdf (reverse order)
	fmt.Println("Extracting pages 3,2,1 from Doc_B.pdf...")
	
	pageSelection, err := api.ParsePageSelection("3,2,1")
	if err != nil {
		log.Fatalf("Failed to parse page selection: %v", err)
	}
	
	err = api.TrimFile("Doc_B.pdf", "output/experiment05_reverse.pdf", pageSelection, conf)
	if err != nil {
		log.Fatalf("Failed to extract pages: %v", err)
	}
	
	fmt.Println("Successfully extracted pages 3,2,1 to output/experiment05_reverse.pdf")
	
	// Verify the result
	pageCount, err := api.PageCountFile("output/experiment05_reverse.pdf")
	if err != nil {
		log.Fatalf("Failed to get page count of result: %v", err)
	}
	
	fmt.Printf("Result file has %d page(s)\n", pageCount)
	
	if pageCount == 3 {
		fmt.Println("✅ Test 05 PASSED - Pages extracted in reverse order!")
	} else {
		fmt.Printf("❌ Test 05 FAILED - Expected 3 pages, got %d\n", pageCount)
	}
}

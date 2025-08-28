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
	fmt.Println("=== Experiment 03: Extract Single Page ===")
	
	// Create default configuration
	conf := model.NewDefaultConfiguration()
	
	// Extract page 1 from Doc_A.pdf
	fmt.Println("Extracting page 1 from Doc_A.pdf...")
	
	pageSelection, err := api.ParsePageSelection("1")
	if err != nil {
		log.Fatalf("Failed to parse page selection: %v", err)
	}
	
	err = api.TrimFile("Doc_A.pdf", "output/experiment03_single_page.pdf", pageSelection, conf)
	if err != nil {
		log.Fatalf("Failed to extract page: %v", err)
	}
	
	fmt.Println("Successfully extracted page 1 to output/experiment03_single_page.pdf")
	
	// Verify the result
	pageCount, err := api.PageCountFile("output/experiment03_single_page.pdf")
	if err != nil {
		log.Fatalf("Failed to get page count of result: %v", err)
	}
	
	fmt.Printf("Result file has %d page(s)\n", pageCount)
	
	if pageCount == 1 {
		fmt.Println("✅ Test 03 PASSED - Single page extracted successfully!")
	} else {
		fmt.Printf("❌ Test 03 FAILED - Expected 1 page, got %d\n", pageCount)
	}
}

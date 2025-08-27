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
	fmt.Println("=== Test 04: Extract Multiple Pages ===")
	
	// Create default configuration
	conf := model.NewDefaultConfiguration()
	
	// Extract pages 1-2 from Doc_A.pdf
	fmt.Println("Extracting pages 1-2 from Doc_A.pdf...")
	
	pageSelection, err := api.ParsePageSelection("1-2")
	if err != nil {
		log.Fatalf("Failed to parse page selection: %v", err)
	}
	
	err = api.TrimFile("Doc_A.pdf", "output/test04_multi_pages.pdf", pageSelection, conf)
	if err != nil {
		log.Fatalf("Failed to extract pages: %v", err)
	}
	
	fmt.Println("Successfully extracted pages 1-2 to output/test04_multi_pages.pdf")
	
	// Verify the result
	pageCount, err := api.PageCountFile("output/test04_multi_pages.pdf")
	if err != nil {
		log.Fatalf("Failed to get page count of result: %v", err)
	}
	
	fmt.Printf("Result file has %d page(s)\n", pageCount)
	
	if pageCount == 2 {
		fmt.Println("✅ Test 04 PASSED - Multiple pages extracted successfully!")
	} else {
		fmt.Printf("❌ Test 04 FAILED - Expected 2 pages, got %d\n", pageCount)
	}
}

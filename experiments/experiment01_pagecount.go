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
)

func experiment01PageCount() {
	fmt.Println("=== Experiment 01: Page Count API ===")
	
	// Test Doc_A.pdf
	fmt.Println("Testing Doc_A.pdf...")
	countA, err := api.PageCountFile("Doc_A.pdf")
	if err != nil {
		log.Fatalf("Error getting page count for Doc_A.pdf: %v", err)
	}
	fmt.Printf("Doc_A.pdf has %d pages\n", countA)
	
	// Test Doc_B.pdf
	fmt.Println("Testing Doc_B.pdf...")
	countB, err := api.PageCountFile("Doc_B.pdf")
	if err != nil {
		log.Fatalf("Error getting page count for Doc_B.pdf: %v", err)
	}
	fmt.Printf("Doc_B.pdf has %d pages\n", countB)
	
	// Summary
	fmt.Printf("\nSummary:\n")
	fmt.Printf("- Doc_A.pdf: %d pages\n", countA)
	fmt.Printf("- Doc_B.pdf: %d pages\n", countB)
	fmt.Printf("- Pages match: %t\n", countA == countB)
	
	fmt.Println("Experiment 01 completed successfully!")
}

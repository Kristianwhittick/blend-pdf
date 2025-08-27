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

func main() {
	fmt.Println("=== Test 09: Memory Context API ===")
	
	// Test loading Doc_A.pdf into memory context
	fmt.Println("Loading Doc_A.pdf into memory context...")
	ctxA, err := api.ReadContextFile("Doc_A.pdf")
	if err != nil {
		log.Fatalf("Error loading Doc_A.pdf into context: %v", err)
	}
	fmt.Printf("Doc_A context loaded successfully\n")
	fmt.Printf("Doc_A page count from context: %d\n", ctxA.PageCount)
	
	// Test loading Doc_B.pdf into memory context
	fmt.Println("Loading Doc_B.pdf into memory context...")
	ctxB, err := api.ReadContextFile("Doc_B.pdf")
	if err != nil {
		log.Fatalf("Error loading Doc_B.pdf into context: %v", err)
	}
	fmt.Printf("Doc_B context loaded successfully\n")
	fmt.Printf("Doc_B page count from context: %d\n", ctxB.PageCount)
	
	// Test validation from context
	fmt.Println("\nValidating contexts...")
	err = api.ValidateContext(ctxA)
	if err != nil {
		log.Printf("Doc_A validation error: %v", err)
	} else {
		fmt.Println("Doc_A context is valid")
	}
	
	err = api.ValidateContext(ctxB)
	if err != nil {
		log.Printf("Doc_B validation error: %v", err)
	} else {
		fmt.Println("Doc_B context is valid")
	}
	
	// Summary
	fmt.Printf("\nSummary:\n")
	fmt.Printf("- Doc_A context pages: %d\n", ctxA.PageCount)
	fmt.Printf("- Doc_B context pages: %d\n", ctxB.PageCount)
	fmt.Printf("- Both contexts loaded successfully\n")
	
	fmt.Println("Test 09 completed successfully!")
}

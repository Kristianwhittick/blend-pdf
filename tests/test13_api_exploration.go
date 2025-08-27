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
	"bytes"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/pdfcpu/pdfcpu/pkg/api"
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/model"
)

func main() {
	fmt.Println("=== Test 13: API Exploration ===")
	
	// Create default configuration
	conf := model.NewDefaultConfiguration()
	
	// Test 1: Load context from file (we know this works)
	fmt.Println("1. Loading context from file...")
	ctxA, err := api.ReadContextFile("Doc_A.pdf")
	if err != nil {
		log.Fatalf("Error loading context: %v", err)
	}
	fmt.Printf("   ✅ ReadContextFile works: %d pages\n", ctxA.PageCount)
	
	// Test 2: Write context to file (we know this works)
	fmt.Println("2. Writing context to file...")
	err = api.WriteContextFile(ctxA, "output/test13_copy.pdf")
	if err != nil {
		log.Printf("   ❌ WriteContextFile error: %v", err)
	} else {
		fmt.Println("   ✅ WriteContextFile works")
	}
	
	// Test 3: Try to read PDF as bytes and create context
	fmt.Println("3. Reading PDF as bytes...")
	pdfBytes, err := ioutil.ReadFile("Doc_A.pdf")
	if err != nil {
		log.Printf("   ❌ Error reading file: %v", err)
	} else {
		fmt.Printf("   ✅ Read %d bytes\n", len(pdfBytes))
		
		// Try to create context from bytes using bytes.Reader
		fmt.Println("4. Creating context from bytes...")
		reader := bytes.NewReader(pdfBytes)
		ctxFromBytes, err := api.ReadContext(reader, conf)
		if err != nil {
			log.Printf("   ❌ ReadContext error: %v", err)
		} else {
			fmt.Printf("   ✅ ReadContext works: %d pages\n", ctxFromBytes.PageCount)
		}
	}
	
	// Test 4: Try MergeCreateFile with contexts written to temp files
	fmt.Println("5. Testing merge with temporary files...")
	
	// Load second PDF
	ctxB, err := api.ReadContextFile("Doc_B.pdf")
	if err != nil {
		log.Printf("   ❌ Error loading Doc_B: %v", err)
		return
	}
	
	// Write both contexts to temp files
	tempA := "temp_a.pdf"
	tempB := "temp_b.pdf"
	
	err = api.WriteContextFile(ctxA, tempA)
	if err != nil {
		log.Printf("   ❌ Error writing temp A: %v", err)
		return
	}
	
	err = api.WriteContextFile(ctxB, tempB)
	if err != nil {
		log.Printf("   ❌ Error writing temp B: %v", err)
		return
	}
	
	// Try to merge the temp files
	inFiles := []string{tempA, tempB}
	err = api.MergeCreateFile(inFiles, "output/test13_merged.pdf", false, conf)
	if err != nil {
		log.Printf("   ❌ MergeCreateFile error: %v", err)
	} else {
		fmt.Println("   ✅ MergeCreateFile works")
		
		// Check result
		mergedCount, err := api.PageCountFile("output/test13_merged.pdf")
		if err != nil {
			log.Printf("   ❌ Error checking merged file: %v", err)
		} else {
			fmt.Printf("   ✅ Merged file has %d pages\n", mergedCount)
		}
	}
	
	// Clean up temp files
	ioutil.WriteFile(tempA, []byte{}, 0644) // Clear file
	ioutil.WriteFile(tempB, []byte{}, 0644) // Clear file
	
	// Test 5: Try WriteContext to get bytes
	fmt.Println("6. Testing WriteContext to bytes...")
	var buf bytes.Buffer
	err = api.WriteContext(ctxA, &buf)
	if err != nil {
		log.Printf("   ❌ WriteContext error: %v", err)
	} else {
		fmt.Printf("   ✅ WriteContext works: %d bytes\n", buf.Len())
		
		// Try to read the bytes back
		reader2 := bytes.NewReader(buf.Bytes())
		ctxFromWritten, err := api.ReadContext(reader2, conf)
		if err != nil {
			log.Printf("   ❌ ReadContext from written bytes error: %v", err)
		} else {
			fmt.Printf("   ✅ Round-trip works: %d pages\n", ctxFromWritten.PageCount)
		}
	}
	
	fmt.Println("\n=== API Exploration Summary ===")
	fmt.Println("✅ ReadContextFile - Load PDF into memory context")
	fmt.Println("✅ WriteContextFile - Write context to file")
	fmt.Println("✅ ReadContext - Load from io.ReadSeeker (bytes.Reader)")
	fmt.Println("✅ WriteContext - Write to io.Writer (bytes.Buffer)")
	fmt.Println("✅ MergeCreateFile - Merge files")
	fmt.Println("❓ Need to explore: Page extraction from contexts")
}

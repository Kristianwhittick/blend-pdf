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
	fmt.Println("=== Experiment 12: Memory Bytes Processing ===")

	// Create default configuration
	conf := model.NewDefaultConfiguration()

	// Load PDF files into memory as bytes
	fmt.Println("Loading Doc_A.pdf into memory...")
	bytesA, err := ioutil.ReadFile("Doc_A.pdf")
	if err != nil {
		log.Fatalf("Error reading Doc_A.pdf: %v", err)
	}
	fmt.Printf("Doc_A loaded: %d bytes\n", len(bytesA))

	fmt.Println("Loading Doc_B.pdf into memory...")
	bytesB, err := ioutil.ReadFile("Doc_B.pdf")
	if err != nil {
		log.Fatalf("Error reading Doc_B.pdf: %v", err)
	}
	fmt.Printf("Doc_B loaded: %d bytes\n", len(bytesB))

	// Try to create contexts from byte arrays
	fmt.Println("Creating context from Doc_A bytes...")
	readerA := bytes.NewReader(bytesA)
	ctxA, err := api.ReadContext(readerA, conf)
	if err != nil {
		log.Fatalf("Error creating context from Doc_A bytes: %v", err)
	}
	fmt.Printf("Doc_A context created: %d pages\n", ctxA.PageCount)

	fmt.Println("Creating context from Doc_B bytes...")
	readerB := bytes.NewReader(bytesB)
	ctxB, err := api.ReadContext(readerB, conf)
	if err != nil {
		log.Fatalf("Error creating context from Doc_B bytes: %v", err)
	}
	fmt.Printf("Doc_B context created: %d pages\n", ctxB.PageCount)

	// Try to write context back to bytes
	fmt.Println("Converting Doc_A context back to bytes...")
	var bufferA bytes.Buffer
	err = api.WriteContext(ctxA, &bufferA)
	if err != nil {
		log.Fatalf("Error writing Doc_A context to bytes: %v", err)
	}
	fmt.Printf("Doc_A context converted: %d bytes\n", bufferA.Len())

	fmt.Println("SUCCESS: Memory bytes processing completed!")
	fmt.Printf("Original Doc_A: %d bytes, Processed: %d bytes\n", len(bytesA), bufferA.Len())
	fmt.Printf("Original Doc_B: %d bytes, Context pages: %d\n", len(bytesB), ctxB.PageCount)
}

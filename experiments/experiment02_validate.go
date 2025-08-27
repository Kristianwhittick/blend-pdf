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
	fmt.Println("=== Experiment 02: PDF Validation API ===")
	
	// Create default configuration
	conf := model.NewDefaultConfiguration()
	conf.ValidationMode = model.ValidationRelaxed
	
	// Test Doc_A.pdf
	fmt.Println("Validating Doc_A.pdf...")
	err := api.ValidateFile("Doc_A.pdf", conf)
	if err != nil {
		log.Printf("Doc_A.pdf validation failed: %v", err)
	} else {
		fmt.Println("Doc_A.pdf is valid!")
	}
	
	// Test Doc_B.pdf
	fmt.Println("Validating Doc_B.pdf...")
	err = api.ValidateFile("Doc_B.pdf", conf)
	if err != nil {
		log.Printf("Doc_B.pdf validation failed: %v", err)
	} else {
		fmt.Println("Doc_B.pdf is valid!")
	}
	
	// Test with strict validation
	fmt.Println("\nTesting with strict validation...")
	conf.ValidationMode = model.ValidationStrict
	
	err = api.ValidateFile("Doc_A.pdf", conf)
	if err != nil {
		fmt.Printf("Doc_A.pdf strict validation failed: %v\n", err)
	} else {
		fmt.Println("Doc_A.pdf passes strict validation!")
	}
	
	fmt.Println("Experiment 02 completed!")
}

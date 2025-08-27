package main

import (
	"fmt"
	"log"

	"github.com/pdfcpu/pdfcpu/pkg/api"
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/model"
)

func main() {
	fmt.Println("=== Test 02: PDF Validation API ===")
	
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
	
	fmt.Println("Test 02 completed!")
}

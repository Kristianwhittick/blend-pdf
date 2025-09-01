package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"

	"github.com/pdfcpu/pdfcpu/pkg/api"
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/model"
)

func main() {
	fmt.Println("=== Experiment 22: Complete Zip Flow with Validation ===")

	// Test files (paths from project root)
	file1 := "archive/Doc_A.pdf" // A1, A2, A3
	file2 := "archive/Doc_B.pdf" // M, 9, f
	reversedFile := "output/test_22_reversed.pdf"
	output := "output/test_22_complete.pdf"

	// Verify input files exist
	if _, err := os.Stat(file1); os.IsNotExist(err) {
		log.Fatalf("Input file 1 not found: %s", file1)
	}
	if _, err := os.Stat(file2); os.IsNotExist(err) {
		log.Fatalf("Input file 2 not found: %s", file2)
	}

	conf := model.NewDefaultConfiguration()

	// Step 1: Validate input files
	fmt.Println("Step 1: Validating input files...")

	count1, err := api.PageCountFile(file1)
	if err != nil {
		log.Fatalf("Failed to get page count for file1: %v", err)
	}

	count2, err := api.PageCountFile(file2)
	if err != nil {
		log.Fatalf("Failed to get page count for file2: %v", err)
	}

	fmt.Printf("✓ File 1: %d pages\n", count1)
	fmt.Printf("✓ File 2: %d pages\n", count2)

	if count1 != count2 {
		log.Fatalf("Page count mismatch: %d != %d", count1, count2)
	}
	fmt.Printf("✓ Page counts match: %d pages each\n", count1)

	// Step 2: Read original content for comparison
	fmt.Println("\nStep 2: Reading original content...")

	originalContent1 := readPDFContent(file1)
	originalContent2 := readPDFContent(file2)

	fmt.Printf("Original File 1 content: %s\n", originalContent1)
	fmt.Printf("Original File 2 content: %s\n", originalContent2)

	// Step 3: Reverse file2 using CollectFile
	fmt.Println("\nStep 3: Reversing file2...")

	// Build reverse page selection dynamically
	var pageNums []string
	for i := count2; i >= 1; i-- {
		pageNums = append(pageNums, fmt.Sprintf("%d", i))
	}
	reverseSelection := strings.Join(pageNums, ",")
	fmt.Printf("Reverse selection: %s\n", reverseSelection)

	pageSelection, err := api.ParsePageSelection(reverseSelection)
	if err != nil {
		log.Fatalf("Failed to parse page selection: %v", err)
	}

	err = api.CollectFile(file2, reversedFile, pageSelection, conf)
	if err != nil {
		log.Fatalf("CollectFile failed: %v", err)
	}

	// Verify reversed content
	reversedContent := readPDFContent(reversedFile)
	fmt.Printf("✓ Reversed file content: %s\n", reversedContent)

	// Step 4: Zip merge
	fmt.Println("\nStep 4: Zip merging...")

	err = api.MergeCreateZipFile(file1, reversedFile, output, conf)
	if err != nil {
		log.Fatalf("MergeCreateZipFile failed: %v", err)
	}

	// Step 5: Validate final result
	fmt.Println("\nStep 5: Validating final result...")

	outputCount, err := api.PageCountFile(output)
	if err != nil {
		log.Fatalf("Failed to get output page count: %v", err)
	}

	finalContent := readPDFContent(output)

	fmt.Printf("✓ Final output pages: %d (expected: %d)\n", outputCount, count1+count2)
	fmt.Printf("✓ Final content: %s\n", finalContent)

	// Step 6: Analyze the pattern
	fmt.Println("\nStep 6: Pattern analysis...")

	expectedPattern := buildExpectedPattern(originalContent1, reversedContent)

	fmt.Printf("Expected pattern: %s\n", expectedPattern)
	fmt.Printf("Actual result:    %s\n", finalContent)

	if strings.ReplaceAll(finalContent, " ", "") == strings.ReplaceAll(expectedPattern, " ", "") {
		fmt.Println("✅ SUCCESS: Pattern matches expected interleaved result!")
	} else {
		fmt.Println("❌ FAILURE: Pattern does not match expected result")
	}

	// Summary
	fmt.Println("\n=== SUMMARY ===")
	fmt.Printf("Input 1: %s → %s\n", file1, originalContent1)
	fmt.Printf("Input 2: %s → %s\n", file2, originalContent2)
	fmt.Printf("Reversed: %s\n", reversedContent)
	fmt.Printf("Final: %s\n", finalContent)
	fmt.Printf("Expected: %s\n", expectedPattern)

	// Cleanup
	fmt.Println("\nCleanup:")
	fmt.Printf("rm %s %s\n", reversedFile, output)
}

func readPDFContent(filename string) string {
	cmd := exec.Command("pdftotext", filename, "-")
	output, err := cmd.Output()
	if err != nil {
		return fmt.Sprintf("ERROR: %v", err)
	}

	// Clean up the output - remove whitespace and newlines
	content := strings.ReplaceAll(string(output), "\n", "")
	content = strings.ReplaceAll(content, " ", "")
	content = strings.ReplaceAll(content, "\f", "") // Form feed

	return content
}

func buildExpectedPattern(content1, reversedContent2 string) string {
	// For interleaved pattern: A1, f, A2, 9, A3, M
	// We need to split the content by characters and interleave

	chars1 := strings.Split(content1, "")
	chars2 := strings.Split(reversedContent2, "")

	var result []string
	maxLen := len(chars1)
	if len(chars2) > maxLen {
		maxLen = len(chars2)
	}

	for i := 0; i < maxLen; i++ {
		if i < len(chars1) {
			result = append(result, chars1[i])
		}
		if i < len(chars2) {
			result = append(result, chars2[i])
		}
	}

	return strings.Join(result, "")
}

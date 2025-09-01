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

package ui

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// FallbackUI provides a basic text interface for unsupported terminals
type FallbackUI struct {
	model   Model
	fileOps FileOperations
	scanner *bufio.Scanner
}

// NewFallbackUI creates a new fallback UI instance
func NewFallbackUI(watchDir, archiveDir, outputDir, errorDir, version string, fileOps FileOperations) *FallbackUI {
	model := NewModel(watchDir, archiveDir, outputDir, errorDir, version)

	return &FallbackUI{
		model:   model,
		fileOps: fileOps,
		scanner: bufio.NewScanner(os.Stdin),
	}
}

// Run starts the fallback interface
func (f *FallbackUI) Run() error {
	fmt.Printf("BlendPDFGo v%s - Basic Mode\n", f.model.Version)
	fmt.Println("Terminal UI not supported, using basic interface...")
	fmt.Println("For the full experience, use Windows Terminal, PowerShell 7+, or Linux/macOS terminal.")
	fmt.Println()

	f.displayDirectories()

	for {
		f.updateFiles()
		f.displayStatus()
		f.displayMenu()

		choice := f.getUserChoice()
		if !f.handleChoice(choice) {
			break
		}
	}

	f.displayStatistics()
	return nil
}

func (f *FallbackUI) displayDirectories() {
	fmt.Printf("Watch Directory:   %s\n", f.model.WatchDir)
	fmt.Printf("Archive Directory: %s\n", f.model.ArchiveDir)
	fmt.Printf("Output Directory:  %s\n", f.model.OutputDir)
	fmt.Printf("Error Directory:   %s\n", f.model.ErrorDir)
	fmt.Println()
}

func (f *FallbackUI) updateFiles() {
	// Update file counts and list
	msg := f.updateFilesMsg()
	f.model.MainFiles = msg.mainFiles
	f.model.ArchiveCount = msg.archiveCount
	f.model.OutputCount = msg.outputCount
	f.model.ErrorCount = msg.errorCount
}

func (f *FallbackUI) updateFilesMsg() fileUpdateMsg {
	var mainFiles []FileInfo

	// Get PDF files from watch directory
	if files, err := f.fileOps.FindPDFFiles(f.model.WatchDir); err == nil {
		for _, file := range files {
			size := f.fileOps.GetHumanReadableSize(file)
			mainFiles = append(mainFiles, FileInfo{
				Name: file,
				Size: size,
			})
		}
	}

	return fileUpdateMsg{
		mainFiles:    mainFiles,
		archiveCount: f.fileOps.CountPDFFiles(f.model.ArchiveDir),
		outputCount:  f.fileOps.CountPDFFiles(f.model.OutputDir),
		errorCount:   f.fileOps.CountPDFFiles(f.model.ErrorDir),
	}
}

func (f *FallbackUI) displayStatus() {
	fmt.Printf("Files: Main(%d) Archive(%d) Output(%d) Error(%d) | Session: %s\n",
		len(f.model.MainFiles), f.model.ArchiveCount, f.model.OutputCount,
		f.model.ErrorCount, f.model.ElapsedTime())

	// Show available files if any
	if len(f.model.MainFiles) > 0 {
		fmt.Println("Available PDF files:")
		for i, file := range f.model.MainFiles {
			if i >= 5 { // Limit display to 5 files
				fmt.Printf("  ... and %d more file(s)\n", len(f.model.MainFiles)-5)
				break
			}
			fmt.Printf("  %s (%s)\n", file.Name, file.Size)
		}
	}
	fmt.Println()
}

func (f *FallbackUI) displayMenu() {
	fmt.Println("Options:")
	fmt.Println("  [S] Single File  - Move a single PDF file to output")
	fmt.Println("  [M] Merge PDFs   - Merge two PDF files with interleaved pattern")
	fmt.Println("  [H] Help         - Show help information")
	fmt.Println("  [Q] Quit         - Exit the program")
	fmt.Print("Enter choice (S/M/H/Q): ")
}

func (f *FallbackUI) getUserChoice() string {
	if f.scanner.Scan() {
		return strings.TrimSpace(strings.ToUpper(f.scanner.Text()))
	}
	return "Q" // Default to quit on EOF
}

func (f *FallbackUI) handleChoice(choice string) bool {
	switch choice {
	case "S":
		return f.handleSingleFile()
	case "M":
		return f.handleMergeFiles()
	case "H":
		f.showHelp()
		return true
	case "Q":
		return false
	default:
		fmt.Println("Invalid choice. Please enter S, M, H, or Q.")
		return true
	}
}

func (f *FallbackUI) handleSingleFile() bool {
	if len(f.model.MainFiles) == 0 {
		fmt.Println("❌ No PDF files found")
		return true
	}

	fmt.Printf("Processing single file: %s\n", f.model.MainFiles[0].Name)

	// Call actual file operations
	if _, err := f.fileOps.ProcessSingleFile(); err != nil {
		fmt.Printf("❌ Error: %v\n", err)
		f.model.ErrorCount2++
	} else {
		fmt.Println("✅ Single file processed successfully")
		f.model.SuccessCount++
	}

	fmt.Println()
	return true
}

func (f *FallbackUI) handleMergeFiles() bool {
	if len(f.model.MainFiles) < 2 {
		fmt.Println("❌ Need at least 2 PDF files for merge")
		return true
	}

	fmt.Printf("Merging: %s and %s\n", f.model.MainFiles[0].Name, f.model.MainFiles[1].Name)

	// Call actual merge operations
	if _, err := f.fileOps.ProcessMergeFiles(); err != nil {
		fmt.Printf("❌ Error: %v\n", err)
		f.model.ErrorCount2++
	} else {
		fmt.Printf("✅ Files merged successfully\n")
		f.model.SuccessCount++
	}

	fmt.Println()
	return true
}

func (f *FallbackUI) showHelp() {
	fmt.Printf("\nBlendPDFGo v%s - Help\n", f.model.Version)
	fmt.Println("===================")
	fmt.Println()
	fmt.Println("This tool merges PDF files with automatic page reversal for double-sided scanning.")
	fmt.Println()
	fmt.Println("Operations:")
	fmt.Println("  Single File: Moves the first PDF file to the output directory")
	fmt.Println("  Merge PDFs:  Merges two PDFs with interleaved pattern (A1, B3, A2, B2, A3, B1)")
	fmt.Println()
	fmt.Println("File Processing:")
	fmt.Println("  - Success: Files moved to archive/ directory")
	fmt.Println("  - Error:   Files moved to error/ directory")
	fmt.Println("  - Output:  Merged files placed in output/ directory")
	fmt.Println()
	fmt.Println("For the full terminal UI experience, use:")
	fmt.Println("  - Windows Terminal with PowerShell 7+")
	fmt.Println("  - Linux or macOS terminal")
	fmt.Println()
}

func (f *FallbackUI) displayStatistics() {
	fmt.Println("\nSession Statistics:")
	fmt.Printf("Successful operations: %d\n", f.model.SuccessCount)
	fmt.Printf("Errors encountered: %d\n", f.model.ErrorCount2)
	fmt.Printf("Time elapsed: %s\n", f.model.ElapsedTime())
	fmt.Println()
}

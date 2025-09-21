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

// LegacyUI provides a basic text interface for legacy terminals
type LegacyUI struct {
	watchDir   string
	archiveDir string
	outputDir  string
	errorDir   string
	version    string
	fileOps    *FileOpsBridge
	scanner    *bufio.Scanner
}

// NewLegacyUI creates a new legacy UI instance
func NewLegacyUI(watchDir, archiveDir, outputDir, errorDir, version string, fileOps *FileOpsBridge) *LegacyUI {
	return &LegacyUI{
		watchDir:   watchDir,
		archiveDir: archiveDir,
		outputDir:  outputDir,
		errorDir:   errorDir,
		version:    version,
		fileOps:    fileOps,
		scanner:    bufio.NewScanner(os.Stdin),
	}
}

// Run starts the legacy UI
func (l *LegacyUI) Run() error {
	fmt.Printf("BlendPDFGo %s - PDF Merging Tool\n", l.version)
	fmt.Println("=================================")
	fmt.Printf("Watching: %s\n", l.watchDir)
	fmt.Println()

	for {
		l.showStatus()
		l.showMenu()

		choice := l.getUserChoice()
		if !l.handleChoice(choice) {
			break
		}
	}

	return nil
}

// showStatus displays current file counts
func (l *LegacyUI) showStatus() {
	mainCount := l.fileOps.CountPDFFiles(l.watchDir)
	archiveCount := l.fileOps.CountPDFFiles(l.archiveDir)
	outputCount := l.fileOps.CountPDFFiles(l.outputDir)
	errorCount := l.fileOps.CountPDFFiles(l.errorDir)

	fmt.Printf("Files: Main(%d) Archive(%d) Output(%d) Error(%d)\n",
		mainCount, archiveCount, outputCount, errorCount)
	fmt.Println()
}

// showMenu displays the menu options
func (l *LegacyUI) showMenu() {
	fmt.Println("Available Options:")
	fmt.Println("  [S] Single File  - Move a single PDF file to output")
	fmt.Println("  [M] Merge PDFs   - Merge two PDF files with interleaved pattern")
	fmt.Println("  [U] Undo         - Reverse last operation")
	fmt.Println("  [A] Archive      - Toggle archive mode")
	fmt.Println("  [H] Help         - Show help information")
	fmt.Println("  [Q] Quit         - Exit the program")
	fmt.Println()
}

// getUserChoice gets user input
func (l *LegacyUI) getUserChoice() string {
	fmt.Print("Enter choice (S/M/U/A/H/Q): ")
	if l.scanner.Scan() {
		input := strings.TrimSpace(l.scanner.Text())
		// Handle keyboard shortcuts
		input = l.processKeyboardShortcuts(input)
		return strings.ToUpper(input)
	}
	return "Q"
}

// processKeyboardShortcuts handles enhanced keyboard shortcuts
func (l *LegacyUI) processKeyboardShortcuts(input string) string {
	switch strings.ToLower(input) {
	case "f1", "help", "?":
		return "H" // Help
	case "ctrl+q", "exit", "quit", "bye":
		return "Q" // Quit
	case "ctrl+z", "undo":
		return "U" // Undo
	case "archive", "toggle":
		return "A" // Archive toggle
	case "single", "1":
		return "S" // Single file
	case "merge", "2":
		return "M" // Merge
	default:
		return input
	}
}

// handleChoice processes user choice
func (l *LegacyUI) handleChoice(choice string) bool {
	switch choice {
	case "S":
		l.handleSingleFile()
	case "M":
		l.handleMerge()
	case "U":
		l.handleUndo()
	case "A":
		l.handleArchiveToggle()
	case "H":
		l.showHelp()
	case "Q":
		fmt.Println("Exiting...")
		return false
	default:
		fmt.Println("Invalid choice. Please try again.")
	}

	fmt.Println()
	return true
}

// handleSingleFile processes single file operation
func (l *LegacyUI) handleSingleFile() {
	fmt.Println("Processing single file...")
	_, err := l.fileOps.ProcessSingleFile()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		fmt.Println("Single file operation completed successfully.")
	}
}

// handleMerge processes merge operation
func (l *LegacyUI) handleMerge() {
	fmt.Println("Processing merge...")
	_, err := l.fileOps.ProcessMergeFiles()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		fmt.Println("Merge operation completed successfully.")
	}
}

// handleUndo processes undo operation
func (l *LegacyUI) handleUndo() {
	fmt.Println("Processing undo...")
	err := l.fileOps.ProcessUndo()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		fmt.Println("Undo operation completed successfully.")
	}
}

// handleArchiveToggle toggles archive mode
func (l *LegacyUI) handleArchiveToggle() {
	l.fileOps.ToggleArchiveMode()
	fmt.Println("Archive mode toggled.")
}

// showHelp displays help information
func (l *LegacyUI) showHelp() {
	fmt.Println("BlendPDFGo Help")
	fmt.Println("===============")
	fmt.Println()
	fmt.Println("Operations:")
	fmt.Println("  Single File: Moves the first PDF file to the output directory")
	fmt.Println("  Merge PDFs:  Merges two PDFs with interleaved pattern")
	fmt.Println("  Undo:        Reverses the last operation")
	fmt.Println()
	fmt.Println("Keyboard Shortcuts:")
	fmt.Println("  S, single, 1     - Single file operation")
	fmt.Println("  M, merge, 2      - Merge operation")
	fmt.Println("  U, undo, Ctrl+Z  - Undo last operation")
	fmt.Println("  A, archive       - Toggle archive mode")
	fmt.Println("  H, help, F1, ?   - Show this help")
	fmt.Println("  Q, quit, Ctrl+Q  - Exit program")
	fmt.Println()
	fmt.Println("Archive Mode:")
	fmt.Println("  ON:  Files are copied to archive before processing")
	fmt.Println("  OFF: Files are processed without archiving")
	fmt.Println()
	fmt.Println("Press Enter to continue...")
	l.scanner.Scan()
}

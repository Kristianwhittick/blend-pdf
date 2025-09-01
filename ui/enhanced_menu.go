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
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
	"time"

	"github.com/fsnotify/fsnotify"
)

// EnhancedMenu provides a better-looking menu without complex TUI
type EnhancedMenu struct {
	watchDir        string
	archiveDir      string
	outputDir       string
	errorDir        string
	version         string
	fileOps         FileOperations
	scanner         *bufio.Scanner
	startTime       time.Time
	successCount    int
	errorCount      int
	recentOps       []RecentOperation
	lastFileCount   int
	lastUpdateTime  time.Time
	isProcessing    bool
	currentOp       string
	progressStep    int
	progressTotal   int
	watcher         *fsnotify.Watcher
	needsRefresh    bool
}

// RecentOperation stores detailed operation information
type RecentOperation struct {
	Timestamp   time.Time
	Description string
	Status      string // "SUCCESS" or "FAILED"
	Details     string
}

// NewEnhancedMenu creates an enhanced menu
func NewEnhancedMenu(watchDir, archiveDir, outputDir, errorDir, version string, fileOps FileOperations) *EnhancedMenu {
	menu := &EnhancedMenu{
		watchDir:       watchDir,
		archiveDir:     archiveDir,
		outputDir:      outputDir,
		errorDir:       errorDir,
		version:        version,
		fileOps:        fileOps,
		scanner:        bufio.NewScanner(os.Stdin),
		startTime:      time.Now(),
		recentOps:      make([]RecentOperation, 0, 5),
		lastUpdateTime: time.Now(),
	}
	
	// Initialize file system watcher
	menu.setupFileWatcher()
	
	return menu
}

// setupFileWatcher initializes real-time file system monitoring
func (e *EnhancedMenu) setupFileWatcher() {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		// Fall back to polling if watcher fails
		return
	}
	
	e.watcher = watcher
	
	// Watch the main directory for PDF file changes
	err = e.watcher.Add(e.watchDir)
	if err != nil {
		e.watcher.Close()
		e.watcher = nil
		return
	}
	
	// Start monitoring in background
	go e.monitorFileChanges()
}

// monitorFileChanges handles file system events in background
func (e *EnhancedMenu) monitorFileChanges() {
	if e.watcher == nil {
		return
	}
	
	for {
		select {
		case event, ok := <-e.watcher.Events:
			if !ok {
				return
			}
			
			// Only care about PDF files
			if filepath.Ext(event.Name) == ".pdf" {
				// Mark that we need to refresh the display
				e.needsRefresh = true
			}
			
		case err, ok := <-e.watcher.Errors:
			if !ok {
				return
			}
			// Ignore watcher errors, fall back to polling
			_ = err
		}
	}
}

// Run starts the enhanced menu with real-time monitoring
func (e *EnhancedMenu) Run() error {
	// Ensure watcher is cleaned up on exit
	defer e.cleanup()
	
	e.clearScreen()
	e.showHeader()
	e.showStatus()
	fmt.Print("Enter choice (S/M/H/Q): ")

	for {
		choice := e.getUserChoice()
		
		// Handle invalid choices by continuing the loop
		if choice != "S" && choice != "M" && choice != "H" && choice != "Q" {
			fmt.Println("❌ Invalid choice.")
			fmt.Print("Enter choice (S/M/H/Q): ")
			continue
		}
		
		if !e.handleChoice(choice) {
			break
		}

		// Refresh full interface for valid operations
		e.clearScreen()
		e.showHeader()
		e.showStatus()
		fmt.Print("Enter choice (S/M/H/Q): ")
	}

	e.showStatistics()
	return nil
}

// cleanup closes the file watcher
func (e *EnhancedMenu) cleanup() {
	if e.watcher != nil {
		e.watcher.Close()
	}
}

func (e *EnhancedMenu) clearScreen() {
	if runtime.GOOS == "windows" {
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	} else {
		fmt.Print("\033[2J\033[H")
	}
}

func (e *EnhancedMenu) showHeader() {
	// Get file counts
	var mainFiles []FileInfo
	if files, err := e.fileOps.FindPDFFiles(e.watchDir); err == nil {
		for _, file := range files {
			size := e.fileOps.GetHumanReadableSize(file)
			mainFiles = append(mainFiles, FileInfo{
				Name: file,
				Size: size,
			})
		}
	}

	archiveCount := e.fileOps.CountPDFFiles(e.archiveDir)
	outputCount := e.fileOps.CountPDFFiles(e.outputDir)
	errorCount := e.fileOps.CountPDFFiles(e.errorDir)

	fmt.Println("┌─────────────────────────────────────────────────────────────────────────────┐")
	title := fmt.Sprintf("BlendPDFGo v%s", e.version)
	padding := (77 - len(title)) / 2
	fmt.Printf("│%*s%s%*s│\n", padding, "", title, 77-padding-len(title), "")
	fmt.Println("├─────────────────────────────────────────────────────────────────────────────┤")

	// Format paths with counts on the right
	fmt.Printf("│ Watch  : %-63s %2d │\n", e.watchDir, len(mainFiles))
	fmt.Printf("│ Archive: %-63s %2d │\n", e.archiveDir, archiveCount)
	fmt.Printf("│ Output : %-63s %2d │\n", e.outputDir, outputCount)
	fmt.Printf("│ Error  : %-63s %2d │\n", e.errorDir, errorCount)

	fmt.Println("└─────────────────────────────────────────────────────────────────────────────┘")
	fmt.Println()
}

func (e *EnhancedMenu) showStatus() {
	// Show available files
	var mainFiles []FileInfo
	if files, err := e.fileOps.FindPDFFiles(e.watchDir); err == nil {
		for _, file := range files {
			size := e.fileOps.GetHumanReadableSize(file)
			mainFiles = append(mainFiles, FileInfo{
				Name: file,
				Size: size,
			})
		}
	}

	if len(mainFiles) > 0 {
		fmt.Println("Available PDF files:")
		for i, file := range mainFiles {
			if i >= 5 {
				fmt.Printf("  ... and %d more file(s)\n", len(mainFiles)-5)
				break
			}
			fmt.Printf("  %s (%s)\n", file.Name, file.Size)
		}
	} else {
		fmt.Println("No PDF files found in watch directory")
	}

	// R5B.3 - Horizontal separator line
	fmt.Println("─────────────────────────────────────────────────────────────────────────────")

	// R5B.4 - Enhanced Recent Output section with detailed operation information
	if len(e.recentOps) > 0 {
		fmt.Println("Recent Operations:")
		for _, op := range e.recentOps {
			statusIcon := "✅"
			if op.Status == "FAILED" {
				statusIcon = "❌"
			}
			fmt.Printf("  %s [%s] %s\n", statusIcon, op.Timestamp.Format("15:04:05"), op.Details)
		}
	} else {
		fmt.Println("Recent Operations:")
		fmt.Println("  No operations performed yet")
	}

	// R5B.5 - Actions bar (persistent during operations)
	e.showActionsBar()

	// R5B.6 - Status/Progress section (2 lines: status + progress, progress overwrites status during operations)
	if e.isProcessing {
		e.showProgressBar()
		fmt.Println() // Second line for progress section
	} else {
		e.showStatusLine(len(mainFiles))
		fmt.Println() // Second line for status section (empty when not processing)
	}
	fmt.Println()
}

func (e *EnhancedMenu) showActionsBar() {
	fmt.Println("┌─────────────────────────────────────────────────────────────────────────────┐")
	fmt.Println("│                                 Actions                                     │")
	fmt.Println("├─────────────────────────────────────────────────────────────────────────────┤")
	fmt.Println("│  [S] Single File  - Move a single PDF file to output directory              │")
	fmt.Println("│  [M] Merge PDFs   - Merge two PDF files with interleaved pattern            │")
	fmt.Println("│  [H] Help         - Show help information                                   │")
	fmt.Println("│  [Q] Quit         - Exit the program                                        │")
	fmt.Println("└─────────────────────────────────────────────────────────────────────────────┘")
}

func (e *EnhancedMenu) getUserChoice() string {
	// Create a channel to receive input
	inputChan := make(chan string, 1)
	
	// Start goroutine to read input
	go func() {
		if e.scanner.Scan() {
			inputChan <- strings.TrimSpace(strings.ToUpper(e.scanner.Text()))
		} else {
			inputChan <- "Q"
		}
	}()
	
	// Check for input or refresh needs every 100ms
	ticker := time.NewTicker(100 * time.Millisecond)
	defer ticker.Stop()
	
	for {
		select {
		case input := <-inputChan:
			return input
		case <-ticker.C:
			if e.needsRefresh {
				e.needsRefresh = false
				e.clearScreen()
				e.showHeader()
				e.showStatus()
				fmt.Print("Enter choice (S/M/H/Q): ")
			}
		}
	}
}

func (e *EnhancedMenu) handleChoice(choice string) bool {
	fmt.Println()

	switch choice {
	case "S":
		return e.handleSingleFile()
	case "M":
		return e.handleMergeFiles()
	case "H":
		e.showHelp()
		return true
	case "Q":
		return false
	}
	return true
}

func (e *EnhancedMenu) showProgressBar() {
	// R5.8 - Progress bar replaces status line during operations
	elapsed := time.Since(e.lastUpdateTime)
	
	// Create animated progress bar
	barWidth := 40
	progress := float64(e.progressStep) / float64(e.progressTotal)
	if e.progressTotal == 0 {
		// Indeterminate progress - use time-based animation
		progress = float64(int(elapsed.Seconds())%barWidth) / float64(barWidth)
	}
	
	filled := int(progress * float64(barWidth))
	
	fmt.Printf("Processing: %s [", e.currentOp)
	for i := 0; i < barWidth; i++ {
		if i < filled {
			fmt.Print("█")
		} else if i == filled && e.progressTotal == 0 {
			// Animated cursor for indeterminate progress
			fmt.Print("▶")
		} else {
			fmt.Print("░")
		}
	}
	
	if e.progressTotal > 0 {
		fmt.Printf("] %d/%d", e.progressStep, e.progressTotal)
	} else {
		fmt.Printf("] %.1fs", elapsed.Seconds())
	}
}

func (e *EnhancedMenu) showStatusLine(fileCount int) {
	// R5.7 - Status line with operation counts
	fmt.Printf("Status: Operations: %d | Errors: %d | Files monitored: %d", 
		e.successCount, e.errorCount, fileCount)
}

func (e *EnhancedMenu) setProcessing(operation string) {
	e.isProcessing = true
	e.currentOp = operation
	e.progressStep = 0
	e.progressTotal = 0
	e.lastUpdateTime = time.Now()
}

func (e *EnhancedMenu) setProgressStep(step, total int) {
	e.progressStep = step
	e.progressTotal = total
}

func (e *EnhancedMenu) clearProcessing() {
	e.isProcessing = false
	e.currentOp = ""
}

func (e *EnhancedMenu) addRecentOperation(description, status, details string) {
	operation := RecentOperation{
		Timestamp:   time.Now(),
		Description: description,
		Status:      status,
		Details:     details,
	}
	
	e.recentOps = append(e.recentOps, operation)
	if len(e.recentOps) > 5 {
		e.recentOps = e.recentOps[1:]
	}
}

func (e *EnhancedMenu) handleSingleFile() bool {
	e.setProcessing("Single file processing")
	
	// Show progress during operation
	e.clearScreen()
	e.showHeader()
	e.showStatus()
	
	// Simulate progress steps
	e.setProgressStep(1, 3)
	time.Sleep(200 * time.Millisecond) // Brief pause to show progress
	
	e.clearScreen()
	e.showHeader()
	e.setProgressStep(2, 3)
	e.showStatus()
	
	if description, err := e.fileOps.ProcessSingleFile(); err != nil {
		fmt.Printf("❌ Error: %v\n", err)
		e.errorCount++
		e.addRecentOperation("Single file processing", "FAILED", err.Error())
	} else {
		e.setProgressStep(3, 3)
		e.clearScreen()
		e.showHeader()
		e.showStatus()
		
		fmt.Println("✅ Single file processed successfully")
		e.successCount++
		e.addRecentOperation("Single file processing", "SUCCESS", description)
	}

	e.clearProcessing()
	return true
}

func (e *EnhancedMenu) handleMergeFiles() bool {
	e.setProcessing("Merge operation")
	
	// Show progress during operation
	e.clearScreen()
	e.showHeader()
	e.showStatus()
	
	// Simulate progress steps for merge operation
	e.setProgressStep(1, 4)
	time.Sleep(200 * time.Millisecond)
	
	e.clearScreen()
	e.showHeader()
	e.setProgressStep(2, 4)
	e.showStatus()
	time.Sleep(200 * time.Millisecond)
	
	e.clearScreen()
	e.showHeader()
	e.setProgressStep(3, 4)
	e.showStatus()

	if description, err := e.fileOps.ProcessMergeFiles(); err != nil {
		fmt.Printf("❌ Error: %v\n", err)
		e.errorCount++
		e.addRecentOperation("Merge operation", "FAILED", err.Error())
	} else {
		e.setProgressStep(4, 4)
		e.clearScreen()
		e.showHeader()
		e.showStatus()
		
		fmt.Println("✅ Files merged successfully")
		e.successCount++
		e.addRecentOperation("Merge operation", "SUCCESS", description)
	}

	e.clearProcessing()
	return true
}

func (e *EnhancedMenu) showHelp() {
	fmt.Printf("BlendPDFGo v%s - Help\n", e.version)
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
}

func (e *EnhancedMenu) showStatistics() {
	fmt.Println()
	fmt.Println("┌─────────────────────────────────────────────────────────────────────────────┐")
	fmt.Println("│                            Session Statistics                               │")
	fmt.Println("├─────────────────────────────────────────────────────────────────────────────┤")
	fmt.Printf("│ Successful operations: %52d │\n", e.successCount)
	fmt.Printf("│ Errors encountered: %55d │\n", e.errorCount)
	fmt.Printf("│ Time elapsed: %61s │\n", e.elapsedTime())
	fmt.Println("└─────────────────────────────────────────────────────────────────────────────┘")
}

func (e *EnhancedMenu) elapsedTime() string {
	elapsed := time.Since(e.startTime)
	if elapsed < time.Minute {
		return fmt.Sprintf("%ds", int(elapsed.Seconds()))
	}
	return fmt.Sprintf("%dm %ds", int(elapsed.Minutes()), int(elapsed.Seconds())%60)
}

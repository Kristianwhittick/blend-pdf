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
	"fmt"
	"os"
	"runtime"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
)

// FileOperations interface for dependency injection
type FileOperations interface {
	FindPDFFiles(dir string) ([]string, error)
	CountPDFFiles(dir string) int
	GetHumanReadableSize(filename string) string
	ProcessSingleFile() (string, error) // Returns operation description
	ProcessMergeFiles() (string, error) // Returns operation description
}

// TUI represents the terminal user interface
type TUI struct {
	model   Model
	program *tea.Program
	fileOps FileOperations
}

// NewTUI creates a new TUI instance
func NewTUI(watchDir, archiveDir, outputDir, errorDir, version string, fileOps FileOperations) *TUI {
	model := NewModel(watchDir, archiveDir, outputDir, errorDir, version)

	return &TUI{
		model:   model,
		fileOps: fileOps,
	}
}

// Run starts the TUI
func (t *TUI) Run() error {
	// Check if we should use TUI or fallback
	if !t.supportsTUI() {
		return t.runFallback()
	}

	// Create a wrapper model that has access to the bridge
	wrapper := &tuiWrapper{
		Model:   t.model,
		fileOps: t.fileOps,
	}

	// Create Bubble Tea program
	t.program = tea.NewProgram(
		wrapper,
		tea.WithAltScreen(),
		tea.WithMouseCellMotion(),
	)

	// Run the program
	_, err := t.program.Run()
	return err
}

// tuiWrapper wraps the Model with access to file operations
type tuiWrapper struct {
	Model
	fileOps FileOperations
}

// Init initializes the wrapper
func (w *tuiWrapper) Init() tea.Cmd {
	return tea.Batch(tickCmd(), w.updateFilesCmd())
}

// Update handles messages with access to file operations
func (w *tuiWrapper) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		w.Width = msg.Width
		w.Height = msg.Height
		return w, nil

	case tea.KeyMsg:
		return w.handleKeyPress(msg)

	case tickMsg:
		// Update files every second
		return w, tea.Batch(tickCmd(), w.updateFilesCmd())

	case fileUpdateMsg:
		w.MainFiles = msg.mainFiles
		w.ArchiveCount = msg.archiveCount
		w.OutputCount = msg.outputCount
		w.ErrorCount = msg.errorCount
		return w, nil

	case operationCompleteMsg:
		w.Processing = false
		if msg.success {
			w.SuccessCount++
		} else {
			w.ErrorCount2++
		}
		w.AddRecentOp(msg.message)
		return w, nil
	}

	return w, nil
}

// updateFilesCmd creates a command to update files using the bridge
func (w *tuiWrapper) updateFilesCmd() tea.Cmd {
	return func() tea.Msg {
		var mainFiles []FileInfo

		// Get PDF files from watch directory
		if files, err := w.fileOps.FindPDFFiles(w.WatchDir); err == nil {
			for _, file := range files {
				size := w.fileOps.GetHumanReadableSize(file)
				mainFiles = append(mainFiles, FileInfo{
					Name: file,
					Size: size,
				})
			}
		}

		return fileUpdateMsg{
			mainFiles:    mainFiles,
			archiveCount: w.fileOps.CountPDFFiles(w.ArchiveDir),
			outputCount:  w.fileOps.CountPDFFiles(w.OutputDir),
			errorCount:   w.fileOps.CountPDFFiles(w.ErrorDir),
		}
	}
}

func (w *tuiWrapper) handleKeyPress(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	switch msg.String() {
	case "ctrl+c", "q", "Q":
		w.Quitting = true
		return w, tea.Quit

	case "s", "S":
		return w.handleSingleFile()

	case "m", "M":
		return w.handleMergeFiles()

	case "t", "T":
		return w.handleToggleMode()

	case "up", "k":
		if w.SelectionMode == UserSelectMode && w.Cursor > 0 {
			w.Cursor--
		}
		return w, nil

	case "down", "j":
		if w.SelectionMode == UserSelectMode && w.Cursor < len(w.MainFiles)-1 {
			w.Cursor++
		}
		return w, nil

	case " ":
		if w.SelectionMode == UserSelectMode && w.Cursor < len(w.MainFiles) {
			w.MainFiles[w.Cursor].Selected = !w.MainFiles[w.Cursor].Selected
		}
		return w, nil
	}

	return w, nil
}

func (w *tuiWrapper) handleSingleFile() (tea.Model, tea.Cmd) {
	if len(w.MainFiles) == 0 {
		w.AddRecentOp("❌ No PDF files found")
		return w, nil
	}

	w.Processing = true
	w.ProgressMsg = "Processing single file..."

	return w, func() tea.Msg {
		_, err := w.fileOps.ProcessSingleFile()
		if err != nil {
			return operationCompleteMsg{
				success: false,
				message: fmt.Sprintf("❌ Error: %v", err),
			}
		}
		return operationCompleteMsg{
			success: true,
			message: "✅ Single file processed successfully",
		}
	}
}

func (w *tuiWrapper) handleMergeFiles() (tea.Model, tea.Cmd) {
	selectedFiles := w.SelectedFiles()
	if len(selectedFiles) < 2 {
		w.AddRecentOp("❌ Need at least 2 files for merge")
		return w, nil
	}

	w.Processing = true
	w.ProgressMsg = fmt.Sprintf("Merging %s and %s...", selectedFiles[0].Name, selectedFiles[1].Name)

	return w, func() tea.Msg {
		_, err := w.fileOps.ProcessMergeFiles()
		if err != nil {
			return operationCompleteMsg{
				success: false,
				message: fmt.Sprintf("❌ Error: %v", err),
			}
		}
		return operationCompleteMsg{
			success: true,
			message: fmt.Sprintf("✅ Merged %s and %s", selectedFiles[0].Name, selectedFiles[1].Name),
		}
	}
}

func (w *tuiWrapper) handleToggleMode() (tea.Model, tea.Cmd) {
	if w.SelectionMode == AlphaOrderMode {
		w.SelectionMode = UserSelectMode
		w.Cursor = 0
		// Clear auto-selections
		for i := range w.MainFiles {
			w.MainFiles[i].Selected = false
		}
	} else {
		w.SelectionMode = AlphaOrderMode
		// Clear manual selections
		for i := range w.MainFiles {
			w.MainFiles[i].Selected = false
		}
	}

	w.AddRecentOp(fmt.Sprintf("🔄 Switched to %s mode", w.SelectionMode.String()))
	return w, nil
}

// supportsTUI checks if the terminal supports TUI
func (t *TUI) supportsTUI() bool {
	// Check for basic terminal capabilities
	if os.Getenv("TERM") == "" {
		return false
	}

	// Check for known problematic terminals on Windows
	if runtime.GOOS == "windows" {
		// PowerShell 5 and CMD have limited TUI support
		if strings.Contains(os.Getenv("PSModulePath"), "WindowsPowerShell") {
			return false
		}
	}

	// Check if we're in a CI environment
	if os.Getenv("CI") != "" || os.Getenv("GITHUB_ACTIONS") != "" {
		return false
	}

	return true
}

// runFallback runs the fallback interface for unsupported terminals
func (t *TUI) runFallback() error {
	fallback := NewFallbackUI(t.model.WatchDir, t.model.ArchiveDir, t.model.OutputDir, t.model.ErrorDir, t.model.Version, t.fileOps)
	return fallback.Run()
}

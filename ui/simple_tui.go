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
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

// SimpleTUI provides a basic full-screen interface
type SimpleTUI struct {
	watchDir   string
	archiveDir string
	outputDir  string
	errorDir   string
	version    string
	fileOps    FileOperations

	// State
	mainFiles    []FileInfo
	archiveCount int
	outputCount  int
	errorCount   int
	startTime    time.Time
	successCount int
	errorCount2  int

	// UI state
	width    int
	height   int
	quitting bool
}

// NewSimpleTUI creates a simple TUI
func NewSimpleTUI(watchDir, archiveDir, outputDir, errorDir, version string, fileOps FileOperations) *SimpleTUI {
	return &SimpleTUI{
		watchDir:   watchDir,
		archiveDir: archiveDir,
		outputDir:  outputDir,
		errorDir:   errorDir,
		version:    version,
		fileOps:    fileOps,
		startTime:  time.Now(),
	}
}

// Run starts the simple TUI
func (s *SimpleTUI) Run() error {
	// Check if terminal supports TUI
	if !s.supportsTUI() {
		return fmt.Errorf("TUI not supported")
	}

	p := tea.NewProgram(s, tea.WithAltScreen())
	_, err := p.Run()
	return err
}

// Init initializes the model
func (s *SimpleTUI) Init() tea.Cmd {
	s.updateFiles()
	return nil
}

// Update handles messages
func (s *SimpleTUI) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		s.width = msg.Width
		s.height = msg.Height
		return s, nil

	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q", "Q":
			s.quitting = true
			return s, tea.Quit

		case "s", "S":
			s.handleSingleFile()
			s.updateFiles()
			return s, nil

		case "m", "M":
			s.handleMergeFiles()
			s.updateFiles()
			return s, nil

		case "r", "R":
			s.updateFiles()
			return s, nil
		}
	}

	return s, nil
}

// View renders the interface
func (s *SimpleTUI) View() string {
	if s.quitting {
		return "Goodbye!\n"
	}

	// Header with version in border
	headerStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color("#00ADD8")).
		Bold(true).
		Align(lipgloss.Center)

	versionStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color("#0080FF")).
		Bold(true)

	borderStyle := lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color("#888888")).
		Padding(1, 2)

	// Build content
	title := fmt.Sprintf("BlendPDFGo %s", versionStyle.Render("v"+s.version))

	dirs := fmt.Sprintf("Watch: %s | Archive: %s | Output: %s | Error: %s",
		s.watchDir, s.archiveDir, s.outputDir, s.errorDir)

	fileCounts := fmt.Sprintf("Files: Main(%d) Archive(%d) Output(%d) Error(%d) | Session: %s",
		len(s.mainFiles), s.archiveCount, s.outputCount, s.errorCount, s.elapsedTime())

	// Available files
	filesSection := "Available PDF files:\n"
	if len(s.mainFiles) == 0 {
		filesSection += "  No PDF files found"
	} else {
		for i, file := range s.mainFiles {
			if i >= 10 {
				filesSection += fmt.Sprintf("  ... and %d more file(s)", len(s.mainFiles)-10)
				break
			}
			filesSection += fmt.Sprintf("  %s (%s)\n", file.Name, file.Size)
		}
	}

	// Actions
	actions := "Actions: [S]ingle File  [M]erge PDFs  [R]efresh  [Q]uit"

	// Status
	status := fmt.Sprintf("Operations: Success(%d) Errors(%d) | Ready for input",
		s.successCount, s.errorCount2)

	content := lipgloss.JoinVertical(
		lipgloss.Left,
		headerStyle.Render(title),
		"",
		dirs,
		"",
		fileCounts,
		"",
		filesSection,
		"",
		actions,
		"",
		status,
	)

	return borderStyle.Width(s.width - 4).Render(content)
}

// Helper methods
func (s *SimpleTUI) updateFiles() {
	var mainFiles []FileInfo

	if files, err := s.fileOps.FindPDFFiles(s.watchDir); err == nil {
		for _, file := range files {
			size := s.fileOps.GetHumanReadableSize(file)
			mainFiles = append(mainFiles, FileInfo{
				Name: file,
				Size: size,
			})
		}
	}

	s.mainFiles = mainFiles
	s.archiveCount = s.fileOps.CountPDFFiles(s.archiveDir)
	s.outputCount = s.fileOps.CountPDFFiles(s.outputDir)
	s.errorCount = s.fileOps.CountPDFFiles(s.errorDir)
}

func (s *SimpleTUI) handleSingleFile() {
	if len(s.mainFiles) == 0 {
		return
	}

	if _, err := s.fileOps.ProcessSingleFile(); err != nil {
		s.errorCount2++
	} else {
		s.successCount++
	}
}

func (s *SimpleTUI) handleMergeFiles() {
	if len(s.mainFiles) < 2 {
		return
	}

	if _, err := s.fileOps.ProcessMergeFiles(); err != nil {
		s.errorCount2++
	} else {
		s.successCount++
	}
}

func (s *SimpleTUI) elapsedTime() string {
	elapsed := time.Since(s.startTime)
	if elapsed < time.Minute {
		return fmt.Sprintf("%ds", int(elapsed.Seconds()))
	}
	return fmt.Sprintf("%dm %ds", int(elapsed.Minutes()), int(elapsed.Seconds())%60)
}

func (s *SimpleTUI) supportsTUI() bool {
	if os.Getenv("TERM") == "" {
		return false
	}

	if runtime.GOOS == "windows" {
		if strings.Contains(os.Getenv("PSModulePath"), "WindowsPowerShell") {
			return false
		}
	}

	if os.Getenv("CI") != "" || os.Getenv("GITHUB_ACTIONS") != "" {
		return false
	}

	return true
}

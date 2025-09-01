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
	"time"

	tea "github.com/charmbracelet/bubbletea"
)

// FileSelectionMode represents the file selection mode
type FileSelectionMode int

const (
	AlphaOrderMode FileSelectionMode = iota
	UserSelectMode
)

func (m FileSelectionMode) String() string {
	switch m {
	case AlphaOrderMode:
		return "Alpha Order"
	case UserSelectMode:
		return "User Select"
	default:
		return "Unknown"
	}
}

// FileInfo represents a PDF file with metadata
type FileInfo struct {
	Name     string
	Size     string
	Selected bool
}

// Model represents the main TUI model
type Model struct {
	// Application state
	WatchDir   string
	ArchiveDir string
	OutputDir  string
	ErrorDir   string
	Version    string

	// File state
	MainFiles    []FileInfo
	ArchiveCount int
	OutputCount  int
	ErrorCount   int

	// UI state
	SelectionMode FileSelectionMode
	Cursor        int
	VerboseMode   bool
	DebugMode     bool

	// Session state
	SuccessCount int
	ErrorCount2  int // Renamed to avoid conflict
	StartTime    time.Time

	// Recent operations
	RecentOps []string

	// Terminal dimensions
	Width  int
	Height int

	// Operation state
	Processing  bool
	ProgressMsg string

	// Quit flag
	Quitting bool
}

// NewModel creates a new TUI model
func NewModel(watchDir, archiveDir, outputDir, errorDir, version string) Model {
	return Model{
		WatchDir:      watchDir,
		ArchiveDir:    archiveDir,
		OutputDir:     outputDir,
		ErrorDir:      errorDir,
		Version:       version,
		SelectionMode: AlphaOrderMode,
		StartTime:     time.Now(),
		RecentOps:     make([]string, 0, 5),
		MainFiles:     make([]FileInfo, 0),
	}
}

// Messages for Bubble Tea
type tickMsg time.Time
type fileUpdateMsg struct {
	mainFiles    []FileInfo
	archiveCount int
	outputCount  int
	errorCount   int
}

type operationCompleteMsg struct {
	success bool
	message string
}

// Commands
func tickCmd() tea.Cmd {
	return tea.Tick(time.Second, func(t time.Time) tea.Msg {
		return tickMsg(t)
	})
}

func (m Model) Init() tea.Cmd {
	return tea.Batch(tickCmd(), m.updateFilesCmd())
}

func (m Model) updateFilesCmd() tea.Cmd {
	return func() tea.Msg {
		// This will be called by the TUI to update files
		// The actual implementation will be in tui.go using the bridge
		return fileUpdateMsg{
			mainFiles:    []FileInfo{}, // Will be populated by TUI
			archiveCount: 0,
			outputCount:  0,
			errorCount:   0,
		}
	}
}

// Helper methods
func (m *Model) AddRecentOp(op string) {
	m.RecentOps = append([]string{op}, m.RecentOps...)
	if len(m.RecentOps) > 5 {
		m.RecentOps = m.RecentOps[:5]
	}
}

func (m Model) ElapsedTime() string {
	elapsed := time.Since(m.StartTime)
	if elapsed < time.Minute {
		return fmt.Sprintf("%ds", int(elapsed.Seconds()))
	}
	return fmt.Sprintf("%dm %ds", int(elapsed.Minutes()), int(elapsed.Seconds())%60)
}

func (m Model) SelectedFiles() []FileInfo {
	if m.SelectionMode == AlphaOrderMode {
		// Return first two files for alpha order mode
		if len(m.MainFiles) >= 2 {
			return m.MainFiles[:2]
		}
		return m.MainFiles
	}

	// Return selected files for user select mode
	var selected []FileInfo
	for _, file := range m.MainFiles {
		if file.Selected {
			selected = append(selected, file)
		}
	}
	return selected
}

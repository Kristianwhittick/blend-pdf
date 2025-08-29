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
	"github.com/charmbracelet/lipgloss"
)

// MinimalTUI provides a basic working interface
type MinimalTUI struct {
	version  string
	watchDir string
	quitting bool
	width    int
	height   int
}

// NewMinimalTUI creates a minimal TUI
func NewMinimalTUI(watchDir, archiveDir, outputDir, errorDir, version string, fileOps FileOperations) *MinimalTUI {
	return &MinimalTUI{
		version:  version,
		watchDir: watchDir,
	}
}

// Run starts the minimal TUI
func (m *MinimalTUI) Run() error {
	// Check if terminal supports TUI
	if !m.supportsTUI() {
		return fmt.Errorf("TUI not supported")
	}
	
	p := tea.NewProgram(m, tea.WithAltScreen())
	_, err := p.Run()
	return err
}

// Init initializes the model
func (m *MinimalTUI) Init() tea.Cmd {
	return nil
}

// Update handles messages
func (m *MinimalTUI) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
		return m, nil
		
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q", "Q":
			m.quitting = true
			return m, tea.Quit
		}
	}
	
	return m, nil
}

// View renders the interface
func (m *MinimalTUI) View() string {
	if m.quitting {
		return "Goodbye!\n"
	}
	
	// Simple bordered interface
	borderStyle := lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color("#888888")).
		Padding(1, 2)
	
	headerStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color("#00ADD8")).
		Bold(true).
		Align(lipgloss.Center)
	
	versionStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color("#0080FF")).
		Bold(true)
	
	title := fmt.Sprintf("BlendPDFGo %s", versionStyle.Render("v"+m.version))
	
	content := lipgloss.JoinVertical(
		lipgloss.Left,
		headerStyle.Render(title),
		"",
		fmt.Sprintf("Watch Directory: %s", m.watchDir),
		"",
		"This is a test of the new TUI interface.",
		"",
		"Press 'q' or Ctrl+C to quit",
	)
	
	if m.width > 0 {
		return borderStyle.Width(m.width-4).Render(content)
	}
	
	return borderStyle.Render(content)
}

func (m *MinimalTUI) supportsTUI() bool {
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

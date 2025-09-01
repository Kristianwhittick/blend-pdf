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
	"strings"

	"github.com/charmbracelet/lipgloss"
)

// Styles
var (
	// Colors
	primaryColor = lipgloss.Color("#00ADD8") // Go blue
	successColor = lipgloss.Color("#00FF00") // Green
	errorColor   = lipgloss.Color("#FF0000") // Red
	warningColor = lipgloss.Color("#FFFF00") // Yellow
	infoColor    = lipgloss.Color("#0080FF") // Blue
	borderColor  = lipgloss.Color("#888888") // Gray

	// Base styles
	baseStyle = lipgloss.NewStyle().
			Padding(0, 1)

	// Header styles
	headerStyle = lipgloss.NewStyle().
			Foreground(primaryColor).
			Bold(true).
			Align(lipgloss.Center)

	versionStyle = lipgloss.NewStyle().
			Foreground(infoColor).
			Bold(true)

	// Border styles
	borderStyle = lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(borderColor)

	// File count styles
	fileCountStyle = lipgloss.NewStyle().
			Foreground(infoColor).
			Bold(true)

	// File list styles
	fileItemStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#FFFFFF"))

	selectedFileStyle = lipgloss.NewStyle().
				Foreground(successColor).
				Bold(true)

	cursorStyle = lipgloss.NewStyle().
			Foreground(primaryColor).
			Bold(true)

	// Status styles
	successStyle = lipgloss.NewStyle().
			Foreground(successColor).
			Bold(true)

	errorStyle = lipgloss.NewStyle().
			Foreground(errorColor).
			Bold(true)

	warningStyle = lipgloss.NewStyle().
			Foreground(warningColor).
			Bold(true)

	// Action styles
	actionStyle = lipgloss.NewStyle().
			Foreground(warningColor).
			Bold(true)
)

// View renders the main TUI view
func (m Model) View() string {
	if m.Quitting {
		return "Goodbye!\n"
	}

	// Calculate available space
	contentHeight := m.Height - 2 // Account for borders

	// Build the UI sections
	header := m.renderHeader()
	fileCounts := m.renderFileCounts()
	availableFiles := m.renderAvailableFiles()
	recentOutput := m.renderRecentOutput()
	actions := m.renderActions()
	status := m.renderStatus()

	// Combine sections
	content := lipgloss.JoinVertical(
		lipgloss.Left,
		header,
		"",
		fileCounts,
		"",
		availableFiles,
		"",
		recentOutput,
		"",
		actions,
		"",
		status,
	)

	// Apply border with version in top border
	return borderStyle.
		Width(m.Width - 2).
		Height(contentHeight).
		Render(content)
}

func (m Model) renderHeader() string {
	title := "BlendPDFGo"
	version := fmt.Sprintf("v%s", m.Version)

	// Create header with version in border
	headerText := fmt.Sprintf("%s %s", title, versionStyle.Render(version))

	dirs := fmt.Sprintf("Watch: %s | Archive: %s | Output: %s | Error: %s",
		m.WatchDir, m.ArchiveDir, m.OutputDir, m.ErrorDir)

	return lipgloss.JoinVertical(
		lipgloss.Center,
		headerStyle.Render(headerText),
		lipgloss.NewStyle().Foreground(infoColor).Render(dirs),
	)
}

func (m Model) renderFileCounts() string {
	counts := fmt.Sprintf("Files: Main(%d) Archive(%d) Output(%d) Error(%d) | Session Timer: %s",
		len(m.MainFiles), m.ArchiveCount, m.OutputCount, m.ErrorCount, m.ElapsedTime())

	return fileCountStyle.Render(counts)
}

func (m Model) renderAvailableFiles() string {
	if len(m.MainFiles) == 0 {
		return warningStyle.Render("No PDF files found")
	}

	// Mode indicator
	modeIndicator := fmt.Sprintf("[Mode: %s ▼]", m.SelectionMode.String())
	header := fmt.Sprintf("Available PDFs: %s", actionStyle.Render(modeIndicator))

	// File list
	var files []string
	for i, file := range m.MainFiles {
		prefix := "  "
		style := fileItemStyle

		// Show cursor in user select mode
		if m.SelectionMode == UserSelectMode && i == m.Cursor {
			prefix = cursorStyle.Render("▶ ")
			style = cursorStyle
		}

		// Show selection marker
		if file.Selected || (m.SelectionMode == AlphaOrderMode && i < 2) {
			prefix += successStyle.Render("✓ ")
		} else {
			prefix += "  "
		}

		fileText := fmt.Sprintf("%s (%s)", file.Name, file.Size)
		files = append(files, prefix+style.Render(fileText))
	}

	// Limit to 10 files max
	if len(files) > 10 {
		files = files[:10]
		files = append(files, fmt.Sprintf("  ... and %d more file(s)", len(m.MainFiles)-10))
	}

	return lipgloss.JoinVertical(
		lipgloss.Left,
		header,
		strings.Join(files, "\n"),
	)
}

func (m Model) renderRecentOutput() string {
	if len(m.RecentOps) == 0 {
		return "Recent Output:\n  No operations yet"
	}

	header := "Recent Output:"
	var ops []string
	for _, op := range m.RecentOps {
		ops = append(ops, "  "+op)
	}

	return lipgloss.JoinVertical(
		lipgloss.Left,
		header,
		strings.Join(ops, "\n"),
	)
}

func (m Model) renderActions() string {
	actions := []string{
		actionStyle.Render("[S]") + "ingle File",
		actionStyle.Render("[M]") + "erge PDFs",
		actionStyle.Render("[T]") + "oggle Mode",
		actionStyle.Render("[Q]") + "uit",
	}

	if m.SelectionMode == UserSelectMode {
		actions = append([]string{
			actionStyle.Render("[↑↓]") + "Navigate",
			actionStyle.Render("[Space]") + "Select",
		}, actions...)
	}

	return "Actions: " + strings.Join(actions, " ")
}

func (m Model) renderStatus() string {
	if m.Processing {
		return warningStyle.Render("Status: ") + m.ProgressMsg
	}

	status := fmt.Sprintf("Ready | Operations: %s%d%s | Errors: %s%d%s",
		successStyle.Render(""), m.SuccessCount, successStyle.Render(""),
		errorStyle.Render(""), m.ErrorCount2, errorStyle.Render(""))

	return status
}

// Progress bar for operations
func (m Model) renderProgressBar(progress float64) string {
	width := 40
	filled := int(progress * float64(width))

	bar := strings.Repeat("█", filled) + strings.Repeat("░", width-filled)
	return fmt.Sprintf("[%s] %.0f%%", bar, progress*100)
}

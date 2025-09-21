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
	"os"
	"runtime"
	"strings"
)

// TerminalCapabilities represents what the terminal can do
type TerminalCapabilities struct {
	SupportsColor   bool
	SupportsBorders bool
	SupportsUTF8    bool
	IsLegacy        bool
	Name            string
}

// DetectTerminalCapabilities detects terminal capabilities for compatibility
func DetectTerminalCapabilities() *TerminalCapabilities {
	caps := &TerminalCapabilities{
		SupportsColor:   true,
		SupportsBorders: true,
		SupportsUTF8:    true,
		IsLegacy:        false,
		Name:            "modern",
	}

	// Windows-specific detection
	if runtime.GOOS == "windows" {
		// Check for PowerShell version
		psVersion := os.Getenv("PSVersionTable")
		psModulePath := os.Getenv("PSModulePath")

		// Check for CMD
		comSpec := os.Getenv("ComSpec")

		// Check for Windows Terminal
		wtSession := os.Getenv("WT_SESSION")

		// PowerShell 5 or CMD detection
		if (psModulePath != "" && psVersion == "") || strings.Contains(comSpec, "cmd.exe") {
			if wtSession == "" { // Not Windows Terminal
				caps.IsLegacy = true
				caps.SupportsBorders = false
				caps.SupportsUTF8 = false
				caps.Name = "legacy-windows"
			}
		}
	}

	// Check TERM environment variable
	term := os.Getenv("TERM")
	if term == "dumb" || term == "" {
		caps.IsLegacy = true
		caps.SupportsColor = false
		caps.SupportsBorders = false
		caps.Name = "basic"
	}

	return caps
}

// ShouldUseFallbackUI determines if we should use the fallback UI
func ShouldUseFallbackUI() bool {
	caps := DetectTerminalCapabilities()
	return caps.IsLegacy || !caps.SupportsBorders
}

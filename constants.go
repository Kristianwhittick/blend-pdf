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

package main

import (
	"log"
	"os"
	"time"
)

// Application constants
const (
	VERSION = "1.3.0"
)

// ANSI color codes for terminal output
const (
	RED    = "\033[0;31m"
	GREEN  = "\033[0;32m"
	YELLOW = "\033[0;33m"
	BLUE   = "\033[0;34m"
	NC     = "\033[0m" // No Color
)

// Log level constants
const (
	LOG_DEBUG = 0
	LOG_INFO  = 1
	LOG_WARN  = 2
	LOG_ERROR = 3
)

// Application state variables
var (
	// Mode flags
	VERBOSE  = false
	DEBUG    = false
	CONTINUE = true

	// Directory paths
	FOLDER    = ""
	ARCHIVE   = ""
	OUTPUT    = ""
	ERROR_DIR = ""
	LOCKFILE  = ""
)

// Session tracking variables
var (
	COUNTER     = 0
	ERROR_COUNT = 0
	START_TIME  = time.Now()
)

// Operation tracking for undo functionality
type LastOperation struct {
	Type          string   // "single" or "merge"
	OriginalFiles []string // Original file paths in main/
	ActualFiles   []string // Actual filenames used (with conflict resolution)
	OutputFolders []string // Output folders used
	ArchiveFiles  []string // Files in archive/ (for merge operations)
	Timestamp     time.Time
}

var LAST_OPERATION *LastOperation

// Structured logging instances
var (
	debugLogger *log.Logger
	infoLogger  *log.Logger
	warnLogger  *log.Logger
	errorLogger *log.Logger
)

// Initialize structured loggers for debug mode
func initLoggers() {
	debugLogger = createLogger(os.Stdout, "[DEBUG] ", log.Ldate|log.Ltime|log.Lshortfile)
	infoLogger = createLogger(os.Stdout, "[INFO] ", log.Ldate|log.Ltime)
	warnLogger = createLogger(os.Stdout, "[WARN] ", log.Ldate|log.Ltime)
	errorLogger = createLogger(os.Stderr, "[ERROR] ", log.Ldate|log.Ltime|log.Lshortfile)
}

// Create logger with specified output and format
func createLogger(output *os.File, prefix string, flags int) *log.Logger {
	return log.New(output, prefix, flags)
}

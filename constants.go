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

const (
	VERSION = "1.0.0"
	// ANSI color codes
	RED    = "\033[0;31m"
	GREEN  = "\033[0;32m"
	YELLOW = "\033[0;33m"
	BLUE   = "\033[0;34m"
	NC     = "\033[0m" // No Color
	
	// Log levels
	LOG_DEBUG = 0
	LOG_INFO  = 1
	LOG_WARN  = 2
	LOG_ERROR = 3
)

var (
	VERBOSE      = false
	DEBUG        = false
	CONTINUE     = true
	FOLDER       = ""
	ARCHIVE      = ""
	OUTPUT       = ""
	ERROR_DIR    = ""
	LOCKFILE     = ""
	
	// Session statistics
	COUNTER      = 0
	ERROR_COUNT  = 0
	START_TIME   = time.Now()
	
	// Structured logging
	debugLogger  *log.Logger
	infoLogger   *log.Logger
	warnLogger   *log.Logger
	errorLogger  *log.Logger
)

// Initialize loggers
func initLoggers() {
	debugLogger = log.New(os.Stdout, "[DEBUG] ", log.Ldate|log.Ltime|log.Lshortfile)
	infoLogger = log.New(os.Stdout, "[INFO] ", log.Ldate|log.Ltime)
	warnLogger = log.New(os.Stdout, "[WARN] ", log.Ldate|log.Ltime)
	errorLogger = log.New(os.Stderr, "[ERROR] ", log.Ldate|log.Ltime|log.Lshortfile)
}
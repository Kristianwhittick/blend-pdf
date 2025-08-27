package main

const (
	VERSION = "1.0.0"
	// ANSI color codes
	RED    = "\033[0;31m"
	GREEN  = "\033[0;32m"
	YELLOW = "\033[0;33m"
	BLUE   = "\033[0;34m"
	NC     = "\033[0m" // No Color
)

var (
	VERBOSE   = false
	CONTINUE  = true
	FOLDER    = ""
	ARCHIVE   = ""
	OUTPUT    = ""
	ERROR_DIR = ""
	LOCKFILE  = ""
)
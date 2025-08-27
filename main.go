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
	"bufio"
	"fmt"
	"io"
	"os"
	"os/signal"
	"path/filepath"
	"strings"
	"syscall"
)

func main() {
	// Set up signal handling for graceful shutdown
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		fmt.Printf("\n%sShutting down gracefully...%s\n", YELLOW, NC)
		cleanup()
		os.Exit(0)
	}()

	// Parse command line arguments
	folder, err := parseArgs()
	if err != nil {
		printError(err.Error())
		os.Exit(1)
	}

	// Setup directories
	if err := setupDirectories(folder); err != nil {
		printError(err.Error())
		os.Exit(1)
	}

	// Main program loop
	for CONTINUE {
		processMenu()
	}

	cleanup()
}

func cleanup() {
	showStatistics()
}

func showHelp() {
	fmt.Printf("BlendPDF v%s - A tool for merging PDF files\n\n", VERSION)
	fmt.Printf("Usage: %s [options] [folder]\n\n", filepath.Base(os.Args[0]))
	fmt.Printf("Command line options:\n")
	fmt.Printf("  -h, --help     Show this help information and exit\n")
	fmt.Printf("  -v, --version  Show version information and exit\n")
	fmt.Printf("  -V, --verbose  Enable verbose mode (show all program output)\n")
	fmt.Printf("  [folder]       Specify folder to watch (default: current directory)\n\n")
	fmt.Printf("Examples:\n")
	fmt.Printf("  %s -h                # Show help\n", filepath.Base(os.Args[0]))
	fmt.Printf("  %s -v                # Show version\n", filepath.Base(os.Args[0]))
	fmt.Printf("  %s -V                # Run in verbose mode\n", filepath.Base(os.Args[0]))
	fmt.Printf("  %s /path/to/pdfs     # Watch specific folder\n", filepath.Base(os.Args[0]))
	fmt.Printf("  %s                   # Watch current directory\n\n", filepath.Base(os.Args[0]))
	fmt.Printf("Interactive options:\n")
	fmt.Printf("  S - Move a single PDF file to the output directory\n")
	fmt.Printf("  M - Merge two PDF files (first file + reversed second file)\n")
	fmt.Printf("  H - Show this help information\n")
	fmt.Printf("  V - Toggle verbose mode\n")
	fmt.Printf("  Q - Quit the program\n\n")
}

func processMenu() {
	fmt.Println()
	displayFileCounts()
	showFilePreview()
	
	fmt.Printf("Options: %s[S]%single, %s[M]%serge, %s[H]%selp, %s[V]%serbose, %s[Q]%suit\n", 
		YELLOW, NC, YELLOW, NC, YELLOW, NC, YELLOW, NC, YELLOW, NC)
	
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter choice (S/M/H/V/Q): ")
	input, err := reader.ReadString('\n')
	if err != nil {
		// Handle EOF gracefully (happens when input is piped)
		if err == io.EOF {
			fmt.Printf("\n%sEnd of input reached. Exiting...%s\n", YELLOW, NC)
			CONTINUE = false
			return
		}
		printError(fmt.Sprintf("Error reading input: %v", err))
		return
	}
	
	input = strings.TrimSpace(strings.ToUpper(input))
	
	switch input {
	case "S":
		processSingleFile()
	case "M":
		processMergeFiles()
	case "H":
		showHelp()
	case "V":
		VERBOSE = !VERBOSE
		if VERBOSE {
			printSuccess("Verbose mode enabled")
		} else {
			printInfo("Verbose mode disabled")
		}
	case "Q":
		fmt.Printf("%sExiting program...%s\n", YELLOW, NC)
		CONTINUE = false
	default:
		printWarning("Invalid choice. Please enter S, M, H, V, or Q.")
	}
}

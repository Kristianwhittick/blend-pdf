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

func parseArgs() (string, bool) {
	args := os.Args[1:]
	var folder string
	
	for i := range args {
		arg := args[i]
		switch arg {
		case "-h", "--help":
			showHelp()
			return "", true
		case "-v", "--version":
			fmt.Printf("BlendPDF v%s\n", VERSION)
			return "", true
		case "-V", "--verbose":
			VERBOSE = true
		default:
			// Assume it's a folder path
			folder = arg
		}
	}
	
	return folder, false
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
	fmt.Printf("Options: %s[S]%s single file move, %s[M]%s merge two files, %s[H]%s help, %s[V]%s verbose, %s[Q]%s quit\n", 
		YELLOW, NC, YELLOW, NC, YELLOW, NC, YELLOW, NC, YELLOW, NC)
	
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter choice: ")
	input, err := reader.ReadString('\n')
	if err != nil {
		// Handle EOF gracefully (happens when input is piped)
		if err == io.EOF {
			fmt.Printf("\n%sEnd of input reached. Exiting...%s\n", YELLOW, NC)
			cleanup()
			return
		}
		fmt.Printf("%sError reading input: %v%s\n", RED, err, NC)
		return
	}
	
	input = strings.TrimSpace(strings.ToUpper(input))
	
	switch input {
	case "S":
		moveSingleFile()
	case "M":
		mergeFiles()
	case "H":
		showHelp()
	case "V":
		VERBOSE = !VERBOSE
		if VERBOSE {
			fmt.Printf("%sVerbose mode enabled.%s\n", GREEN, NC)
		} else {
			fmt.Printf("%sVerbose mode disabled.%s\n", YELLOW, NC)
		}
	case "Q":
		cleanup()
	case "":
		// Empty input, just continue
		return
	default:
		fmt.Printf("%sInvalid option '%s'. Please enter S, M, H, V, or Q.%s\n", RED, input, NC)
	}
}

func main() {
	// Parse command line arguments
	folder, shouldExit := parseArgs()
	if shouldExit {
		return
	}

	// Setup lock file to prevent multiple instances
	if err := setupLock(); err != nil {
		return
	}
	defer os.Remove(LOCKFILE)

	// Handle Ctrl+C gracefully
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		cleanup()
	}()

	// Setup directories
	if err := setupDirectories(folder); err != nil {
		return
	}

	// Main program loop
	fmt.Printf("%sBlendPDF v%s started.%s\n", GREEN, VERSION, NC)
	fmt.Printf("Press Ctrl+C to exit.\n")

	for CONTINUE {
		processMenu()
	}
}
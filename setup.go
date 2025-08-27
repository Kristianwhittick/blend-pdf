package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func setupLock() error {
	LOCKFILE = filepath.Join(os.TempDir(), filepath.Base(os.Args[0])+".lock")

	if _, err := os.Stat(LOCKFILE); err == nil {
		fmt.Printf("%sError:%s Script is already running\n", RED, NC)
		return fmt.Errorf("script is already running")
	}

	file, err := os.Create(LOCKFILE)
	if err != nil {
		return err
	}
	file.Close()

	return nil
}

func cleanup() {
	fmt.Printf("%sShutting down...%s\n", YELLOW, NC)
	os.Remove(LOCKFILE)
	CONTINUE = false
	os.Exit(0)
}

func setupDirectories(folder string) error {
	// Use provided folder or current directory
	if folder == "" {
		var err error
		FOLDER, err = os.Getwd()
		if err != nil {
			return err
		}
	} else {
		FOLDER = folder
	}

	ARCHIVE = filepath.Join(FOLDER, "archive")
	OUTPUT = filepath.Join(FOLDER, "output")
	ERROR_DIR = filepath.Join(FOLDER, "error")

	// Validate folder exists
	if _, err := os.Stat(FOLDER); os.IsNotExist(err) {
		fmt.Printf("%sError:%s Directory '%s' does not exist.\n", RED, NC, FOLDER)
		return err
	}

	// Create required folders
	os.MkdirAll(ARCHIVE, 0755)
	os.MkdirAll(OUTPUT, 0755)
	os.MkdirAll(ERROR_DIR, 0755)

	fmt.Printf("Watching folder: %s%s%s\n", BLUE, FOLDER, NC)
	fmt.Printf("Archive  folder: %s%s%s\n", BLUE, ARCHIVE, NC)
	fmt.Printf("Output   folder: %s%s%s\n", BLUE, OUTPUT, NC)
	fmt.Printf("Error    folder: %s%s%s\n", BLUE, ERROR_DIR, NC)
	fmt.Println()

	return nil
}

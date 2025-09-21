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
	"crypto/md5"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
)

// Lock file management functions

// Setup lock file to prevent multiple instances
func setupLock() error {
	watchDir := determineWatchDirectory()
	LOCKFILE = generateLockFileName(watchDir)

	if err := checkExistingLockFile(); err != nil {
		return err
	}

	return createLockFile()
}

// Determine watch directory from command line arguments
func determineWatchDirectory() string {
	watchDir := "." // Default to current directory

	if len(os.Args) > 1 {
		lastArg := os.Args[len(os.Args)-1]
		if !strings.HasPrefix(lastArg, "-") {
			watchDir = lastArg
		}
	}

	return watchDir
}

// Generate directory-specific lock file name using hash
func generateLockFileName(watchDir string) string {
	normalizedPath := normalizeDirectoryPath(watchDir)
	hashStr := generateDirectoryHash(normalizedPath)
	lockFileName := fmt.Sprintf("blendpdf-%s.lock", hashStr)

	return createPlatformSpecificLockPath(watchDir, lockFileName)
}

// Normalize directory path for consistent hashing
func normalizeDirectoryPath(watchDir string) string {
	absPath, err := filepath.Abs(watchDir)
	if err != nil {
		absPath = watchDir // Fallback to original path
	}

	cleanPath := filepath.Clean(absPath)
	return strings.ToLower(filepath.ToSlash(cleanPath))
}

// Generate 8-character MD5 hash from path
func generateDirectoryHash(normalizedPath string) string {
	hash := md5.Sum([]byte(normalizedPath))
	return fmt.Sprintf("%x", hash)[:8]
}

// Create platform-specific lock file path
func createPlatformSpecificLockPath(watchDir, lockFileName string) string {
	if runtime.GOOS == "windows" {
		// On Windows, store in the watch directory to avoid permission issues
		return filepath.Join(watchDir, lockFileName)
	}
	// On Unix systems, use /tmp directory
	return filepath.Join("/tmp", lockFileName)
}

// Check if lock file already exists
func checkExistingLockFile() error {
	if _, err := os.Stat(LOCKFILE); err == nil {
		return createLockFileExistsError()
	}
	return nil
}

// Create error for existing lock file
func createLockFileExistsError() error {
	printError("Another instance is already running")
	printInfo(fmt.Sprintf("Lock file exists: %s", LOCKFILE))
	printInfo("If you're sure no other instance is running, remove the lock file manually")
	return fmt.Errorf("already running (exit code 6)")
}

// Create new lock file with process ID
func createLockFile() error {
	file, err := os.Create(LOCKFILE)
	if err != nil {
		return fmt.Errorf("failed to create lock file: %v", err)
	}
	defer file.Close()

	if err := writePIDToLockFile(file); err != nil {
		cleanupFailedLockFile()
		return err
	}

	logLockFileCreation()
	return nil
}

// Write process ID to lock file
func writePIDToLockFile(file *os.File) error {
	_, err := file.WriteString(strconv.Itoa(os.Getpid()))
	if err != nil {
		return fmt.Errorf("failed to write to lock file: %v", err)
	}
	return nil
}

// Clean up failed lock file creation
func cleanupFailedLockFile() {
	if err := os.Remove(LOCKFILE); err != nil {
		printError(fmt.Sprintf("Failed to remove lock file: %v", err))
	}
}

// Log lock file creation in verbose mode
func logLockFileCreation() {
	if VERBOSE {
		printInfo(fmt.Sprintf("Created lock file: %s", LOCKFILE))
	}
}

// Clean up lock file and resources
func cleanupLock() {
	if LOCKFILE == "" {
		return
	}

	if err := os.Remove(LOCKFILE); err != nil {
		if VERBOSE {
			printWarning(fmt.Sprintf("Failed to remove lock file: %v", err))
		}
	} else if VERBOSE {
		printInfo("Removed lock file")
	}
}

// Command line argument parsing functions

// Global configuration
var CONFIG *Config

// Parse command line arguments with enhanced error handling
func parseArgs() (string, error) {
	args := os.Args[1:]
	folder := ""

	for i, arg := range args {
		if err := processArgument(arg, args, i, &folder); err != nil {
			return "", err
		}
	}

	resolvedFolder, err := resolveFolderPath(folder)
	if err != nil {
		return "", err
	}

	// Load configuration after resolving folder
	CONFIG, err = loadConfig(resolvedFolder)
	if err != nil && VERBOSE {
		printWarning(fmt.Sprintf("Failed to load config: %v, using defaults", err))
		CONFIG = getDefaultConfig()
	}

	// Override config with command line flags
	applyCommandLineOverrides(args)

	return resolvedFolder, nil
}

// Apply command line overrides to configuration
func applyCommandLineOverrides(args []string) {
	for i, arg := range args {
		switch arg {
		case "-V", "--verbose":
			CONFIG.VerboseMode = true
		case "-D", "--debug":
			CONFIG.DebugMode = true
			CONFIG.VerboseMode = true
		case "--no-archive":
			CONFIG.ArchiveMode = false
		case "-o", "--output":
			// Handle multiple output folders: -o folder1,folder2,folder3
			if i+1 < len(args) {
				folders := strings.Split(args[i+1], ",")
				CONFIG.OutputFolders = folders
			}
		}
	}

	// Apply config to global variables
	VERBOSE = CONFIG.VerboseMode
	DEBUG = CONFIG.DebugMode
}

// Process individual command line argument
func processArgument(arg string, args []string, index int, folder *string) error {
	switch arg {
	case "-h", "--help":
		showHelp()
		os.Exit(0)
	case "-v", "--version":
		showVersion()
		os.Exit(0)
	case "-V", "--verbose":
		enableVerboseMode()
	case "-D", "--debug":
		enableDebugMode()
	case "--no-archive":
		// Handled in applyCommandLineOverrides
	case "-o", "--output":
		// Skip the next argument (it's the folder list)
		// Handled in applyCommandLineOverrides
	default:
		return handleNonFlagArgument(arg, args, index, folder)
	}
	return nil
}

// Show version information
func showVersion() {
	fmt.Printf("BlendPDFGo v%s\n", VERSION)
}

// Enable verbose mode
func enableVerboseMode() {
	VERBOSE = true
	printSuccess("Verbose mode enabled")
}

// Enable debug mode
func enableDebugMode() {
	DEBUG = true
	VERBOSE = true // Debug mode implies verbose
	initLoggers()
	printSuccess("Debug mode enabled (includes verbose)")
}

// Handle non-flag arguments (folder paths)
func handleNonFlagArgument(arg string, args []string, index int, folder *string) error {
	if strings.HasPrefix(arg, "-") {
		return fmt.Errorf("unknown flag: %s", arg)
	}

	if *folder != "" {
		return fmt.Errorf("multiple folder paths specified")
	}

	*folder = arg
	return nil
}

// Resolve folder path to absolute path
func resolveFolderPath(folder string) (string, error) {
	if folder == "" {
		return getCurrentDirectory()
	}

	absFolder, err := filepath.Abs(folder)
	if err != nil {
		return "", fmt.Errorf("failed to resolve absolute path for '%s': %v", folder, err)
	}

	logParsedArguments(absFolder)
	return absFolder, nil
}

// Get current working directory
func getCurrentDirectory() (string, error) {
	folder, err := os.Getwd()
	if err != nil {
		return "", fmt.Errorf("failed to get current directory: %v", err)
	}
	return folder, nil
}

// Log parsed arguments in debug mode
func logParsedArguments(absFolder string) {
	if DEBUG {
		printDebug(fmt.Sprintf("Parsed arguments: folder=%s, verbose=%t, debug=%t", absFolder, VERBOSE, DEBUG))
	}
}

// PDF file validation functions

// Validate PDF file with enhanced error reporting
func validatePDFFile(file string) error {
	if err := checkFileExists(file); err != nil {
		return err
	}

	if err := checkFileProperties(file); err != nil {
		return err
	}

	if err := checkPDFStructure(file); err != nil {
		return err
	}

	return nil
}

// Check if file exists and is accessible
func checkFileExists(file string) error {
	info, err := os.Stat(file)
	if os.IsNotExist(err) {
		return fmt.Errorf("file does not exist")
	}
	if err != nil {
		return fmt.Errorf("cannot access file: %v", err)
	}

	// Store file info for further checks
	return checkFileType(info)
}

// Check file type and properties
func checkFileType(info os.FileInfo) error {
	if !info.Mode().IsRegular() {
		return fmt.Errorf("not a regular file")
	}

	if info.Size() == 0 {
		return fmt.Errorf("file is empty")
	}

	return nil
}

// Check file properties (extension, size, etc.)
func checkFileProperties(file string) error {
	if !strings.HasSuffix(strings.ToLower(file), ".pdf") {
		return fmt.Errorf("file does not have .pdf extension")
	}
	return nil
}

// Check PDF structure using pdfcpu
func checkPDFStructure(file string) error {
	if !validatePDF(file) {
		return fmt.Errorf("invalid PDF structure")
	}
	return nil
}

// File operation utilities

// Enhanced file operation with error recovery
// Move file with error recovery and conflict resolution
func moveFileWithRecovery(src, dst string) error {
	if err := ensureDestinationDirectory(dst); err != nil {
		return err
	}

	dst = resolveDestinationConflicts(dst)
	return performFileMove(src, dst)
}

// Ensure destination directory exists
func ensureDestinationDirectory(dst string) error {
	dstDir := filepath.Dir(dst)
	if err := os.MkdirAll(dstDir, 0750); err != nil {
		return fmt.Errorf("failed to create destination directory: %v", err)
	}
	return nil
}

// Resolve destination file conflicts by generating unique names
func resolveDestinationConflicts(dst string) string {
	if _, err := os.Stat(dst); os.IsNotExist(err) {
		return dst // No conflict
	}

	return generateUniqueFileName(dst)
}

// Generate unique filename if destination exists
func generateUniqueFileName(dst string) string {
	dstDir := filepath.Dir(dst)
	base := strings.TrimSuffix(filepath.Base(dst), filepath.Ext(dst))
	ext := filepath.Ext(dst)

	for counter := 1; counter <= 1000; counter++ {
		newDst := filepath.Join(dstDir, fmt.Sprintf("%s_%d%s", base, counter, ext))
		if _, err := os.Stat(newDst); os.IsNotExist(err) {
			if VERBOSE {
				printWarning(fmt.Sprintf("Destination exists, using: %s", filepath.Base(newDst)))
			}
			return newDst
		}
	}

	// If we can't find a unique name after 1000 attempts, use original
	return dst
}

// Perform the actual file move operation
func performFileMove(src, dst string) error {
	err := os.Rename(src, dst)
	if err != nil {
		return attemptCopyAndDelete(src, dst, err)
	}
	return nil
}

// Attempt copy and delete as fallback for move operations
func attemptCopyAndDelete(src, dst string, originalErr error) error {
	if copyErr := copyFile(src, dst); copyErr != nil {
		return fmt.Errorf("move failed: %v, copy fallback failed: %v", originalErr, copyErr)
	}

	if deleteErr := os.Remove(src); deleteErr != nil {
		printWarning(fmt.Sprintf("Original file not deleted: %v", deleteErr))
	}

	return nil
}

// Copy file as fallback for move operations
func copyFile(src, dst string) error {
	if err := validateFilePaths(src, dst); err != nil {
		return err
	}

	return performFileCopy(src, dst)
}

// Validate file paths to prevent directory traversal
func validateFilePaths(src, dst string) error {
	cleanSrc := filepath.Clean(src)
	cleanDst := filepath.Clean(dst)

	if strings.Contains(cleanSrc, "..") || strings.Contains(cleanDst, "..") {
		return fmt.Errorf("invalid file path: directory traversal not allowed")
	}

	return nil
}

// Perform the actual file copy operation
func performFileCopy(src, dst string) error {
	sourceFile, err := os.Open(src) // #nosec G304 - path validated above
	if err != nil {
		return err
	}
	defer sourceFile.Close()

	destFile, err := os.Create(dst) // #nosec G304 - path validated above
	if err != nil {
		return err
	}
	defer destFile.Close()

	_, err = destFile.ReadFrom(sourceFile)
	return err
}

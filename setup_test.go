package main

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseArgs(t *testing.T) {
	tests := []struct {
		name string
		args []string
	}{
		{"No arguments", []string{"blendpdf"}},
		{"Verbose flag", []string{"blendpdf", "-V"}},
		{"Debug flag", []string{"blendpdf", "-D"}},
		{"Folder argument", []string{"blendpdf", "/tmp/test"}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Save original os.Args
			originalArgs := os.Args
			defer func() { os.Args = originalArgs }()
			
			// Set test args
			os.Args = tt.args
			
			// Test that parseArgs doesn't panic
			assert.NotPanics(t, func() {
				folder, err := parseArgs()
				// Basic validation that it returns something
				assert.NotEmpty(t, folder)
				assert.NoError(t, err)
			})
		})
	}
}

func TestGenerateDirectoryHash(t *testing.T) {
	tests := []struct {
		name     string
		path     string
		expected int // Expected length
	}{
		{"Simple path", "/tmp/test", 8},
		{"Complex path", "/home/user/projects/blendpdf", 8},
		{"Current directory", ".", 8},
		{"Relative path", "../test", 8},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			hash := generateDirectoryHash(tt.path)
			assert.Len(t, hash, tt.expected, "Hash should be 8 characters")
			assert.Regexp(t, "^[a-f0-9]+$", hash, "Hash should be lowercase hex")
		})
	}
}

func TestGenerateDirectoryHashConsistency(t *testing.T) {
	// Same path should generate same hash
	path := "/tmp/test"
	hash1 := generateDirectoryHash(path)
	hash2 := generateDirectoryHash(path)
	assert.Equal(t, hash1, hash2, "Same path should generate same hash")
}

func TestNormalizeDirectoryPath(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		contains []string // Strings that should be in the result
	}{
		{"Simple path", "/tmp/test", []string{"/tmp/test"}},
		{"Path with trailing slash", "/tmp/test/", []string{"/tmp/test"}},
		{"Current directory", ".", []string{}}, // Will be converted to absolute path
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := normalizeDirectoryPath(tt.input)
			assert.NotEmpty(t, result, "Normalized path should not be empty")
			
			for _, expected := range tt.contains {
				assert.Contains(t, result, expected, "Result should contain expected substring")
			}
		})
	}
}

func TestSetupLock(t *testing.T) {
	tempDir := t.TempDir()
	
	// Save original global variables and os.Args
	originalFolder := FOLDER
	originalLockfile := LOCKFILE
	originalArgs := os.Args
	defer func() {
		FOLDER = originalFolder
		LOCKFILE = originalLockfile
		os.Args = originalArgs
		cleanupLock() // Ensure cleanup
	}()
	
	// Set up test environment with unique directory
	FOLDER = tempDir
	LOCKFILE = "" // Reset
	// Set os.Args to use our temp directory
	os.Args = []string{"blendpdf.test", tempDir}
	
	// Test lock setup
	err := setupLock()
	assert.NoError(t, err, "setupLock should succeed")
	
	// Verify lock file was created
	assert.NotEmpty(t, LOCKFILE, "LOCKFILE should be set")
	assert.FileExists(t, LOCKFILE, "Lock file should exist")
}

func TestSetupLockAlreadyExists(t *testing.T) {
	tempDir := t.TempDir()
	
	// Save original global variables and os.Args
	originalFolder := FOLDER
	originalLockfile := LOCKFILE
	originalArgs := os.Args
	defer func() {
		FOLDER = originalFolder
		LOCKFILE = originalLockfile
		os.Args = originalArgs
		cleanupLock() // Ensure cleanup
	}()
	
	// Set up test environment with unique directory
	FOLDER = tempDir
	LOCKFILE = "" // Reset
	// Set os.Args to use our temp directory
	os.Args = []string{"blendpdf.test", tempDir}
	
	// Create first lock
	err := setupLock()
	assert.NoError(t, err)
	
	// Save the first lock file path
	firstLockFile := LOCKFILE
	
	// Try to create second lock (should fail)
	err = setupLock()
	assert.Error(t, err, "Second setupLock should fail")
	assert.Contains(t, strings.ToLower(err.Error()), "already running", "Error should mention already running")
	
	// Clean up the first lock
	LOCKFILE = firstLockFile
	cleanupLock()
}

func TestCleanupLock(t *testing.T) {
	tempDir := t.TempDir()
	
	// Save original global variables and os.Args
	originalFolder := FOLDER
	originalLockfile := LOCKFILE
	originalArgs := os.Args
	defer func() {
		FOLDER = originalFolder
		LOCKFILE = originalLockfile
		os.Args = originalArgs
	}()
	
	// Set up test environment with unique directory
	FOLDER = tempDir
	LOCKFILE = "" // Reset
	// Set os.Args to use our temp directory
	os.Args = []string{"blendpdf.test", tempDir}
	
	// Create lock file
	err := setupLock()
	assert.NoError(t, err)
	
	// Verify it exists
	assert.FileExists(t, LOCKFILE)
	
	// Clean up
	cleanupLock()
	
	// Verify it's gone
	assert.NoFileExists(t, LOCKFILE, "Lock file should be removed")
}

func TestValidatePDFFile(t *testing.T) {
	tempDir := t.TempDir()
	
	tests := []struct {
		name      string
		filename  string
		content   string
		shouldErr bool
	}{
		{
			name:      "Valid file exists",
			filename:  "test.pdf",
			content:   "fake pdf content",
			shouldErr: true, // Will error with fake PDF content
		},
		{
			name:      "Non-existent file",
			filename:  "nonexistent.pdf",
			content:   "",
			shouldErr: true,
		},
		{
			name:      "Empty file",
			filename:  "empty.pdf",
			content:   "",
			shouldErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var filePath string
			if tt.content != "" {
				filePath = filepath.Join(tempDir, tt.filename)
				err := os.WriteFile(filePath, []byte(tt.content), 0644)
				assert.NoError(t, err)
			} else if tt.filename != "nonexistent.pdf" {
				filePath = filepath.Join(tempDir, tt.filename)
				err := os.WriteFile(filePath, []byte(tt.content), 0644)
				assert.NoError(t, err)
			} else {
				filePath = filepath.Join(tempDir, tt.filename)
			}
			
			err := validatePDFFile(filePath)
			if tt.shouldErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestMoveFileWithRecovery(t *testing.T) {
	tempDir := t.TempDir()
	
	// Create source file
	sourceFile := filepath.Join(tempDir, "test.pdf")
	err := os.WriteFile(sourceFile, []byte("test content"), 0644)
	assert.NoError(t, err)
	
	// Create destination directory
	destDir := filepath.Join(tempDir, "destination")
	err = os.MkdirAll(destDir, 0755)
	assert.NoError(t, err)
	
	destFile := filepath.Join(destDir, "test.pdf")
	
	// Move file
	err = moveFileWithRecovery(sourceFile, destFile)
	assert.NoError(t, err)
	
	// Verify source file is gone
	assert.NoFileExists(t, sourceFile, "Source file should be removed")
	
	// Verify destination file exists
	assert.FileExists(t, destFile, "Destination file should exist")
	
	// Verify content
	content, err := os.ReadFile(destFile)
	assert.NoError(t, err)
	assert.Equal(t, "test content", string(content))
}

func TestMoveFileWithRecoveryNonExistentSource(t *testing.T) {
	tempDir := t.TempDir()
	destDir := filepath.Join(tempDir, "destination")
	err := os.MkdirAll(destDir, 0755)
	assert.NoError(t, err)
	
	err = moveFileWithRecovery("/non/existent/file.pdf", filepath.Join(destDir, "test.pdf"))
	assert.Error(t, err)
}

func TestEnableDebugMode(t *testing.T) {
	// Save original state
	originalDebug := DEBUG
	originalVerbose := VERBOSE
	defer func() { 
		DEBUG = originalDebug
		VERBOSE = originalVerbose
	}()
	
	DEBUG = false
	VERBOSE = false
	enableDebugMode()
	assert.True(t, DEBUG, "Should enable debug mode")
	assert.True(t, VERBOSE, "Should enable verbose mode when debug is enabled")
}

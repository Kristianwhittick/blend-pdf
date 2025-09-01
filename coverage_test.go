package main

import (
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

// Minimal tests to increase coverage

func TestBasicFunctions(t *testing.T) {
	// Test basic display functions
	assert.NotPanics(t, func() {
		showHelp()
		showCommandLineOptions()
		showUsageExamples()
		showInteractiveOptions()
		displayMenuOptions()
		exitApplication()
		cleanup()
		showStatistics()
	})
}

func TestToggleFunctions(t *testing.T) {
	originalVerbose := VERBOSE
	originalDebug := DEBUG
	defer func() {
		VERBOSE = originalVerbose
		DEBUG = originalDebug
	}()
	
	VERBOSE = false
	toggleVerboseMode()
	assert.True(t, VERBOSE)
	
	DEBUG = false
	toggleDebugMode()
	assert.True(t, DEBUG)
}

func TestExecuteUserChoiceBasic(t *testing.T) {
	assert.NotPanics(t, func() {
		executeUserChoice("S")
		executeUserChoice("M")
		executeUserChoice("H")
		executeUserChoice("V")
		executeUserChoice("D")
		executeUserChoice("Q")
		executeUserChoice("invalid")
	})
}

func TestFileOperations(t *testing.T) {
	tempDir := t.TempDir()
	
	// Test file existence
	testFile := filepath.Join(tempDir, "test.txt")
	assert.False(t, fileExists(testFile))
	
	_ = os.WriteFile(testFile, []byte("test content"), 0644)
	assert.True(t, fileExists(testFile))
	
	// Test file size
	size := getFileSize(testFile)
	assert.Equal(t, int64(12), size)
	
	size = getFileSize("nonexistent")
	assert.Equal(t, int64(0), size)
}

func TestBasicValidation(t *testing.T) {
	tempDir := t.TempDir()
	
	// Test with invalid PDF
	invalidPDF := filepath.Join(tempDir, "invalid.pdf")
	_ = os.WriteFile(invalidPDF, []byte("not a pdf"), 0644)
	
	// These should return errors but test the code paths
	assert.NotPanics(t, func() {
		_ = validatePDF(invalidPDF)
		_, _ = getPageCount(invalidPDF)
	})
	
	config := createValidationConfig()
	assert.NotNil(t, config)
}

func TestCleanupOperations(t *testing.T) {
	tempDir := t.TempDir()
	
	// Create test files
	file1 := filepath.Join(tempDir, "temp1.pdf")
	file2 := filepath.Join(tempDir, "temp2.pdf")
	_ = os.WriteFile(file1, []byte("test1"), 0644)
	_ = os.WriteFile(file2, []byte("test2"), 0644)
	
	// Test cleanup
	cleanupTempFiles([]string{file1, file2})
	
	assert.False(t, fileExists(file1))
	assert.False(t, fileExists(file2))
}

func TestReversedFileName(t *testing.T) {
	result := createReversedFileName("test.pdf")
	assert.Contains(t, result, "reverse")
}

func TestLogOperations(t *testing.T) {
	originalDebug := DEBUG
	defer func() { DEBUG = originalDebug }()
	
	DEBUG = true
	assert.NotPanics(t, func() {
		logDebugOperation("TEST", "operation")
	})
	
	DEBUG = false
	assert.NotPanics(t, func() {
		logDebugOperation("TEST", "operation")
	})
}

func TestProcessOperations(t *testing.T) {
	tempDir := t.TempDir()
	originalFolder := FOLDER
	defer func() { FOLDER = originalFolder }()
	FOLDER = tempDir
	
	// Create directories
	_ = os.MkdirAll(filepath.Join(tempDir, "archive"), 0755)
	_ = os.MkdirAll(filepath.Join(tempDir, "output"), 0755)
	_ = os.MkdirAll(filepath.Join(tempDir, "error"), 0755)
	
	assert.NotPanics(t, func() {
		processSingleFileOperation()
		processMergeOperation()
		showApplicationHelp()
		processSingleFileWithValidation()
		processMergeFilesWithValidation()
		displayApplicationStatus()
	})
}

func TestStatisticsAndCounters(t *testing.T) {
	// Set test values
	originalCounter := COUNTER
	originalErrorCount := ERROR_COUNT
	originalStartTime := START_TIME
	defer func() {
		COUNTER = originalCounter
		ERROR_COUNT = originalErrorCount
		START_TIME = originalStartTime
	}()
	
	COUNTER = 10
	ERROR_COUNT = 3
	START_TIME = time.Now().Add(-2 * time.Minute)
	
	assert.NotPanics(t, func() {
		showStatistics()
	})
}

func TestRemoveFile(t *testing.T) {
	tempDir := t.TempDir()
	testFile := filepath.Join(tempDir, "test-reverse.pdf")
	os.WriteFile(testFile, []byte("test"), 0644)
	
	assert.NotPanics(t, func() {
		removeReversedFile(testFile)
	})
	
	assert.False(t, fileExists(testFile))
}

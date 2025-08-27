package main

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSetupDirectories(t *testing.T) {
	tempDir := t.TempDir()
	
	// Save original global variables
	originalFolder := FOLDER
	originalArchive := ARCHIVE
	originalOutput := OUTPUT
	originalErrorDir := ERROR_DIR
	
	defer func() {
		FOLDER = originalFolder
		ARCHIVE = originalArchive
		OUTPUT = originalOutput
		ERROR_DIR = originalErrorDir
	}()
	
	err := setupDirectories(tempDir)
	assert.NoError(t, err)
	
	// Verify global variables are set
	assert.Equal(t, tempDir, FOLDER)
	assert.Equal(t, filepath.Join(tempDir, "archive"), ARCHIVE)
	assert.Equal(t, filepath.Join(tempDir, "output"), OUTPUT)
	assert.Equal(t, filepath.Join(tempDir, "error"), ERROR_DIR)
	
	// Verify directories were created
	assert.DirExists(t, ARCHIVE, "archive directory should exist")
	assert.DirExists(t, OUTPUT, "output directory should exist")
	assert.DirExists(t, ERROR_DIR, "error directory should exist")
}

func TestFindPDFFiles(t *testing.T) {
	tempDir := t.TempDir()
	
	// Set up the global FOLDER variable
	originalFolder := FOLDER
	defer func() { FOLDER = originalFolder }()
	
	err := setupDirectories(tempDir)
	assert.NoError(t, err)
	
	// Create test files
	pdfFiles := []string{"doc1.pdf", "doc2.pdf", "doc3.pdf"}
	nonPdfFiles := []string{"readme.txt", "image.jpg", "data.csv"}
	
	for _, file := range pdfFiles {
		filePath := filepath.Join(tempDir, file)
		err := os.WriteFile(filePath, []byte("fake pdf content"), 0644)
		assert.NoError(t, err)
	}
	
	for _, file := range nonPdfFiles {
		filePath := filepath.Join(tempDir, file)
		err := os.WriteFile(filePath, []byte("other content"), 0644)
		assert.NoError(t, err)
	}
	
	// Test getting PDF files
	files, err := findPDFFiles()
	assert.NoError(t, err)
	assert.Len(t, files, 3, "Should find 3 PDF files")
	
	// Verify files are sorted alphabetically
	expectedFiles := []string{"doc1.pdf", "doc2.pdf", "doc3.pdf"}
	for i, file := range files {
		assert.Equal(t, expectedFiles[i], filepath.Base(file))
	}
}

func TestFindPDFFilesNoPDFs(t *testing.T) {
	tempDir := t.TempDir()
	
	originalFolder := FOLDER
	defer func() { FOLDER = originalFolder }()
	
	err := setupDirectories(tempDir)
	assert.NoError(t, err)
	
	// Create non-PDF files
	nonPdfFiles := []string{"readme.txt", "image.jpg", "data.csv"}
	for _, file := range nonPdfFiles {
		filePath := filepath.Join(tempDir, file)
		err := os.WriteFile(filePath, []byte("content"), 0644)
		assert.NoError(t, err)
	}
	
	files, err := findPDFFiles()
	assert.NoError(t, err)
	assert.Len(t, files, 0, "Should find no PDF files")
}

func TestFindPDFFilesEmptyDirectory(t *testing.T) {
	tempDir := t.TempDir()
	
	originalFolder := FOLDER
	defer func() { FOLDER = originalFolder }()
	
	err := setupDirectories(tempDir)
	assert.NoError(t, err)
	
	files, err := findPDFFiles()
	assert.NoError(t, err)
	assert.Len(t, files, 0, "Should find no files in empty directory")
}

func TestFormatFileSize(t *testing.T) {
	tests := []struct {
		name     string
		size     int64
		expected string
	}{
		{"Bytes", 512, "512B"},
		{"Kilobytes", 1024, "1.0K"},
		{"Kilobytes with decimal", 1536, "1.5K"},
		{"Megabytes", 1048576, "1.0M"},
		{"Megabytes with decimal", 1572864, "1.5M"},
		{"Gigabytes", 1073741824, "1.0G"},
		{"Large size", 2147483648, "2.0G"},
		{"Zero size", 0, "0B"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := formatFileSize(tt.size)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestCountPDFFiles(t *testing.T) {
	tempDir := t.TempDir()
	
	// Create subdirectories
	archiveDir := filepath.Join(tempDir, "archive")
	outputDir := filepath.Join(tempDir, "output")
	errorDir := filepath.Join(tempDir, "error")
	
	err := os.MkdirAll(archiveDir, 0755)
	assert.NoError(t, err)
	err = os.MkdirAll(outputDir, 0755)
	assert.NoError(t, err)
	err = os.MkdirAll(errorDir, 0755)
	assert.NoError(t, err)
	
	// Create PDF files in different directories
	mainFiles := []string{"main1.pdf", "main2.pdf"}
	archiveFiles := []string{"archive1.pdf"}
	outputFiles := []string{"output1.pdf", "output2.pdf", "output3.pdf"}
	errorFiles := []string{"error1.pdf", "error2.pdf"}
	
	for _, file := range mainFiles {
		filePath := filepath.Join(tempDir, file)
		err := os.WriteFile(filePath, []byte("content"), 0644)
		assert.NoError(t, err)
	}
	
	for _, file := range archiveFiles {
		filePath := filepath.Join(archiveDir, file)
		err := os.WriteFile(filePath, []byte("content"), 0644)
		assert.NoError(t, err)
	}
	
	for _, file := range outputFiles {
		filePath := filepath.Join(outputDir, file)
		err := os.WriteFile(filePath, []byte("content"), 0644)
		assert.NoError(t, err)
	}
	
	for _, file := range errorFiles {
		filePath := filepath.Join(errorDir, file)
		err := os.WriteFile(filePath, []byte("content"), 0644)
		assert.NoError(t, err)
	}
	
	// Test counting using the actual function
	mainCount := countPDFFiles(tempDir)
	archiveCount := countPDFFiles(archiveDir)
	outputCount := countPDFFiles(outputDir)
	errorCount := countPDFFiles(errorDir)
	
	assert.Equal(t, 2, mainCount, "Should count 2 main PDF files")
	assert.Equal(t, 1, archiveCount, "Should count 1 archive PDF file")
	assert.Equal(t, 3, outputCount, "Should count 3 output PDF files")
	assert.Equal(t, 2, errorCount, "Should count 2 error PDF files")
}

func TestGetHumanReadableSize(t *testing.T) {
	tempDir := t.TempDir()
	
	// Create test file with known size
	testFile := filepath.Join(tempDir, "test.pdf")
	content := make([]byte, 1024) // 1KB
	err := os.WriteFile(testFile, content, 0644)
	assert.NoError(t, err)
	
	size := getHumanReadableSize(testFile)
	assert.Equal(t, "1.0K", size)
}

func TestGetHumanReadableSizeNonExistent(t *testing.T) {
	size := getHumanReadableSize("/non/existent/file.pdf")
	assert.Equal(t, "unknown", size, "Should return 'unknown' for non-existent file")
}

func TestMoveProcessedFiles(t *testing.T) {
	tempDir := t.TempDir()
	
	// Set up directories
	err := setupDirectories(tempDir)
	assert.NoError(t, err)
	
	// Create test files
	testFiles := []string{"test1.pdf", "test2.pdf"}
	var filePaths []string
	
	for _, file := range testFiles {
		filePath := filepath.Join(tempDir, file)
		err := os.WriteFile(filePath, []byte("content"), 0644)
		assert.NoError(t, err)
		filePaths = append(filePaths, filePath)
	}
	
	// Save original counters
	originalCounter := COUNTER
	originalErrorCount := ERROR_COUNT
	defer func() {
		COUNTER = originalCounter
		ERROR_COUNT = originalErrorCount
	}()
	
	// Test moving files (this will test the function without panicking)
	assert.NotPanics(t, func() {
		moveProcessedFiles(ARCHIVE, "Test move", filePaths...)
	})
}

func TestPrintFunctions(t *testing.T) {
	// Test that print functions don't panic
	assert.NotPanics(t, func() {
		printSuccess("Test success message")
	})
	
	assert.NotPanics(t, func() {
		printError("Test error message")
	})
	
	assert.NotPanics(t, func() {
		printWarning("Test warning message")
	})
	
	assert.NotPanics(t, func() {
		printInfo("Test info message")
	})
	
	assert.NotPanics(t, func() {
		printDebug("Test debug message")
	})
}

func TestLogFunctions(t *testing.T) {
	// Test that log functions don't panic
	assert.NotPanics(t, func() {
		logOperation("TEST", "file1.pdf", "file2.pdf", "SUCCESS")
	})
	
	assert.NotPanics(t, func() {
		logPerformance("TEST", 1000000, 1024) // 1ms, 1KB
	})
}

func TestDisplayFunctions(t *testing.T) {
	tempDir := t.TempDir()
	
	// Set up directories
	err := setupDirectories(tempDir)
	assert.NoError(t, err)
	
	// Test display functions don't panic
	assert.NotPanics(t, func() {
		displayFileCounts()
	})
	
	assert.NotPanics(t, func() {
		showFilePreview()
	})
	
	assert.NotPanics(t, func() {
		showStatistics()
	})
}

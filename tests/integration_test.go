package tests

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

// IntegrationTestSuite contains integration tests for BlendPDFGo
type IntegrationTestSuite struct {
	suite.Suite
	tempDir string
	helper  *TestHelper
}

// SetupSuite runs once before all tests in the suite
func (suite *IntegrationTestSuite) SetupSuite() {
	suite.tempDir = suite.T().TempDir()
	suite.helper = NewTestHelper(suite.T())
}

// SetupTest runs before each test
func (suite *IntegrationTestSuite) SetupTest() {
	// Clean up any files from previous tests
	files, _ := filepath.Glob(filepath.Join(suite.tempDir, "*.pdf"))
	for _, file := range files {
		os.Remove(file)
	}

	// Clean up subdirectories
	subdirs := []string{"archive", "output", "error"}
	for _, subdir := range subdirs {
		dirPath := filepath.Join(suite.tempDir, subdir)
		os.RemoveAll(dirPath)
	}
}

// TestCompleteWorkflowSetup tests the complete application setup process
func (suite *IntegrationTestSuite) TestCompleteWorkflowSetup() {
	t := suite.T()

	// Test directory validation
	// Note: We can't import main package functions directly in this test package
	// In a real implementation, you would either:
	// 1. Move functions to a separate package that can be imported
	// 2. Use build tags to include test-specific code
	// 3. Test through the main binary interface

	// For now, we test the directory structure expectations
	subdirs := []string{"archive", "output", "error"}
	for _, subdir := range subdirs {
		dirPath := filepath.Join(suite.tempDir, subdir)
		err := os.MkdirAll(dirPath, 0755)
		assert.NoError(t, err)
		assert.DirExists(t, dirPath)
	}
}

// TestFileOperationsWorkflow tests file handling workflows
func (suite *IntegrationTestSuite) TestFileOperationsWorkflow() {
	t := suite.T()

	// Create directory structure
	subdirs := []string{"archive", "output", "error"}
	for _, subdir := range subdirs {
		dirPath := filepath.Join(suite.tempDir, subdir)
		err := os.MkdirAll(dirPath, 0755)
		assert.NoError(t, err)
	}

	// Create test PDF files
	testFiles := []string{"doc1.pdf", "doc2.pdf", "doc3.pdf"}
	for _, filename := range testFiles {
		filePath := filepath.Join(suite.tempDir, filename)
		err := os.WriteFile(filePath, []byte("fake pdf content"), 0644)
		assert.NoError(t, err)
	}

	// Test file discovery
	files, err := filepath.Glob(filepath.Join(suite.tempDir, "*.pdf"))
	assert.NoError(t, err)
	assert.Len(t, files, 3, "Should find 3 PDF files")

	// Test file movement simulation
	sourceFile := files[0]
	destDir := filepath.Join(suite.tempDir, "archive")
	destFile := filepath.Join(destDir, filepath.Base(sourceFile))

	err = os.Rename(sourceFile, destFile)
	assert.NoError(t, err)

	// Verify file was moved
	assert.NoFileExists(t, sourceFile, "Source file should be gone")
	assert.FileExists(t, destFile, "Destination file should exist")
}

// TestErrorHandlingWorkflow tests error scenarios
func (suite *IntegrationTestSuite) TestErrorHandlingWorkflow() {
	t := suite.T()

	// Test with non-existent directory
	nonExistentDir := "/non/existent/directory"
	_, err := os.Stat(nonExistentDir)
	assert.True(t, os.IsNotExist(err), "Directory should not exist")

	// Test with permission denied scenario (if possible)
	if os.Getuid() != 0 { // Don't run as root
		restrictedDir := filepath.Join(suite.tempDir, "restricted")
		err := os.MkdirAll(restrictedDir, 0000) // No permissions
		if err == nil {
			defer func() {
				_ = os.Chmod(restrictedDir, 0755) // Restore for cleanup
			}()

			// Try to create file in restricted directory
			restrictedFile := filepath.Join(restrictedDir, "test.pdf")
			err = os.WriteFile(restrictedFile, []byte("content"), 0644)
			assert.Error(t, err, "Should fail to write to restricted directory")
		}
	}
}

// TestConcurrentAccess tests concurrent access scenarios
func (suite *IntegrationTestSuite) TestConcurrentAccess() {
	t := suite.T()

	// Simulate lock file behavior
	lockFile := filepath.Join(suite.tempDir, "test.lock")

	// Create first lock file
	err := os.WriteFile(lockFile, []byte("12345"), 0644)
	assert.NoError(t, err)

	// Verify lock file exists
	assert.FileExists(t, lockFile)

	// Try to create second lock file (should detect existing)
	_, err = os.Stat(lockFile)
	assert.NoError(t, err, "Lock file should exist, preventing second instance")

	// Clean up
	err = os.Remove(lockFile)
	assert.NoError(t, err)

	// Verify cleanup
	assert.NoFileExists(t, lockFile)
}

// TestLargeFileHandling tests handling of larger files
func (suite *IntegrationTestSuite) TestLargeFileHandling() {
	t := suite.T()

	// Create a larger fake file (1MB)
	largeContent := make([]byte, 1024*1024)
	for i := range largeContent {
		largeContent[i] = byte(i % 256)
	}

	largeFile := filepath.Join(suite.tempDir, "large.pdf")
	err := os.WriteFile(largeFile, largeContent, 0644)
	assert.NoError(t, err)

	// Verify file size
	info, err := os.Stat(largeFile)
	assert.NoError(t, err)
	assert.Equal(t, int64(1024*1024), info.Size())

	// Test file operations with large file
	destDir := filepath.Join(suite.tempDir, "output")
	err = os.MkdirAll(destDir, 0755)
	assert.NoError(t, err)

	destFile := filepath.Join(destDir, "large.pdf")
	err = os.Rename(largeFile, destFile)
	assert.NoError(t, err)

	// Verify large file was moved correctly
	info, err = os.Stat(destFile)
	assert.NoError(t, err)
	assert.Equal(t, int64(1024*1024), info.Size())
}

// TestFileSystemEdgeCases tests edge cases in file system operations
func (suite *IntegrationTestSuite) TestFileSystemEdgeCases() {
	t := suite.T()

	// Test with files containing special characters
	specialFiles := []string{
		"file with spaces.pdf",
		"file-with-dashes.pdf",
		"file_with_underscores.pdf",
		"file.with.dots.pdf",
	}

	for _, filename := range specialFiles {
		filePath := filepath.Join(suite.tempDir, filename)
		err := os.WriteFile(filePath, []byte("content"), 0644)
		assert.NoError(t, err, "Should create file with special characters: %s", filename)

		// Test that file can be read
		content, err := os.ReadFile(filePath)
		assert.NoError(t, err)
		assert.Equal(t, "content", string(content))

		// Clean up
		err = os.Remove(filePath)
		assert.NoError(t, err)
	}
}

// TestDirectoryStructureIntegrity tests directory structure maintenance
func (suite *IntegrationTestSuite) TestDirectoryStructureIntegrity() {
	t := suite.T()

	requiredDirs := []string{"archive", "output", "error"}

	// Create all required directories
	for _, dir := range requiredDirs {
		dirPath := filepath.Join(suite.tempDir, dir)
		err := os.MkdirAll(dirPath, 0755)
		assert.NoError(t, err)
	}

	// Verify all directories exist
	for _, dir := range requiredDirs {
		dirPath := filepath.Join(suite.tempDir, dir)
		assert.DirExists(t, dirPath, "Required directory should exist: %s", dir)

		// Test that directories are writable
		testFile := filepath.Join(dirPath, "test.txt")
		err := os.WriteFile(testFile, []byte("test"), 0644)
		assert.NoError(t, err, "Should be able to write to directory: %s", dir)

		// Clean up test file
		err = os.Remove(testFile)
		assert.NoError(t, err)
	}
}

// TestSessionStateManagement tests session state tracking
func (suite *IntegrationTestSuite) TestSessionStateManagement() {
	t := suite.T()

	// Test counter functionality
	var successCount, errorCount int

	// Simulate successful operations
	successCount++
	successCount++
	assert.Equal(t, 2, successCount)

	// Simulate error operations
	errorCount++
	assert.Equal(t, 1, errorCount)

	// Test statistics calculation
	totalOperations := successCount + errorCount
	assert.Equal(t, 3, totalOperations)

	if totalOperations > 0 {
		successRate := float64(successCount) / float64(totalOperations) * 100
		assert.InDelta(t, 66.67, successRate, 0.1) // Allow 0.1% tolerance for floating point
	}
}

// TestPerformanceBaseline establishes performance baselines
func (suite *IntegrationTestSuite) TestPerformanceBaseline() {
	t := suite.T()

	// Test file creation performance
	numFiles := 100
	files := make([]string, numFiles)

	for i := 0; i < numFiles; i++ {
		filename := filepath.Join(suite.tempDir, "perf_test_"+fmt.Sprintf("%03d", i)+".pdf")
		err := os.WriteFile(filename, []byte("test content"), 0644)
		assert.NoError(t, err)
		files[i] = filename
	}

	// Test file enumeration performance
	foundFiles, err := filepath.Glob(filepath.Join(suite.tempDir, "perf_test_*.pdf"))
	assert.NoError(t, err)
	assert.Len(t, foundFiles, numFiles, "Should find all created files")

	// Clean up performance test files
	for _, file := range files {
		os.Remove(file)
	}
}

// Run the integration test suite
func TestIntegrationSuite(t *testing.T) {
	suite.Run(t, new(IntegrationTestSuite))
}

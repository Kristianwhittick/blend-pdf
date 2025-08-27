package tests

import (
	"os"
	"path/filepath"
	"testing"
)

// TestHelper provides utilities for testing
type TestHelper struct {
	t       *testing.T
	tempDir string
}

// NewTestHelper creates a new test helper
func NewTestHelper(t *testing.T) *TestHelper {
	tempDir := t.TempDir()
	return &TestHelper{
		t:       t,
		tempDir: tempDir,
	}
}

// TempDir returns the temporary directory for this test
func (h *TestHelper) TempDir() string {
	return h.tempDir
}

// CreateTempFile creates a temporary file with given content
func (h *TestHelper) CreateTempFile(name, content string) string {
	filePath := filepath.Join(h.tempDir, name)
	err := os.WriteFile(filePath, []byte(content), 0644)
	if err != nil {
		h.t.Fatalf("Failed to create temp file %s: %v", name, err)
	}
	return filePath
}

// CreateTempDir creates a temporary directory
func (h *TestHelper) CreateTempDir(name string) string {
	dirPath := filepath.Join(h.tempDir, name)
	err := os.MkdirAll(dirPath, 0755)
	if err != nil {
		h.t.Fatalf("Failed to create temp dir %s: %v", name, err)
	}
	return dirPath
}

// FileExists checks if a file exists
func (h *TestHelper) FileExists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}

// DirExists checks if a directory exists
func (h *TestHelper) DirExists(path string) bool {
	info, err := os.Stat(path)
	return err == nil && info.IsDir()
}

// Cleanup performs any necessary cleanup (called automatically by t.TempDir())
func (h *TestHelper) Cleanup() {
	// Cleanup is handled automatically by t.TempDir()
}

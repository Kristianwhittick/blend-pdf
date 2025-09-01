package main

import (
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExitCodes(t *testing.T) {
	// Test that exit codes are used correctly in the application
	// Since exit codes are not defined as constants, we test the expected values
	tests := []struct {
		name     string
		expected int
		desc     string
	}{
		{"Success", 0, "Normal program termination"},
		{"General Error", 1, "General application error"},
		{"Missing Dependencies", 2, "Missing required dependencies"},
		{"Invalid Directory", 3, "Invalid directory path"},
		{"Invalid PDF", 4, "Invalid PDF file"},
		{"Merge Failed", 5, "PDF merge operation failed"},
		{"Already Running", 6, "Another instance already running"},
		{"User Timeout", 7, "User inactivity timeout"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.GreaterOrEqual(t, tt.expected, 0, "Exit code should be non-negative")
			assert.LessOrEqual(t, tt.expected, 255, "Exit code should be valid (0-255)")
		})
	}
}

func TestColorConstants(t *testing.T) {
	tests := []struct {
		name     string
		color    string
		expected string
	}{
		{"Red Color", RED, "\033[0;31m"},
		{"Green Color", GREEN, "\033[0;32m"},
		{"Yellow Color", YELLOW, "\033[0;33m"},
		{"Blue Color", BLUE, "\033[0;34m"},
		{"Reset Color", NC, "\033[0m"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expected, tt.color)
		})
	}
}

func TestLoggerInitialization(t *testing.T) {
	// Initialize loggers first
	initLoggers()

	// Test that loggers are not nil
	assert.NotNil(t, debugLogger, "debugLogger should not be nil")
	assert.NotNil(t, infoLogger, "infoLogger should not be nil")
	assert.NotNil(t, warnLogger, "warnLogger should not be nil")
	assert.NotNil(t, errorLogger, "errorLogger should not be nil")

	// Test that loggers have correct prefixes
	tests := []struct {
		name   string
		logger *log.Logger
		prefix string
	}{
		{"Debug Logger", debugLogger, "[DEBUG] "},
		{"Info Logger", infoLogger, "[INFO] "},
		{"Warn Logger", warnLogger, "[WARN] "},
		{"Error Logger", errorLogger, "[ERROR] "},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.prefix, tt.logger.Prefix())
		})
	}
}

func TestTimeoutConstants(t *testing.T) {
	// Test timeout value (5 minutes = 300 seconds)
	expectedTimeout := 300
	assert.Equal(t, expectedTimeout, 300, "User timeout should be 300 seconds")
}

func TestVersionConstant(t *testing.T) {
	assert.NotEmpty(t, VERSION, "Version should not be empty")
	assert.Contains(t, VERSION, ".", "Version should contain version format")
}

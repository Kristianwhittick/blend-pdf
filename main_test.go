package main

import (
	"path/filepath"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestExecuteUserChoice(t *testing.T) {
	tempDir := t.TempDir()
	
	// Set up directories
	err := setupDirectories(tempDir)
	assert.NoError(t, err)
	
	tests := []struct {
		name     string
		choice   string
		expected bool // true if should continue, false if should exit
	}{
		{"Single file option", "S", true},
		{"Single file option lowercase", "s", true},
		{"Merge files option", "M", true},
		{"Merge files option lowercase", "m", true},
		{"Help option", "H", true},
		{"Help option lowercase", "h", true},
		{"Verbose toggle", "V", true},
		{"Verbose toggle lowercase", "v", true},
		{"Debug toggle", "D", true},
		{"Debug toggle lowercase", "d", true},
		{"Quit option", "Q", false},
		{"Quit option lowercase", "q", false},
		{"Invalid option", "X", true},
		{"Empty input", "", true},
		{"Multiple characters", "ABC", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Reset counters for each test
			COUNTER = 0
			ERROR_COUNT = 0
			
			// Test that executeUserChoice doesn't panic
			assert.NotPanics(t, func() {
				executeUserChoice(tt.choice)
			})
		})
	}
}

func TestToggleVerboseMode(t *testing.T) {
	// Save original state
	originalVerbose := VERBOSE
	defer func() { VERBOSE = originalVerbose }()
	
	// Test toggling from false to true
	VERBOSE = false
	toggleVerboseMode()
	assert.True(t, VERBOSE, "Should toggle to true")
	
	// Test toggling from true to false
	toggleVerboseMode()
	assert.False(t, VERBOSE, "Should toggle to false")
}

func TestToggleDebugMode(t *testing.T) {
	// Save original state
	originalDebug := DEBUG
	originalVerbose := VERBOSE
	defer func() { 
		DEBUG = originalDebug
		VERBOSE = originalVerbose
	}()
	
	// Test enabling debug mode
	DEBUG = false
	VERBOSE = false
	toggleDebugMode()
	assert.True(t, DEBUG, "Should enable debug mode")
	assert.True(t, VERBOSE, "Should enable verbose mode when debug is enabled")
	
	// Test disabling debug mode
	toggleDebugMode()
	assert.False(t, DEBUG, "Should disable debug mode")
	// Note: VERBOSE might remain true, which is acceptable behavior
}

func TestShowStatistics(t *testing.T) {
	// Save original values
	originalSuccess := COUNTER
	originalErrors := ERROR_COUNT
	originalStart := START_TIME
	
	defer func() {
		COUNTER = originalSuccess
		ERROR_COUNT = originalErrors
		START_TIME = originalStart
	}()
	
	// Set test values
	COUNTER = 5
	ERROR_COUNT = 2
	START_TIME = time.Now().Add(-2 * time.Minute)
	
	// Test that function doesn't panic
	assert.NotPanics(t, func() {
		showStatistics()
	})
}

func TestShowHelp(t *testing.T) {
	// Test that help display doesn't panic
	assert.NotPanics(t, func() {
		showHelp()
	})
}

func TestShowVersion(t *testing.T) {
	// Test that version display doesn't panic
	assert.NotPanics(t, func() {
		showVersion()
	})
}

func TestProcessSingleFileOperation(t *testing.T) {
	tempDir := t.TempDir()
	
	// Set up directories
	err := setupDirectories(tempDir)
	assert.NoError(t, err)
	
	// Test that function doesn't panic with no files
	assert.NotPanics(t, func() {
		processSingleFileOperation()
	})
}

func TestProcessMergeOperation(t *testing.T) {
	tempDir := t.TempDir()
	
	// Set up directories
	err := setupDirectories(tempDir)
	assert.NoError(t, err)
	
	// Test that function doesn't panic with no files
	assert.NotPanics(t, func() {
		processMergeOperation()
	})
}

func TestDisplayApplicationStatus(t *testing.T) {
	tempDir := t.TempDir()
	
	// Set up directories
	err := setupDirectories(tempDir)
	assert.NoError(t, err)
	
	// Test that menu display doesn't panic
	assert.NotPanics(t, func() {
		displayApplicationStatus()
	})
}

func TestProcessMenuInput_EdgeCases(t *testing.T) {
	tests := []struct {
		name  string
		input string
	}{
		{"Whitespace input", "  "},
		{"Tab input", "\t"},
		{"Newline input", "\n"},
		{"Mixed whitespace", " \t\n "},
		{"Special characters", "!@#$%"},
		{"Numbers", "123"},
		{"Long input", strings.Repeat("A", 100)},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tempDir := t.TempDir()
			err := setupDirectories(tempDir)
			assert.NoError(t, err)
			
			// Test that invalid inputs are handled gracefully
			assert.NotPanics(t, func() {
				executeUserChoice(tt.input)
			})
		})
	}
}

func TestSessionTimingCalculation(t *testing.T) {
	// Test elapsed time calculation
	testStart := time.Now().Add(-5 * time.Minute)
	
	// Save original start time
	originalStart := START_TIME
	defer func() { START_TIME = originalStart }()
	
	START_TIME = testStart
	
	// Calculate elapsed time (approximate)
	elapsed := time.Since(START_TIME)
	assert.True(t, elapsed >= 4*time.Minute, "Elapsed time should be at least 4 minutes")
	assert.True(t, elapsed <= 6*time.Minute, "Elapsed time should be at most 6 minutes")
}

func TestGlobalVariableInitialization(t *testing.T) {
	// Test that global variables are properly initialized
	assert.NotNil(t, START_TIME, "START_TIME should be initialized")
	assert.GreaterOrEqual(t, COUNTER, 0, "COUNTER should be non-negative")
	assert.GreaterOrEqual(t, ERROR_COUNT, 0, "ERROR_COUNT should be non-negative")
}

func TestMenuChoiceValidation(t *testing.T) {
	validChoices := []string{"S", "s", "M", "m", "H", "h", "V", "v", "D", "d", "Q", "q"}
	invalidChoices := []string{"A", "B", "C", "1", "2", "3", "!", "@", "#"}
	
	tempDir := t.TempDir()
	err := setupDirectories(tempDir)
	assert.NoError(t, err)
	
	// Test valid choices don't panic
	for _, choice := range validChoices {
		t.Run("Valid choice: "+choice, func(t *testing.T) {
			assert.NotPanics(t, func() {
				executeUserChoice(choice)
			})
		})
	}
	
	// Test invalid choices don't panic
	for _, choice := range invalidChoices {
		t.Run("Invalid choice: "+choice, func(t *testing.T) {
			assert.NotPanics(t, func() {
				executeUserChoice(choice)
			})
		})
	}
}

// Integration-style tests for main workflow components

func TestMainWorkflowComponents(t *testing.T) {
	tempDir := t.TempDir()
	
	// Test complete setup workflow
	t.Run("Complete setup workflow", func(t *testing.T) {
		// Setup directories
		err := setupDirectories(tempDir)
		assert.NoError(t, err)
		
		// Verify setup is complete
		assert.DirExists(t, filepath.Join(tempDir, "archive"))
		assert.DirExists(t, filepath.Join(tempDir, "output"))
		assert.DirExists(t, filepath.Join(tempDir, "error"))
	})
}

func TestErrorRecoveryScenarios(t *testing.T) {
	tempDir := t.TempDir()
	
	tests := []struct {
		name string
		test func(t *testing.T)
	}{
		{
			name: "Handle missing directories gracefully",
			test: func(t *testing.T) {
				// Don't create directories first
				assert.NotPanics(t, func() {
					executeUserChoice("S")
				})
			},
		},
		{
			name: "Handle empty directory gracefully",
			test: func(t *testing.T) {
				err := setupDirectories(tempDir)
				assert.NoError(t, err)
				
				assert.NotPanics(t, func() {
					executeUserChoice("M")
				})
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, tt.test)
	}
}

// Benchmark tests for main workflow functions
func BenchmarkExecuteUserChoice(b *testing.B) {
	tempDir := b.TempDir()
	err := setupDirectories(tempDir)
	if err != nil {
		b.Fatal(err)
	}
	
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		executeUserChoice("H") // Help is fast and doesn't modify state
	}
}

func BenchmarkDisplayApplicationStatus(b *testing.B) {
	tempDir := b.TempDir()
	err := setupDirectories(tempDir)
	if err != nil {
		b.Fatal(err)
	}
	
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		displayApplicationStatus()
	}
}

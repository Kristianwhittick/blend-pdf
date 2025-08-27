# BlendPDFGo Testing Guide

## Overview
This document provides comprehensive testing procedures for BlendPDFGo, a tool for merging PDF files with special handling for double-sided scanning workflows.

## Test Environment Setup

### Prerequisites
- Go 1.19+ installed
- pdfcpu library available
- poppler-utils installed (for pdftotext verification)
- Test PDF files: Doc_A.pdf (A1, A2, A3) and Doc_B.pdf (M, 9, f)

### Test Files Location
- Unit tests: `*_test.go` files alongside source code
- Integration tests: `tests/` directory
- API tests: `experiments/experiment01_pagecount.go` through `experiment16_final_memory_approach.go` (see api_experiments.md)
- Documentation: `docs/` folder
- Working directories: `archive/`, `output/`, `error/`

## Unit Testing Framework

### Go Testing Standards
BlendPDFGo uses Go's built-in testing framework with additional libraries for enhanced testing capabilities.

#### Testing Libraries
- **Built-in `testing`**: Core Go testing framework
- **`testify/assert`**: Enhanced assertions and test utilities
- **`testify/mock`**: Mock generation and verification
- **`testify/suite`**: Test suite organization

#### Test File Organization
```
blendpdfgo/
├── main.go
├── main_test.go           # Main function tests
├── constants.go
├── constants_test.go      # Constants and logger tests
├── setup.go
├── setup_test.go          # CLI parsing and setup tests
├── pdfops.go
├── pdfops_test.go         # PDF operations tests
├── fileops.go
├── fileops_test.go        # File operations tests
└── tests/
    ├── integration_test.go # Full workflow tests
    ├── testdata/          # Test PDF files
    └── helpers.go         # Test utilities
```

### Unit Test Categories

#### 1. Core Function Tests
- **PDF Operations**: Page counting, validation, merging, extraction
- **File Operations**: Moving, copying, directory management
- **CLI Parsing**: Argument validation, flag handling
- **Error Handling**: Graceful failure scenarios
- **Utility Functions**: String formatting, path handling

#### 2. Integration Tests
- **Complete Workflows**: Single file move, merge operations
- **Directory Management**: Archive, output, error folder handling
- **Lock File Management**: Multi-instance prevention
- **Session Statistics**: Operation counting, timing

#### 3. Mock Testing
- **File System Operations**: Mock file I/O for isolated testing
- **PDF Library Calls**: Mock pdfcpu API responses
- **System Calls**: Mock signal handling, process management

### Test Coverage Requirements
- **Minimum Coverage**: 90% code coverage across all packages
- **Critical Functions**: 100% coverage for PDF operations and file handling
- **Error Paths**: All error conditions must be tested
- **Edge Cases**: Boundary conditions and unusual inputs

### Unit Test Implementation Guide

#### Setting Up Test Environment
```bash
# Install testing dependencies
go mod tidy
go get github.com/stretchr/testify/assert
go get github.com/stretchr/testify/mock
go get github.com/stretchr/testify/suite

# Create test directory structure
mkdir -p tests/testdata
mkdir -p tests/mocks
```

#### Test File Structure
Each Go source file should have a corresponding test file:
- `main.go` → `main_test.go`
- `constants.go` → `constants_test.go`
- `setup.go` → `setup_test.go`
- `pdfops.go` → `pdfops_test.go`
- `fileops.go` → `fileops_test.go`

#### Test Naming Conventions
```go
// Function: validatePDF
func TestValidatePDF(t *testing.T) { ... }

// Function: validatePDF with error case
func TestValidatePDF_InvalidFile(t *testing.T) { ... }

// Function: validatePDF with edge case
func TestValidatePDF_EmptyFile(t *testing.T) { ... }
```

#### Test Categories to Implement

##### 1. PDF Operations Tests (`pdfops_test.go`)
```go
func TestGetPageCount_ValidPDF(t *testing.T)
func TestGetPageCount_InvalidPDF(t *testing.T)
func TestValidatePDF_ValidFile(t *testing.T)
func TestValidatePDF_CorruptedFile(t *testing.T)
func TestCreateInterleavedMerge_Success(t *testing.T)
func TestCreateInterleavedMerge_PageCountMismatch(t *testing.T)
func TestExtractPageInReverse_SinglePage(t *testing.T)
func TestExtractPageInReverse_MultiPage(t *testing.T)
```

##### 2. File Operations Tests (`fileops_test.go`)
```go
func TestMoveFileToDirectory_Success(t *testing.T)
func TestMoveFileToDirectory_PermissionDenied(t *testing.T)
func TestEnsureDirectoriesExist_NewDirectories(t *testing.T)
func TestEnsureDirectoriesExist_ExistingDirectories(t *testing.T)
func TestGetPDFFiles_MultipleFiles(t *testing.T)
func TestGetPDFFiles_NoFiles(t *testing.T)
func TestFormatFileSize_VariousSizes(t *testing.T)
```

##### 3. CLI and Setup Tests (`setup_test.go`)
```go
func TestParseArguments_ValidFlags(t *testing.T)
func TestParseArguments_InvalidFlags(t *testing.T)
func TestCreateLockFile_Success(t *testing.T)
func TestCreateLockFile_AlreadyExists(t *testing.T)
func TestGenerateDirectoryHash_ConsistentOutput(t *testing.T)
func TestValidateWatchDirectory_ValidPath(t *testing.T)
func TestValidateWatchDirectory_InvalidPath(t *testing.T)
```

##### 4. Main Function Tests (`main_test.go`)
```go
func TestProcessSingleFile_Success(t *testing.T)
func TestProcessSingleFile_NoFiles(t *testing.T)
func TestProcessMergeFiles_Success(t *testing.T)
func TestProcessMergeFiles_PageCountMismatch(t *testing.T)
func TestHandleMenuChoice_ValidOptions(t *testing.T)
func TestHandleMenuChoice_InvalidOptions(t *testing.T)
```

##### 5. Constants and Utilities Tests (`constants_test.go`)
```go
func TestLoggerInitialization(t *testing.T)
func TestExitCodes_ValidValues(t *testing.T)
func TestColorConstants_ValidANSI(t *testing.T)
```

#### Mock Testing Strategy
```go
// Example mock for file system operations
type MockFileSystem struct {
    mock.Mock
}

func (m *MockFileSystem) ReadFile(filename string) ([]byte, error) {
    args := m.Called(filename)
    return args.Get(0).([]byte), args.Error(1)
}

// Example test using mocks
func TestProcessPDF_WithMockFS(t *testing.T) {
    mockFS := new(MockFileSystem)
    mockFS.On("ReadFile", "test.pdf").Return([]byte("pdf content"), nil)
    
    // Test implementation
    result := processPDFWithFS(mockFS, "test.pdf")
    
    assert.NoError(t, result)
    mockFS.AssertExpectations(t)
}
```

#### Test Data Management
```go
// Create test PDF files for testing
func setupTestPDFs(t *testing.T) (string, string) {
    testDir := t.TempDir()
    
    // Create test PDF files
    docA := filepath.Join(testDir, "Doc_A.pdf")
    docB := filepath.Join(testDir, "Doc_B.pdf")
    
    // Copy test files or create minimal PDFs
    // Implementation depends on test requirements
    
    return docA, docB
}

// Cleanup test environment
func teardownTest(t *testing.T, testDir string) {
    os.RemoveAll(testDir)
}
```

#### Running Unit Tests
```bash
# Run all tests
go test ./...

# Run tests with coverage
go test -cover ./...

# Run tests with detailed coverage report
go test -coverprofile=coverage.out ./...
go tool cover -html=coverage.out

# Run specific test file
go test -v ./pdfops_test.go

# Run specific test function
go test -run TestValidatePDF ./...

# Run tests with race detection
go test -race ./...
```

#### Test Coverage Goals
- **Overall Coverage**: Minimum 90%
- **Critical Functions**: 100% coverage
  - `validatePDF()`
  - `getPageCount()`
  - `createInterleavedMerge()`
  - `moveFileToDirectory()`
  - `createLockFile()`
- **Error Handling**: All error paths tested
- **Edge Cases**: Boundary conditions covered

## Main Program Testing

### Build and Basic Tests
```bash
# Build the program
go build

# Test merge functionality
echo -e "M\nQ" | ./blendpdfgo

# Verify output
pdftotext output/Doc_A_Doc_B.pdf -
# Expected: A1, f, A2, 9, A3, M
```

### Interactive Menu Tests

#### Command Line Arguments
1. `./blendpdfgo -h` - Show help
2. `./blendpdfgo -v` - Show version  
3. `./blendpdfgo -V` - Enable verbose mode
4. `./blendpdfgo /path/to/folder` - Watch specific folder
5. `./blendpdfgo` - Watch current directory

#### Menu Options
- **S** - Move single PDF to output
- **M** - Merge two PDFs with interleaved pattern
- **H** - Show help information
- **V** - Toggle verbose mode
- **Q** - Quit program

#### Error Handling Tests
1. Run with no PDF files present
2. Run with only one PDF file
3. Run with invalid PDF files
4. Test with mismatched page counts
5. Test with insufficient permissions

## Comprehensive Test Plan (Based on Bash Version)

### Command Line Arguments Tests
1. Run with `-h` flag to verify help display
2. Run with `--help` flag to verify help display
3. Run with `-v` flag to verify version display (should show v1.0.1)
4. Run with `--version` flag to verify version display (should show v1.0.1)
5. Run with `-V` flag to verify verbose mode enabling
6. Run with `--verbose` flag to verify verbose mode enabling
7. Run with valid folder path
8. Run with non-existent folder path
9. Run with no arguments (should use current directory)
10. Run with multiple command line options combined (e.g., `-V /path/to/folder`)
11. Run with invalid command line options (should show error)

### Dependency Checks
12. Run with pdfcpu missing
13. Run with all dependencies present
14. Verify dependency error messages are clear and helpful

### Lock File Tests
15. Run while another instance is running (should exit with code 6)
16. Verify lock file is removed after normal exit
17. Verify lock file is removed after signal interruption
18. Test lock file creation in /tmp directory
19. Verify lock file prevents multiple instances correctly

### Directory Setup Tests
20. Run with existing folder that has no subdirectories
21. Run with existing folder that already has archive/output/error subdirectories
22. Verify permissions for creating subdirectories
23. Test with read-only parent directory (should fail gracefully)
24. Test with symbolic links to directories
25. Verify path resolution works correctly

### File Count Display Tests
26. Verify file counts are accurate for each directory (Main/Archive/Output/Error)
27. Test file count display with empty directories
28. Test file count display with files in all directories
29. Verify counts update correctly after operations
30. Test with non-PDF files in directories (should not be counted)

### File Preview Tests (Verbose Mode)
31. Test file preview with no PDF files (should not show preview)
32. Test file preview with 1-3 PDF files (should show all)
33. Test file preview with exactly 5 PDF files (should show all 5)
34. Test file preview with more than 5 PDF files (should show 5 + "... and X more")
35. Verify file sizes are displayed correctly in preview
36. Test file preview with files having special characters in names
37. Verify file preview only shows in verbose mode
38. Test file preview sorting (should be alphabetical)

### Single File Move Tests
39. Move a single PDF file when only one exists
40. Attempt to move when no PDF files exist
41. Move a single PDF file when multiple exist (should take first alphabetically)
42. Move an invalid PDF file (should fail validation and move to error)
43. Test file size display in verbose mode during single file move
44. Verify success counter increments correctly
45. Verify files move to correct output directory

### Merge Files Tests
46. Merge two PDF files with single-page second file (no reversal)
47. Merge two PDF files with multi-page second file (tests page reversal)
48. Attempt to merge when only one PDF file exists
49. Attempt to merge when no PDF files exist
50. Merge with invalid first PDF file
51. Merge with invalid second PDF file
52. Merge with successful pdfcpu execution
53. Merge with failed pdfcpu execution
54. Test page count detection using pdfcpu
55. Verify temporary reverse files are created and cleaned up
56. Test merge output filename format (file1-file2.pdf)
57. Verify original files move to archive after successful merge
58. Verify files move to error directory after failed merge

### Interactive Menu Tests
59. Test 'S' option (single file move)
60. Test 's' option (lowercase)
61. Test 'M' option (merge files)
62. Test 'm' option (lowercase)
63. Test 'H' option (help)
64. Test 'h' option (lowercase)
65. Test 'V' option (toggle verbose)
66. Test 'v' option (lowercase)
67. Test 'Q' option (quit)
68. Test 'q' option (lowercase)
69. Test invalid option (should show warning)
70. Test timeout (no input for 300 seconds, should exit with code 7)
71. Verify menu prompt format and colors
72. Test file count display before each menu prompt

### Verbose Mode Tests
73. Test command output with verbose mode enabled
74. Test command output with verbose mode disabled
75. Test toggling verbose mode multiple times
76. Verify verbose mode shows file sizes during operations
77. Verify verbose mode shows page counts and reversal information
78. Verify verbose mode shows detailed pdfcpu command output
79. Test file preview display in verbose mode
80. Verify verbose mode toggle messages (enabled/disabled)

### Session Statistics Tests
81. Verify statistics display on normal exit
82. Verify statistics display on signal interruption
83. Test successful operations counter accuracy
84. Test error counter accuracy
85. Test elapsed time calculation
86. Verify statistics format and colors
87. Test statistics with zero operations
88. Test statistics with mixed success/error operations

### Error Handling Tests
89. Test PDF validation with corrupted files
90. Test operations with insufficient disk space
91. Test operations with read-only files
92. Test operations with files being deleted during processing
93. Verify error messages are clear and helpful
94. Test graceful handling of pdfcpu failures
95. Verify proper cleanup on errors
96. Test error directory functionality

### Signal Handling Tests
97. Send SIGINT signal (Ctrl+C) - verify graceful shutdown
98. Send SIGTERM signal - verify graceful shutdown
99. Verify statistics display on signal interruption
100. Verify lock file cleanup on signal interruption
101. Test signal handling during file operations

### Temporary File Management Tests
102. Verify temporary reverse files are created during multi-page merges
103. Verify temporary files are cleaned up after successful operations
104. Verify temporary files are cleaned up after failed operations
105. Test temporary file naming convention (*-reverse.pdf)
106. Verify lock file creation and cleanup

### Color and Output Formatting Tests
107. Verify color codes work correctly in terminal
108. Test output formatting with different terminal widths
109. Verify color consistency across different message types
110. Test output readability in verbose vs normal mode

### Edge Cases and Stress Tests
111. Test with very large PDF files (>100MB)
112. Test with PDF files with special characters in names
113. Test with PDF files with spaces in names
114. Test with very long filenames
115. Test with read-only PDF files
116. Test with insufficient disk space scenarios
117. Test with full paths vs. relative paths
118. Test with symbolic links to PDF files
119. Test rapid successive operations
120. Test with corrupted PDF files
121. Test with password-protected PDFs
122. Test with zero-byte files
123. Test with non-PDF files with .pdf extension

### Performance Tests
124. Test processing time with multiple large files
125. Test memory usage during operations
126. Test concurrent access scenarios (multiple terminals)
127. Verify timeout functionality works correctly
128. Test responsiveness during large file operations

### Integration Tests
129. Test complete workflow: startup → single move → merge → quit
130. Test complete workflow with verbose mode throughout
131. Test error recovery scenarios
132. Test mixed success/failure operations in single session
133. Verify all directories are created and used correctly
134. Test session continuity across multiple operations

### Regression Tests
135. Verify all original functionality still works after updates
136. Test backward compatibility with existing file structures
137. Verify no performance degradation from new features
138. Test that new features don't interfere with existing operations

## Expected Merge Results

### Input Files
- **Doc_A.pdf**: A1, A2, A3 (pages 1, 2, 3)
- **Doc_B.pdf**: M, 9, f (pages 1, 2, 3)

### Expected Output
**Interleaved Pattern**: A1, f, A2, 9, A3, M

This represents:
- Doc1_Page1 (A1) + Doc2_Page3 (f)
- Doc1_Page2 (A2) + Doc2_Page2 (9)  
- Doc1_Page3 (A3) + Doc2_Page1 (M)

**Bug Fix Note (2025-08-27)**: Fixed page reversal logic where `api.TrimFile` with comma-separated selections was not reordering pages correctly. Now uses individual page extraction and manual merging for proper reversal.

### File Movement
- **Success**: Input files moved to `archive/`, output in `output/`
- **Error**: Input files moved to `error/` with error message

## Verification Commands

### Check Page Content
```bash
# Verify individual pages
pdftotext Doc_A.pdf -
pdftotext Doc_B.pdf -

# Verify merged result
pdftotext output/Doc_A_Doc_B.pdf -

# Check file locations
ls -la archive/
ls -la output/
ls -la error/
```

### Page Count Verification
```bash
# Using pdfinfo (if available)
pdfinfo Doc_A.pdf | grep Pages
pdfinfo Doc_B.pdf | grep Pages
pdfinfo output/Doc_A_Doc_B.pdf | grep Pages
```

## Troubleshooting

### Common Issues
1. **Lock file error**: Remove `/tmp/blendpdfgo.lock`
2. **EOF error**: Normal when using piped input
3. **Invalid PDF**: Check file integrity with `pdfinfo`
4. **Permission denied**: Check file/directory permissions

### Debug Mode
Run with verbose flag to see detailed output:
```bash
./blendpdfgo -V
```

## Test Checklist

### ✅ Core Functionality
- [ ] Page counting works correctly
- [ ] PDF validation works
- [ ] Single page extraction works
- [ ] Multiple page extraction works
- [ ] Reverse page extraction works
- [ ] Simple merge works
- [ ] Interleaved merge works correctly
- [ ] File movement to archive/output/error works

### ✅ User Interface
- [ ] Menu displays correctly
- [ ] All menu options work
- [ ] Invalid input handled gracefully
- [ ] EOF handled gracefully
- [ ] Help information displays
- [ ] Verbose mode toggles correctly

### ✅ Error Handling
- [ ] Missing files handled
- [ ] Invalid PDFs rejected
- [ ] Page count mismatches detected
- [ ] Insufficient permissions handled
- [ ] Lock file conflicts handled

## Performance Testing

### Large File Tests
Test with larger PDF files to ensure:
- Memory usage remains reasonable
- Processing time is acceptable
- Temporary files are cleaned up properly

### Stress Testing
- Multiple rapid operations
- Very large page counts
- Files with special characters
- Network-mounted directories

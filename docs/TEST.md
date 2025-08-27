# BlendPDFGo Testing Guide

## Overview
This document provides comprehensive testing procedures for BlendPDFGo, a tool for merging PDF files with special handling for double-sided scanning workflows.

## Test Environment Setup

### Prerequisites
- Go 1.19+ installed
- pdfcpu library available
- poppler-utils installed (for pdftotext verification)
- Test PDF files: Doc_A.pdf (A1, A2, A3) and Doc_B.pdf (M, 9, *)

### Test Files Location
- API tests: `tests/test01_pagecount.go` through `test08_interleaved.go`
- Documentation: `docs/` folder
- Working directories: `archive/`, `output/`, `error/`

## API Testing

### Running Individual API Tests
```bash
cd /home/kris/scan/blendpdfgo
go run tests/test01_pagecount.go    # Test page counting
go run tests/test02_validate.go     # Test PDF validation
go run tests/test03_extract.go      # Test single page extraction
go run tests/test04_extract_multi.go # Test multiple page extraction
go run tests/test05_reverse.go      # Test reverse page extraction
go run tests/test06_merge.go        # Test simple merge
go run tests/test07_page_merge.go   # Test individual page merge
go run tests/test08_interleaved.go  # Test complete interleaved pattern
```

### Expected API Test Results
- **Test 01**: Both PDFs should have 3 pages
- **Test 02**: Both PDFs should validate successfully
- **Test 03**: Extract A1 to single page PDF
- **Test 04**: Extract A1, A2 to two-page PDF
- **Test 05**: Extract pages in reverse order
- **Test 06**: Simple concatenation: A1, A2, A3, M, 9, *
- **Test 07**: Partial interleaved: A1, *, A2, 9
- **Test 08**: Full interleaved: A1, *, A2, 9, A3, M ✅

## Main Program Testing

### Build and Basic Tests
```bash
# Build the program
go build

# Test merge functionality
echo -e "M\nQ" | ./blendpdfgo

# Verify output
pdftotext output/Doc_A_Doc_B.pdf -
# Expected: A1, *, A2, 9, A3, M
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

## Expected Merge Results

### Input Files
- **Doc_A.pdf**: A1, A2, A3 (pages 1, 2, 3)
- **Doc_B.pdf**: M, 9, * (pages 1, 2, 3)

### Expected Output
**Interleaved Pattern**: A1, *, A2, 9, A3, M

This represents:
- Doc1_Page1 (A1) + Doc2_Page3 (*)
- Doc1_Page2 (A2) + Doc2_Page2 (9)  
- Doc1_Page3 (A3) + Doc2_Page1 (M)

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

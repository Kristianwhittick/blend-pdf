# BlendPDFGo Specification

## Overview
A tool for merging PDF files with special handling for double-sided scanning workflows.

## Core Requirements

### PDF Merging
- Merge exactly 2 PDF documents
- Both documents must have the exact same number of pages
- If page counts don't match, error out immediately and move files to error directory
- The second document's pages are in reverse order and need to be processed in reverse during merging

### Merging Pattern
- Use interleaved merging pattern: Doc1_Page1, Doc2_Page3, Doc1_Page2, Doc2_Page2, Doc1_Page3, Doc2_Page1
- Final output should alternate between pages from first document and corresponding reversed pages from second document
- The second document's pages are processed in reverse order (last page first, first page last)

### File Handling
- **Success**: Move both input files to `archive/` directory, place merged PDF in `output/` directory
- **Error**: Move both input files to `error/` directory if page counts don't match or processing fails
- **File Selection**: Automatically select the first two PDF files found in the directory (sorted alphabetically)

### Output Naming
- Combine both input filenames without "_merged" suffix
- Format: `FirstFileName_SecondFileName.pdf`
- Example: `Doc_A.pdf` + `Doc_B.pdf` → `Doc_A_Doc_B.pdf`

### Validation
- Validate that both files are valid PDF documents before processing
- Check exact page count match (no tolerance for differences)
- Provide clear error messages for validation failures

### User Interface
- Interactive command-line interface
- Option "M" for merge functionality
- Display file information including page counts during processing
- Show merging pattern information to user
- Provide verbose output when enabled

## Directory Structure
```
project/
├── archive/     # Successfully processed input files
├── output/      # Final merged PDF files
├── error/       # Files that couldn't be processed
└── [input PDFs] # Source PDF files to be processed
```

## Error Handling
- Page count mismatch: Move files to error/ with descriptive message
- Invalid PDF files: Stop processing and report error
- File I/O errors: Move files to error/ and report issue
- Merge operation failures: Move files to error/ and report failure

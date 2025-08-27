# PDFcpu API Knowledge Base

## Overview
Documentation of pdfcpu API functions based on experimental testing.

## API Functions Tested

### `api.PageCountFile(filename string) (int, error)`
- **Status**: ✅ TESTED & WORKING
- **Purpose**: Get the number of pages in a PDF file
- **Parameters**: 
  - `filename`: Path to PDF file
- **Returns**: Page count and error
- **Notes**: 
  - Works reliably with valid PDF files
  - Returns exact page count
  - Simple to use, no configuration needed

### `api.ReadContextFile(filename string) (*model.Context, error)`
- **Status**: ✅ TESTED & WORKING
- **Purpose**: Load PDF file into memory context
- **Parameters**:
  - `filename`: Path to PDF file
- **Returns**: Context object and error
- **Notes**: 
  - Loads entire PDF structure into memory
  - Context contains PageCount and other metadata
  - More reliable than ReadContext from bytes

### `api.WriteContextFile(ctx *model.Context, filename string) error`
- **Status**: ✅ TESTED & WORKING
- **Purpose**: Write context to PDF file
- **Parameters**:
  - `ctx`: Context object
  - `filename`: Output file path
- **Returns**: Error if failed
- **Notes**: 
  - Successfully writes context back to file
  - Preserves PDF structure and content

### `api.ReadContext(rs io.ReadSeeker, conf *model.Configuration) (*model.Context, error)`
- **Status**: ⚠️ TESTED & PARTIALLY WORKING
- **Purpose**: Load PDF from byte stream into context
- **Parameters**:
  - `rs`: ReadSeeker (use bytes.NewReader(pdfBytes))
  - `conf`: Configuration object
- **Returns**: Context object and error
- **Notes**: 
  - Works but may return PageCount = 0 for some PDFs
  - Less reliable than ReadContextFile
  - Use bytes.NewReader() to convert []byte to ReadSeeker

### `api.WriteContext(ctx *model.Context, w io.Writer) error`
- **Status**: ✅ TESTED & WORKING
- **Purpose**: Write context to byte stream
- **Parameters**:
  - `ctx`: Context object
  - `w`: Writer (use bytes.Buffer)
- **Returns**: Error if failed
- **Notes**: 
  - Successfully converts context to bytes
  - Use bytes.Buffer to capture output
  - Can round-trip with ReadContext

### `api.ValidateContext(ctx *model.Context) error`
- **Status**: ✅ TESTED & WORKING
- **Purpose**: Validate PDF context structure
- **Parameters**:
  - `ctx`: Context object
- **Returns**: Error if invalid, nil if valid
- **Notes**: 
  - Takes only context parameter (no configuration)
  - Validates PDF structure in memory

### `api.TrimFile(inFile, outFile string, selectedPages []string, conf *model.Configuration) error`
- **Status**: ⚠️ TESTED & PARTIALLY WORKING
- **Purpose**: Extract specific pages from PDF
- **Parameters**:
  - `inFile`: Input PDF file path
  - `outFile`: Output PDF file path
  - `selectedPages`: Page selection (from ParsePageSelection)
  - `conf`: Configuration object
- **Returns**: Error if failed
- **Notes**: 
  - Works for some pages but may fail with "cannot dereference pageNodeDict" error
  - Requires ParsePageSelection to format page numbers
  - File-based operation only

### `api.ParsePageSelection(pageStr string) ([]string, error)`
- **Status**: ✅ TESTED & WORKING
- **Purpose**: Parse page selection string into format for TrimFile
- **Parameters**:
  - `pageStr`: Page selection string (e.g., "1", "1-3", "1,3,5")
- **Returns**: Parsed page selection and error
- **Notes**: 
  - Required for TrimFile operations
  - Handles single pages, ranges, and lists

### `api.MergeCreateFile(inFiles []string, outFile string, divider bool, conf *model.Configuration) error`
- **Status**: ✅ TESTED & WORKING
- **Purpose**: Merge multiple PDF files into one
- **Parameters**:
  - `inFiles`: Array of input PDF file paths
  - `outFiles`: Output merged PDF file path
  - `divider`: Whether to add divider pages (use false)
  - `conf`: Configuration object
- **Returns**: Error if failed
- **Notes**: 
  - Works reliably for merging multiple PDFs
  - Files are merged in order provided
  - Set divider to false for seamless merge

## In-Memory Processing Approach

### What Works ✅
1. **Load PDFs into memory contexts** using `ReadContextFile`
2. **Validate page counts** using context.PageCount
3. **Keep PDF data as bytes** in memory for minimal I/O
4. **Write contexts to temporary files** for operations requiring file paths
5. **Extract pages using TrimFile** (with error handling for problematic pages)
6. **Merge extracted pages** using MergeCreateFile
7. **Clean up temporary files** after processing

### What Doesn't Work ❌
1. **Direct context-based page extraction** - No ExtractPages or TrimContext functions
2. **Direct context merging** - No MergeContext function
3. **ReadContext from bytes** - Unreliable, often returns 0 pages
4. **Some page extractions** - TrimFile fails on certain pages with validation errors

### Recommended Hybrid Approach ✅
```go
// 1. Load PDFs into memory contexts for validation
ctxA, err := api.ReadContextFile("doc_a.pdf")
ctxB, err := api.ReadContextFile("doc_b.pdf")

// 2. Validate in memory
if ctxA.PageCount != ctxB.PageCount {
    return errors.New("page count mismatch")
}

// 3. Use temporary files for page operations
tempDir := "temp_processing"
defer os.RemoveAll(tempDir)

// 4. Extract pages with error handling
for i := 1; i <= ctxA.PageCount; i++ {
    // Write context to temp file
    tempFile := fmt.Sprintf("%s/temp_%d.pdf", tempDir, i)
    api.WriteContextFile(ctxA, tempFile)
    
    // Extract page with TrimFile
    pageFile := fmt.Sprintf("%s/page_%d.pdf", tempDir, i)
    pageSelection, _ := api.ParsePageSelection(fmt.Sprintf("%d", i))
    err := api.TrimFile(tempFile, pageFile, pageSelection, conf)
    
    // Handle extraction errors gracefully
    if err != nil {
        log.Printf("Skipping page %d due to error: %v", i, err)
        continue
    }
    
    // Read extracted page into memory if needed
    pageBytes, _ := ioutil.ReadFile(pageFile)
    // Store or process pageBytes...
}

// 5. Merge using file-based operations
api.MergeCreateFile(pageFiles, outputFile, false, conf)
```

## Configuration Notes
- **Default Config**: `model.NewDefaultConfiguration()`
- **Required for**: TrimFile, MergeCreateFile, ReadContext, WriteContext
- **Not required for**: PageCountFile, ReadContextFile, WriteContextFile, ValidateContext

## Best Practices
1. **Use ReadContextFile** instead of ReadContext for reliability
2. **Handle TrimFile errors** gracefully - some pages may fail extraction
3. **Use temporary directories** for intermediate files
4. **Clean up temp files** with defer statements
5. **Validate contexts** before processing
6. **Keep original data in memory** as bytes for minimal I/O
7. **Use hybrid approach** - memory for validation, files for operations

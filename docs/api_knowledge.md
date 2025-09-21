# PDFcpu API Knowledge Base

## Overview
Documentation of pdfcpu API functions based on experimental testing.

## Related Documentation
- **[API Experiments Procedures](api_experiments_procedures.md)**: Step-by-step testing procedures
- **[Memory Processing Research](memory_processing_research.md)**: In-memory processing research and conclusions
- **[Testing Guide](testing.md)**: Comprehensive testing procedures for the application

## Breakthrough: Stream-Based APIs Discovered ✅

### Response from pdfcpu Maintainer (Issue #1219)
The maintainer pointed us to existing stream-based APIs that we had missed:

1. **extract_test.go: func TestExtractPagesLowLevel(t *testing.T)**
2. **merge_test.go: TestMergeRaw(t *testing.T)**  
3. **trim.go: func Trim(rs io.ReadSeeker, w io.Writer, selectedPages []string, conf *model.Configuration) error**

### Newly Discovered Stream-Based Functions ✅

#### `api.Trim(rs io.ReadSeeker, w io.Writer, selectedPages []string, conf *model.Configuration) error`
- **Status**: ✅ TESTED & WORKING (Experiment 29)
- **Purpose**: Extract/trim pages from PDF stream to output stream
- **Parameters**:
  - `rs`: Input PDF as ReadSeeker (use bytes.NewReader)
  - `w`: Output writer (use bytes.Buffer)
  - `selectedPages`: Page selection from ParsePageSelection
  - `conf`: Configuration object
- **Returns**: Error if failed
- **Notes**: 
  - TRUE in-memory processing - no temporary files!
  - Can reverse page order: "3,2,1" works correctly
  - Perfect for page extraction and reordering

#### `api.MergeRaw(rsc []io.ReadSeeker, w io.Writer, dividerPage bool, conf *model.Configuration) error`
- **Status**: ✅ TESTED & WORKING (Experiment 29)
- **Purpose**: Merge multiple PDF streams into single output stream
- **Parameters**:
  - `rsc`: Slice of ReadSeekers (multiple PDFs)
  - `w`: Output writer (use bytes.Buffer)
  - `dividerPage`: Whether to add divider pages (use false)
  - `conf`: Configuration object
- **Returns**: Error if failed
- **Notes**: 
  - TRUE in-memory merging - no temporary files!
  - Handles multiple PDFs in sequence
  - Perfect for combining extracted pages

#### `api.MergeCreateZip(rs1, rs2 io.ReadSeeker, w io.Writer, conf *model.Configuration) error`
- **Status**: ✅ TESTED & WORKING (Experiment 29)
- **Purpose**: Merge two PDF streams with interleaved (zip) pattern
- **Parameters**:
  - `rs1`: First PDF as ReadSeeker
  - `rs2`: Second PDF as ReadSeeker  
  - `w`: Output writer (use bytes.Buffer)
  - `conf`: Configuration object
- **Returns**: Error if failed
- **Notes**: 
  - TRUE in-memory interleaved merging!
  - Creates perfect zip pattern: A1, B1, A2, B2, A3, B3
  - Combined with Trim for complete solution

## Complete In-Memory Workflow ✅

### Perfect Solution Achieved
Using the discovered APIs, we can now achieve 100% in-memory processing:

```go
// Complete in-memory interleaved merge workflow
func createInterleavedMergeInMemory(file1, file2, output string) error {
    // 1. Load PDFs into memory
    bytes1, _ := os.ReadFile(file1)
    bytes2, _ := os.ReadFile(file2)
    
    // 2. Validate page counts (using file API for simplicity)
    pageCount1, _ := api.PageCountFile(file1)
    pageCount2, _ := api.PageCountFile(file2)
    if pageCount1 != pageCount2 {
        return errors.New("page count mismatch")
    }
    
    // 3. Reverse second document in memory
    reader2 := bytes.NewReader(bytes2)
    var reversedBuffer bytes.Buffer
    reversePages := "3,2,1" // For 3-page document
    reverseSelection, _ := api.ParsePageSelection(reversePages)
    err := api.Trim(reader2, &reversedBuffer, reverseSelection, conf)
    
    // 4. Zip merge for perfect interleaving
    reader1 := bytes.NewReader(bytes1)
    reversedReader := bytes.NewReader(reversedBuffer.Bytes())
    var finalBuffer bytes.Buffer
    err = api.MergeCreateZip(reader1, reversedReader, &finalBuffer, conf)
    
    // 5. Write final result
    return os.WriteFile(output, finalBuffer.Bytes(), 0644)
}
```

### Benefits of Stream-Based Approach ✅
- **Zero Temporary Files**: Complete in-memory processing
- **Better Performance**: No disk I/O during processing
- **Simpler Code**: Eliminates temp file management
- **More Reliable**: Fewer failure points
- **Memory Efficient**: Process only what's needed

### Comparison with Previous Approaches

#### Old Zip Merge Approach (Experiment 21)
- **Temp Files**: 1 (for reversed document)
- **API Calls**: 2 (CollectFile + MergeCreateZipFile)
- **Complexity**: Medium

#### New Stream-Based Approach (Experiment 29) ✅
- **Temp Files**: 0 (pure in-memory)
- **API Calls**: 2 (Trim + MergeCreateZip)
- **Complexity**: Low
- **Performance**: Best

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
  - **LIMITATION**: Comma-separated selections like "3,2,1" extract pages in document order, not specified order
  - **By Design**: The `trim` command automatically sorts pages (confirmed in pdfcpu issue #950)
  - **Workaround**: Use `api.CollectFile()` for order-preserving extraction or extract pages individually

### `api.CollectFile(inFile, outFile string, selectedPages []string, conf *model.Configuration) error`
- **Status**: ✅ TESTED & CONFIRMED WORKING
- **Purpose**: Extract specific pages from PDF while preserving specified order
- **Parameters**:
  - `inFile`: Input PDF file path
  - `outFile`: Output PDF file path
  - `selectedPages`: Page selection (from ParsePageSelection)
  - `conf`: Configuration object
- **Returns**: Error if failed
- **Notes**: 
  - **Preserves Order**: Unlike `TrimFile`, this maintains the specified page order
  - **Solution**: Use this instead of `TrimFile` for proper page reversal (3,2,1 → pages 3,2,1)
  - **Source**: pdfcpu maintainer recommendation from issue #950
  - **Tested**: Experiment 17 confirms function availability and identical signature to TrimFile
  - **Drop-in Replacement**: Can replace TrimFile calls with no parameter changes

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

### `api.MergeCreateZipFile(inFile1, inFile2, outFile string, conf *model.Configuration) error`
- **Status**: ✅ TESTED & WORKING
- **Purpose**: Merge two PDF files with interleaved (zip) pattern
- **Parameters**:
  - `inFile1`: First PDF file path
  - `inFile2`: Second PDF file path
  - `outFile`: Output merged PDF file path
  - `conf`: Configuration object
- **Returns**: Error if failed
- **Notes**: 
  - Creates interleaved pattern: Page1_File1, Page1_File2, Page2_File1, Page2_File2, etc.
  - Perfect for double-sided scanning workflows
  - Combines with `CollectFile` for complete interleaved solution
  - **Breakthrough Discovery**: Eliminates need for individual page extraction loops

### `api.MergeCreateZip(rs1, rs2 io.ReadSeeker, w io.Writer, conf *model.Configuration) error`
- **Status**: ✅ AVAILABLE (not tested)
- **Purpose**: Merge two PDF streams with interleaved (zip) pattern
- **Parameters**:
  - `rs1`: First PDF as ReadSeeker
  - `rs2`: Second PDF as ReadSeeker
  - `w`: Output writer
  - `conf`: Configuration object
- **Returns**: Error if failed
- **Notes**: 
  - Stream-based version of MergeCreateZipFile
  - Useful for in-memory processing

## Recommended Zip Merge Approach ✅

### Strategy
Use the elegant 2-step zip merge solution for perfect interleaved patterns with minimal temporary files.

### Implementation Pattern
```go
// 1. Validate page counts match
ctxA, _ := api.ReadContextFile("doc_a.pdf")
ctxB, _ := api.ReadContextFile("doc_b.pdf")
if ctxA.PageCount != ctxB.PageCount {
    return errors.New("page count mismatch")
}

// 2. Reverse second document using CollectFile
pageSelection, _ := api.ParsePageSelection("3,2,1") // For 3-page document
reversedFile := "temp-reversed.pdf"
err := api.CollectFile("doc_b.pdf", reversedFile, pageSelection, conf)
defer os.Remove(reversedFile)

// 3. Zip merge for perfect interleaving
err = api.MergeCreateZipFile("doc_a.pdf", reversedFile, "output.pdf", conf)

// Result: A1, f, A2, 9, A3, M (perfect interleaved pattern)
```

### Benefits of Zip Merge Approach ✅
- **Dramatic Simplification**: 2 API calls instead of 6+ individual page extractions
- **Reduced Temporary Files**: 1 temp file instead of 6+ temp files
- **Perfect Interleaving**: Native zip merge provides exact pattern needed
- **Better Performance**: Fewer I/O operations and API calls
- **True Solution**: Uses intended pdfcpu APIs, not workarounds

## In-Memory Processing Approach (Legacy)

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

## Experimental Testing Results

### Core API Experiments

#### Experiment 01: Page Count API ✅
- **Status**: ✅ COMPLETED
- **Goal**: Test `api.PageCountFile()` function
- **Result**: SUCCESS - Both files return 3 pages
- **File**: experiment01_pagecount.go

#### Experiment 02: PDF Validation API ✅
- **Status**: ✅ COMPLETED  
- **Goal**: Test `api.ValidateFile()` function
- **Result**: SUCCESS - Both PDFs validate with relaxed and strict modes
- **File**: experiment02_validate.go

#### Experiment 03: Extract Single Page ✅
- **Status**: ✅ COMPLETED
- **Goal**: Test `api.TrimFile()` to extract one page
- **Result**: SUCCESS - Page 1 extracted successfully
- **File**: experiment03_extract.go

#### Experiment 04: Extract Multiple Pages ✅
- **Status**: ✅ COMPLETED
- **Goal**: Test `api.TrimFile()` with page range
- **Result**: SUCCESS - Multiple pages extracted
- **File**: experiment04_extract_multi.go

#### Experiment 05: Extract Pages in Reverse Order ⚠️
- **Status**: ⚠️ LIMITATION DISCOVERED
- **Goal**: Test `api.TrimFile()` with reverse page selection
- **Result**: LIMITATION - Comma-separated selections extract pages in document order, not specified order
- **Workaround**: Extract pages individually and merge manually for proper reordering
- **File**: experiment05_reverse.go

#### Experiment 06: Simple Merge Two Files ✅
- **Status**: ✅ COMPLETED
- **Goal**: Test `api.MergeCreateFile()` basic functionality
- **Result**: SUCCESS - Simple merge works
- **File**: experiment06_merge.go

#### Experiment 07: Merge Individual Pages ✅
- **Status**: ✅ COMPLETED
- **Goal**: Test merging extracted individual pages
- **Result**: SUCCESS - Individual page merging works
- **File**: experiment07_page_merge.go

#### Experiment 08: Complete Interleaved Pattern ✅
- **Status**: ✅ COMPLETED
- **Goal**: Test full interleaved merge implementation
- **Result**: SUCCESS - Interleaved pattern works
- **File**: experiment08_interleaved.go

### Zip Merge Experiments

#### Experiment 20: Basic Zip Merge ✅
- **Status**: ✅ COMPLETED
- **Goal**: Test `api.MergeCreateZipFile()` basic functionality
- **Result**: SUCCESS - Perfect interleaved pattern: A1, M, A2, 9, A3, f
- **File**: experiment20_zip_merge_basic.go

#### Experiment 21: CollectFile + Zip Merge ✅
- **Status**: ✅ COMPLETED
- **Goal**: Test complete solution: CollectFile reversal + zip merge
- **Result**: SUCCESS - Perfect target pattern: A1, f, A2, 9, A3, M
- **File**: experiment21_collect_zip_merge.go

#### Experiment 22: Complete Zip Flow Validation ✅
- **Status**: ✅ COMPLETED
- **Goal**: Test complete workflow with validation and content verification
- **Result**: SUCCESS - Confirmed 2-step solution works perfectly
- **File**: experiment22_complete_zip_flow.go

### Memory Processing Experiments

#### Experiment 09: Memory Context Loading ✅
- **Status**: ✅ COMPLETED
- **Goal**: Test loading PDFs into memory contexts
- **Result**: SUCCESS - `api.ReadContextFile()` and `api.ValidateContext()` work
- **File**: experiment09_memory_context.go

#### Experiment 10: Memory Page Extraction ⚠️
- **Status**: ⚠️ PARTIALLY WORKING
- **Goal**: Test extracting pages using memory contexts
- **Result**: PARTIAL - Context operations work, but some page extractions fail
- **File**: experiment10_memory_extract_simple.go

#### Experiment 13: API Exploration ✅
- **Status**: ✅ COMPLETED
- **Goal**: Explore available in-memory APIs
- **Result**: SUCCESS - Identified working and non-working functions
- **File**: experiment13_api_exploration.go

#### Experiment 14: Working Memory Approach ⚠️
- **Status**: ⚠️ PARTIALLY WORKING
- **Goal**: Test pure in-memory processing
- **Result**: PARTIAL - ReadContext from bytes unreliable (returns 0 pages)
- **File**: experiment14_working_memory.go

#### Experiment 15: Hybrid Memory Approach ✅
- **Status**: ✅ COMPLETED
- **Goal**: Test hybrid approach (memory + minimal temp files)
- **Result**: SUCCESS - Works with some page extraction failures
- **File**: experiment15_hybrid_memory.go

#### Experiment 17: CollectFile API Availability ✅
- **Status**: ✅ COMPLETED
- **Goal**: Test `api.CollectFile()` function availability and signature
- **Result**: SUCCESS - CollectFile function exists with identical signature to TrimFile
- **File**: experiment17_collect.go

#### Experiment 18: CollectFile Strategy Analysis ✅
- **Status**: ✅ COMPLETED
- **Goal**: Analyse CollectFile-based interleaved merge strategy
- **Result**: SUCCESS - Confirmed drop-in replacement capability and implementation strategy
- **File**: experiment18_collect_interleaved.go

## Best Practices
1. **Use ReadContextFile** instead of ReadContext for reliability
2. **Handle TrimFile errors** gracefully - some pages may fail extraction
3. **Use temporary directories** for intermediate files
4. **Clean up temp files** with defer statements
5. **Validate contexts** before processing
6. **Keep original data in memory** as bytes for minimal I/O
7. **Use hybrid approach** - memory for validation, files for operations
8. **For page reordering** - extract pages individually and merge manually instead of comma-separated selections

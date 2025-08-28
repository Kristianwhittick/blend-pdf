# Feature Request: New API Functions for In-Memory PDF Processing

## Summary
Request for new context-based API functions to enable true in-memory PDF operations without requiring temporary files.

## Background
Currently, pdfcpu's API requires file paths for most operations (page extraction, merging, trimming), which limits applications that want to process PDFs entirely in memory. While `api.ReadContextFile()` and `api.WriteContext()` provide some in-memory capabilities, key operations still require file-based workarounds.

## Current Limitations

### Missing Context-Based Operations
The following operations currently require file paths and cannot be performed directly on `*model.Context` objects:

1. **Page Extraction**: No `api.ExtractPagesContext()` equivalent
2. **Page Trimming**: No `api.TrimContext()` equivalent  
3. **Context Merging**: No `api.MergeContexts()` equivalent

### Current Workaround Pattern
Applications must use a hybrid approach with temporary files:

```go
// Load PDF into memory context
ctx, err := api.ReadContextFile("input.pdf")

// But then must write back to file for operations
tempFile := "temp.pdf"
api.WriteContextFile(ctx, tempFile)
defer os.Remove(tempFile)

// Extract pages using file-based API
pageSelection, _ := api.ParsePageSelection("1,3,5")
err = api.TrimFile(tempFile, "output.pdf", pageSelection, conf)
```

## Proposed API Functions

### 1. Context-Based Page Extraction
```go
// Extract specific pages from context, returning new context
func ExtractPagesContext(ctx *model.Context, pages []int) (*model.Context, error)

// Alternative: Extract pages and return as separate contexts
func ExtractPagesContexts(ctx *model.Context, pages []int) ([]*model.Context, error)
```

### 2. Context-Based Trimming
```go
// Trim context to specific pages (equivalent to TrimFile)
func TrimContext(ctx *model.Context, pageSelection []string) (*model.Context, error)
```

### 3. Context Merging
```go
// Merge multiple contexts into single context
func MergeContexts(contexts []*model.Context, divider bool) (*model.Context, error)
```

## Use Cases

### 1. PDF Processing Applications
Applications that merge PDFs with specific page reordering patterns:
- Load two PDFs into memory contexts
- Extract individual pages from each context
- Merge pages in interleaved pattern with reordering
- Output final result without intermediate files

**Example**: Merging two 3-page documents where the second document's pages need to be reversed:
- Document A: pages A1, A2, A3
- Document B: pages B1, B2, B3 (needs reversal to B3, B2, B1)
- Desired output: A1, B3, A2, B2, A3, B1 (interleaved pattern)

### 2. Web Services
HTTP APIs that process PDFs without touching disk:
- Receive PDF as HTTP request body
- Process entirely in memory
- Return processed PDF in HTTP response

### 3. Microservices
Containerized services with minimal disk I/O:
- Process PDFs in memory for better performance
- Reduce temporary file management complexity
- Improve resource utilization

## Benefits

### Performance
- **Reduced I/O**: Eliminate temporary file creation/deletion
- **Better Memory Efficiency**: Keep data in memory throughout processing
- **Faster Processing**: Avoid disk write/read cycles

### Architecture
- **Cleaner Code**: Eliminate temporary file management
- **Better Error Handling**: Fewer failure points (no disk space issues)
- **Simplified Deployment**: Reduced temporary directory requirements

### Resource Usage
- **Lower Disk Usage**: No temporary files
- **Predictable Memory**: Clear memory usage patterns
- **Container Friendly**: Better for containerized environments

## Implementation Suggestions

### Approach 1: Context Manipulation
Build on existing context infrastructure:
- Leverage existing `*model.Context` structure
- Implement page manipulation at context level
- Maintain compatibility with existing file-based APIs

### Approach 2: Internal Refactoring
Extract core logic from file-based functions:
- Separate file I/O from PDF processing logic
- Create internal context-based implementations
- Expose both file-based and context-based APIs

## Real-World Example

### Current Implementation (Typical Application)
```go
// Must use 6+ temporary files for interleaved merge
func createInterleavedMerge(file1, file2, output string) error {
    // Load contexts for validation
    ctx1, _ := api.ReadContextFile(file1)
    ctx2, _ := api.ReadContextFile(file2)
    
    // But then create temporary files for each page
    var tempFiles []string
    for i := 1; i <= ctx1.PageCount; i++ {
        tempFile := fmt.Sprintf("temp_page_%d.pdf", i)
        // Extract page using file-based API
        api.TrimFile(file1, tempFile, pageSelection, conf)
        tempFiles = append(tempFiles, tempFile)
    }
    
    // Merge temporary files
    api.MergeCreateFile(tempFiles, output, false, conf)
    
    // Cleanup temporary files
    for _, file := range tempFiles {
        os.Remove(file)
    }
}
```

### Proposed Implementation
```go
// Pure in-memory processing
func createInterleavedMerge(file1, file2, output string) error {
    // Load contexts
    ctx1, _ := api.ReadContextFile(file1)
    ctx2, _ := api.ReadContextFile(file2)
    
    // Extract pages in memory
    var contexts []*model.Context
    for i := 1; i <= ctx1.PageCount; i++ {
        pageCtx1, _ := api.ExtractPagesContext(ctx1, []int{i})
        pageCtx2, _ := api.ExtractPagesContext(ctx2, []int{ctx2.PageCount-i+1})
        contexts = append(contexts, pageCtx1, pageCtx2)
    }
    
    // Merge contexts in memory
    mergedCtx, _ := api.MergeContexts(contexts, false)
    
    // Write final result
    api.WriteContextFile(mergedCtx, output)
}
```

## Compatibility

### Backward Compatibility
- Existing file-based APIs remain unchanged
- New context-based APIs are additive
- Applications can migrate incrementally

### Migration Path
1. **Phase 1**: Introduce context-based APIs alongside existing ones
2. **Phase 2**: Update documentation with in-memory examples

## Community Impact

### Existing Applications
- **PDF Merging Tools**: Eliminate 6+ temporary files per merge operation
- **Web Services**: Enable true stateless PDF processing
- **CLI Tools**: Provide memory-efficient processing options

### New Possibilities
- **Stream Processing**: Process PDFs in data pipelines
- **Cloud Functions**: Serverless PDF processing without disk
- **Mobile Apps**: In-memory processing for resource-constrained environments

## References

### Related Issues
- Memory usage concerns in various issues
- Performance optimization requests
- Temporary file management complexity

### Similar Libraries
Other PDF libraries provide in-memory processing capabilities, making this a competitive feature for pdfcpu adoption.

## Conclusion

Context-based PDF processing would significantly enhance pdfcpu's capabilities for modern applications requiring in-memory operations. The proposed APIs build naturally on existing context infrastructure while providing substantial benefits for performance, architecture, and resource usage.

This enhancement would position pdfcpu as a more complete solution for applications requiring efficient, memory-based PDF processing workflows.

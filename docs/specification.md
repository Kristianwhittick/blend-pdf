# BlendPDF Specification

## Overview
A tool for merging PDF files with special handling for double-sided scanning workflows.

---

## Section 1: Requirements

### R1. Core PDF Processing Requirements
- **R1.1** Merge exactly 2 PDF documents
- **R1.2** Both documents must have the exact same number of pages
- **R1.3** Error out immediately and move files to error directory if page counts don't match
- **R1.4** Process second document's pages in reverse order during merging
- **R1.5** Use interleaved merging pattern: Doc1_Page1, Doc2_PageLast, Doc1_Page2, Doc2_PageSecondLast, etc.
- **R1.6** Validate that both files are valid PDF documents before processing
- **R1.7** Provide clear error messages for validation failures

### R2. Smart Page Reversal Requirements
- **R2.1** Single-page second file: Direct merge (no reversal needed)
- **R2.2** Multi-page second file: Extract pages individually in reverse order, then merge
- **R2.3** Use separate `api.TrimFile` calls for each page to ensure proper ordering
- **R2.4** Clean up temporary extracted files after processing

### R3. File Handling Requirements
- **R3.1** Success: Move both input files to `archive/` directory
- **R3.2** Success: Place merged PDF in `output/` directory
- **R3.3** Error: Move both input files to `error/` directory if processing fails
- **R3.4** Automatically select the first two PDF files found in the directory (sorted alphabetically)
- **R3.5** Create archive/, output/, error/ directories if missing
- **R3.6** Validate directory permissions before operations
- **R3.7** Handle permission errors gracefully

### R4. Output Naming Requirements
- **R4.1** Combine both input filenames with hyphen separator
- **R4.2** Format: `FirstFileName-SecondFileName.pdf`
- **R4.3** Example: `Doc_A.pdf` + `Doc_B.pdf` → `Doc_A-Doc_B.pdf`

### R5. User Interface Requirements
- **R5.1** Full-screen terminal takeover with segmented layout
- **R5.2** Header section displaying watch directory and relative paths with integrated file counts
- **R5.3** File counts integrated into header: Watch directory shows main count, Archive/Output/Error show respective counts on right side
- **R5.4** Available PDFs section with file selection mode toggle
- **R5.5** Separated Recent Output section showing processed files with detailed operation information
- **R5.6** Actions bar with keyboard shortcuts: [S]ingle File [M]erge PDFs [T]oggle Mode [Q]uit
- **R5.7** Status line showing current state, operation counts, and error counts
- **R5.8** Progress bar replaces status line during operations
- **R5.9** Real-time updates without user input (file monitoring)
- **R5.10** Graceful fallback to basic interface on legacy terminals (PowerShell 5, CMD)
- **R5.11** Invalid choice handling with interface redraw to prevent stacking of error messages

### R5A. File Selection Mode Requirements
- **R5A.1** Alpha Order Mode: Automatically select first two PDFs alphabetically (current behaviour)
- **R5A.2** User Selection Mode: Arrow key navigation with visual selection indicators
- **R5A.3** Mode toggle with [T] key between Alpha Order and User Selection
- **R5A.4** Visual mode indicator: `[Mode: Alpha Order ▼]` or `[Mode: User Select ▼]`
- **R5A.5** In User Selection: Arrow keys navigate, Space/Enter to select files
- **R5A.6** Selected files marked with visual indicators (▶ or ✓)

### R5B. Layout Structure Requirements
- **R5B.1** Header: Application title, version, and directory paths with integrated file counts
- **R5B.2** Available PDFs section with mode selector and file list
- **R5B.3** Horizontal separator line
- **R5B.4** Recent Output section showing completed operations with single-line format
  - Single file operations: Show filename (e.g., "✅ [15:04:05] Single file move - Document.pdf")
  - Merge operations: Show input and output files (e.g., "✅ [15:04:05] Merge - DocA.pdf + DocB.pdf → DocA-DocB.pdf")
  - Error operations: Show clear warnings (e.g., "❌ [15:04:05] Warning: 2 PDF files required, found 1")
  - Display last 5 operations with timestamps and status icons
  - Single line per operation for compact display
- **R5B.5** Actions bar with keyboard shortcuts (persistent during operations)
- **R5B.6** Status/Progress section (2 lines: status line + progress line, progress overwrites status during operations)
- **R5B.7** Session timer display only on program exit

### R6. Verbose Mode Requirements
- **R6.1** File Preview: Show up to 5 PDF files with sizes when verbose enabled
- **R6.2** File Size Display: Show file sizes during operations
- **R6.3** Page Information: Display page counts and reversal details
- **R6.4** Command Output: Show detailed pdfcpu command execution
- **R6.5** Toggle Support: Allow runtime toggling of verbose mode

### R7. Output Formatting Requirements
- **R7.1** Red: Error messages and failures
- **R7.2** Green: Success messages and confirmations
- **R7.3** Yellow: Warnings and informational messages
- **R7.4** Blue: File names, paths, and data values
- **R7.5** No Colour: Default text and prompts

### R8. Session Statistics Requirements
- **R8.1** Track successful operations during session
- **R8.2** Track failed operations and errors
- **R8.3** Calculate elapsed time from start to exit
- **R8.4** Display comprehensive statistics on program exit

### R9. Command Line Interface Requirements
- **R9.1** `-h, --help`: Show comprehensive help information and exit
- **R9.2** `-v, --version`: Show version information and exit
- **R9.3** `-V, --verbose`: Enable verbose mode from startup
- **R9.4** `-D, --debug`: Enable debug mode (includes verbose + structured logging)
- **R9.5** `[folder]`: Specify folder path to watch (default: current directory)
- **R9.6** Support multiple flags together: `-V /path/to/folder`, `-D /path/to/folder`
- **R9.7** Validate folder paths and show clear error messages
- **R9.8** Handle both relative and absolute paths

### R10. Error Handling Requirements
- **R10.1** Validate PDF structure before processing
- **R10.2** Move invalid PDFs to error/ directory with descriptive messages
- **R10.3** Continue operation after individual file failures
- **R10.4** Exact page count match required between two PDFs
- **R10.5** Move files to error/ with clear mismatch message
- **R10.6** Display both page counts in error message
- **R10.7** Handle pdfcpu command failures gracefully
- **R10.8** Provide specific error messages for different failure types

### R11. Lock File Protection Requirements
- **R11.1** Create directory-specific lock file to prevent multiple instances in same folder
- **R11.2** Location: `/tmp/blendpdf-<8-char-hash>.lock` (Unix) or `<watch-folder>/blendpdf-<8-char-hash>.lock` (Windows)
- **R11.3** Hash generated from absolute watch directory path using MD5 (8 characters)
- **R11.4** Clean up lock file on normal exit and signal interruption
- **R11.5** Show clear error message if already running in same directory
- **R11.6** Allow multiple instances in different directories simultaneously

### R12. Signal Handling Requirements
- **R12.1** Handle SIGINT (Ctrl+C) and SIGTERM signals
- **R12.2** Display session statistics on interruption
- **R12.3** Clean up temporary files and lock file
- **R12.4** Proper cleanup of any in-progress operations

### R13. File Operations Requirements
- **R13.1** Process files in alphabetical order
- **R13.2** Select first two PDFs for merge operations
- **R13.3** Select first PDF for single file operations
- **R13.4** Create temporary reversed PDFs during multi-page merges
- **R13.5** Use naming convention: `originalname-reverse.pdf`
- **R13.6** Clean up temporary files after successful operations
- **R13.7** Clean up temporary files on failures and interruptions
- **R13.8** Display file sizes in human-readable format (KB, MB, GB)
- **R13.9** Show sizes during verbose operations
- **R13.10** Include size information in file previews

### R14. Performance Requirements
- **R14.1** Auto-exit after 5 minutes (300 seconds) of user inactivity
- **R14.2** Show timeout warning before exit
- **R14.3** Exit with specific timeout exit code (7)
- **R14.4** Efficient handling of large PDF files
- **R14.5** Minimal memory footprint during operations
- **R14.6** Proper cleanup of resources
- **R14.7** Optimize for common use cases (small to medium PDFs)
- **R14.8** Provide progress indicators for large file operations
- **R14.9** Performance monitoring in debug mode

### R15. Exit Code Requirements
- **R15.1** `0`: Success
- **R15.2** `1`: General error
- **R15.3** `2`: Missing dependencies
- **R15.4** `3`: Invalid directory
- **R15.5** `4`: Invalid PDF file
- **R15.6** `5`: Merge operation failed
- **R15.7** `6`: Already running (lock file exists)
- **R15.8** `7`: User timeout

### R16. Compatibility Requirements
- **R16.1** Support standard PDF formats
- **R16.2** Handle password-protected PDFs gracefully
- **R16.3** Validate PDF structure before processing
- **R16.4** Cross-platform compatibility (Linux, macOS, Windows)
- **R16.5** Single binary deployment
- **R16.6** No external dependencies beyond pdfcpu
- **R16.7** Go 1.24 or higher
- **R16.8** pdfcpu library (latest stable version)
- **R16.9** Standard library dependencies only

---
---

## Section 2: Implementation Details

### Directory Structure
```
project/
├── archive/     # Successfully processed input files
├── output/      # Final merged PDF files and single file moves
├── error/       # Files that couldn't be processed
└── [input PDFs] # Source PDF files to be processed
```

### Command Line Interface Examples
```
BlendPDFGo v1.0.0 - A tool for merging PDF files

Usage: blendpdf [options] [folder]

Command line options:
  -h, --help     Show this help information and exit
  -v, --version  Show version information and exit
  -V, --verbose  Enable verbose mode (show all program output)
  -D, --debug    Enable debug mode (includes verbose + structured logging)
  [folder]       Specify folder to watch (default: current directory)

Examples:
  blendpdf -h                # Show help
  blendpdf -v                # Show version
  blendpdf -V                # Run in verbose mode
  blendpdf -D                # Run in debug mode
  blendpdf /path/to/pdfs     # Watch specific folder
  blendpdf -V /path/to/pdfs  # Verbose mode with specific folder
  blendpdf                   # Watch current directory

Interactive options:
  S - Move a single PDF file to the output directory
  M - Merge two PDF files (first file + reversed second file)
  H - Show this help information
  V - Toggle verbose mode
  D - Toggle debug mode
  Q - Quit the program
```

### Expected Merge Results
**Input Files**:
- Doc_A.pdf: A1, A2, A3 (pages 1, 2, 3)
- Doc_B.pdf: M, 9, f (pages 1, 2, 3)

**Expected Output**:
Interleaved Pattern: A1, f, A2, 9, A3, M

This represents:
- Doc1_Page1 (A1) + Doc2_Page3 (f)
- Doc1_Page2 (A2) + Doc2_Page2 (9)  
- Doc1_Page3 (A3) + Doc2_Page1 (M)

### File Movement behaviour
- **Success**: Input files moved to `archive/`, output in `output/`
- **Error**: Input files moved to `error/` with error message

---

## Section 3: Implementation Status

### ✅ All Requirements Implemented
All requirements R1.1 through R16.9 have been successfully implemented and tested.

### Current Status: PRODUCTION READY
- Complete feature parity with original bash version achieved
- Professional user interface with comprehensive feedback
- Robust error handling and recovery mechanisms
- Structured logging and debug capabilities
- Multi-platform build system with automated releases

### Optional Enhancement Available
- **Phase 4: In-Memory Processing** - 52.9% memory efficiency improvement
- **Status**: Ready for implementation with complete research and test code
- **Priority**: Performance optimisation (optional)

---

## Section 4: Long-term Roadmap

### Maintenance
- Regular dependency updates
- Bug fixes and improvements
- Performance optimisations
- Security updates

### Potential Enhancements
- Configuration file support
- Batch processing capabilities
- Additional PDF manipulation features
- Integration with other tools

### Community
- User feedback incorporation
- Feature requests evaluation
- Documentation improvements
- Example usage scenarios

---

## Section 5: Design Decisions and Rejected Features

### Core Design Philosophy
BlendPDF is designed to solve the specific problem of non-duplex scanner workflows with minimal changes to PDF files. This focused approach drives many design decisions.

### Rejected Features and Rationale

#### Batch Processing Mode
**Rejected**: Single and merge files can be in any order, making batch processing impossible to implement reliably.

#### File Preview Integration (PDF Thumbnails)
**Rejected**: Not possible in CLI program. Would only work with window-based program.

#### Custom Page Ordering
**Rejected**: No added benefit. This program solves non-duplex scanner issue; other merge requirements can be solved with the pdfcpu program directly.

#### PDF Optimisation
**Rejected**: Outside of scope. This program is designed to make minimal changes to PDFs.

#### Multi-Document Merging (>2 PDFs)
**Rejected**: Outside of scope. This program solves non-duplex scanner issue specifically.

#### Watermark/Metadata Support
**Rejected**: Outside of scope. This program is designed to make minimal changes to PDFs.

#### Smart File Detection (Front/Back Patterns)
**Rejected**: Not possible as file names are generated by the scanner and do not include front/back indicators.

#### Operation Scheduling/Auto-Processing
**Rejected**: Batch processing is not possible as single and merge files can be in any order.

#### Progress Persistence/Session Recovery
**Rejected**: No benefit. Each operation is already atomic.

#### Parallel Processing
**Rejected**: Not doing batch operations, so parallel processing is not applicable.

#### File Encryption Support
**Rejected**: Outside scope. Program designed to make minimal changes to PDFs.

#### Digital Signature Preservation
**Rejected**: Out of scope, as this is not possible with non-duplex scanned documents.

#### Health Monitoring/System Metrics
**Rejected**: Not appropriate for interactive CLI tool. BlendPDF processes individual files with immediate user feedback. Health monitoring is designed for long-running services, batch processing systems, or unattended operations. Current session statistics already provide sufficient visibility for this application's scope.

# BlendPDFGo Specification

## Overview
A tool for merging PDF files with special handling for double-sided scanning workflows.

## Core Requirements

### PDF Merging
- Merge exactly 2 PDF documents
- Both documents must have the exact same number of pages
- If page counts don't match, error out immediately and move files to error directory
- The second document's pages are in reverse order and need to be processed in reverse during merging

### Smart Page Reversal Logic
- **Single-page second file**: Direct merge (no reversal needed)
- **Multi-page second file**: Create temporary reversed copy, then merge
- Use `pdfcpu collect -p pages,in,reverse,order` for page reversal
- Clean up temporary reversed files after processing

### Merging Pattern
- Use interleaved merging pattern: Doc1_Page1, Doc2_Page3, Doc1_Page2, Doc2_Page2, Doc1_Page3, Doc2_Page1
- Final output should alternate between pages from first document and corresponding reversed pages from second document
- The second document's pages are processed in reverse order (last page first, first page last)

### File Handling
- **Success**: Move both input files to `archive/` directory, place merged PDF in `output/` directory
- **Error**: Move both input files to `error/` directory if page counts don't match or processing fails
- **File Selection**: Automatically select the first two PDF files found in the directory (sorted alphabetically)

### Output Naming
- Combine both input filenames with hyphen separator
- Format: `FirstFileName-SecondFileName.pdf`
- Example: `Doc_A.pdf` + `Doc_B.pdf` → `Doc_A-Doc_B.pdf`

### Validation
- Validate that both files are valid PDF documents before processing
- Check exact page count match (no tolerance for differences)
- Provide clear error messages for validation failures

## User Interface Requirements

### File Count Display
- Show real-time file counts before each menu prompt
- Format: `Files: Main(X) Archive(Y) Output(Z) Error(W)`
- Update counts dynamically after each operation

### Verbose Mode Features
- **File Preview**: Show up to 5 PDF files with sizes when verbose enabled
- **File Size Display**: Show file sizes during operations
- **Page Information**: Display page counts and reversal details
- **Command Output**: Show detailed pdfcpu command execution
- **Toggle Support**: Allow runtime toggling of verbose mode

### Colored Output
- **Red**: Error messages and failures
- **Green**: Success messages and confirmations
- **Yellow**: Warnings and informational messages
- **Blue**: File names, paths, and data values
- **No Color**: Default text and prompts

### Session Statistics
- Track successful operations during session
- Track failed operations and errors
- Calculate elapsed time from start to exit
- Display comprehensive statistics on program exit

### Interactive Menu
- **S** - Move single PDF to output directory
- **M** - Merge two PDFs with smart reversal logic
- **H** - Show comprehensive help information
- **V** - Toggle verbose mode on/off
- **Q** - Quit program with statistics display

## Command Line Interface

### Flags and Arguments
- `-h, --help`: Show comprehensive help information and exit
- `-v, --version`: Show version information and exit
- `-V, --verbose`: Enable verbose mode from startup
- `[folder]`: Specify folder path to watch (default: current directory)

### Combined Options
- Support multiple flags together: `-V /path/to/folder`
- Validate folder paths and show clear error messages
- Handle both relative and absolute paths

### Help Output
```
BlendPDFGo v1.0.0 - A tool for merging PDF files

Usage: blendpdfgo [options] [folder]

Command line options:
  -h, --help     Show this help information and exit
  -v, --version  Show version information and exit
  -V, --verbose  Enable verbose mode (show all program output)
  [folder]       Specify folder to watch (default: current directory)

Examples:
  blendpdfgo -h                # Show help
  blendpdfgo -v                # Show version
  blendpdfgo -V                # Run in verbose mode
  blendpdfgo /path/to/pdfs     # Watch specific folder
  blendpdfgo                   # Watch current directory

Interactive options:
  S - Move a single PDF file to the output directory
  M - Merge two PDF files (first file + reversed second file)
  H - Show this help information
  V - Toggle verbose mode
  Q - Quit the program
```

## Directory Structure
```
project/
├── archive/     # Successfully processed input files
├── output/      # Final merged PDF files and single file moves
├── error/       # Files that couldn't be processed
└── [input PDFs] # Source PDF files to be processed
```

## Error Handling

### PDF Validation
- Validate PDF structure before processing
- Move invalid PDFs to error/ directory with descriptive messages
- Continue operation after individual file failures

### Page Count Validation
- Exact page count match required between two PDFs
- Move files to error/ with clear mismatch message
- Display both page counts in error message

### Processing Failures
- Handle pdfcpu command failures gracefully
- Move files to error/ directory on merge failures
- Provide specific error messages for different failure types

### Lock File Protection
- Create lock file to prevent multiple instances
- Location: `/tmp/blendpdfgo.lock`
- Clean up lock file on normal exit and signal interruption
- Show clear error message if already running

### Graceful Shutdown
- Handle SIGINT (Ctrl+C) and SIGTERM signals
- Display session statistics on interruption
- Clean up temporary files and lock file
- Proper cleanup of any in-progress operations

## File Operations

### Automatic Directory Creation
- Create archive/, output/, error/ directories if missing
- Validate directory permissions before operations
- Handle permission errors gracefully

### File Sorting and Selection
- Process files in alphabetical order
- Select first two PDFs for merge operations
- Select first PDF for single file operations

### Temporary File Management
- Create temporary reversed PDFs during multi-page merges
- Use naming convention: `originalname-reverse.pdf`
- Clean up temporary files after successful operations
- Clean up temporary files on failures and interruptions

### File Size Reporting
- Display file sizes in human-readable format (KB, MB, GB)
- Show sizes during verbose operations
- Include size information in file previews

## Performance Requirements

### Timeout Protection
- Auto-exit after 5 minutes (300 seconds) of user inactivity
- Show timeout warning before exit
- Exit with specific timeout exit code

### Memory Management
- Efficient handling of large PDF files
- Minimal memory footprint during operations
- Proper cleanup of resources

### Processing Speed
- Optimize for common use cases (small to medium PDFs)
- Provide progress indicators for large file operations
- Minimize disk I/O where possible

## Exit Codes
- `0`: Success
- `1`: General error
- `2`: Missing dependencies
- `3`: Invalid directory
- `4`: Invalid PDF file
- `5`: Merge operation failed
- `6`: Already running (lock file exists)
- `7`: User timeout

## Compatibility

### PDF Support
- Support standard PDF formats
- Handle password-protected PDFs gracefully
- Validate PDF structure before processing

### Platform Support
- Cross-platform compatibility (Linux, macOS, Windows)
- Single binary deployment
- No external dependencies beyond pdfcpu

### Version Requirements
- Go 1.19 or higher
- pdfcpu library (latest stable version)
- Standard library dependencies only

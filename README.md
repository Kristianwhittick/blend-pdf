# BlendPDFGo


## Purpose

Turn your single-sided scanner into a double-sided scanning powerhouse!

BlendPDFGo was created from a real-world frustration: digitizing double-sided documents when all you have is a single-sided scanner.
If you’ve ever wasted time manually reordering and merging pages, this tool is built for you.

### The Problem

When scanning double-sided documents with a single-sided scanner, you usually end up with two files:

**Front PDF** → pages 1, 3, 5, 7 … (in correct order)

**Back PDF** → pages 8, 6, 4, 2 … (in reverse order)

Combining them into a clean, properly ordered document is tedious, error-prone, and easy to get wrong.

### The Solution

BlendPDFGo does the heavy lifting automatically:

✅ Detects and validates your scanned PDFs

✅ Reverses back-side pages when needed

✅ Merges both files into a perfectly ordered PDF (1, 2, 3, 4 …)

✅ Archives processed files and handles errors gracefully

✅ Shows progress and useful stats in real time

Who It’s For

🏠 Home users digitizing bills, letters, or family documents

🏢 Small offices that don’t have costly duplex scanners

🎓 Students scanning textbooks and research material

📚 Anyone who values speed and accuracy in document management

Why You’ll Love It

No more manual page shuffling.
No more mistakes.
No more wasted time.

Just scan, run BlendPDFGo, and enjoy perfectly ordered documents every single time.


## History

A Go-based tool for merging and managing PDF files with automatic page reversal and file organization. This is a modern reimplementation of the original bash-based BlendPDF tool with enhanced performance and reliability.

## Features

### Core Functionality ✅
- **Single File Processing**: Move individual PDF files to output directory
- **Smart PDF Merging**: Merge two PDFs with automatic page reversal of the second file
- **Intelligent Page Handling**: Only reverse multi-page PDFs (single-page PDFs merge directly)
- **File Validation**: Automatic PDF validation with comprehensive error handling
- **Directory Organization**: Automatic creation and management of archive, output, and error folders

### User Interface ✅
- **File Count Display**: Real-time display of PDF counts in each directory
- **Verbose Mode**: Detailed output including file sizes and processing information
- **Colored Output**: Color-coded messages for better readability
- **File Preview**: Show available PDF files with sizes in verbose mode
- **Session Statistics**: Track successful operations, errors, and elapsed time
- **Progress Indicators**: Detailed operation progress in verbose mode

### Advanced Features ✅
- **Lock File Protection**: Prevents multiple instances from running simultaneously
- **Graceful Shutdown**: Proper cleanup and statistics display on exit
- **Timeout Protection**: Automatic exit after period of inactivity (5 minutes)
- **Command Line Interface**: Comprehensive CLI flags and arguments
- **Debug Mode**: Structured logging with performance monitoring

## Current Status

**Production Ready** - All core features implemented and tested.

### ✅ Implemented Features
- Interactive command-line menu with enhanced UI
- Smart PDF merging with interleaved pattern
- File movement to archive/output/error directories
- Page count validation and smart reversal logic
- Comprehensive error handling and recovery
- Session statistics and performance monitoring
- Lock file protection and timeout handling
- Debug mode with structured logging

### 🔄 Optional Enhancement
- **In-Memory Processing**: Hybrid approach for performance optimization (See [docs/backlog.md](docs/backlog.md))

## Installation

### Prerequisites
- Go 1.19 or higher
- pdfcpu library (automatically installed via Go modules)

### Build from Source
```bash
# Clone the repository
git clone https://github.com/Kristianwhittick/blend_pdf.git
cd blend_pdf

# Build the application
go build

# Run the application
./blendpdfgo
```

## Usage

### Basic Usage
```bash
# Run in current directory
./blendpdfgo

# Show help
./blendpdfgo -h

# Run in verbose mode
./blendpdfgo -V

# Run in debug mode (includes verbose + structured logging)
./blendpdfgo -D

# Watch specific folder
./blendpdfgo /path/to/pdfs

# Combined options
./blendpdfgo -V /path/to/pdfs
```

### Interactive Menu Options

Once running, use these options:

- **S** - Move a single PDF file to the output directory
- **M** - Merge two PDF files (first file + reversed second file)
- **H** - Show help information
- **V** - Toggle verbose mode
- **D** - Toggle debug mode (structured logging)
- **Q** - Quit the program

### Expected Behavior

#### Single File Mode (S)
1. Finds the first PDF file in the main directory
2. Validates the PDF file structure
3. **Valid PDF**: Moves to `output/` directory
4. **Invalid PDF**: Moves to `error/` directory

#### Merge Mode (M)
1. Finds the first two PDF files in the main directory (alphabetically sorted)
2. Validates both PDF files and checks page count match
3. **Smart Processing**:
   - **Single-page second file**: Direct merge (no reversal)
   - **Multi-page second file**: Creates temporary reversed copy, then merges
4. Creates merged file: `file1-file2.pdf` in `output/`
5. **Success**: Moves original files to `archive/`
6. **Failure**: Moves original files to `error/`

## Directory Structure

The tool automatically creates and manages these directories:

```
your-folder/
├── file1.pdf          # Input files
├── file2.pdf
├── archive/            # Successfully processed files
├── output/             # Single files and merged results
└── error/              # Invalid or problematic files
```

## User Interface Examples

### File Count Display
```
Files: Main(2) Archive(0) Output(0) Error(0)
Enter choice (S/M/H/V/D/Q): 
```

### Verbose Mode Output
```
Files: Main(6) Archive(0) Output(0) Error(0)
Available PDF files:
  document1.pdf (2.3M)
  document2.pdf (1.8M)
  report.pdf (4.1M)
  scan001.pdf (856K)
  manual.pdf (12M)
  ... and 1 more file(s)

Enter choice (S/M/H/V/D/Q): M
Merging: document1.pdf document2.pdf -> document1-document2.pdf
File 1 size: 2.3M
File 2 size: 1.8M
pages = 5
Success: Files merged and moved. (3)
```

### Session Statistics
```
Session Statistics:
Successful operations: 5
Errors encountered: 1
Time elapsed: 2m 34s
```

### Debug Mode Output
```
[DEBUG] Processing merge: document1.pdf + document2.pdf
[INFO] OPERATION: MERGE | Files: document1.pdf, document2.pdf | Result: SUCCESS
[INFO] PERFORMANCE: MERGE | Duration: 1.2s | Size: 4194304 bytes | Speed: 3.33 MB/s
```

## Development

### Project Structure
```
blend_pdf/
├── main.go              # Main application entry point
├── constants.go         # Application constants and loggers
├── setup.go            # Initialization and CLI parsing
├── pdfops.go           # PDF processing operations
├── fileops.go          # File management operations
├── docs/               # Comprehensive documentation
│   ├── backlog.md      # Development roadmap (Phase 4 remaining)
│   ├── COMMIT_PLAN.md  # Git history and workflow
│   ├── api_knowledge.md # pdfcpu API research
│   ├── specification.md # Project requirements
│   └── TEST.md         # Testing procedures
└── tests/              # API research and test programs
```

### Development Status

**All Phases Complete** except optional performance enhancement:

1. **Phase 1**: ✅ Core Functionality Parity
   - Enhanced UI with file counts and colors
   - Smart page reversal logic
   - Comprehensive error handling

2. **Phase 2**: ✅ Interface and Management
   - Command line interface enhancements
   - Session management and statistics
   - Advanced file operations with timeout

3. **Phase 3**: ✅ Polish and Enhancement
   - Structured logging and debug mode
   - Performance monitoring and optimization

4. **Phase 4**: 🔄 Optional Performance Enhancement
   - Hybrid in-memory processing (52.9% memory efficiency)

### Research and Documentation

This project includes extensive research and documentation:

- **API Knowledge Base**: Complete pdfcpu API function reference
- **Memory Processing Research**: 8 test programs exploring in-memory processing
- **Bash Version Analysis**: Comprehensive feature comparison and gap analysis
- **Test Plan**: 140+ test cases for full functionality validation

## Dependencies

- **[pdfcpu](https://github.com/pdfcpu/pdfcpu)**: PDF processor and toolkit
- **Go Standard Library**: File operations, CLI handling, etc.

## Comparison with Original Bash Version

This Go implementation provides **complete feature parity** with the original bash version plus enhancements:

### Advantages of Go Version
- **Better Performance**: Compiled binary vs interpreted script
- **Enhanced UI**: File counts, colors, verbose mode, debug mode
- **Cross-Platform**: Single binary deployment
- **Better Error Handling**: Structured error management with recovery
- **Maintainability**: Strongly typed, well-structured code
- **Advanced Features**: Lock files, timeout protection, performance monitoring

### Feature Parity Achieved ✅
- All interactive menu options
- Smart page reversal logic
- Comprehensive file validation
- Session statistics and tracking
- Colored output and verbose mode
- Lock file protection
- Graceful shutdown handling
- Timeout protection

## Contributing

1. Fork the repository
2. Create a feature branch (`git checkout -b feature/amazing-feature`)
3. Make your changes following the development phases in [docs/backlog.md](docs/backlog.md)
4. Test thoroughly using the procedures in [docs/TEST.md](docs/TEST.md)
5. Commit your changes (`git commit -m 'Add amazing feature'`)
6. Push to the branch (`git push origin feature/amazing-feature`)
7. Open a Pull Request

## License

This project is licensed under the Apache License 2.0.  
See the [LICENSE](LICENSE) file for details.

## Acknowledgments

- Original bash implementation provided the foundation and requirements
- [pdfcpu](https://github.com/pdfcpu/pdfcpu) library for robust PDF processing
- Extensive API research and testing informed the hybrid memory approach
# BlendPDFGo

A Go-based tool for merging and managing PDF files with automatic page reversal and file organization. This is a modern reimplementation of the original bash-based BlendPDF tool with enhanced performance and reliability.

## Features

### Core Functionality
- **Single File Processing**: Move individual PDF files to output directory
- **Smart PDF Merging**: Merge two PDFs with automatic page reversal of the second file
- **Intelligent Page Handling**: Only reverse multi-page PDFs (single-page PDFs merge directly)
- **File Validation**: Automatic PDF validation with comprehensive error handling
- **Directory Organization**: Automatic creation and management of archive, output, and error folders

### User Interface (Planned)
- **File Count Display**: Real-time display of PDF counts in each directory
- **Verbose Mode**: Detailed output including file sizes and processing information
- **Colored Output**: Color-coded messages for better readability
- **File Preview**: Show available PDF files with sizes in verbose mode
- **Session Statistics**: Track successful operations, errors, and elapsed time
- **Progress Indicators**: Detailed operation progress in verbose mode

### Advanced Features (Planned)
- **Lock File Protection**: Prevents multiple instances from running simultaneously
- **Graceful Shutdown**: Proper cleanup and statistics display on exit
- **Timeout Protection**: Automatic exit after period of inactivity
- **Command Line Interface**: Comprehensive CLI flags and arguments
- **Memory Optimization**: Hybrid in-memory processing for better performance

## Current Status

This is an active development project. The current Go implementation provides basic functionality with the following features:

### ✅ Implemented
- Interactive command-line menu
- Basic PDF merging with interleaved pattern
- File movement to archive/output/error directories
- Page count validation
- Basic error handling

### 🔄 In Development (See [docs/backlog.md](docs/backlog.md))
- Enhanced user interface with file counts and colors
- Smart page reversal logic (critical missing feature)
- Comprehensive error handling and validation
- Command line interface enhancements
- Session management and statistics

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

# Show help (planned)
./blendpdfgo -h

# Run in verbose mode (planned)
./blendpdfgo -V

# Watch specific folder (planned)
./blendpdfgo /path/to/pdfs
```

### Interactive Menu Options

Once running, use these options:

- **S** - Move a single PDF file to the output directory
- **M** - Merge two PDF files (first file + reversed second file)
- **H** - Show help information
- **V** - Toggle verbose mode (planned enhancement)
- **Q** - Quit the program

### Expected Behavior

#### Single File Mode (S)
1. Finds the first PDF file in the main directory
2. Validates the PDF file
3. **Valid PDF**: Moves to `output/` directory
4. **Invalid PDF**: Moves to `error/` directory

#### Merge Mode (M)
1. Finds the first two PDF files in the main directory (alphabetically sorted)
2. Validates both PDF files
3. **Smart Processing**:
   - **Single-page second file**: Direct merge (no reversal)
   - **Multi-page second file**: Creates temporary reversed copy, then merges
4. Creates merged file: `file1_file2.pdf` in `output/`
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

## Planned User Interface Enhancements

### File Count Display (Coming Soon)
```
Files: Main(2) Archive(0) Output(0) Error(0)
Enter choice (S/M/H/V/Q): 
```

### Verbose Mode Output (Coming Soon)
```
Files: Main(6) Archive(0) Output(0) Error(0)
Available PDF files:
  document1.pdf (2.3M)
  document2.pdf (1.8M)
  report.pdf (4.1M)
  scan001.pdf (856K)
  manual.pdf (12M)
  ... and 1 more file(s)

Enter choice (S/M/H/V/Q): M
Merging: document1.pdf document2.pdf -> document1_document2.pdf
File 1 size: 2.3M
File 2 size: 1.8M
Pages = 5
Success: Files merged and moved. (3)
```

### Session Statistics (Coming Soon)
```
Session Statistics:
Successful operations: 5
Errors encountered: 1
Time elapsed: 2m 34s
```

## Development

### Project Structure
```
blend_pdf/
├── main.go              # Main application entry point
├── constants.go         # Application constants
├── setup.go            # Initialization and setup
├── pdfops.go           # PDF processing operations
├── fileops.go          # File management operations
├── docs/               # Comprehensive documentation
│   ├── backlog.md      # Development roadmap
│   ├── api_knowledge.md # pdfcpu API research
│   ├── specification.md # Project requirements
│   └── TEST.md         # Testing procedures
└── tests/              # API research and test programs
```

### Development Phases

1. **Phase 1**: Core Functionality Parity (High Priority)
   - Enhanced UI with file counts and colors
   - Smart page reversal logic
   - Comprehensive error handling

2. **Phase 2**: Interface and Management (Medium Priority)
   - Command line interface enhancements
   - Session management and statistics
   - Advanced file operations

3. **Phase 3**: Polish and Enhancement (Low Priority)
   - Structured logging and formatting
   - Performance optimizations

4. **Phase 4**: Advanced Features (Performance)
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

This Go implementation aims to provide feature parity with the original bash version while adding:

### Advantages of Go Version
- **Better Performance**: Compiled binary vs interpreted script
- **Memory Efficiency**: Hybrid in-memory processing approach
- **Cross-Platform**: Single binary deployment
- **Better Error Handling**: Structured error management
- **Maintainability**: Strongly typed, well-structured code

### Feature Parity Goals
- All interactive menu options
- Smart page reversal logic
- Comprehensive file validation
- Session statistics and tracking
- Colored output and verbose mode
- Lock file protection
- Graceful shutdown handling

## Contributing

1. Fork the repository
2. Create a feature branch (`git checkout -b feature/amazing-feature`)
3. Make your changes following the development phases in [docs/backlog.md](docs/backlog.md)
4. Test thoroughly using the procedures in [docs/TEST.md](docs/TEST.md)
5. Commit your changes (`git commit -m 'Add amazing feature'`)
6. Push to the branch (`git push origin feature/amazing-feature`)
7. Open a Pull Request

## License

This project is open source. Feel free to modify and distribute.

## Acknowledgments

- Original bash implementation provided the foundation and requirements
- [pdfcpu](https://github.com/pdfcpu/pdfcpu) library for robust PDF processing
- Extensive API research and testing informed the hybrid memory approach
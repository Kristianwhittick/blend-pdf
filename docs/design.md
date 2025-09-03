# Technical Design - BlendPDFGo

## Architecture Overview
BlendPDFGo is a command-line tool built in Go that processes PDF files for double-sided scanning workflows. The architecture follows a modular design with clear separation of concerns for PDF processing, file management, user interface, and cross-platform compatibility.

## Technology Decisions

### Core Technologies
- **Language**: Go (Golang) 1.24+
- **PDF Processing**: pdfcpu library (latest stable)
- **CLI Framework**: Custom implementation with color support
- **Build System**: Go modules with cross-platform builds
- **Testing**: Go testing framework with comprehensive coverage

### Key Libraries/Dependencies
- **pdfcpu**: Primary PDF processing library for validation, page extraction, and merging
- **Standard Library**: File operations, signal handling, cross-platform compatibility
- **Color Output**: Custom terminal color implementation for user feedback

### Architecture Rationale
- **Single Binary**: No external dependencies for easy deployment
- **Cross-Platform**: Native Go compilation for Windows, macOS, Linux
- **Memory Efficient**: Streaming processing where possible to handle large files
- **Robust Error Handling**: Comprehensive error detection and recovery

---

## System Components

### Component 1: PDF Processing Engine
**Purpose**: Core PDF validation, page extraction, and merging operations
**Responsibilities**: 
- PDF file validation and structure checking
- Page count extraction and validation
- Intelligent page reversal for multi-page documents
- Interleaved merging pattern implementation
- Temporary file management and cleanup

**Interfaces**: 
- `validatePDF(filename string) error`
- `getPageCount(filename string) (int, error)`
- `createInterleavedMerge(file1, file2, output string) error`
- `extractPagesInReverse(filename string) ([]string, error)`

### Component 2: File Management System
**Purpose**: Handles file organization, directory management, and file operations
**Responsibilities**: 
- Automatic PDF file discovery and selection
- Directory structure creation (archive/, output/, error/)
- File movement based on processing outcomes
- Lock file management for concurrent usage prevention
- Temporary file cleanup and resource management

**Interfaces**: 
- `selectPDFFiles(directory string) ([]string, error)`
- `createDirectories(basePath string) error`
- `moveToArchive(files []string) error`
- `moveToError(files []string, reason string) error`
- `createLockFile(directory string) error`

### Component 3: User Interface Controller
**Purpose**: Manages command-line interface, user feedback, and session management
**Responsibilities**: 
- Command-line argument parsing and validation
- Real-time progress indicators and status updates
- Color-coded message display (success, warning, error, info)
- Session statistics tracking and reporting
- Keyboard shortcut handling and user interaction

**Interfaces**: 
- `parseArguments(args []string) (*Config, error)`
- `displayProgress(operation string, progress float64)`
- `showMessage(level LogLevel, message string)`
- `trackOperation(operation string, success bool, duration time.Duration)`
- `displaySessionStats(stats *SessionStats)`

### Component 4: Cross-Platform Compatibility Layer
**Purpose**: Handles platform-specific operations and ensures consistent behavior
**Responsibilities**: 
- File path handling across different operating systems
- Terminal capability detection and fallback options
- Signal handling for graceful shutdown
- Platform-specific optimizations and configurations

**Interfaces**: 
- `normalizePath(path string) string`
- `detectTerminalCapabilities() *TerminalInfo`
- `setupSignalHandlers(cleanup func())`
- `getPlatformTempDir() string`

---

## Data Design

### Data Models
```go
type PDFFile struct {
    Path      string
    Name      string
    Size      int64
    PageCount int
    Valid     bool
}

type ProcessingSession struct {
    StartTime     time.Time
    FilesProcessed int
    SuccessCount  int
    ErrorCount    int
    Operations    []Operation
}

type Operation struct {
    Type      string    // "merge", "validate", "move"
    Files     []string
    Result    string    // "success", "error"
    Duration  time.Duration
    Message   string
}

type Config struct {
    WatchDirectory string
    VerboseMode    bool
    DebugMode      bool
    OutputDir      string
    ArchiveDir     string
    ErrorDir       string
}
```

### Data Flow
1. **File Discovery**: Scan directory for PDF files, validate and collect metadata
2. **Processing Pipeline**: Validate → Extract → Merge → Organize
3. **Result Tracking**: Record operations, timing, and outcomes
4. **Cleanup**: Remove temporary files, update statistics, release resources

### Storage Strategy
- **Configuration**: JSON files for user preferences and settings
- **Temporary Files**: System temp directory with unique naming
- **Lock Files**: Directory-specific lock files with MD5 hash naming
- **Logs**: Structured logging with configurable levels (debug, info, warn, error)

---

## API Design

### Core Processing APIs
```go
// PDF Processing
func ValidatePDF(filename string) error
func GetPageCount(filename string) (int, error)
func CreateInterleavedMerge(file1, file2, output string) error
func ExtractPagesInReverse(filename string) ([]string, error)

// File Management
func SelectPDFFiles(directory string) ([]PDFFile, error)
func CreateDirectoryStructure(basePath string) error
func MoveFiles(files []string, destination string) error
func CleanupTempFiles(pattern string) error

// User Interface
func DisplayProgress(operation string, current, total int)
func ShowColorMessage(level LogLevel, format string, args ...interface{})
func GetUserInput(prompt string) (string, error)
func HandleKeyboardShortcuts() <-chan KeyEvent

// Configuration
func LoadConfig(path string) (*Config, error)
func SaveConfig(config *Config, path string) error
func ValidateConfig(config *Config) error
```

### Error Handling Strategy
- **Validation Errors**: Clear messages with suggested solutions
- **Processing Errors**: Graceful degradation with file safety
- **System Errors**: Platform-specific error handling and recovery
- **User Errors**: Helpful guidance and correction suggestions

---

## Performance Considerations

### Performance Requirements
- **File Size**: Support PDFs up to 1GB efficiently
- **Memory Usage**: Maximum 100MB RAM for typical operations
- **Processing Speed**: <5 seconds for documents under 50 pages
- **Startup Time**: <1 second application initialization
- **Concurrent Usage**: Multiple instances in different directories

### Optimization Strategy
- **Streaming Processing**: Process large files in chunks to minimize memory usage
- **Lazy Loading**: Load PDF metadata only when needed
- **Efficient Algorithms**: Optimized page extraction and merging algorithms
- **Resource Pooling**: Reuse temporary file handles and processing resources
- **Caching**: Cache frequently accessed file metadata

### Memory Management
```go
// Memory-efficient processing approach
func ProcessLargePDF(filename string) error {
    // Stream processing with limited memory footprint
    // Chunk-based page extraction
    // Immediate cleanup of intermediate results
}

// Resource management
defer func() {
    cleanupTempFiles()
    releaseMemory()
    closePDFHandles()
}()
```

---

## Security Design

### Security Requirements
- **File Safety**: No data loss during processing failures
- **Input Validation**: Comprehensive PDF structure validation
- **Path Security**: Prevention of directory traversal attacks
- **Resource Limits**: Protection against resource exhaustion

### Security Measures
- **Input Sanitization**: Validate all file paths and PDF structures
- **Safe File Operations**: Atomic file moves and rollback capabilities
- **Resource Limits**: Maximum file size and processing time limits
- **Error Isolation**: Contain errors to prevent system-wide issues

### Access Controls
- **File Permissions**: Respect system file permissions and ownership
- **Directory Access**: Validate read/write permissions before processing
- **Lock File Security**: Prevent lock file manipulation by other processes

---

## Deployment Architecture

### Environment Strategy
- **Development**: Local development with comprehensive testing
- **Testing**: Automated testing across multiple platforms and PDF types
- **Production**: Single binary deployment with configuration files

### Build Process
```bash
# Cross-platform builds
GOOS=windows GOARCH=amd64 go build -o blendpdf.exe
GOOS=darwin GOARCH=amd64 go build -o blendpdf-macos
GOOS=linux GOARCH=amd64 go build -o blendpdf-linux

# Release packaging
tar -czf blendpdf-v1.0.0-linux.tar.gz blendpdf-linux
zip blendpdf-v1.0.0-windows.zip blendpdf.exe
```

### Deployment Process
1. **Build Validation**: Automated testing on target platforms
2. **Performance Testing**: Benchmark testing with various PDF types
3. **Integration Testing**: End-to-end workflow validation
4. **Release Packaging**: Platform-specific binary packaging
5. **Distribution**: GitHub releases with checksums and signatures

---

## Monitoring and Logging

### Logging Strategy
```go
type LogLevel int

const (
    DEBUG LogLevel = iota
    INFO
    WARN
    ERROR
)

// Structured logging with context
func LogOperation(level LogLevel, operation string, context map[string]interface{}) {
    // Timestamp, level, operation, context details
    // Configurable output format (JSON, text)
    // File and console output options
}
```

### Monitoring Requirements
- **Performance Metrics**: Processing time, memory usage, file sizes
- **Error Tracking**: Error frequency, types, and resolution
- **Usage Statistics**: File types processed, success rates, user patterns
- **System Health**: Resource usage, lock file status, cleanup effectiveness

### Debug Mode Features
- **Verbose Output**: Detailed operation logging and progress tracking
- **Performance Monitoring**: Real-time resource usage and timing
- **State Inspection**: Internal state dumps and processing pipeline visibility
- **Error Analysis**: Detailed error context and stack traces

---

## Integration Points

### External Dependencies
- **pdfcpu Library**: PDF processing operations and validation
- **Operating System**: File system operations and signal handling
- **Terminal**: Color output and keyboard input handling

### Future Integration Points
- **Web Interface**: REST API endpoints for remote operation
- **Cloud Storage**: Integration with Google Drive, Dropbox, OneDrive
- **Notification Systems**: Email, webhook, and messaging integrations
- **Plugin System**: Extensible processing pipeline for custom operations

### API Compatibility
- **Backward Compatibility**: Maintain CLI interface stability
- **Configuration Migration**: Automatic config file version upgrades
- **Data Format Stability**: Consistent output file naming and organization

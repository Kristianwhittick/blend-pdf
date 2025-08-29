# BlendPDFGo Tasks

## 1. Development Process Summary

### Project Status: ✅ PRODUCTION READY
- All core functionality is complete and production-ready
- Complete feature parity with original bash version achieved
- Professional user interface with comprehensive feedback
- Robust error handling and recovery mechanisms
- Structured logging and debug capabilities

### Development Phases Completed
- **Phase 1**: Core Functionality Parity (Tasks 1-3)
- **Phase 2**: Interface and Management (Tasks 4-6) 
- **Phase 3**: Polish and Enhancement (Tasks 7-8)
- **Phase 5**: Production Enhancement (Tasks 10-11)

### Key Achievements
- Enhanced features beyond original bash version
- Debug mode with performance monitoring
- Structured logging with multiple levels
- Enhanced CLI with comprehensive options
- Timeout protection and lock file management
- Performance metrics and operation tracking
- Multi-platform build system with automated releases
- Directory-specific lock files for concurrent usage

### ✅ Completed Features Summary
- **File Count Display**: Real-time PDF counts in each directory
- **Colored Output**: Red/Green/Yellow/Blue message types
- **File Preview**: Shows up to 5 PDF files with sizes in verbose mode
- **Session Statistics**: Tracks operations, errors, and elapsed time
- **Smart Page Reversal**: Only reverses multi-page PDFs
- **Enhanced PDF Validation**: Comprehensive validation before processing
- **Lock File Protection**: Prevents multiple instances
- **Timeout Protection**: Auto-exit after 5 minutes of inactivity
- **Debug Mode**: Structured logging with performance monitoring
- **CLI Enhancements**: Complete command line interface
- **Error Recovery**: Graceful handling of all failure scenarios

### Implementation Status Overview
- **Feature Parity**: ✅ Complete with bash version
- **User Interface**: ✅ Professional with comprehensive feedback
- **Error Handling**: ✅ Robust with detailed logging
- **Performance**: ✅ Monitoring and optimization ready
- **Documentation**: ✅ Comprehensive with testing procedures
- **Testing**: ✅ API function testing (tests 01-16), core functionality validation
- **Build System**: ✅ Multi-platform builds with automated releases

### Development Guidelines
- Future development should maintain backward compatibility
- Comprehensive research and test code available for reference in `/experiments/`
- All API knowledge documented in `/docs/api_knowledge.md`
- Memory processing research documented in `/docs/memory_processing_research.md`
- Follow git workflow and commit conventions in `/docs/git_flow.md`

---

## 2. Completed Tasks (Git Commit Order)

### Task 1: Fix getPageCount Function
- **Status**: ✅ COMPLETED
- **Requirements**: R1.6 (PDF validation), R10.1 (PDF structure validation)
- **Description**: Fixed the `getPageCount` function to use correct pdfcpu API (`api.PageCountFile`)
- **Issue**: Was trying to use `api.PDFInfo` with incorrect parameters
- **Solution**: Replaced with `api.PageCountFile(file)` which directly returns page count

### Task 2: Implement Page Count Validation
- **Status**: ✅ COMPLETED
- **Requirements**: R1.2 (exact page count match), R1.3 (error handling), R10.4-R10.6 (page count validation)
- **Description**: Added exact page count validation between two PDFs
- **Implementation**: 
  - Get page counts for both files
  - Compare for exact match (no tolerance)
  - Move files to error/ directory if counts don't match
  - Display clear error messages

### Task 3: Fix Merging Logic for Interleaved Pattern
- **Status**: ✅ COMPLETED
- **Requirements**: R1.1 (merge 2 PDFs), R1.4-R1.5 (interleaved pattern), R2.1-R2.4 (smart page reversal)
- **Description**: Rewrote merging logic to create interleaved pattern (Doc1_Page1, Doc2_Page3, Doc1_Page2, Doc2_Page2, Doc1_Page3, Doc2_Page1)
- **Implementation**:
  - Created `createInterleavedMerge` function
  - Extract individual pages from both documents
  - Second document pages processed in reverse order
  - Merge pages in alternating pattern
  - Clean up temporary files
- **Bug Fix (2025-08-27)**: Fixed page reversal logic
  - **Issue**: `api.TrimFile` with comma-separated selection "3,2,1" was not reordering pages
  - **Solution**: Extract pages individually in reverse order and merge manually
  - **Result**: Correct interleaved pattern: A1, f, A2, 9, A3, M

### Task 4: Update Filename Generation
- **Status**: ✅ COMPLETED
- **Requirements**: R4.1-R4.3 (output naming requirements)
- **Description**: Updated output filename to combine both input names with hyphen separator
- **Format**: `FirstFileName-SecondFileName.pdf`
- **Example**: `Doc_A.pdf` + `Doc_B.pdf` → `Doc_A-Doc_B.pdf`

### Task 5: Auto-select First Two PDFs
- **Status**: ✅ COMPLETED
- **Requirements**: R3.4 (automatic file selection), R13.1-R13.3 (file sorting and selection)
- **Description**: Modified file selection to automatically pick first two PDF files
- **Implementation**: Sorts files alphabetically and selects first two

### Task 6: Enhanced User Interface and Display
- **Status**: ✅ COMPLETED - Commit: 61eb72c
- **Requirements**: R5.1-R5.3 (file count display), R6.1-R6.2 (verbose mode), R7.1-R7.5 (colored output), R8.1-R8.4 (session statistics)
- **Features Implemented**:
  - File count display: "Files: Main(X) Archive(Y) Output(Z) Error(W)"
  - File preview in verbose mode showing up to 5 files with sizes
  - Colored output with Red/Green/Yellow/Blue message types
  - Session statistics tracking operations, errors, and elapsed time
  - Human-readable file size display functions
  - Progress indicators in verbose mode

### Task 7: Smart PDF Processing with Page Reversal Logic
- **Status**: ✅ COMPLETED - Commit: a949421
- **Requirements**: R2.1-R2.4 (smart page reversal), R1.6-R1.7 (PDF validation), R13.4-R13.7 (temporary file management)
- **Features Implemented**:
  - Smart page reversal: only reverse multi-page PDFs
  - Enhanced PDF validation before processing
  - Temporary file management with proper cleanup
  - Merge mode selection using pdfcpu
  - Page count detection and validation

### Task 8: Robust Error Handling and File Management
- **Status**: ✅ COMPLETED - Commit: 8710720
- **Requirements**: R11.1-R11.6 (lock file protection), R10.1-R10.8 (error handling), R9.1-R9.8 (CLI parsing), R3.5-R3.7 (directory management)
- **Features Implemented**:
  - Lock file protection to prevent multiple instances
  - PDF validation with detailed error reporting
  - Enhanced command line argument parsing
  - Graceful failure recovery for individual operations
  - File conflict resolution with automatic renaming

### Task 9: Command Line Interface Enhancements
- **Status**: ✅ COMPLETED - Commit: 5c26188
- **Requirements**: R9.1-R9.8 (CLI interface requirements)
- **Features Implemented**:
  - Version display with -v/--version flag
  - Help display with -h/--help flag
  - Verbose flag with -V/--verbose
  - Debug flag with -D/--debug
  - Folder argument support
  - Combined options support

### Task 10: Session Management and Statistics
- **Status**: ✅ COMPLETED - Commit: 5c26188
- **Requirements**: R8.1-R8.4 (session statistics), R12.1-R12.4 (signal handling)
- **Features Implemented**:
  - Operation counter tracking successful operations
  - Error counter tracking failed operations
  - Session timer tracking elapsed time
  - Statistics display on program exit
  - Graceful shutdown with Ctrl+C handling

### Task 11: Advanced File Operations
- **Status**: ✅ COMPLETED - Commit: 5c26188
- **Requirements**: R3.5 (directory creation), R13.1-R13.10 (file operations), R14.1-R14.3 (timeout protection)
- **Features Implemented**:
  - Automatic directory creation
  - File sorting in alphabetical order
  - File size reporting in human-readable format
  - Timeout protection: Auto-exit after 5 minutes
  - Real-time file monitoring with dynamic counts

### Task 12: Output and Logging Improvements
- **Status**: ✅ COMPLETED - Commit: 9a94010
- **Requirements**: R9.4 (debug mode), R7.1-R7.5 (output formatting), R6.4 (command output)
- **Features Implemented**:
  - Structured logging with separate loggers for DEBUG/INFO/WARN/ERROR
  - Debug mode with comprehensive operation logging
  - Consistent message formatting across all operations
  - Interactive debug mode toggle
  - Enhanced help text with debug mode documentation

### Task 13: Performance and Reliability
- **Status**: ✅ COMPLETED - Commit: 9a94010
- **Requirements**: R14.4-R14.9 (performance requirements), R10.3 (error recovery)
- **Features Implemented**:
  - Large file handling with performance metrics
  - Memory usage monitoring through file size tracking
  - Operation logging for troubleshooting and analysis
  - Enhanced error recovery with detailed logging
  - Performance metrics showing duration, file size, and processing speed

### Task 15: Multi-Platform Build System
- **Status**: ✅ COMPLETED
- **Requirements**: R16.4-R16.5 (cross-platform compatibility, single binary deployment)
- **Priority**: Production Enhancement (Required)
- **Description**: Implement automated build system for multiple platforms and architectures
- **Platforms**: Windows (amd64), Linux (amd64), Linux (arm64), macOS (amd64)
- **Features Implemented**:
  - Cross-compilation build script for Windows, Linux, macOS
  - Comprehensive Makefile with build targets
  - GitHub Actions workflow for automated builds
  - Release automation with checksums and artifacts
  - Support individual and batch platform builds

### Task 16: Directory-Specific Lock Files
- **Status**: ✅ COMPLETED
- **Requirements**: R11.1-R11.6 (lock file protection requirements)
- **Priority**: Enhancement (User Requested)
- **Description**: Implement directory-specific lock files to allow multiple instances in different folders
- **Features Implemented**:
  - Hash-based lock file names using 8-character MD5 hash
  - Cross-platform lock file placement (tmp on Unix, watch folder on Windows)
  - Path normalization for case-insensitive filesystem compatibility
  - Allow multiple instances in different directories simultaneously
  - Maintain single-directory instance prevention

### Task 22: Code Refactoring Implementation
- **Status**: ✅ COMPLETED - Commit: 62fab9a
- **Requirements**: Code quality improvement (supports all requirements through better maintainability)
- **Priority**: Code Quality (Medium)
- **Description**: Refactor codebase for improved maintainability and clarity
- **Scope**: All code files (main.go, constants.go, setup.go, pdfops.go, fileops.go)
- **Refactoring Implemented**:
  - Break down large functions into smaller, focused functions
  - Apply single responsibility principle throughout codebase
  - Eliminate code duplication between similar functions
  - Improve naming conventions for better clarity
  - Enhance error handling and separation of concerns
- **Key Improvements**:
  - main.go: Split processMenu into multiple focused functions
  - fileops.go: Separate file operations, display, and logging functions
  - pdfops.go: Organize PDF operations by functionality (validation, reversal, merging)
  - setup.go: Break down argument parsing and lock file management
  - constants.go: Better organization and documentation of constants

### Task 18: Documentation Review and Cleanup
- **Status**: ✅ COMPLETED
- **Requirements**: Documentation quality improvement (supports all requirements through better documentation)
- **Priority**: Documentation Quality (Medium)
- **Description**: Review all documentation files for overlaps, duplicates, and clarity improvements
- **Scope**: All files in `/docs/` directory (treat README.md separately)
- **Analysis Completed**:
  - Analyzed all 10 documentation files for content overlaps (73,779 bytes total)
  - Identified critical overlaps in API documentation across 3 files
  - Found implementation status duplicated in 3 files
  - Discovered testing information scattered across multiple files
  - Identified 2 redundant files (`status.md`, `development_references.md`)
- **Deliverables**:
  - Complete analysis in `documentation_review_analysis.md`
  - Specific recommendations for Task 19 implementation
  - File consolidation and renaming proposals
  - Content reorganization strategy

### Task 17: Unit Testing Implementation
- **Status**: ✅ COMPLETED
- **Requirements**: Quality assurance for all requirements (R1.1-R16.9)
- **Priority**: Quality Assurance (High)
- **Description**: Implement comprehensive unit testing framework for Go code
- **Prerequisites**: Move experiment content from testing.md to api_experiments_procedures.md ✅
- **Implementation**: Applied Go built-in testing + Testify framework
- **Deliverables**:
  - Research document with framework comparison and recommendations
  - Complete unit test suite covering all main components
  - Integration test suite with workflow testing
  - Test helper utilities and mock strategies
  - 55.9% code coverage achieved as foundation (updated after test fixes)
  - Comprehensive testing documentation and procedures
- **Coverage Achieved**:
  - Constants Tests: Exit codes, colors, logger initialization ✅
  - Setup Tests: CLI parsing, lock file management, directory hashing ✅
  - File Operations Tests: Directory setup, file discovery, display functions ✅
  - PDF Operations Tests: Validation, page counting, merge operations ✅
  - Main Function Tests: Menu handling, session management, error recovery ✅
  - Integration Tests: Complete workflows, error handling, performance baselines ✅

### Task 19: Documentation Cleanup Implementation
- **Status**: ✅ COMPLETED
- **Requirements**: Documentation quality improvement (supports all requirements through better documentation)
- **Priority**: Documentation Quality (Medium)
- **Description**: Implement the recommendations from Task 18 documentation review
- **Prerequisites**: Complete Task 18 and review recommendations
- **Implementation**: Applied approved changes from documentation review analysis
- **Deliverables**:
  - Reorganized API documentation structure
  - Moved experiment summaries to api_knowledge.md
  - Streamlined api_experiments_procedures.md for testing procedures only
  - Updated all cross-references and file links
  - Added consistent navigation sections
  - Removed documentation_review_analysis.md after completion

---

## 3. Next Stages

### Task 26: Remove User Timeout
- **Status**: ✅ COMPLETED
- **Requirements**: R14.1-R14.3 (timeout protection removal)
- **Priority**: User Experience Enhancement (High)
- **Description**: Remove the 5-minute user inactivity timeout to allow indefinite operation
- **Implementation**: Remove timeout logic from main menu loop and signal handling
- **Files Modified**: main.go
- **Benefits**: Users can leave application running without forced exit
- **Changes Made**:
  - Removed timeout goroutine and channel logic from getUserChoice()
  - Simplified user input to direct ReadString() call
  - Eliminated waitForUserInput() and readUserInput() helper functions
  - Removed 5-minute timeout exit code (7)

### Task 20: CLI Library Research
- **Status**: ✅ COMPLETED
- **Requirements**: R5.4-R5.10 (interactive menu), R9.1-R9.8 (CLI interface), user experience enhancement
- **Priority**: User Experience Enhancement (Medium)
- **Description**: Research Go CLI libraries for enhanced user experience
- **Implementation Completed**:
  - Researched 6 major Go CLI libraries (Cobra, Bubble Tea, Survey, Termui, Progressbar, Fatih/Color)
  - Analyzed pros/cons and use cases for each library
  - Identified 3 implementation approaches with effort estimates
  - Created comprehensive findings document in `docs/cli_library_research.md`
  - Recommended phased approach: Survey + Progressbar + Fatih/Color for quick wins
- **Deliverables**:
  - ✅ Complete library comparison with maintenance status
  - ✅ Specific UI pattern examples and code snippets
  - ✅ Implementation priority recommendations
  - ✅ Risk assessment and effort estimates

### Task 21: UI Interface Recommendations
- **Status**: ✅ COMPLETED
- **Requirements**: R5.1-R5B.7 (full-screen UI), R5A.1-R5A.6 (file selection modes), user experience enhancement
- **Priority**: User Experience Enhancement (Medium)
- **Description**: Design enhanced full-screen UI interface to replace scrolling input
- **Prerequisites**: Complete Task 20 CLI library research ✅
- **Implementation Completed**:
  - Designed full-screen terminal takeover layout with segmented sections
  - Specified header with watch directory and relative folder paths
  - Created file selection mode toggle (Alpha Order vs User Selection)
  - Defined progress bar placement (replaces status line during operations)
  - Updated specification with detailed UI requirements (R5.1-R5B.7, R5A.1-R5A.6)
  - Recommended Bubble Tea + Lipgloss implementation with PowerShell 5 fallback
- **Layout Design**:
  - Header: Application title, version, directory paths
  - File counts: Real-time counts with session timer
  - Available PDFs: Mode toggle and file list with selection indicators
  - Recent Output: Separated section showing completed operations
  - Actions: Keyboard shortcuts including new [T]oggle mode
  - Status/Progress: Dynamic line for status and progress bars
- **Key Features**:
  - Real-time file monitoring and updates
  - Visual file selection with arrow keys
  - Mode switching between automatic and manual selection
  - Progress visualization during operations
  - Cross-platform compatibility with graceful fallback

### Task 23: GitHub Issue for pdfcpu Library
- **Status**: ❌ CANCELLED
- **Priority**: Community Contribution (Low)
- **Description**: ~~Create GitHub issue/feature request for pdfcpu page extraction ordering~~
- **Cancellation Reason**: Issue already reported and resolved in pdfcpu issue #950
- **Resolution**: pdfcpu maintainer confirmed `trim` command sorts pages by design; `collect` command preserves order
- **Action**: Use `api.CollectFile()` instead of `api.TrimFile()` for order-preserving page extraction

### Task 25: pdfcpu Feature Request for In-Memory Processing
- **Status**: ✅ COMPLETED
- **Priority**: Community Contribution (Low)
- **Description**: Create GitHub feature request for in-memory PDF processing capabilities in pdfcpu
- **Background**: Current pdfcpu API requires file-based operations, limiting true in-memory processing

#### Feature Request Details
- **Missing Functions**: Context-based page extraction and merging
- **Proposed APIs**:
  - `api.ExtractPagesContext(ctx *model.Context, pages []int) ([]*model.Context, error)`
  - `api.MergeContexts(contexts []*model.Context) (*model.Context, error)`
  - `api.TrimContext(ctx *model.Context, pages []string) (*model.Context, error)`
- **Use Case**: Enable zero-temp-file PDF processing for better performance and resource usage
- **Benefits**: Reduced disk I/O, better memory efficiency, cleaner application architecture

#### Implementation Completed
1. **Research existing issues** - ✅ Analyzed existing GitHub issues for similar requests
2. **Prepare technical specification** - ✅ Detailed proposed API functions with signatures
3. **Include use case examples** - ✅ BlendPDFGo use case with before/after code examples
4. **Reference current limitations** - ✅ Documented file-based workarounds currently needed
5. **Provide implementation suggestions** - ✅ Two implementation approaches suggested

#### Deliverables
- ✅ GitHub issue research and analysis completed
- ✅ Technical specification document created (`docs/pdfcpu_feature_request.md`)
- ✅ Feature request markdown ready for review
- ✅ GitHub issue submitted
- [ ] Follow-up on maintainer feedback (pending response)

#### Acceptance Criteria
- ✅ Comprehensive feature request document prepared
- ✅ Technical details clearly explained with examples
- ✅ Use case demonstrates real-world benefit (BlendPDFGo interleaved merge)
- ✅ Current limitations and workarounds documented
- ✅ Implementation suggestions provided
- ✅ GitHub issue submitted
- [ ] Maintainer feedback received and documented (pending response)

#### Research Results
- **Existing Issues**: Found several memory-related issues but none specifically requesting context-based APIs
- **Common Patterns**: Memory usage concerns, PageCount=0 issues with ReadContext, performance regressions
- **Opportunity**: No existing feature request for context-based processing APIs

### Task 24: Implement Zip Merge Solution for Interleaved Pattern
- **Status**: 📋 READY FOR IMPLEMENTATION
- **Requirements**: R2.1-R2.4 (smart page reversal), R1.4-R1.5 (interleaved pattern)
- **Priority**: Code Quality Enhancement (High) - **UPGRADED FROM MEDIUM**
- **Description**: Replace complex loop-based interleaving with simple 2-step zip merge solution
- **Background**: Discovered `MergeCreateZipFile` API provides perfect interleaved merging when combined with `CollectFile`

#### Implementation Details
- **Research Completed**: ✅ (Experiments 20-22 confirm zip merge solution)
- **API Functions**: `CollectFile` + `MergeCreateZipFile` 
- **Breakthrough**: Eliminates all individual page extraction and complex merging logic

#### Technical Requirements
1. **Replace `createInterleavedMerge()` function** with 2-step zip merge approach
2. **Step 1**: Use `CollectFile` to reverse second document pages
3. **Step 2**: Use `MergeCreateZipFile` to interleave first document with reversed second document
4. **Maintain existing error handling** and validation
5. **Update function comments** to reflect the elegant solution

#### Files to Modify
- `pdfops.go` - Replace `createInterleavedMerge()` function completely
- `docs/api_knowledge.md` - Add zip merge API documentation
- Tests - Update to reflect new implementation

#### Benefits
- **Dramatic Simplification**: 2 API calls instead of 6+ individual page extractions
- **Reduced Temporary Files**: 1 temp file instead of 6+ temp files
- **Perfect Interleaving**: Native zip merge provides exact pattern needed
- **Better Performance**: Fewer I/O operations and API calls
- **Cleaner Code**: Eliminates complex loop-based merging logic
- **True Solution**: Uses intended pdfcpu APIs, not workarounds

#### Implementation Strategy
```go
// OLD APPROACH (complex workaround):
func createInterleavedMerge(file1, file2, output string) error {
    // Extract pages individually in reverse order (loop)
    for i := pageCount; i >= 1; i-- {
        extractSinglePage(file2, i, conf) // 3+ API calls
    }
    // Extract pages individually from first file (loop)  
    for i := 1; i <= pageCount; i++ {
        extractSinglePage(file1, i, conf) // 3+ API calls
    }
    // Manually merge in interleaved pattern
    api.MergeCreateFile(interleavedFiles, output, false, conf)
    // Total: 6+ temp files, 7+ API calls
}

// NEW APPROACH (elegant solution):
func createInterleavedMerge(file1, file2, output string) error {
    // Step 1: Reverse second document
    pageSelection, _ := api.ParsePageSelection("3,2,1")
    reversedFile := "temp-reversed.pdf"
    err := api.CollectFile(file2, reversedFile, pageSelection, conf)
    defer os.Remove(reversedFile)
    
    // Step 2: Zip merge for perfect interleaving
    err = api.MergeCreateZipFile(file1, reversedFile, output, conf)
    
    // Total: 1 temp file, 2 API calls
    // Result: A1, f, A2, 9, A3, M (perfect interleaved pattern)
}
```

#### Important Notes
- **Major Breakthrough**: This is NOT a workaround - it's the proper solution using intended APIs
- **Temp Files**: Reduces from 6+ files to 1 temporary reversed file
- **Performance**: Dramatic improvement in speed and resource usage
- **API Discovery**: `MergeCreateZipFile` was the missing piece for elegant interleaving
- **Validation**: Experiments 20-22 confirm perfect interleaved output pattern

#### Acceptance Criteria
- [ ] Replace complex `createInterleavedMerge()` with 2-step zip merge
- [ ] Maintain exact same interleaved pattern output (A1, f, A2, 9, A3, M)
- [ ] Reduce temporary files from 6+ to 1
- [ ] Reduce API calls from 7+ to 2
- [ ] All existing tests pass with new implementation
- [ ] Performance improvement measurable
- [ ] Documentation updated to reflect elegant solution

#### Estimated Effort
- **Development**: 2-3 hours (complete rewrite of merge function)
- **Testing**: 1-2 hours (verify pattern and performance)
- **Documentation**: 1 hour

#### Suggested Commit Message
```
feat: Implement zip merge solution for interleaved PDF pattern

- Replace complex loop-based interleaving with elegant 2-step solution
- Use CollectFile + MergeCreateZipFile for perfect interleaved pattern
- Reduce from 6+ temp files to 1 temp file
- Reduce from 7+ API calls to 2 API calls  
- Maintain exact same output pattern (A1, f, A2, 9, A3, M)
- Major performance and code quality improvement
```

### Task 27: Full-Screen UI Implementation
- **Status**: 📋 READY FOR IMPLEMENTATION
- **Requirements**: R5.1-R5B.7 (full-screen UI), R5A.1-R5A.6 (file selection modes)
- **Priority**: User Experience Enhancement (High)
- **Description**: Implement full-screen terminal UI to replace current scrolling menu interface
- **Prerequisites**: Task 20 (CLI Library Research) ✅, Task 21 (UI Design) ✅

#### Implementation Details
- **Research Completed**: ✅ (Task 20 - CLI library analysis)
- **Design Completed**: ✅ (Task 21 - full-screen layout specification)
- **Recommended Stack**: Bubble Tea + Lipgloss with PowerShell 5 fallback
- **Cross-Platform**: Modern terminals (full UI) + legacy Windows (basic fallback)

#### Technical Requirements
1. **Terminal Takeover**: Full-screen interface replacing scrolling menu
2. **Segmented Layout**: Header, file counts, available PDFs, recent output, actions, status
3. **Real-Time Updates**: File monitoring without user input
4. **File Selection Modes**: Toggle between Alpha Order and User Selection
5. **Visual Indicators**: Arrow keys navigation, selection markers (▶, ✓)
6. **Progress Visualization**: Progress bars during operations
7. **Cross-Platform Compatibility**: Graceful fallback for legacy terminals

#### Files to Create/Modify
- `ui/` - New directory for UI components
- `ui/tui.go` - Main TUI implementation
- `ui/models.go` - Bubble Tea models
- `ui/views.go` - Layout and rendering
- `ui/fallback.go` - Legacy terminal fallback
- `main.go` - Integration with existing menu system
- `go.mod` - Add Bubble Tea and Lipgloss dependencies

#### Implementation Phases
1. **Phase 1**: Basic Bubble Tea setup with terminal detection
2. **Phase 2**: Layout implementation (header, sections, status)
3. **Phase 3**: File monitoring and real-time updates
4. **Phase 4**: File selection modes and navigation
5. **Phase 5**: Progress bars and operation feedback
6. **Phase 6**: Cross-platform testing and fallback

#### Benefits
- **Modern Interface**: Professional full-screen terminal UI
- **Better UX**: Real-time updates, visual selection, progress bars
- **File Selection**: Both automatic and manual selection modes
- **Cross-Platform**: Works on all target platforms with fallbacks

#### Acceptance Criteria
- [ ] Full-screen terminal takeover implemented
- [ ] All layout sections rendering correctly
- [ ] Real-time file monitoring working
- [ ] File selection modes functional (Alpha Order + User Selection)
- [ ] Progress bars during operations
- [ ] Cross-platform compatibility verified
- [ ] Legacy terminal fallback working
- [ ] All existing functionality preserved
- [ ] Tests updated for new UI components

#### Estimated Effort
- **Development**: 2-3 weeks
- **Testing**: 1 week
- **Documentation**: 2-3 days

#### Suggested Commit Message
```
feat: Implement full-screen terminal UI with Bubble Tea

- Add modern TUI replacing scrolling menu interface
- Implement real-time file monitoring and updates
- Add file selection modes (Alpha Order vs User Selection)
- Include progress bars and visual feedback
- Cross-platform support with legacy terminal fallback
- Maintain all existing functionality with enhanced UX
```

### Maintenance Activities
- Regular dependency updates
- Bug fixes and improvements as reported
- Security updates as needed
- Documentation improvements based on user feedback

---

## 4. Backlog Items (On Hold)

> **⚠️ IMPORTANT**: These backlog items are **NOT TO BE WORKED ON** until explicitly moved to "Next Stages" section by request. They are documented here for future reference only.

### Task 14: Implement In-Memory Processing Approach (Phase 4)
- **Status**: 🔄 READY FOR IMPLEMENTATION (ON HOLD)
- **Requirements**: R14.4-R14.6 (performance requirements), R2.1-R2.4 (smart page reversal), R13.4-R13.7 (temporary file management)
- **Priority**: Performance Enhancement (Optional)
- **Description**: Replace current file-based merging with hybrid in-memory approach to reduce temporary file usage
- **Benefits**: 
  - 52.9% memory efficiency vs original files
  - Reduced disk I/O operations
  - Better error handling for problematic PDF pages
  - Faster processing with minimal temporary files

#### Implementation Details
- **Research Completed**: ✅ (Experiments 09-16 in `/experiments/` directory)
- **API Knowledge**: ✅ (Documented in `/docs/api_knowledge.md`)
- **Approach Validated**: ✅ (Hybrid approach in `experiment16_final_memory_approach.go`)

#### Technical Requirements
1. **Load PDFs into memory** as byte arrays for validation
2. **Use `api.ReadContextFile()`** for reliable context creation
3. **Validate page counts** in memory before processing
4. **Extract pages with minimal temp files** using error handling
5. **Keep extracted pages in memory** as byte arrays
6. **Final merge from memory** with proper cleanup

#### Files to Modify
- `main.go` - Update merge logic in interactive menu
- `pdfops.go` - Replace `createInterleavedMerge()` function
- `fileops.go` - May need updates for temp file handling

#### Reference Implementation
- See `experiments/experiment16_final_memory_approach.go` for working example
- See `docs/memory_processing_research.md` for implementation pattern
- See `docs/api_knowledge.md` for API reference

#### Important Notes
- **Hybrid Approach**: Not pure in-memory due to pdfcpu API limitations
- **Temp Files**: Reduces from 7 to 2-3 files (not zero due to API constraints)
- **API Limitation**: pdfcpu requires file paths for page extraction and merging
- **For Zero Temp Files**: See Task 25 (pdfcpu feature request) for true in-memory processing

#### Acceptance Criteria
- [ ] Merging uses minimal temporary files (only during page extraction)
- [ ] Original PDF data kept in memory throughout process
- [ ] Graceful handling of pages that fail extraction
- [ ] Memory usage ~50% of original file sizes
- [ ] Proper cleanup of all temporary files
- [ ] Maintains existing interleaved merge pattern (A1, B3, A2, B2, A3, B1)

#### Estimated Effort
- **Development**: 4-6 hours
- **Testing**: 2-3 hours
- **Documentation**: 1 hour

#### Suggested Commit Message
```
perf: Add hybrid in-memory PDF processing

- Load PDFs into memory for validation
- Use minimal temporary files during operations
- Implement graceful handling of extraction failures
- Achieve ~50% memory efficiency vs original approach
```

### Future Enhancement Ideas (On Hold)
- Configuration file support
- Batch processing capabilities
- Additional PDF manipulation features
- Integration with other tools
- Web interface for remote processing
- API endpoint for programmatic access
- Plugin system for custom processing
- Advanced PDF optimization features

---

## Development Notes
- All backlog items require explicit approval before moving to "Next Stages"
- Maintain backward compatibility for all future changes
- Follow established git workflow and commit conventions
- Comprehensive testing required for all new features
- Update documentation for any changes made

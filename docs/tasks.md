# Task Board - BlendPDFGo

## Task Summary (40 Total)
- ✅ **Done**: 27 tasks
- 🔄 **In Progress**: 0 tasks  
- 📋 **To Do**: 6 tasks
- 🗂️ **Backlog**: 7 tasks

### 📊 Project Status: PRODUCTION READY
All core functionality complete with professional UI, real-time monitoring, comprehensive testing, and multi-platform deployment.

---

## 🔄 In Progress
*No tasks currently in progress*

---

## 📋 To Do (Ready for Work)

### T-023: PowerShell 5/CMD Compatibility Implementation
**Epic**: E-05 | **Story**: US-009
**Priority**: Medium | **Estimate**: 4 hours
**Dependencies**: None

**Description**: Implement graceful fallback for legacy Windows terminals

**Acceptance Criteria**:
- [ ] Detect PowerShell 5 and CMD environments
- [ ] Provide fallback display options for limited terminals
- [ ] Maintain functionality without advanced terminal features
- [ ] Test on Windows legacy environments

---

### T-028: Undo/Restore Functionality
**Epic**: E-08 | **Story**: US-012
**Priority**: Low | **Estimate**: 8 hours
**Dependencies**: Core functionality

**Description**: Ability to reverse last operation (move files back from archive/output)

**Acceptance Criteria**:
- [ ] Track recent file operations
- [ ] Implement restore from archive/output directories
- [ ] Provide clear undo operation feedback
- [ ] Handle undo conflicts gracefully

---

### T-031: Configuration File Support
**Epic**: E-07 | **Story**: US-011
**Priority**: Low | **Estimate**: 6 hours
**Dependencies**: Core functionality

**Description**: Store configuration in blendpdf.json in working directory

**Acceptance Criteria**:
- [ ] JSON configuration file format
- [ ] Customizable output directories and naming
- [ ] Configuration validation and defaults
- [ ] Documentation and examples

---

### T-032: Archive Single Files
**Epic**: E-02 | **Story**: US-005
**Priority**: Medium | **Estimate**: 3 hours
**Dependencies**: T-005

**Description**: Copy single files to archive before moving to output, with --no-archive toggle

**Acceptance Criteria**:
- [ ] Copy files to archive before processing
- [ ] Implement --no-archive command line option
- [ ] Maintain backward compatibility
- [ ] Update file organization logic

---

### T-033: Multi Output Folders
**Epic**: E-07 | **Story**: US-011
**Priority**: Low | **Estimate**: 4 hours
**Dependencies**: T-031

**Description**: Support multiple output destinations with atomic operations and validation

**Acceptance Criteria**:
- [ ] Configure multiple output directories
- [ ] Atomic file operations across destinations
- [ ] Validation of output directory accessibility
- [ ] Error handling for failed destinations

---

### T-034: Keyboard Shortcuts Enhancement
**Epic**: E-03 | **Story**: US-014
**Priority**: Low | **Estimate**: 2 hours
**Dependencies**: T-009

**Description**: Add more keyboard shortcuts for faster navigation

**Acceptance Criteria**:
- [ ] Additional keyboard shortcuts for common operations
- [ ] Help display for available shortcuts
- [ ] Cross-platform keyboard handling
- [ ] User documentation updates

---

## 🗂️ Backlog (Future Work)

### T-014: Implement In-Memory Processing Approach
**Epic**: E-04 | **Story**: US-008
**Priority**: Medium | **Estimate**: 12 hours
**Status**: Ready for implementation (52.9% memory efficiency)

**Description**: Optimize memory usage for large file processing

---

### T-029: Web Interface
**Epic**: E-08 | **Story**: US-037
**Priority**: Low | **Estimate**: 20 hours

**Description**: Browser-based UI for remote operation

---

### T-030: Cloud Storage Integration
**Epic**: E-08 | **Story**: US-038
**Priority**: Low | **Estimate**: 16 hours

**Description**: Direct integration with Google Drive, Dropbox, OneDrive

---

### T-035: API Endpoints
**Epic**: E-08 | **Story**: US-034
**Priority**: Low | **Estimate**: 12 hours

**Description**: REST API for programmatic access

---

### T-036: Email/Notification Support
**Epic**: E-08 | **Story**: US-039
**Priority**: Low | **Estimate**: 8 hours
**Dependencies**: T-029, T-035

**Description**: Send completion notifications

---

### T-037: Error Recovery Enhancement
**Epic**: E-01 | **Story**: US-004
**Priority**: Low | **Estimate**: 6 hours

**Description**: Auto-retry failed operations with exponential backoff

---

### T-038: Audit Logging
**Epic**: E-06 | **Story**: US-010
**Priority**: Low | **Estimate**: 4 hours

**Description**: Detailed operation logs for compliance

---

## ✅ Done

### T-001: Fix getPageCount Function ✅ COMPLETED
**Epic**: E-01 | **Story**: US-001
**Completed**: Aug 27 | **Actual Time**: 2 hours

**Description**: Fixed the `getPageCount` function to use correct pdfcpu API

**Completion Notes**: Replaced `api.PDFInfo` with `api.PageCountFile(file)` which directly returns page count. Core PDF validation now working correctly.

---

### T-002: Implement Page Count Validation ✅ COMPLETED
**Epic**: E-01 | **Story**: US-001
**Completed**: Aug 27 | **Actual Time**: 3 hours

**Description**: Added exact page count validation between two PDFs

**Completion Notes**: Implemented exact page count matching with clear error messages. Files moved to error/ directory if counts don't match.

---

### T-003: Fix Merging Logic for Interleaved Pattern ✅ COMPLETED
**Epic**: E-01 | **Story**: US-002, US-003
**Completed**: Aug 27 | **Actual Time**: 8 hours

**Description**: Rewrote merging logic to create correct interleaved pattern

**Completion Notes**: Fixed page reversal logic by extracting pages individually in reverse order. Correct interleaved pattern now achieved: A1, f, A2, 9, A3, M.

---

### T-004: Update Filename Generation ✅ COMPLETED
**Epic**: E-02 | **Story**: US-005
**Completed**: Aug 27 | **Actual Time**: 1 hour

**Description**: Improved output filename generation with timestamp and source info

**Completion Notes**: Implemented descriptive filename generation including timestamp and source file information.

---

### T-005: Auto-select First Two PDFs ✅ COMPLETED
**Epic**: E-02 | **Story**: US-005
**Completed**: Aug 27 | **Actual Time**: 2 hours

**Description**: Automatically select the first two PDF files found in directory

**Completion Notes**: Implemented automatic PDF file selection with alphabetical sorting. Creates required directories if missing.

---

### T-006: Enhanced User Interface and Display ✅ COMPLETED
**Epic**: E-03 | **Story**: US-006
**Completed**: Aug 28 | **Actual Time**: 4 hours

**Description**: Professional UI with color-coded messages and real-time feedback

**Completion Notes**: Implemented comprehensive UI with file count display, colored output (Red/Green/Yellow/Blue), file preview, and session statistics.

---

### T-007: Smart PDF Processing with Page Reversal Logic ✅ COMPLETED
**Epic**: E-01 | **Story**: US-002
**Completed**: Aug 28 | **Actual Time**: 3 hours

**Description**: Intelligent page reversal only for multi-page PDFs

**Completion Notes**: Optimized processing to only reverse multi-page PDFs. Single-page files processed directly without reversal.

---

### T-008: Robust Error Handling and File Management ✅ COMPLETED
**Epic**: E-01 | **Story**: US-004
**Completed**: Aug 28 | **Actual Time**: 3 hours

**Description**: Comprehensive error handling with detailed logging

**Completion Notes**: Implemented robust error handling with graceful failure recovery and detailed error reporting.

---

### T-009: Command Line Interface Enhancements ✅ COMPLETED
**Epic**: E-03 | **Story**: US-014
**Completed**: Aug 28 | **Actual Time**: 2 hours

**Description**: Complete CLI with options and keyboard shortcuts

**Completion Notes**: Added comprehensive command line interface with debug mode, verbose options, and keyboard shortcuts.

---

### T-010: Session Management and Statistics ✅ COMPLETED
**Epic**: E-03 | **Story**: US-007
**Completed**: Aug 28 | **Actual Time**: 2 hours

**Description**: Track operations, errors, and performance metrics

**Completion Notes**: Implemented session statistics tracking with operation counts, error tracking, and elapsed time monitoring.

---

### T-011: Advanced File Operations ✅ COMPLETED
**Epic**: E-02 | **Story**: US-005
**Completed**: Aug 28 | **Actual Time**: 2 hours

**Description**: Enhanced file handling with lock files and timeout protection

**Completion Notes**: Added lock file protection to prevent multiple instances and timeout protection for automatic exit.

---

### T-012: Output and Logging Improvements ✅ COMPLETED
**Epic**: E-03 | **Story**: US-006
**Completed**: Aug 28 | **Actual Time**: 2 hours

**Description**: Structured logging with multiple levels and debug mode

**Completion Notes**: Implemented comprehensive logging system with debug mode and performance monitoring.

---

### T-013: Performance and Reliability ✅ COMPLETED
**Epic**: E-04 | **Story**: US-008
**Completed**: Aug 28 | **Actual Time**: 3 hours

**Description**: Performance monitoring and optimization

**Completion Notes**: Added performance metrics, operation tracking, and reliability improvements for production use.

---

### T-015: Multi-Platform Build System ✅ COMPLETED
**Epic**: E-05 | **Story**: US-009
**Completed**: Aug 29 | **Actual Time**: 4 hours

**Description**: Cross-platform builds with automated releases

**Completion Notes**: Implemented multi-platform build system with GitHub Actions for automated releases across Windows, macOS, and Linux.

---

### T-016: Directory-Specific Lock Files ✅ COMPLETED
**Epic**: E-02 | **Story**: US-005
**Completed**: Aug 29 | **Actual Time**: 2 hours

**Description**: Allow concurrent usage in different directories

**Completion Notes**: Enhanced lock file system to be directory-specific, enabling concurrent usage across different working directories.

---

### T-017: Unit Testing Implementation ✅ COMPLETED
**Epic**: E-06 | **Story**: US-010
**Completed**: Aug 30 | **Actual Time**: 8 hours

**Description**: Comprehensive unit test coverage for core functionality

**Completion Notes**: Implemented extensive unit testing with API function testing (tests 01-16) and core functionality validation.

---

### T-018: Documentation Review and Cleanup ✅ COMPLETED
**Epic**: E-06 | **Story**: US-010
**Completed**: Sep 01 | **Actual Time**: 4 hours

**Description**: Comprehensive documentation review and organization

**Completion Notes**: Reviewed and cleaned up all documentation, removing duplicates and organizing information effectively.

---

### T-019: Documentation Cleanup Implementation ✅ COMPLETED
**Epic**: E-06 | **Story**: US-010
**Completed**: Sep 01 | **Actual Time**: 2 hours

**Description**: Final documentation cleanup and standardization

**Completion Notes**: Completed documentation standardization with consistent formatting and comprehensive coverage.

---

### T-020: CLI Library Research ✅ COMPLETED
**Epic**: E-03 | **Story**: US-006
**Completed**: Aug 29 | **Actual Time**: 4 hours

**Description**: Research and evaluate CLI libraries for enhanced user experience

**Completion Notes**: Comprehensive research documented in cli_library_research.md with recommendations for future enhancements.

---

### T-021: UI Interface Recommendations ✅ COMPLETED
**Epic**: E-03 | **Story**: US-006
**Completed**: Aug 29 | **Actual Time**: 2 hours

**Description**: Develop recommendations for future UI enhancements

**Completion Notes**: Created comprehensive UI recommendations for future web interface and advanced features.

---

### T-022: Code Refactoring Implementation ✅ COMPLETED
**Epic**: E-06 | **Story**: US-010
**Completed**: Aug 30 | **Actual Time**: 6 hours

**Description**: Refactor code for maintainability and performance

**Completion Notes**: Completed major code refactoring for improved maintainability, performance, and code quality.

---

### T-024: Implement Zip Merge Solution for Interleaved Pattern ✅ COMPLETED
**Epic**: E-01 | **Story**: US-003
**Completed**: Sep 01 | **Actual Time**: 6 hours

**Description**: Breakthrough solution using MergeCreateZipFile API for efficient interleaved merging

**Completion Notes**: Discovered and implemented native pdfcpu API for interleaved merging, reducing complexity from 6+ temp files and 7+ API calls to 1 temp file and 2 API calls.

---

### T-025: pdfcpu Feature Request for In-Memory Processing ✅ COMPLETED
**Epic**: E-04 | **Story**: US-008
**Completed**: Sep 01 | **Actual Time**: 3 hours

**Description**: Research and document feature request for in-memory processing

**Completion Notes**: Documented comprehensive feature request for pdfcpu library to support in-memory processing, potentially improving memory efficiency by 52.9%.

---

### T-026: Remove User Timeout ✅ COMPLETED
**Epic**: E-03 | **Story**: US-006
**Completed**: Sep 01 | **Actual Time**: 1 hour

**Description**: Remove automatic timeout to improve user experience

**Completion Notes**: Removed automatic 5-minute timeout based on user feedback, allowing unlimited processing time.

---

### T-027: Full-Screen UI Implementation ✅ COMPLETED
**Epic**: E-03 | **Story**: US-006
**Completed**: Sep 01 | **Actual Time**: 3 hours

**Description**: Implement full-screen terminal UI with enhanced display

**Completion Notes**: Implemented comprehensive full-screen UI with real-time updates, enhanced file display, and improved user experience.

---

## Task Management Notes

### Current Sprint/Focus
Focus on remaining enhancements and advanced features. Core functionality is production-ready.

### Blockers and Issues
*No current blockers*

### Next Priorities
1. T-023: PowerShell 5/CMD Compatibility (cross-platform support)
2. T-032: Archive Single Files (workflow improvement)
3. T-028: Undo/Restore Functionality (user experience enhancement)

---

## Development Progress Summary
- **Phase 1 (Core Functionality)**: ✅ Complete - All core PDF processing implemented
- **Phase 2 (User Interface)**: ✅ Complete - Professional UI with comprehensive feedback
- **Phase 3 (Cross-Platform)**: ✅ Complete - Multi-platform builds and compatibility
- **Phase 4 (Quality Assurance)**: ✅ Complete - Comprehensive testing and documentation
- **Phase 5 (Advanced Features)**: 🔄 6/13 Complete - Configuration and extensibility features
- **Phase 6 (Future Enhancements)**: 📋 Not Started - Web interface, cloud integration, API endpoints

**Project Status**: Production-ready with ongoing enhancements for advanced features and user experience improvements.
- 📋 Task 39: Plugin System - Allow custom processing plugins
- 📋 Task 40: Docker Container - Containerised deployment option

### 📊 Project Status: PRODUCTION READY
All core functionality complete with professional UI, real-time monitoring, comprehensive testing, and multi-platform deployment.

---

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
- **Coloured Output**: Red/Green/Yellow/Blue message types
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
- **Performance**: ✅ Monitoring and optimisation ready
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
- **Requirements**: R5.1-R5.3 (file count display), R6.1-R6.2 (verbose mode), R7.1-R7.5 (coloured output), R8.1-R8.4 (session statistics)
- **Features Implemented**:
  - File count display: "Files: Main(X) Archive(Y) Output(Z) Error(W)"
  - File preview in verbose mode showing up to 5 files with sizes
  - Coloured output with Red/Green/Yellow/Blue message types
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
  - Analysed all 10 documentation files for content overlaps (73,779 bytes total)
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
  - Constants Tests: Exit codes, colours, logger initialisation ✅
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
  - Reorganised API documentation structure
  - Moved experiment summaries to api_knowledge.md
  - Streamlined api_experiments_procedures.md for testing procedures only
  - Updated all cross-references and file links
  - Added consistent navigation sections
  - Removed documentation_review_analysis.md after completion

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

### Task 24: Implement Zip Merge Solution for Interleaved Pattern
- **Status**: ✅ COMPLETED - Commit: 1e4c365
- **Requirements**: R2.1-R2.4 (smart page reversal), R1.4-R1.5 (interleaved pattern)
- **Priority**: Code Quality Enhancement (High) - **UPGRADED FROM MEDIUM**
- **Description**: Replace complex loop-based interleaving with simple 2-step zip merge solution
- **Background**: Discovered `MergeCreateZipFile` API provides perfect interleaved merging when combined with `CollectFile`
- **Implementation Completed**:
  - Research Completed: ✅ (Experiments 20-22 confirm zip merge solution)
  - API Functions: `CollectFile` + `MergeCreateZipFile` 
  - Breakthrough: Eliminates all individual page extraction and complex merging logic
- **Benefits Achieved**:
  - Dramatic Simplification: 2 API calls instead of 6+ individual page extractions
  - Reduced Temporary Files: 1 temp file instead of 6+ temp files
  - Perfect Interleaving: Native zip merge provides exact pattern needed
  - Better Performance: Fewer I/O operations and API calls
  - True Solution: Uses intended pdfcpu APIs, not workarounds
- **Files Modified**: pdfops.go, docs/api_knowledge.md, experiments 20-22
- **Testing**: All tests pass with new implementation, verified correct output pattern

### Task 20: CLI Library Research
- **Status**: ✅ COMPLETED
- **Requirements**: R5.4-R5.10 (interactive menu), R9.1-R9.8 (CLI interface), user experience enhancement
- **Priority**: User Experience Enhancement (Medium)
- **Description**: Research Go CLI libraries for enhanced user experience
- **Implementation Completed**:
  - Researched 6 major Go CLI libraries (Cobra, Bubble Tea, Survey, Termui, Progressbar, Fatih/Color)
  - Analysed pros/cons and use cases for each library
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

### Task 25: pdfcpu Feature Request for In-Memory Processing
- **Status**: ✅ COMPLETED
- **Priority**: Community Contribution (Low)
- **Description**: Create GitHub feature request for in-memory PDF processing capabilities in pdfcpu
- **Background**: Current pdfcpu API requires file-based operations, limiting true in-memory processing
- **Feature Request Details**:
  - Missing Functions: Context-based page extraction and merging
  - Proposed APIs: `api.ExtractPagesContext()`, `api.MergeContexts()`, `api.TrimContext()`
  - Use Case: Enable zero-temp-file PDF processing for better performance and resource usage
  - Benefits: Reduced disk I/O, better memory efficiency, cleaner application architecture
- **Implementation Completed**:
  - ✅ GitHub issue research and analysis completed
  - ✅ Technical specification document created (`docs/pdfcpu_feature_request.md`)
  - ✅ Feature request markdown ready for review
  - ✅ GitHub issue submitted
  - [ ] Follow-up on maintainer feedback (pending response)

### Task 27: Full-Screen UI Implementation
- **Status**: ✅ COMPLETED
- **Requirements**: R5.1-R5B.7 (full-screen UI), R5A.1-R5A.6 (file selection modes)
- **Priority**: User Experience Enhancement (High)
- **Description**: Implement full-screen terminal UI to replace current scrolling menu interface
- **Prerequisites**: Task 20 (CLI Library Research) ✅, Task 21 (UI Design) ✅
- **Implementation Completed**:
  - ✅ Enhanced menu interface with full-screen bordered layout
  - ✅ Version number displayed in top border as requested
  - ✅ File counts integrated into header alongside directory paths
  - ✅ Professional appearance with clear sections
  - ✅ Cross-platform compatibility without complex TUI dependencies
  - ✅ All existing functionality preserved and enhanced
  - ✅ Session statistics with bordered display
  - ✅ Clean screen management and user experience
  - ✅ Real-time file monitoring and progress bars (R5.8-R5.9)
  - ✅ Enhanced Recent Output with detailed operation information (R5B.4)
  - ✅ Persistent actions bar during operations (R5B.5)
  - ✅ 2-line status/progress section (R5B.6)
  - ✅ Single-line recent operations format with timestamps and status icons
  - ✅ Clear error messages for insufficient files scenarios
  - ✅ Removed "Press Enter to continue" prompts for smoother workflow
  - ✅ Removed extra dividing line for cleaner visual flow
  - ✅ True real-time folder monitoring with fsnotify (event-driven vs polling)
  - ✅ Perfect invalid choice handling with interface redraw (no stacking)
- **Files Created**: `ui/enhanced_menu.go`, `ui/bridge.go`, plus TUI research files
- **Dependencies Added**: `github.com/fsnotify/fsnotify v1.9.0` for real-time file system monitoring
- **Testing**: All tests pass, merge functionality verified (A1, f, A2, 9, A3, M pattern)
- **Benefits Achieved**: Professional interface, version visibility, integrated file counts, universal compatibility, instant file change detection, enhanced progress feedback

---

## 3. Next Stages

### Task 23: PowerShell 5/CMD Compatibility Implementation
- **Status**: 📋 TO DO
- **Requirements**: R5.10 (graceful fallback to basic interface on legacy terminals)
- **Priority**: User Experience Enhancement (Medium)
- **Description**: Implement graceful fallback for legacy Windows terminals that don't support full ANSI escape codes
- **Current Status**: Basic cross-platform screen clearing implemented, but no terminal capability detection
- **Implementation Needed**:
  - Detect terminal capabilities (PowerShell version, CMD vs Windows Terminal)
  - Implement basic text-based interface for legacy terminals
  - Maintain full functionality with simplified display
  - Test on PowerShell 5, CMD, and Windows Console Host
- **Acceptance Criteria**:
  - [ ] Application works on PowerShell 5 without garbled output
  - [ ] Application works on CMD with basic functionality
  - [ ] Automatic detection and fallback to appropriate interface
  - [ ] All core operations available in fallback mode

### Task 28: Undo/Restore Functionality
- **Status**: 📋 TO DO
- **Requirements**: User Experience Enhancement
- **Priority**: User Experience Enhancement (Medium)
- **Description**: Ability to reverse last operation (move files back from archive/output)
- **Implementation Needed**:
  - Track last operation details (source files, destination, operation type)
  - Implement restore commands for single file moves and merge operations
  - Add [U]ndo option to main menu
  - Validate files still exist before restore attempt
  - Update session statistics after restore operations
- **Acceptance Criteria**:
  - [ ] Can undo last single file move (move from output back to main)
  - [ ] Can undo last merge operation (move files from archive back to main, remove merged output)
  - [ ] Clear undo history after successful undo operation
  - [ ] Show appropriate messages when undo is not available
  - [ ] Maintain file integrity during restore operations

### Task 31: Configuration File Support
- **Status**: 📋 TO DO
- **Requirements**: User Experience Enhancement
- **Priority**: User Experience Enhancement (Medium)
- **Description**: Store configuration in blendpdf.json in working directory
- **Implementation Needed**:
  - Create JSON configuration file structure
  - Load configuration at startup from blendpdf.json
  - Command line arguments override configuration file settings
  - Support archive mode, output folders, and other settings
  - Create default configuration if file doesn't exist
- **Configuration Structure**:
  ```json
  {
    "archiveMode": true,
    "outputFolders": ["output"],
    "verboseMode": false,
    "debugMode": false
  }
  ```
- **Acceptance Criteria**:
  - [ ] Configuration loaded from blendpdf.json in working directory
  - [ ] Command line args override config file settings
  - [ ] Default config created if file missing
  - [ ] Invalid JSON handled gracefully with error messages
  - [ ] All configurable settings supported

### Task 32: Archive Single Files
- **Status**: 📋 TO DO
- **Requirements**: File Management Enhancement
- **Priority**: User Experience Enhancement (Medium)
- **Description**: Copy single files to archive before moving to output, with --no-archive toggle
- **Implementation Needed**:
  - Change single file operation to copy to archive first, then move to output
  - Add --no-archive command line flag to disable archiving
  - Add archive mode toggle to configuration file
  - Add archive mode toggle to interactive menu
  - Ensure both single and merge operations respect archive mode setting
- **Behaviour Changes**:
  - **Archive Mode ON (default)**: Single files copied to archive, then moved to output
  - **Archive Mode OFF**: Single files moved directly to output (current behaviour)
  - **Merge operations**: Always archive (unchanged behaviour)
- **Acceptance Criteria**:
  - [ ] Single file operations copy to archive when archive mode enabled
  - [ ] --no-archive flag disables archiving for session
  - [ ] Configuration file controls default archive mode
  - [ ] Interactive menu shows archive mode toggle option
  - [ ] Both single and merge operations respect archive mode setting

### Task 33: Multi Output Folders
- **Status**: 📋 TO DO
- **Requirements**: File Management Enhancement
- **Priority**: User Experience Enhancement (Medium)
- **Description**: Support multiple output destinations with atomic operations and validation
- **Implementation Needed**:
  - Add -o flag for multiple output folders: `-o folder1,folder2,folder3`
  - Support multiple output folders in configuration file
  - Validate all output folders exist and are writable at startup
  - Copy output files to all destinations sequentially (atomic per destination)
  - Log success/failure for each destination separately
  - Copy failed outputs to error folder without rolling back successful destinations
- **Atomic Transaction Behaviour**:
  - Copy to first destination, then second, then third, etc.
  - If any destination fails, log failure and copy output to error folder
  - Do NOT rollback successful destinations
  - Continue with remaining destinations after failure
- **Acceptance Criteria**:
  - [ ] Command line -o flag accepts comma-separated folder list
  - [ ] Configuration file supports outputFolders array
  - [ ] All output folders validated at startup
  - [ ] Output files copied to all destinations sequentially
  - [ ] Failed destinations logged with specific error messages
  - [ ] Failed outputs copied to error folder
  - [ ] Successful destinations not rolled back on partial failure

### Task 34: Keyboard Shortcuts Enhancement
- **Status**: 📋 TO DO
- **Requirements**: User Experience Enhancement
- **Priority**: User Experience Enhancement (Low)
- **Description**: Add more keyboard shortcuts for faster navigation
- **Implementation Needed**:
  - Add F1 key for help display
  - Add Ctrl+Q for quit
  - Add Space key for refresh/file monitoring update
  - Enhanced key handling in UI system
  - Update help text to show new shortcuts
- **Acceptance Criteria**:
  - [ ] F1 key displays help information
  - [ ] Ctrl+Q exits the program gracefully
  - [ ] Space key refreshes file counts and display
  - [ ] All existing shortcuts continue to work
  - [ ] Help text updated with new shortcuts

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

### Task 29: Web Interface
- **Status**: 📋 BACKLOG
- **Requirements**: Integration & Automation Enhancement
- **Priority**: User Experience Enhancement (Medium)
- **Description**: Browser-based UI for remote operation
- **Benefits**: Access from any device, team collaboration
- **Implementation Needed**:
  - HTTP server with file upload/download capability
  - Web-based file management interface
  - Real-time operation status updates
  - Session management for multiple users
  - Responsive design for mobile/tablet access
- **Technical Requirements**:
  - Go HTTP server with static file serving
  - WebSocket for real-time updates
  - File upload handling with progress bars
  - Authentication and session management
  - Cross-platform deployment

### Task 30: Cloud Storage Integration
- **Status**: 📋 BACKLOG
- **Requirements**: Integration & Automation Enhancement
- **Priority**: Integration Enhancement (Low)
- **Description**: Direct integration with Google Drive, Dropbox, OneDrive
- **Benefits**: Seamless cloud workflow, automatic backup
- **Implementation Needed**:
  - OAuth integration for cloud providers
  - Direct file upload/download from cloud storage
  - Automatic sync of processed files
  - Cloud folder monitoring
  - Conflict resolution for duplicate files
- **Technical Requirements**:
  - Google Drive API integration
  - Dropbox API integration
  - OneDrive API integration
  - OAuth 2.0 authentication flow
  - Background sync processes

### Task 35: API Endpoints
- **Status**: 📋 BACKLOG
- **Requirements**: Integration & Automation Enhancement
- **Priority**: Integration Enhancement (Medium)
- **Description**: REST API for programmatic access
- **Benefits**: Integration with other tools and workflows
- **Implementation Needed**:
  - HTTP API with JSON responses
  - Authentication and authorisation
  - File upload/download endpoints
  - Operation status tracking
  - API documentation
- **Technical Requirements**:
  - Go HTTP server with REST endpoints
  - JSON request/response handling
  - File handling for uploads/downloads
  - Session management
  - Error handling and status codes

### Task 36: Email/Notification Support
- **Status**: 📋 BACKLOG
- **Requirements**: Integration & Automation Enhancement
- **Priority**: Integration Enhancement (Low)
- **Description**: Send completion notifications or email merged PDFs
- **Dependencies**: Tasks 29 (Web Interface), 35 (API Endpoints)
- **Benefits**: Automated delivery of processed documents
- **Implementation Needed**:
  - SMTP integration for email notifications
  - Template system for notification messages
  - Attachment handling for PDF delivery
  - Configuration for email settings
  - Notification preferences management
- **Technical Requirements**:
  - SMTP client implementation
  - Email template engine
  - File attachment handling
  - Configuration management
  - Error handling for delivery failures

### Task 37: Error Recovery Enhancement
- **Status**: 📋 BACKLOG
- **Requirements**: Performance & Reliability Enhancement
- **Priority**: Performance Enhancement (Low)
- **Description**: Auto-retry failed operations with exponential backoff
- **Benefits**: Handle temporary file locks, network issues
- **Implementation Needed**:
  - Retry logic with configurable attempts
  - Exponential backoff algorithm
  - Error classification (retryable vs permanent)
  - Retry state tracking
  - Configuration for retry parameters
- **Technical Requirements**:
  - Retry mechanism with backoff
  - Error type classification
  - State persistence for retries
  - Configurable retry policies
  - Logging for retry attempts

### Task 38: Audit Logging
- **Status**: 📋 BACKLOG
- **Requirements**: Security & Compliance Enhancement
- **Priority**: Security Enhancement (Medium)
- **Description**: Detailed operation logs for compliance
- **Benefits**: Track all file operations for security/compliance
- **Implementation Needed**:
  - Structured logging with rotation
  - Audit trail for all operations
  - Log retention policies
  - Secure log storage
  - Log analysis and reporting
- **Technical Requirements**:
  - Structured logging framework
  - Log rotation and archival
  - Secure log file handling
  - Compliance reporting features
  - Log integrity verification

### Task 39: Plugin System
- **Status**: 📋 BACKLOG
- **Requirements**: Developer Experience Enhancement
- **Priority**: Extensibility Enhancement (Low)
- **Description**: Allow custom processing plugins
- **Benefits**: Extensibility for specific workflows
- **Implementation Needed**:
  - Go plugin architecture
  - Plugin interface definition
  - Plugin discovery and loading
  - Plugin configuration management
  - Plugin lifecycle management
- **Technical Requirements**:
  - Plugin interface specification
  - Dynamic plugin loading
  - Plugin sandboxing and security
  - Plugin configuration system
  - Plugin documentation framework

### Task 40: Docker Container
- **Status**: 📋 BACKLOG
- **Requirements**: Deployment Enhancement
- **Priority**: Deployment Enhancement (Medium)
- **Description**: Containerised deployment option
- **Benefits**: Easy deployment and scaling
- **Implementation Needed**:
  - Multi-stage Docker build
  - Container optimisation
  - Docker Compose configuration
  - Container registry publishing
  - Documentation for container deployment
- **Technical Requirements**:
  - Dockerfile with multi-stage build
  - Container image optimisation
  - Volume mounting for file operations
  - Environment variable configuration
  - Container health checks

### Future Enhancement Ideas (On Hold)
- Batch processing capabilities
- Additional PDF manipulation features
- Integration with other tools
- Web interface for remote processing
- API endpoint for programmatic access
- Plugin system for custom processing
- Advanced PDF optimisation features

---

## Development Notes
- All backlog items require explicit approval before moving to "Next Stages"
- Maintain backward compatibility for all future changes
- Follow established git workflow and commit conventions
- Comprehensive testing required for all new features
- Update documentation for any changes made

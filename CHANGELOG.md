# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

### Added
- Undo/Restore functionality with [U] option in interactive menu
- Operation tracking for single file and merge operations
- Consistent file conflict resolution across all operations
- Archive preservation during undo operations
- PowerShell 5/CMD compatibility with terminal detection and legacy UI fallback
- Enhanced keyboard shortcuts for faster navigation (F1, Ctrl+Q, Ctrl+Z, etc.)
- Multiple input methods: descriptive words, numbers, and control key combinations
- Terminal capability detection for automatic UI selection

### Fixed
- Merge operations now respect --no-archive flag (previously only single file operations did)
- File conflict handling: all operations now generate unique names instead of overwriting
- Consistent behavior between single file and merge operations for archiving
- PowerShell 7 terminal detection now correctly uses modern UI with Unicode box-drawing characters
- TERM environment variable no longer overrides Windows terminal detection on Windows systems

### Changed
- Both single file and merge operations now consistently respect CONFIG.ArchiveMode setting
- Undo operations restore clean "pre-operation" state with files only in main directory
- Enhanced conflict resolution tracks actual filenames used (including _1, _2 suffixes)
- Automatic UI selection: enhanced for modern terminals, basic for legacy environments
- Cross-platform compatibility with graceful fallback for limited terminals

## [1.2.0] - 2025-09-11

### Changed
- Updated module name to match repository path (github.com/Kristianwhittick/blend-pdf)
- Changed README title to "Blend PDF (GO)" for better clarity

### Fixed
- Module path now matches actual repository location for proper go get functionality

## [1.1.0] - 2025-09-01

### Added
- **NEW UI**: Enhanced full-screen terminal interface with professional bordered layout
- Version number displayed in top border for better visibility
- File counts integrated into header alongside directory paths
- Real-time file monitoring with fsnotify for instant updates
- Enhanced Recent Output section with detailed operation information
- Persistent actions bar during operations for better user experience
- 2-line status/progress section with animated progress bars
- Single-line recent operations format with timestamps and status icons
- Invalid choice handling with clean interface redraw (no stacking)
- True event-driven file system monitoring (no polling)

### Changed
- **Application renamed**: "blendpdfgo" → "blendpdf" for shorter, cleaner name
- Updated all build scripts, documentation, and examples
- Lock file naming updated to use "blendpdf" prefix
- Cross-platform compatibility maintained with universal screen clearing

### Added - Documentation & Planning
- Added comprehensive design decisions documentation (Section 5 in specification)
- Added 3 new tasks: Configuration File Support, Archive Single Files, Multi Output Folders
- Added 6 new backlog items: API Endpoints, Email/Notifications, Error Recovery, Audit Logging, Plugin System, Docker
- Added Keyboard Shortcuts Enhancement task
- Documented rejected features with clear rationale

### Technical
- Added fsnotify dependency for real-time file system monitoring
- Enhanced UI system with bridge pattern for better separation
- Improved error handling and user feedback
- Maintained backward compatibility and all existing functionality

## [1.0.5] - 2025-08-29

### Changed
- **MAJOR**: Replaced complex loop-based PDF interleaving with elegant 2-step zip merge solution
- Reduced temporary files from 6+ to 1 (reversed document only)
- Reduced API calls from 7+ to 2 (CollectFile + MergeCreateZipFile)
- Improved performance and code maintainability significantly

### Technical
- Implemented `api.CollectFile()` for order-preserving page reversal
- Implemented `api.MergeCreateZipFile()` for perfect interleaved merging
- Maintains exact same output pattern (A1, f, A2, 9, A3, M)
- All existing tests pass with new implementation
- Task 24 completed with breakthrough zip merge approach

## [1.0.4] - 2025-08-29

### Fixed
- Fixed GitHub Actions to automatically extract and include version-specific changelog content in releases
- Replaced generic "See commit history" message with actual changelog entries

### Added
- Automatic changelog extraction from CHANGELOG.md during release process
- Version-specific changelog content in GitHub release descriptions

### Technical
- Enhanced GitHub Actions workflow with changelog parsing
- Improved release automation and documentation

## [1.0.3] - 2025-08-29

### Fixed
- Fixed GitHub Actions version handling to use correct version numbers from git tags instead of hardcoded v1.0.0
- Fixed build system version synchronization between git tags and build artifacts
- Fixed release asset naming to show correct version-specific binary names

### Added
- Automatic version sync script (`scripts/sync-version.sh`) for version consistency
- Enhanced build process with automatic git tag version detection
- Makefile integration with `sync-version` target

### Technical
- Updated build system to use `git describe --tags` for version detection
- Enhanced GitHub Actions workflow for proper version handling
- Improved cross-platform build consistency

## [1.0.2] - 2025-08-29

### Changed
- Updated Go version requirement from 1.24 to 1.24.6
- Updated all dependencies to latest compatible versions
- Updated build system to use Go 1.24.6

### Technical
- Verified library compatibility with Go 1.24.6
- All unit tests and integration tests passing
- Cross-platform build compatibility maintained

## [1.0.1] - 2025-08-29

### Fixed
- Fixed merge output filename display format - now shows `file1-file2.pdf` instead of `file1.pdf-file2.pdf` to match actual created filename

### Changed
- Removed 5-minute user inactivity timeout - application now runs indefinitely until user quits
- Simplified user input handling by removing timeout goroutines and channels

### Added
- CHANGELOG.md to track changes between releases
- Comprehensive CLI library research documentation
- Full-screen UI specifications and implementation requirements

## [1.0.0] - 2025-08-27

### Added
- Initial production release
- Complete feature parity with original bash version
- Interactive command-line menu with S/M/H/V/D/Q options
- Smart PDF merging with automatic page reversal
- File count display and colored output
- Verbose mode with file preview and detailed information
- Debug mode with structured logging
- Session statistics tracking
- Directory-specific lock file protection
- Timeout protection (5 minutes)
- Cross-platform build system
- Comprehensive error handling and recovery
- Multi-platform binaries (Windows, Linux, macOS)

### Technical Features
- PDF validation before processing
- Smart page reversal (only for multi-page PDFs)
- Interleaved merge pattern (A1, B3, A2, B2, A3, B1)
- Automatic directory creation (archive/, output/, error/)
- Graceful shutdown with cleanup
- Performance monitoring in debug mode

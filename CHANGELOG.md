# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

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

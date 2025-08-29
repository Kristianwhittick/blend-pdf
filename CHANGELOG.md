# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

### Fixed
- Fixed merge output filename display format - now shows `file1-file2.pdf` instead of `file1.pdf-file2.pdf` to match actual created filename

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

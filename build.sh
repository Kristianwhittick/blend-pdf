#!/bin/bash

# BlendPDFGo Multi-Platform Build Script
# Copyright 2025 Kristian Whittick
# Licensed under the Apache License, Version 2.0

set -e

# Configuration
APP_NAME="blendpdfgo"
VERSION=$(grep 'VERSION = ' constants.go | cut -d'"' -f2)
BUILD_DIR="dist"
SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"

# Colors for output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[0;33m'
BLUE='\033[0;34m'
NC='\033[0m' # No Color

# Platform configurations
declare -A PLATFORMS=(
    ["windows-amd64"]="windows amd64 .exe"
    ["linux-amd64"]="linux amd64 "
    ["linux-arm64"]="linux arm64 "
    ["darwin-amd64"]="darwin amd64 "
)

# Functions
print_header() {
    echo -e "${BLUE}================================${NC}"
    echo -e "${BLUE}  BlendPDFGo Build System v${VERSION}${NC}"
    echo -e "${BLUE}================================${NC}"
    echo
}

print_usage() {
    echo "Usage: $0 [OPTIONS] [PLATFORM]"
    echo
    echo "Options:"
    echo "  -h, --help     Show this help message"
    echo "  -c, --clean    Clean build directory before building"
    echo "  -v, --verbose  Enable verbose output"
    echo "  -a, --all      Build for all platforms"
    echo "  --checksums    Generate SHA256 checksums"
    echo
    echo "Platforms:"
    for platform in "${!PLATFORMS[@]}"; do
        echo "  $platform"
    done
    echo
    echo "Examples:"
    echo "  $0 --all                    # Build for all platforms"
    echo "  $0 linux-amd64             # Build for Linux x86_64"
    echo "  $0 --clean --all           # Clean and build all"
    echo "  $0 --checksums linux-amd64 # Build with checksums"
}

log_info() {
    echo -e "${GREEN}[INFO]${NC} $1"
}

log_warn() {
    echo -e "${YELLOW}[WARN]${NC} $1"
}

log_error() {
    echo -e "${RED}[ERROR]${NC} $1"
}

log_debug() {
    if [[ "$VERBOSE" == "true" ]]; then
        echo -e "${BLUE}[DEBUG]${NC} $1"
    fi
}

clean_build_dir() {
    if [[ -d "$BUILD_DIR" ]]; then
        log_info "Cleaning build directory: $BUILD_DIR"
        rm -rf "$BUILD_DIR"
    fi
    mkdir -p "$BUILD_DIR"
}

validate_environment() {
    log_debug "Validating build environment..."
    
    # Check if we're in the right directory
    if [[ ! -f "go.mod" ]] || [[ ! -f "main.go" ]]; then
        log_error "Must be run from the project root directory"
        exit 1
    fi
    
    # Check Go installation
    if ! command -v go &> /dev/null; then
        log_error "Go is not installed or not in PATH"
        exit 1
    fi
    
    # Check Go version
    GO_VERSION=$(go version | cut -d' ' -f3 | sed 's/go//')
    log_debug "Go version: $GO_VERSION"
    
    # Verify dependencies
    log_debug "Verifying Go dependencies..."
    go mod verify
    
    log_info "Environment validation complete"
}

build_platform() {
    local platform=$1
    local platform_info=(${PLATFORMS[$platform]})
    local goos=${platform_info[0]}
    local goarch=${platform_info[1]}
    local ext=${platform_info[2]}
    
    local output_name="${APP_NAME}-${VERSION}-${platform}${ext}"
    local output_path="${BUILD_DIR}/${output_name}"
    
    log_info "Building for $platform ($goos/$goarch)..."
    log_debug "Output: $output_path"
    
    # Set build environment
    export GOOS=$goos
    export GOARCH=$goarch
    export CGO_ENABLED=0
    
    # Build flags
    local ldflags="-s -w -X main.VERSION=${VERSION}"
    
    # Build command
    local build_start=$(date +%s)
    
    if [[ "$VERBOSE" == "true" ]]; then
        go build -v -ldflags "$ldflags" -o "$output_path" .
    else
        go build -ldflags "$ldflags" -o "$output_path" .
    fi
    
    local build_end=$(date +%s)
    local build_time=$((build_end - build_start))
    
    # Verify build
    if [[ -f "$output_path" ]]; then
        local file_size=$(stat -f%z "$output_path" 2>/dev/null || stat -c%s "$output_path" 2>/dev/null || echo "unknown")
        log_info "✓ Built $platform successfully (${file_size} bytes, ${build_time}s)"
        
        # Add to built platforms list
        BUILT_PLATFORMS+=("$platform:$output_path:$file_size")
    else
        log_error "✗ Failed to build $platform"
        return 1
    fi
    
    # Reset environment
    unset GOOS GOARCH CGO_ENABLED
}

generate_checksums() {
    local checksum_file="${BUILD_DIR}/checksums.sha256"
    
    log_info "Generating SHA256 checksums..."
    
    cd "$BUILD_DIR"
    
    # Generate checksums for all built binaries
    if command -v sha256sum &> /dev/null; then
        sha256sum ${APP_NAME}-${VERSION}-* > checksums.sha256
    elif command -v shasum &> /dev/null; then
        shasum -a 256 ${APP_NAME}-${VERSION}-* > checksums.sha256
    else
        log_warn "No SHA256 utility found, skipping checksums"
        cd "$SCRIPT_DIR"
        return 1
    fi
    
    cd "$SCRIPT_DIR"
    
    log_info "✓ Checksums generated: $checksum_file"
    
    if [[ "$VERBOSE" == "true" ]]; then
        echo
        log_debug "Checksum contents:"
        cat "$checksum_file" | while read line; do
            log_debug "  $line"
        done
    fi
}

print_build_summary() {
    echo
    log_info "Build Summary:"
    echo "  Version: $VERSION"
    echo "  Build Directory: $BUILD_DIR"
    echo "  Platforms Built: ${#BUILT_PLATFORMS[@]}"
    echo
    
    for platform_info in "${BUILT_PLATFORMS[@]}"; do
        IFS=':' read -r platform path size <<< "$platform_info"
        local filename=$(basename "$path")
        printf "  %-20s %s (%s bytes)\n" "$platform" "$filename" "$size"
    done
    
    if [[ -f "${BUILD_DIR}/checksums.sha256" ]]; then
        echo
        echo "  Checksums: checksums.sha256"
    fi
    
    echo
    log_info "Build completed successfully!"
}

# Main script
main() {
    local CLEAN=false
    local VERBOSE=false
    local BUILD_ALL=false
    local GENERATE_CHECKSUMS=false
    local TARGET_PLATFORM=""
    local BUILT_PLATFORMS=()
    
    # Parse command line arguments
    while [[ $# -gt 0 ]]; do
        case $1 in
            -h|--help)
                print_header
                print_usage
                exit 0
                ;;
            -c|--clean)
                CLEAN=true
                shift
                ;;
            -v|--verbose)
                VERBOSE=true
                shift
                ;;
            -a|--all)
                BUILD_ALL=true
                shift
                ;;
            --checksums)
                GENERATE_CHECKSUMS=true
                shift
                ;;
            -*)
                log_error "Unknown option: $1"
                print_usage
                exit 1
                ;;
            *)
                if [[ -n "$TARGET_PLATFORM" ]]; then
                    log_error "Multiple platforms specified. Use --all for all platforms."
                    exit 1
                fi
                TARGET_PLATFORM=$1
                shift
                ;;
        esac
    done
    
    # Validate target platform
    if [[ -n "$TARGET_PLATFORM" ]] && [[ -z "${PLATFORMS[$TARGET_PLATFORM]}" ]]; then
        log_error "Unknown platform: $TARGET_PLATFORM"
        echo "Available platforms: ${!PLATFORMS[*]}"
        exit 1
    fi
    
    # Start build process
    print_header
    
    # Clean if requested
    if [[ "$CLEAN" == "true" ]]; then
        clean_build_dir
    else
        mkdir -p "$BUILD_DIR"
    fi
    
    # Validate environment
    validate_environment
    
    # Determine what to build
    local platforms_to_build=()
    
    if [[ "$BUILD_ALL" == "true" ]]; then
        platforms_to_build=(${!PLATFORMS[@]})
        log_info "Building for all platforms..."
    elif [[ -n "$TARGET_PLATFORM" ]]; then
        platforms_to_build=("$TARGET_PLATFORM")
        log_info "Building for platform: $TARGET_PLATFORM"
    else
        log_error "No platform specified. Use --all or specify a platform."
        print_usage
        exit 1
    fi
    
    # Build platforms
    local build_start=$(date +%s)
    local failed_builds=0
    
    for platform in "${platforms_to_build[@]}"; do
        if ! build_platform "$platform"; then
            ((failed_builds++))
        fi
    done
    
    local build_end=$(date +%s)
    local total_time=$((build_end - build_start))
    
    # Generate checksums if requested
    if [[ "$GENERATE_CHECKSUMS" == "true" ]] && [[ ${#BUILT_PLATFORMS[@]} -gt 0 ]]; then
        generate_checksums
    fi
    
    # Print summary
    print_build_summary
    
    echo "Total build time: ${total_time}s"
    
    if [[ $failed_builds -gt 0 ]]; then
        log_warn "$failed_builds platform(s) failed to build"
        exit 1
    fi
}

# Run main function with all arguments
main "$@"

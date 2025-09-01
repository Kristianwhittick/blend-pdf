#!/bin/bash

# BlendPDFGo Release Script
# Copyright 2025 Kristian Whittick
# Licensed under the Apache License, Version 2.0

set -e

# Configuration
SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
PROJECT_DIR="$(dirname "$SCRIPT_DIR")"
APP_NAME="blendpdf"

# Colors
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[0;33m'
BLUE='\033[0;34m'
NC='\033[0m'

# Functions
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
    echo -e "${BLUE}[DEBUG]${NC} $1"
}

print_usage() {
    echo "Usage: $0 [OPTIONS] VERSION"
    echo
    echo "Create a new release with the specified version."
    echo
    echo "Options:"
    echo "  -h, --help     Show this help message"
    echo "  -d, --dry-run  Show what would be done without making changes"
    echo "  -f, --force    Force release even if version already exists"
    echo "  -p, --push     Push changes and tags to remote"
    echo
    echo "Examples:"
    echo "  $0 1.1.0                    # Create release v1.1.0"
    echo "  $0 --dry-run 1.1.0          # Preview release process"
    echo "  $0 --force --push 1.1.0     # Force create and push release"
}

validate_version() {
    local version=$1
    
    # Check semantic versioning format
    if [[ ! $version =~ ^[0-9]+\.[0-9]+\.[0-9]+$ ]]; then
        log_error "Version must follow semantic versioning (e.g., 1.0.0)"
        return 1
    fi
    
    # Check if version already exists as tag
    if git tag -l | grep -q "^v$version$"; then
        if [[ "$FORCE" != "true" ]]; then
            log_error "Version v$version already exists. Use --force to override."
            return 1
        else
            log_warn "Version v$version already exists but --force specified"
        fi
    fi
    
    return 0
}

update_version_file() {
    local version=$1
    local constants_file="$PROJECT_DIR/constants.go"
    
    log_info "Updating version in constants.go to $version"
    
    if [[ "$DRY_RUN" == "true" ]]; then
        log_debug "DRY RUN: Would update VERSION = \"$version\" in $constants_file"
        return 0
    fi
    
    # Create backup
    cp "$constants_file" "$constants_file.bak"
    
    # Update version
    sed -i.tmp "s/VERSION = \"[^\"]*\"/VERSION = \"$version\"/" "$constants_file"
    rm "$constants_file.tmp"
    
    # Verify update
    local new_version=$(grep 'VERSION = ' "$constants_file" | cut -d'"' -f2)
    if [[ "$new_version" != "$version" ]]; then
        log_error "Failed to update version in constants.go"
        mv "$constants_file.bak" "$constants_file"
        return 1
    fi
    
    rm "$constants_file.bak"
    log_info "✓ Version updated successfully"
}

run_tests() {
    log_info "Running tests before release..."
    
    if [[ "$DRY_RUN" == "true" ]]; then
        log_debug "DRY RUN: Would run tests"
        return 0
    fi
    
    cd "$PROJECT_DIR"
    
    # Run tests
    if ! go test -v ./...; then
        log_error "Tests failed. Aborting release."
        return 1
    fi
    
    # Run linter if available
    if command -v golangci-lint >/dev/null 2>&1; then
        log_info "Running linter..."
        if ! golangci-lint run; then
            log_error "Linter failed. Aborting release."
            return 1
        fi
    fi
    
    log_info "✓ All tests passed"
}

build_release() {
    local version=$1
    
    log_info "Building release binaries for version $version..."
    
    if [[ "$DRY_RUN" == "true" ]]; then
        log_debug "DRY RUN: Would build release binaries"
        return 0
    fi
    
    cd "$PROJECT_DIR"
    
    # Clean and build all platforms
    if ! make clean build-all; then
        log_error "Build failed. Aborting release."
        return 1
    fi
    
    log_info "✓ Release binaries built successfully"
}

create_git_tag() {
    local version=$1
    local tag="v$version"
    
    log_info "Creating Git tag $tag..."
    
    if [[ "$DRY_RUN" == "true" ]]; then
        log_debug "DRY RUN: Would create tag $tag"
        return 0
    fi
    
    cd "$PROJECT_DIR"
    
    # Add changes
    git add constants.go
    
    # Commit version update
    git commit -m "chore: bump version to $version"
    
    # Create annotated tag
    git tag -a "$tag" -m "Release version $version"
    
    log_info "✓ Git tag $tag created"
}

push_changes() {
    local version=$1
    local tag="v$version"
    
    if [[ "$PUSH" != "true" ]]; then
        log_info "Skipping push (use --push to push changes)"
        return 0
    fi
    
    log_info "Pushing changes and tags to remote..."
    
    if [[ "$DRY_RUN" == "true" ]]; then
        log_debug "DRY RUN: Would push changes and tag $tag"
        return 0
    fi
    
    cd "$PROJECT_DIR"
    
    # Push changes and tags
    git push origin main
    git push origin "$tag"
    
    log_info "✓ Changes and tags pushed to remote"
}

generate_release_notes() {
    local version=$1
    local notes_file="$PROJECT_DIR/RELEASE_NOTES_$version.md"
    
    log_info "Generating release notes..."
    
    if [[ "$DRY_RUN" == "true" ]]; then
        log_debug "DRY RUN: Would generate release notes"
        return 0
    fi
    
    # Get previous tag
    local prev_tag=$(git tag -l | grep -E '^v[0-9]+\.[0-9]+\.[0-9]+$' | sort -V | tail -n 2 | head -n 1)
    
    cat > "$notes_file" << EOF
# BlendPDFGo Release $version

## Changes

EOF
    
    if [[ -n "$prev_tag" ]]; then
        echo "Changes since $prev_tag:" >> "$notes_file"
        echo "" >> "$notes_file"
        git log --pretty=format:"- %s" "$prev_tag"..HEAD >> "$notes_file"
    else
        echo "Initial release" >> "$notes_file"
    fi
    
    cat >> "$notes_file" << EOF

## Downloads

Choose the appropriate binary for your platform:

- **Windows (64-bit)**: \`${APP_NAME}-${version}-windows-amd64.exe\`
- **Linux (64-bit)**: \`${APP_NAME}-${version}-linux-amd64\`
- **Linux (ARM64)**: \`${APP_NAME}-${version}-linux-arm64\`
- **macOS (64-bit)**: \`${APP_NAME}-${version}-darwin-amd64\`

## Installation

1. Download the appropriate binary for your platform
2. Make it executable (Linux/macOS): \`chmod +x ${APP_NAME}-*\`
3. Move to your PATH or run directly

## Verification

Verify the integrity of your download using the provided checksums:
\`\`\`bash
sha256sum -c checksums.sha256
\`\`\`
EOF
    
    log_info "✓ Release notes generated: $notes_file"
}

cleanup() {
    local version=$1
    
    # Clean up temporary files
    rm -f "$PROJECT_DIR/RELEASE_NOTES_$version.md"
}

main() {
    local DRY_RUN=false
    local FORCE=false
    local PUSH=false
    local VERSION=""
    
    # Parse arguments
    while [[ $# -gt 0 ]]; do
        case $1 in
            -h|--help)
                print_usage
                exit 0
                ;;
            -d|--dry-run)
                DRY_RUN=true
                shift
                ;;
            -f|--force)
                FORCE=true
                shift
                ;;
            -p|--push)
                PUSH=true
                shift
                ;;
            -*)
                log_error "Unknown option: $1"
                print_usage
                exit 1
                ;;
            *)
                if [[ -n "$VERSION" ]]; then
                    log_error "Multiple versions specified"
                    exit 1
                fi
                VERSION=$1
                shift
                ;;
        esac
    done
    
    # Validate arguments
    if [[ -z "$VERSION" ]]; then
        log_error "Version is required"
        print_usage
        exit 1
    fi
    
    # Validate version format
    if ! validate_version "$VERSION"; then
        exit 1
    fi
    
    # Show what we're doing
    echo -e "${BLUE}BlendPDFGo Release Process${NC}"
    echo "Version: $VERSION"
    echo "Dry Run: $DRY_RUN"
    echo "Force: $FORCE"
    echo "Push: $PUSH"
    echo
    
    # Confirm if not dry run
    if [[ "$DRY_RUN" != "true" ]]; then
        read -p "Continue with release? (y/N): " -n 1 -r
        echo
        if [[ ! $REPLY =~ ^[Yy]$ ]]; then
            log_info "Release cancelled"
            exit 0
        fi
    fi
    
    # Execute release steps
    local start_time=$(date +%s)
    
    # Step 1: Update version
    if ! update_version_file "$VERSION"; then
        log_error "Failed to update version"
        exit 1
    fi
    
    # Step 2: Run tests
    if ! run_tests; then
        log_error "Tests failed"
        exit 1
    fi
    
    # Step 3: Build release
    if ! build_release "$VERSION"; then
        log_error "Build failed"
        exit 1
    fi
    
    # Step 4: Create Git tag
    if ! create_git_tag "$VERSION"; then
        log_error "Failed to create Git tag"
        exit 1
    fi
    
    # Step 5: Generate release notes
    if ! generate_release_notes "$VERSION"; then
        log_error "Failed to generate release notes"
        exit 1
    fi
    
    # Step 6: Push changes
    if ! push_changes "$VERSION"; then
        log_error "Failed to push changes"
        exit 1
    fi
    
    local end_time=$(date +%s)
    local duration=$((end_time - start_time))
    
    # Success message
    echo
    log_info "✓ Release $VERSION completed successfully in ${duration}s"
    
    if [[ "$DRY_RUN" != "true" ]]; then
        echo
        echo "Next steps:"
        echo "1. Check the GitHub Actions build status"
        echo "2. Verify the release artifacts"
        echo "3. Update documentation if needed"
        
        if [[ "$PUSH" == "true" ]]; then
            echo "4. GitHub release will be created automatically"
        else
            echo "4. Push changes with: git push origin main && git push origin v$VERSION"
        fi
    fi
    
    # Cleanup
    cleanup "$VERSION"
}

# Run main function
main "$@"

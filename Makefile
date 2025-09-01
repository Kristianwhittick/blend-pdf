# BlendPDFGo Makefile
# Copyright 2025 Kristian Whittick
# Licensed under the Apache License, Version 2.0

# Configuration
APP_NAME := blendpdf
VERSION := $(shell grep 'VERSION = ' constants.go | cut -d'"' -f2)
BUILD_DIR := dist
GO_VERSION := $(shell go version | cut -d' ' -f3)

# Build flags
LDFLAGS := -s -w -X main.VERSION=$(VERSION)
BUILD_FLAGS := -ldflags "$(LDFLAGS)"

# Platform configurations
PLATFORMS := windows-amd64 linux-amd64 linux-arm64 darwin-amd64

# Colors
RED := \033[0;31m
GREEN := \033[0;32m
YELLOW := \033[0;33m
BLUE := \033[0;34m
NC := \033[0m

# Default target
.DEFAULT_GOAL := help

# Phony targets
.PHONY: help clean build build-all test lint fmt vet deps check install uninstall release version platforms

## Help target
help: ## Show this help message
	@echo "$(BLUE)BlendPDFGo Build System$(NC)"
	@echo "Version: $(VERSION)"
	@echo "Go Version: $(GO_VERSION)"
	@echo ""
	@echo "$(YELLOW)Available targets:$(NC)"
	@awk 'BEGIN {FS = ":.*##"} /^[a-zA-Z_-]+:.*##/ { printf "  $(GREEN)%-15s$(NC) %s\n", $$1, $$2 }' $(MAKEFILE_LIST)
	@echo ""
	@echo "$(YELLOW)Platform targets:$(NC)"
	@for platform in $(PLATFORMS); do \
		printf "  $(GREEN)%-15s$(NC) Build for $$platform\n" "$$platform"; \
	done

## Clean build artifacts
clean: ## Clean build directory and artifacts
	@echo "$(YELLOW)Cleaning build artifacts...$(NC)"
	@rm -rf $(BUILD_DIR)
	@rm -f $(APP_NAME)
	@echo "$(GREEN)✓ Clean completed$(NC)"

## Build for current platform
sync-version: ## Sync version from git tags to constants.go
	@echo "$(YELLOW)Syncing version from git tags...$(NC)"
	@./scripts/sync-version.sh
	@echo "$(GREEN)✓ Version sync completed$(NC)"

build: deps sync-version ## Build for current platform
	@echo "$(YELLOW)Building $(APP_NAME) for current platform...$(NC)"
	@go build $(BUILD_FLAGS) -o $(APP_NAME) .
	@echo "$(GREEN)✓ Build completed: $(APP_NAME)$(NC)"

## Build for all platforms
build-all: clean sync-version ## Build for all supported platforms
	@echo "$(YELLOW)Building for all platforms...$(NC)"
	@./build.sh --all --checksums
	@echo "$(GREEN)✓ Multi-platform build completed$(NC)"

## Run tests
test: ## Run all tests
	@echo "$(YELLOW)Running tests...$(NC)"
	@go test -v -race -cover ./...
	@echo "$(GREEN)✓ Tests completed$(NC)"

## Run linter
lint: ## Run golangci-lint
	@echo "$(YELLOW)Running linter...$(NC)"
	@if command -v golangci-lint >/dev/null 2>&1; then \
		golangci-lint run; \
		echo "$(GREEN)✓ Linting completed$(NC)"; \
	else \
		echo "$(RED)golangci-lint not found. Install with: go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest$(NC)"; \
	fi

## Format code
fmt: ## Format Go code
	@echo "$(YELLOW)Formatting code...$(NC)"
	@go fmt ./...
	@echo "$(GREEN)✓ Code formatted$(NC)"

## Run go vet
vet: ## Run go vet
	@echo "$(YELLOW)Running go vet...$(NC)"
	@go vet ./...
	@echo "$(GREEN)✓ Vet completed$(NC)"

## Download dependencies
deps: ## Download and verify dependencies
	@echo "$(YELLOW)Downloading dependencies...$(NC)"
	@go mod download
	@go mod verify
	@echo "$(GREEN)✓ Dependencies ready$(NC)"

## Run all checks
check: fmt vet lint test ## Run all code quality checks
	@echo "$(GREEN)✓ All checks completed$(NC)"

## Install binary to GOPATH/bin
install: build ## Install binary to GOPATH/bin
	@echo "$(YELLOW)Installing $(APP_NAME)...$(NC)"
	@go install $(BUILD_FLAGS) .
	@echo "$(GREEN)✓ Installed to $(shell go env GOPATH)/bin/$(APP_NAME)$(NC)"

## Uninstall binary from GOPATH/bin
uninstall: ## Remove binary from GOPATH/bin
	@echo "$(YELLOW)Uninstalling $(APP_NAME)...$(NC)"
	@rm -f $(shell go env GOPATH)/bin/$(APP_NAME)
	@echo "$(GREEN)✓ Uninstalled$(NC)"

## Create release build
release: clean check build-all ## Create a complete release build
	@echo "$(YELLOW)Creating release build...$(NC)"
	@echo "$(GREEN)✓ Release build completed$(NC)"
	@echo ""
	@echo "$(BLUE)Release artifacts in $(BUILD_DIR):$(NC)"
	@ls -la $(BUILD_DIR)/

## Show version information
version: ## Show version information
	@echo "$(BLUE)BlendPDFGo Version Information$(NC)"
	@echo "App Version: $(VERSION)"
	@echo "Go Version: $(GO_VERSION)"
	@echo "Build Date: $(shell date)"
	@echo "Git Commit: $(shell git rev-parse --short HEAD 2>/dev/null || echo 'unknown')"

## Show supported platforms
platforms: ## Show supported build platforms
	@echo "$(BLUE)Supported Build Platforms:$(NC)"
	@for platform in $(PLATFORMS); do \
		echo "  - $$platform"; \
	done

# Platform-specific targets
windows-amd64: deps ## Build for Windows x86_64
	@echo "$(YELLOW)Building for Windows x86_64...$(NC)"
	@./build.sh windows-amd64
	@echo "$(GREEN)✓ Windows build completed$(NC)"

linux-amd64: deps ## Build for Linux x86_64
	@echo "$(YELLOW)Building for Linux x86_64...$(NC)"
	@./build.sh linux-amd64
	@echo "$(GREEN)✓ Linux x86_64 build completed$(NC)"

linux-arm64: deps ## Build for Linux ARM64
	@echo "$(YELLOW)Building for Linux ARM64...$(NC)"
	@./build.sh linux-arm64
	@echo "$(GREEN)✓ Linux ARM64 build completed$(NC)"

darwin-amd64: deps ## Build for macOS x86_64
	@echo "$(YELLOW)Building for macOS x86_64...$(NC)"
	@./build.sh darwin-amd64
	@echo "$(GREEN)✓ macOS build completed$(NC)"

# Development targets
dev: build ## Quick development build
	@echo "$(GREEN)Development build ready$(NC)"

dev-test: build test ## Development build with tests
	@echo "$(GREEN)Development build and tests completed$(NC)"

# Docker targets (future enhancement)
docker-build: ## Build Docker image (placeholder)
	@echo "$(YELLOW)Docker build not yet implemented$(NC)"

# Benchmark target
benchmark: ## Run benchmarks
	@echo "$(YELLOW)Running benchmarks...$(NC)"
	@go test -bench=. -benchmem ./...
	@echo "$(GREEN)✓ Benchmarks completed$(NC)"

# Coverage target
coverage: ## Generate test coverage report
	@echo "$(YELLOW)Generating coverage report...$(NC)"
	@go test -coverprofile=coverage.out ./...
	@go tool cover -html=coverage.out -o coverage.html
	@echo "$(GREEN)✓ Coverage report generated: coverage.html$(NC)"

# Tidy dependencies
tidy: ## Clean up go.mod and go.sum
	@echo "$(YELLOW)Tidying dependencies...$(NC)"
	@go mod tidy
	@echo "$(GREEN)✓ Dependencies tidied$(NC)"

# Update dependencies
update: ## Update all dependencies
	@echo "$(YELLOW)Updating dependencies...$(NC)"
	@go get -u ./...
	@go mod tidy
	@echo "$(GREEN)✓ Dependencies updated$(NC)"

# Security check
security: ## Run security checks
	@echo "$(YELLOW)Running security checks...$(NC)"
	@if command -v gosec >/dev/null 2>&1; then \
		gosec ./...; \
		echo "$(GREEN)✓ Security check completed$(NC)"; \
	else \
		echo "$(RED)gosec not found. Install with: go install github.com/securecodewarrior/gosec/v2/cmd/gosec@latest$(NC)"; \
	fi

# Show build info
info: ## Show build environment information
	@echo "$(BLUE)Build Environment Information$(NC)"
	@echo "App Name: $(APP_NAME)"
	@echo "Version: $(VERSION)"
	@echo "Go Version: $(GO_VERSION)"
	@echo "Build Directory: $(BUILD_DIR)"
	@echo "GOOS: $(shell go env GOOS)"
	@echo "GOARCH: $(shell go env GOARCH)"
	@echo "CGO_ENABLED: $(shell go env CGO_ENABLED)"
	@echo "Go Root: $(shell go env GOROOT)"
	@echo "Go Path: $(shell go env GOPATH)"
	@echo "Module Path: $(shell go list -m)"
	@echo "Working Directory: $(shell pwd)"
	@echo "Git Branch: $(shell git branch --show-current 2>/dev/null || echo 'unknown')"
	@echo "Git Commit: $(shell git rev-parse --short HEAD 2>/dev/null || echo 'unknown')"

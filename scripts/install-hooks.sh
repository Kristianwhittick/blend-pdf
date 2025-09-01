#!/bin/bash
# Install git hooks for BlendPDF development

set -e

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
PROJECT_ROOT="$(cd "$SCRIPT_DIR/.." && pwd)"
HOOKS_DIR="$PROJECT_ROOT/.git/hooks"

echo "🔧 Installing git hooks..."

# Create pre-commit hook
cat > "$HOOKS_DIR/pre-commit" << 'EOF'
#!/bin/bash
# Git pre-commit hook for BlendPDF
# Runs linting and tests before allowing commits

set -e

echo "🔍 Running pre-commit checks..."

# Check if golangci-lint is available
if ! command -v golangci-lint >/dev/null 2>&1; then
    echo "❌ golangci-lint not found. Install with:"
    echo "   go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest"
    exit 1
fi

# Run linting
echo "🧹 Running linter..."
if ! golangci-lint run; then
    echo "❌ Linting failed. Please fix the issues above."
    exit 1
fi

# Run tests
echo "🧪 Running tests..."
if ! go test ./...; then
    echo "❌ Tests failed. Please fix the issues above."
    exit 1
fi

echo "✅ Pre-commit checks passed!"
EOF

# Make executable
chmod +x "$HOOKS_DIR/pre-commit"

echo "✅ Git hooks installed successfully!"
echo ""
echo "To skip pre-commit checks (not recommended):"
echo "  git commit --no-verify"
echo ""
echo "To install golangci-lint:"
echo "  go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest"

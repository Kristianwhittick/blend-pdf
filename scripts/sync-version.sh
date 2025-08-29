#!/bin/bash

# Sync version between git tags and constants.go
# Usage: ./scripts/sync-version.sh

set -e

# Get the latest git tag
GIT_VERSION=$(git describe --tags --abbrev=0 2>/dev/null || echo "v1.0.0")

# Remove 'v' prefix if present
VERSION=${GIT_VERSION#v}

echo "Syncing version to: $VERSION"

# Update constants.go
sed -i "s/VERSION = \".*\"/VERSION = \"$VERSION\"/" constants.go

echo "✓ Updated constants.go to version $VERSION"

# Verify the change
CURRENT_VERSION=$(grep 'VERSION = ' constants.go | cut -d'"' -f2)
if [ "$CURRENT_VERSION" = "$VERSION" ]; then
    echo "✓ Version sync successful: $VERSION"
else
    echo "✗ Version sync failed. Expected: $VERSION, Got: $CURRENT_VERSION"
    exit 1
fi

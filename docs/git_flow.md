# Git Workflow and Development Process

## 📝 Commit Message Conventions

### Prefixes
- `feat:` - New features
- `fix:` - Bug fixes
- `docs:` - Documentation changes
- `perf:` - Performance improvements
- `refactor:` - Code refactoring
- `test:` - Adding or updating tests
- `chore:` - Maintenance tasks

### Format
```
<type>: <short description>

<detailed description>
- Bullet point 1
- Bullet point 2
- Bullet point 3
```

### Examples
```bash
feat: Add smart PDF processing with page reversal logic

- Only reverse multi-page PDFs (single-page merge directly)
- Add page count detection before processing
- Implement proper temporary file management
- Use pdfcpu merge -mode zip for better merging
```

```bash
fix: Handle corrupted PDF files gracefully

- Add comprehensive PDF validation before processing
- Move invalid files to error directory with descriptive messages
- Continue operation after individual file failures
```

```bash
docs: Update API usage examples

- Add new function examples for memory processing
- Update parameter descriptions
- Include error handling patterns
```

---

## 🎯 Branching Strategy

### Main Branch
- **`main`**: Production-ready code
- All commits should be stable and tested
- Direct commits only for hotfixes and documentation updates

### Feature Development
```bash
# Create feature branch from main
git checkout main
git pull origin main
git checkout -b feature/feature-name

# Make changes and commits
git add .
git commit -m "feat: implement new feature"

# Push and create PR
git push origin feature/feature-name
# Create Pull Request on GitHub
```

### Bug Fixes
```bash
# Create fix branch from main
git checkout main
git pull origin main
git checkout -b fix/bug-description

# Make changes and commits
git add .
git commit -m "fix: resolve specific issue"

# Push and create PR
git push origin fix/bug-description
# Create Pull Request on GitHub
```

### Documentation Updates
```bash
# Create docs branch from main
git checkout main
git pull origin main
git checkout -b docs/update-description

# Make changes and commits
git add .
git commit -m "docs: update documentation"

# Push and create PR (or direct commit for minor updates)
git push origin docs/update-description
```

### Hotfixes
```bash
# Create hotfix branch from main
git checkout main
git pull origin main
git checkout -b hotfix/critical-fix

# Make changes and commits
git add .
git commit -m "fix: critical production issue"

# Push and merge immediately
git push origin hotfix/critical-fix
git checkout main
git merge hotfix/critical-fix
git push origin main
git branch -d hotfix/critical-fix
```

---

## 📋 Pull Request Guidelines

### PR Title Format
```
<type>: <short description>
```

### PR Description Template
```markdown
## Description
Brief description of changes made.

## Type of Change
- [ ] Bug fix (non-breaking change which fixes an issue)
- [ ] New feature (non-breaking change which adds functionality)
- [ ] Breaking change (fix or feature that would cause existing functionality to not work as expected)
- [ ] Documentation update

## Testing
- [ ] Unit tests pass
- [ ] Integration tests pass
- [ ] Manual testing completed

## Checklist
- [ ] Code follows project style guidelines
- [ ] Self-review completed
- [ ] Documentation updated
- [ ] No new warnings introduced
```

---

## 🏷️ Release Process

### Version Numbering
Follow Semantic Versioning (SemVer):
- **MAJOR**: Breaking changes
- **MINOR**: New features (backward compatible)
- **PATCH**: Bug fixes (backward compatible)

### Automated Release Process (Current)

#### Prerequisites
- All changes committed to main branch
- CHANGELOG.md updated with new version entry
- All tests passing

#### Release Steps
1. **Update CHANGELOG.md**
   ```bash
   # Add new version entry to CHANGELOG.md
   ## [1.1.0] - 2025-XX-XX
   
   ### Added
   - New feature descriptions
   
   ### Fixed
   - Bug fix descriptions
   
   ### Changed
   - Change descriptions
   ```

2. **Commit CHANGELOG.md Changes**
   ```bash
   git add CHANGELOG.md
   git commit -m "docs: Add v1.1.0 changelog entry"
   git push origin main
   ```

3. **Sync Version in constants.go**
   ```bash
   # Update constants.go to match the version you're about to tag
   ./scripts/sync-version.sh  # This will sync to latest tag, so run after step 4
   # OR manually update constants.go VERSION = "1.1.0"
   git add constants.go
   git commit -m "chore: Sync constants.go version to v1.1.0"
   git push origin main
   ```

4. **Create and Push Git Tag**
   ```bash
   git checkout main
   git pull origin main
   git tag -a v1.1.0 -m "Release version 1.1.0"
   git push origin main --tags
   ```

5. **Monitor GitHub Actions** (requires GitHub CLI)
   ```bash
   # Check recent workflow runs
   gh run list --limit 5
   
   # View specific run details
   gh run view <run-id>
   
   # Check release status
   gh release list
   gh release view <tag>
   
   # View workflow logs if needed
   gh run logs <run-id>
   ```

#### ⚠️ Critical: Tag Must Include CHANGELOG Entry and Version Sync
- **The git tag must point to a commit that includes both the CHANGELOG.md entry AND constants.go version update**
- **GitHub Actions checks out code at the tag commit, not latest main**
- **If CHANGELOG.md entry is missing at tag commit, release will show generic message**
- **If constants.go version is wrong at tag commit, build artifacts will show incorrect version**

3. **Automated GitHub Actions**
   - GitHub Actions automatically triggers on tag push
   - Extracts version from git tag (not constants.go)
   - Builds binaries for all platforms (Windows, Linux, macOS)
   - Extracts changelog content for the specific version
   - Creates GitHub release with proper version and changelog
   - Uploads binaries and checksums

#### Important Notes
- **Version Source**: Git tags are the single source of truth for versions
- **Automatic Sync**: `scripts/sync-version.sh` syncs constants.go with git tags
- **Changelog Extraction**: GitHub Actions automatically extracts version-specific changelog
- **No Manual GitHub Release**: GitHub releases are created automatically

### Troubleshooting Releases

#### Version Mismatch Issues
- **Problem**: Build artifacts show wrong version (e.g., v1.0.0 instead of v1.0.3)
- **Cause**: Version not synced between git tags and constants.go
- **Solution**: Run `./scripts/sync-version.sh` or `make sync-version`

#### Missing Changelog in Release
- **Problem**: GitHub release shows "See commit history" instead of actual changes
- **Cause**: Missing or incorrectly formatted CHANGELOG.md entry
- **Solution**: Ensure CHANGELOG.md has proper `## [X.Y.Z] - YYYY-MM-DD` format

#### GitHub Actions Not Triggering
- **Problem**: No automated release created after pushing tag
- **Cause**: Tag not pushed or GitHub Actions workflow issues
- **Solution**: Check Actions tab on GitHub, ensure tag format is `vX.Y.Z`
- **GitHub CLI**: Use `gh run list` to check recent runs, `gh run view <run-id>` for details

### GitHub Actions Monitoring Commands
```bash
# Check recent workflow runs
gh run list --limit 5

# View specific run details
gh run view <run-id>

# Check release status
gh release list
gh release view <tag>

# View workflow logs if needed
gh run logs <run-id>
```

### Critical Release Requirements
- Git tag must include both CHANGELOG.md entry and VERSION constant update
- GitHub Actions extracts changelog content from the tagged commit
- Version sync must happen before tagging, not after

#### Local Build Version Issues
- **Problem**: Local builds show wrong version
- **Cause**: Build script not using git tags
- **Solution**: Use `make build-all` or `./build.sh --all` (includes automatic sync)

---

## 🛠️ Development Tools

### Recommended Git Configuration
```bash
# Set up user information
git config --global user.name "Your Name"
git config --global user.email "your.email@example.com"

# Set up useful aliases
git config --global alias.co checkout
git config --global alias.br branch
git config --global alias.ci commit
git config --global alias.st status
git config --global alias.unstage 'reset HEAD --'
git config --global alias.last 'log -1 HEAD'
git config --global alias.visual '!gitk'

# Set up default branch
git config --global init.defaultBranch main
```

### Useful Git Commands
```bash
# View commit history
git log --oneline --graph --decorate --all

# View changes in staging area
git diff --staged

# Interactive rebase for cleaning up commits
git rebase -i HEAD~3

# Cherry-pick specific commit
git cherry-pick <commit-hash>

# Stash changes temporarily
git stash push -m "work in progress"
git stash pop

# View file history
git log --follow -p -- <filename>
```

---

## 📊 Code Quality Standards

### Go Formatting
```bash
# Format code
go fmt ./...

# Run linter
golangci-lint run

# Run tests with coverage
go test -v -cover ./...
```

### Commit Quality
- Each commit should be atomic (single logical change)
- Commit messages should be clear and descriptive
- Include context and reasoning in commit body
- Reference issues when applicable (#123)

### Code Review Checklist
- [ ] Code is readable and well-documented
- [ ] No hardcoded values or magic numbers
- [ ] Error handling is comprehensive
- [ ] Tests cover new functionality
- [ ] Performance considerations addressed
- [ ] Security implications considered

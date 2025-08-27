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

## 🔄 Development Workflow

### 1. Planning Phase
- Create or update issues in GitHub
- Define acceptance criteria
- Estimate effort and priority
- Assign to milestone if applicable

### 2. Development Phase
- Create feature branch from main
- Implement changes following coding standards
- Write tests for new functionality
- Update documentation as needed

### 3. Testing Phase
- Run all existing tests
- Test new functionality thoroughly
- Verify no regressions introduced
- Test on different platforms if applicable

### 4. Review Phase
- Create Pull Request with detailed description
- Request review from team members
- Address feedback and make necessary changes
- Ensure CI/CD checks pass

### 5. Merge Phase
- Squash commits if multiple small commits
- Use descriptive merge commit message
- Delete feature branch after merge
- Update local main branch

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

### Release Steps
1. **Prepare Release**
   ```bash
   git checkout main
   git pull origin main
   git checkout -b release/v1.1.0
   ```

2. **Update Version**
   - Update version in `constants.go`
   - Update CHANGELOG.md
   - Update documentation if needed

3. **Create Release**
   ```bash
   git add .
   git commit -m "chore: prepare release v1.1.0"
   git push origin release/v1.1.0
   ```

4. **Merge and Tag**
   ```bash
   git checkout main
   git merge release/v1.1.0
   git tag -a v1.1.0 -m "Release version 1.1.0"
   git push origin main --tags
   ```

5. **Create GitHub Release**
   - Go to GitHub Releases
   - Create new release from tag
   - Add release notes
   - Attach binaries if applicable

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

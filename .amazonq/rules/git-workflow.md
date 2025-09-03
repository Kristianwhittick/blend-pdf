# Git Workflow Standards

## Branch Naming Convention
- **Feature branches**: `feature/TASK-123-short-description` or `feature/<feature-name>`
- **Bug fixes**: `bugfix/TASK-456-issue-description` or `fix/<issue-name>`
- **Hotfixes**: `hotfix/critical-issue-description`
- **Release branches**: `release/v1.2.3`
- Use lowercase, hyphens for separation, include ticket number when available
- Create a new branch for each feature or issue from `main`

## Commit Message Standards
- Format: `type(scope): description` or `<type>: <description>`
- Types: `feat`, `fix`, `docs`, `style`, `refactor`, `test`, `chore`
- Keep subject line under 50 characters
- Use imperative mood ("Add feature" not "Added feature")
- Include ticket reference in body when applicable
- Make frequent commits after each significant change
- Each commit should represent a single logical change
- Keep commits small and focused on a specific task

## Merge Workflow
- Always create pull requests for feature branches
- Require at least one code review before merging
- Address all review comments before proceeding
- Rebase feature branch on latest `main` to ensure clean history
- Use `--ff-only` merge to maintain linear history or squash and merge for feature branches
- Delete feature branches after successful merge
- Keep main branch always deployable and stable
- Direct commits to `main` are prohibited except for hotfixes

## Q Developer Git Actions
When working on code changes, Q Developer should:

1. **Before making changes**: Check current branch and suggest creating feature branch if on main
2. **After code modifications**: Automatically stage relevant files and suggest appropriate commit messages
3. **When completing features**: Remind about creating pull request and proper branch cleanup
4. **For commit messages**: Generate semantic commit messages following the established format
5. **Branch management**: Suggest switching to appropriate branch based on task type
6. **Commit timing**: Suggest commits after implementing features, fixing bugs, refactoring, writing tests, or updating documentation

## Detailed Merge Example
```bash
git checkout feature/my-feature
git rebase main
git checkout main
git merge --ff-only feature/my-feature
```

## Best Practices & Automation Guidelines
- Always run tests before committing
- **Run linting before committing** - Use pre-commit hooks where possible
- Pull latest changes before pushing to remote
- Check for merge conflicts before suggesting merges
- Validate branch naming when creating new branches
- Ensure no sensitive data in commits
- Suggest rebasing when branch is behind main
- Each commit should leave the codebase in a working state
- Never rewrite public history - only rebase branches that haven't been shared
- Make atomic commits and commit early and often

## Development Environment Setup
- **Install linting tools** appropriate for the technology stack
- **Set up pre-commit hooks** to run linting and tests automatically
- **Configure quality gates** to ensure code standards are maintained
- **Use automated tools** like golangci-lint, ESLint, pylint, etc.
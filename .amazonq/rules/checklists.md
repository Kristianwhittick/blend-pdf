# Operation Checklists

## Before Every Response
- [ ] Use UK English spelling (behaviour, colour, optimise, realise, centre, licence, analyse)
- [ ] Check if task documentation needs updating
- [ ] Verify if CHANGELOG.md needs updating
- [ ] Consider if README.md needs updating
- [ ] Follow project coding standards and conventions

## Code Changes Checklist
- [ ] Use UK English in all comments and documentation
- [ ] Update relevant task in docs/tasks.md if completing/fixing something
- [ ] Add entry to CHANGELOG.md if user-facing change
- [ ] Update README.md if new features or significant changes
- [ ] Run tests before committing
- [ ] Use proper commit message format with UK spelling
- [ ] Follow project-specific coding standards

## Bug Fixes Checklist
- [ ] Document the fix in the relevant task (T-XXX completion notes)
- [ ] Add fix to CHANGELOG.md Fixed section
- [ ] Use UK English in all documentation updates
- [ ] Update any affected documentation
- [ ] Test the fix thoroughly
- [ ] Verify fix doesn't break existing functionality

## New Features Checklist
- [ ] Update task status in docs/tasks.md
- [ ] Add feature to CHANGELOG.md Added section
- [ ] Update README.md features section if significant
- [ ] Use UK English throughout documentation
- [ ] Include proper documentation and comments
- [ ] Add appropriate tests for new functionality

## Documentation Updates Checklist
- [ ] Use UK English spelling throughout
- [ ] Use hyphens (-) not underscores (_) in markdown filenames
- [ ] Follow project documentation standards
- [ ] Update all affected files (tasks.md, CHANGELOG.md, README.md)
- [ ] Ensure consistency across all documentation
- [ ] Check cross-references are still valid
- [ ] Verify formatting and structure

## Release Process Checklist
- [ ] Update CHANGELOG.md with new version entry using UK English
- [ ] Sync version in constants.go to match git tag (Go projects)
- [ ] Create version tag following project convention (vX.Y.Z format)
- [ ] Push tag to trigger GitHub Actions
- [ ] Verify release artifacts and checksums
- [ ] Test release packages

## Debugging Process Checklist
- [ ] Create debug programs for complex issues
- [ ] Document findings in task completion notes
- [ ] Update relevant documentation with solutions
- [ ] Test fixes across all affected platforms
- [ ] Commit fixes with descriptive messages

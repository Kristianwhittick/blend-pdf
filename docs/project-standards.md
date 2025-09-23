# Blend-PDF Project Standards

## Go Project Standards

### Experiment Organisation
- **Individual Folders**: Each experiment in separate folder (`exp01/`, `exp02/`, etc.) to prevent Go linting warnings
- **Unified Runner**: `experiments/run_experiments.go` handles execution of all experiments
- **Naming Convention**: `experiment##_description.go` within each folder
- **Usage**: `go run experiments/run_experiments.go ##`

### API Research Process
1. **Create Experiment**: New experiment in next available `exp##/` folder
2. **Update Runner**: Add case to `run_experiments.go` switch statement
3. **Document Results**: Update `docs/api-knowledge.md` with findings
4. **Commit Changes**: Follow git workflow for experiment commits

### Testing Integration
- Experiments excluded from main test suite (no test files in experiment folders)
- API validation through experimental testing rather than unit tests
- Results documented in knowledge base for reference

## Development Workflow Additions

### Experiment-Driven Development
- Use experiments to validate API assumptions before implementation
- Document breakthrough discoveries in task updates
- Reference experiment numbers in commit messages and documentation

### Knowledge Management
- `docs/api-knowledge.md`: Comprehensive API function reference
- `docs/api-experiments-procedures.md`: Step-by-step testing procedures
- Experiment results inform implementation decisions

## Release Process (Project-Specific)
- [ ] Update CHANGELOG.md with new version entry
- [ ] Sync version in constants.go to match git tag
- [ ] Create git tag (vX.Y.Z format)
- [ ] Push tag to trigger GitHub Actions
- [ ] Verify release artifacts and checksums
- [ ] Test downloaded binaries

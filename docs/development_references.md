# Development References

## 📋 Phase 5 Implementation References

### Task 16: Directory-Specific Lock Files
- **Hash Algorithm**: MD5 truncated to 8 characters (32 bits)
- **Collision Analysis**: 4.3 billion combinations, 50% collision at ~65,000 paths
- **Cross-Platform**: Unix uses `/tmp/`, Windows uses watch folder
- **Path Normalization**: Absolute, cleaned, lowercase, forward slashes
- **Implementation**: Update `setup.go` lock functions

### Documentation
- **Detailed specifications**: See `docs/tasks.md` (Task 14)
- **Git workflow**: See `docs/git_flow.md` for branching and commit strategies
- **API reference**: See `docs/api_knowledge.md` for pdfcpu functions
- **Implementation pattern**: See `docs/memory_processing_summary.md`

### Code Examples
- **Working example**: See `tests/test16_final_memory_approach.go`
- **API tests**: See `tests/test09_memory_context.go` through `test16_final_memory_approach.go`
- **Research code**: All test files in `/tests/` directory
- **Page reversal fix**: Individual page extraction instead of comma-separated selections

### Implementation Guidelines
- Follow hybrid approach from research
- Maintain existing user interface and error handling
- Preserve all functionality while optimizing performance
- Use comprehensive test coverage for validation

## 📚 General Development Resources

### Project Structure
- **Main code**: `main.go`, `constants.go`, `setup.go`, `pdfops.go`, `fileops.go`
- **Documentation**: `/docs/` directory with specialized files
- **Tests**: `/tests/` directory with API research and validation
- **Examples**: Working code examples for all major features

### Development Workflow
- See `docs/git_flow.md` for complete branching and commit strategies
- Follow semantic versioning for releases
- Maintain backward compatibility for all changes
- Use structured commit messages with appropriate prefixes

### Testing Resources
- **Test plan**: `docs/TEST.md` with 140+ comprehensive test cases
- **API validation**: Individual test programs for each pdfcpu function
- **Integration tests**: Full workflow validation procedures
- **Performance benchmarks**: Memory and speed optimization validation

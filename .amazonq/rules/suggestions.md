# Methodology Suggestions and Improvements

## New Working Practices
When you discover effective new practices during project work:

1. **Document the practice** in your project's AmazonQ.md
2. **Test effectiveness** over multiple uses
3. **Evaluate applicability** to other project types
4. **Suggest addition** to main methodology if valuable

## Suggestion Categories
- **Process Improvements**: Better task management or workflow approaches
- **Tool Recommendations**: Useful development tools and configurations  
- **Quality Practices**: Testing, documentation, and code quality methods
- **Collaboration Techniques**: Team communication and coordination methods

## Current Suggestions for Evaluation

### Validated Practices from Real Projects
1. **Experiment-Driven Development**: Create `/experiments/` directory for API research and proof-of-concepts
   - Use consistent naming: `experimentXX_description.go`
   - Document results in `docs/api_knowledge.md`
   - Unified runner for consistent execution
   
2. **Breakthrough Pattern Recognition**: Look for native API functions that eliminate complex workarounds
   - Research newer/better APIs before implementing complex solutions
   - Document breakthrough discoveries that simplify implementation
   
3. **CI/CD Integration**: Real-time monitoring and release management
   - Use CLI tools for workflow status monitoring
   - Automated build and deployment processes
   - Release verification and validation
   
4. **Comprehensive API Testing**: Systematic validation of all external library functions
   - Number experiments sequentially
   - Document success/failure status clearly
   - Create unified test runner for consistency
   
5. **Task Summary with Counts**: Maintain real-time task counts in Kanban headers
   - Format: "Task Summary (X Total)" with breakdown by status
   - Update counts as tasks move between columns
   - Include project status indicator
   
6. **Production Readiness Indicators**: Clear project status communication
   - Use status indicators: "PRODUCTION READY", "IN DEVELOPMENT", etc.
   - Document feature parity achievements
   - Track implementation status overview
   
7. **Performance Research**: Document optimization opportunities
   - Research and document potential improvements
   - Create feature requests for external libraries
   - Track optimization opportunities in backlog
   
8. **Structured Completion Notes**: Detailed task completion documentation
   - Include actual time vs. estimate
   - Document key implementation decisions
   - Note breakthrough discoveries and solutions
   
9. **Cross-Platform Build Integration**: Automated multi-platform releases
   - Automated builds for multiple platforms
   - Platform-specific optimizations
   - Consistent deployment across platforms

## Migration Guide for Existing Projects
1. **Backup current structure** before making changes
2. **Choose appropriate methodology** (Light vs Full)
3. **Map existing files** to new naming convention
4. **Convert requirements to user stories** format
5. **Restructure tasks** into Kanban format
6. **Update AmazonQ.md** to new template
7. **Test new structure** with a few development cycles

### File Mapping Examples
- `requirements.md` → Keep as initial requirements, create new `stories.md`
- `tasks.md` → Restructure into Kanban format with Epic/Story references
- `technical-architecture.md` → Rename to `design.md`
- `user-stories.md` → Rename to `stories.md`

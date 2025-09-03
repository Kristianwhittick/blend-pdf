# Documentation and Code Commenting Standards

## Code Comments
- Write comments that explain **why**, not **what**
- Use clear, concise language in comments
- Comment complex algorithms and business logic
- Avoid obvious comments that restate the code
- Update comments when code changes
- Use consistent comment style within each language

## Function/Method Documentation
- Document all public functions and methods
- Include purpose, parameters, return values, and exceptions
- Use language-specific documentation formats (JSDoc, docstrings, etc.)
- Provide usage examples for complex functions
- Document side effects and assumptions

## README Files
- Every project must have a README.md in the root directory
- Include project purpose and description
- List key dependencies and versions
- Provide installation and setup instructions
- Include usage examples and basic API documentation
- Document environment variables and configuration
- Add troubleshooting section for common issues

## Q Developer Documentation Actions
When creating or modifying code, Q Developer should:

1. **Code Creation**: Add appropriate comments for complex logic and public interfaces
2. **README Generation**: Create or update README files when adding new features or dependencies
3. **Dependency Documentation**: Document new dependencies in README and package files
4. **API Documentation**: Generate function/method documentation using appropriate formats
5. **Comment Maintenance**: Update existing comments when modifying related code
6. **Documentation Commits**: Suggest separate commits for documentation updates using `docs:` type

## Documentation Standards by Language
- **JavaScript/TypeScript**: Use JSDoc format
- **Python**: Use docstrings following PEP 257
- **Java**: Use Javadoc format
- **Go**: Use standard Go doc comments
- **Markdown**: Use consistent heading structure and formatting

## Best Practices
- Keep documentation close to the code it describes
- Use examples to illustrate complex concepts
- Maintain documentation as part of the development process
- Review documentation during code reviews
- Ensure documentation is accessible to new team members
- **UK Spelling**: Use British English spelling where possible (colour, optimise, realise, behaviour, etc.)
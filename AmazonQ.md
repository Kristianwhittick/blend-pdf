# Amazon Q Development Methodology

## ⚠️ CRITICAL: UK ENGLISH SPELLING REQUIREMENT

**MANDATORY**: Use UK English spelling in ALL text, documentation, comments, and commit messages:
- behaviour (not behavior)
- colour (not color)
- optimise (not optimize)
- realise (not realize)
- centre (not center)
- licence (not license as noun)
- analyse (not analyze)

**PROCESS**: Check every response and commit for US spellings before submitting.

## Operation Checklists

### Before Every Response
- [ ] Use UK English spelling (behaviour, colour, optimise, realise, centre, licence, analyse)
- [ ] Check if task documentation needs updating
- [ ] Verify if CHANGELOG.md needs updating
- [ ] Consider if README.md needs updating

### Code Changes Checklist
- [ ] Use UK English in all comments and documentation
- [ ] Update relevant task in docs/tasks.md if completing/fixing something
- [ ] Add entry to CHANGELOG.md if user-facing change
- [ ] Update README.md if new features or significant changes
- [ ] Run tests before committing
- [ ] Use proper commit message format with UK spelling

### Bug Fixes Checklist
- [ ] Document the fix in the relevant task (T-XXX completion notes)
- [ ] Add fix to CHANGELOG.md Fixed section
- [ ] Use UK English in all documentation updates
- [ ] Test the fix thoroughly
- [ ] Update any affected documentation

### New Features Checklist
- [ ] Update task status in docs/tasks.md
- [ ] Add feature to CHANGELOG.md Added section
- [ ] Update README.md features section if significant
- [ ] Use UK English throughout
- [ ] Include proper documentation and comments

### Documentation Updates Checklist
- [ ] Use UK English spelling throughout
- [ ] Update all affected files (tasks.md, CHANGELOG.md, README.md)
- [ ] Ensure consistency across all documentation
- [ ] Check cross-references are still valid

## Project-Specific Methodologies

### Go Project Standards

#### UK English Spelling - CRITICAL REQUIREMENT ⚠️
**ALWAYS use UK English spelling in ALL documentation, comments, and user-facing text:**
- behaviour (not behavior)
- colour (not color) 
- optimise (not optimize)
- realise (not realize)
- centre (not center)
- licence (not license as noun)
- analyse (not analyze)

**Check EVERY commit for US spellings before committing.**

#### Experiment Organization
- **Individual Folders**: Each experiment in separate folder (`exp01/`, `exp02/`, etc.) to prevent Go linting warnings
- **Unified Runner**: `experiments/run_experiments.go` handles execution of all experiments
- **Naming Convention**: `experiment##_description.go` within each folder
- **Usage**: `go run experiments/run_experiments.go ##`

#### API Research Process
1. **Create Experiment**: New experiment in next available `exp##/` folder
2. **Update Runner**: Add case to `run_experiments.go` switch statement
3. **Document Results**: Update `docs/api_knowledge.md` with findings
4. **Commit Changes**: Follow git workflow for experiment commits

#### Testing Integration
- Experiments excluded from main test suite (no test files in experiment folders)
- API validation through experimental testing rather than unit tests
- Results documented in knowledge base for reference

### Development Workflow Additions

#### Experiment-Driven Development
- Use experiments to validate API assumptions before implementation
- Document breakthrough discoveries in task updates
- Reference experiment numbers in commit messages and documentation

#### Knowledge Management
- `docs/api_knowledge.md`: Comprehensive API function reference
- `docs/api_experiments_procedures.md`: Step-by-step testing procedures
- Experiment results inform implementation decisions

---

## Overview
This document defines the standardized approach to development using Amazon Q Developer. Choose between Light and Full methodologies based on your project needs.

---

## Quick Start

When starting a new project, run the setup process:
1. Navigate to your project directory
2. Request "spec-driven development setup"
3. Choose Light or Full methodology
4. Answer setup questions
5. Begin development with generated structure

---

## Methodology Selection

### Light Methodology
**Best for**: Scripts, prototypes, personal projects, quick experiments
- **Setup time**: 5-10 minutes
- **Files created**: 6-7 core files
- **Focus**: Get started quickly with minimal overhead

### Full Methodology  
**Best for**: Applications, libraries, team projects, production systems
- **Setup time**: 15-30 minutes
- **Files created**: 8-12 comprehensive files
- **Focus**: Comprehensive planning and documentation

---

## Core Agile Concepts

### Essential Terms
- **Epic**: Large feature or capability (collection of related user stories)
- **User Story**: Specific user requirement ("As a [user], I want [goal] so that [benefit]")
- **Definition of Done**: Clear acceptance criteria for completion
- **Task**: Implementation work item with specific deliverables
- **Backlog**: Prioritized list of work items (Epics, Stories, Tasks)

### Kanban Board Structure
- **📋 To Do**: Ready for work, prioritized
- **🔄 In Progress**: Currently being worked on
- **✅ Done**: Completed and validated
- **🗂️ Backlog**: Future work, not yet prioritized

---

## File Structure Standards

### Light Methodology Structure
```
project-name/
├── README.md
├── AmazonQ.md
├── summary.txt                 # Daily progress summary
├── .amazonq/
│   └── rules/
│       ├── git-workflow.md
│       └── documentation-standards.md
├── docs/
│   ├── requirements.md         # Initial ideas and needs
│   ├── overview.md            # Project vision and goals
│   ├── stories.md             # User stories
│   ├── tasks.md               # Kanban task board
│   └── testing.md             # Basic test checklist
├── src/                       # Source code
└── tests/                     # Test files
```

### Full Methodology Structure
```
project-name/
├── README.md
├── AmazonQ.md
├── summary.txt                 # Daily progress summary
├── .amazonq/
│   └── rules/
│       ├── git-workflow.md
│       └── documentation-standards.md
├── docs/
│   ├── requirements.md         # Initial ideas and needs
│   ├── overview.md            # Project vision and goals
│   ├── features.md            # Epics and major features
│   ├── stories.md             # User stories
│   ├── tasks.md               # Kanban task board
│   ├── design.md              # Technical architecture
│   ├── testing.md             # Test strategy and procedures
│   └── standards/
│       ├── coding.md
│       ├── workflow.md
│       └── quality.md
├── completion-summaries/       # Excluded from git
│   ├── T-XXX-COMPLETION-SUMMARY.md
│   └── EPIC-X-COMPLETION-SUMMARY.md
├── src/                       # Source code (varies by project type)
└── tests/                     # Test files (varies by project type)
```

---

## Setup Process

### Interactive Questions

#### Initial Questions (Both Versions)
1. **Methodology Type**: Light or Full?
2. **Project Name**: What should we call this project?
3. **Technology Stack**: Primary languages/frameworks?
4. **Team Size**: Solo developer or team project?

#### Light Version Questions
5. **Primary Goal**: Main purpose of this project?
6. **Timeline**: Quick prototype or ongoing development?
7. **Documentation Level**: Minimal or standard?

#### Full Version Questions  
5. **Project Category**: Web app, mobile app, library, API, etc.?
6. **Development Approach**: Agile sprints or continuous development?
7. **Quality Requirements**: Standard or high-compliance?
8. **Team Standards**: Need coding standards and review processes?
9. **Testing Strategy**: Unit tests only or comprehensive testing?
10. **Documentation Depth**: Standard or comprehensive?

---

## Development Workflow

### Phase 0: Project Setup
1. **Run interactive setup** to create project structure
2. **Initialize version control** with proper .gitignore
3. **Copy default standards** from `~/.aws/amazonq/rules/` to `.amazonq/rules/`:
   - `git-workflow.md` - Git branching, commits, and merge standards
   - `documentation-standards.md` - Code commenting and README requirements
4. **Create initial documentation** from templates
5. **Set up development environment** and tools

### Phase 1: Requirements and Planning
1. **Capture initial requirements** in requirements.md
2. **Define project overview** with vision and goals
3. **Break down into Epics** (Full methodology) or go directly to stories (Light)
4. **Write user stories** with acceptance criteria
5. **Create initial task backlog** with priorities

### Phase 2: Development Cycles
1. **Select tasks** from backlog for current work
2. **Move tasks** through Kanban board (To Do → In Progress → Done)
3. **Implement and test** each task completely
4. **Update documentation** as you learn and evolve
5. **Review and retrospect** regularly to improve process

### Phase 3: Completion and Handoff
1. **Validate all Definition of Done** criteria met
2. **Complete final testing** and quality checks
3. **Update all documentation** to reflect final state
4. **Create completion summaries** for future reference
5. **Archive or deploy** as appropriate

---

## Quality Standards

### Definition of Done (Minimum)
- [ ] Functionality works as specified in user story
- [ ] Code follows project coding standards
- [ ] Basic tests written and passing
- [ ] Documentation updated to reflect changes
- [ ] Code reviewed (team projects)

### Definition of Done (Comprehensive - Full Methodology)
- [ ] All acceptance criteria met
- [ ] Unit tests written with good coverage
- [ ] Integration tests passing
- [ ] Performance requirements met
- [ ] Security considerations addressed
- [ ] Accessibility requirements met (if applicable)
- [ ] Documentation comprehensive and current
- [ ] Code review completed and approved

---

## Task Management

### Task Lifecycle
1. **Creation**: Document in tasks.md with Epic/Story reference
2. **Planning**: Add acceptance criteria and estimates
3. **Implementation**: Move to In Progress, implement solution
4. **Testing**: Validate against Definition of Done
5. **Review**: Code review and quality checks (team projects)
6. **Completion**: Move to Done, update documentation

### Task Format
```markdown
### T-XXX: Task Title [Status]
**Epic**: E-XX | **Story**: US-XX
**Estimate**: X hours/days
**Priority**: High/Medium/Low
**Dependencies**: T-YYY, T-ZZZ

**Description**: What needs to be done

**Acceptance Criteria**:
- [ ] Specific measurable outcome 1
- [ ] Specific measurable outcome 2
- [ ] Specific measurable outcome 3

**Definition of Done**:
- [ ] Functionality implemented and tested
- [ ] Documentation updated
- [ ] Code reviewed (if team project)
```

---

## Daily Progress Management

### Summary.txt File
- **Auto-update**: Created/updated with "Perfect!" responses
- **Content**: Current status, achievements, next steps, constraints
- **Format**: Structured sections for quick project overview

### Progress Tracking
- **Task Board Updates**: Move tasks between Kanban columns
- **Completion Summaries**: Detailed task completion documentation
- **Regular Reviews**: Weekly retrospectives to improve process

---

## Suggestions and Improvements

### New Working Practices
When you discover effective new practices during project work:

1. **Document the practice** in your project's AmazonQ.md
2. **Test effectiveness** over multiple uses
3. **Evaluate applicability** to other project types
4. **Suggest addition** to main methodology if valuable

### Suggestion Categories
- **Process Improvements**: Better task management or workflow approaches
- **Tool Recommendations**: Useful development tools and configurations  
- **Quality Practices**: Testing, documentation, and code quality methods
- **Collaboration Techniques**: Team communication and coordination methods

### Current Suggestions for Evaluation
*This section will be populated as new practices are discovered and tested*

---

## Migration Guide

### For Existing Projects
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

---

## Success Metrics

### Process Effectiveness
- **Setup Speed**: Light ≤ 10 min, Full ≤ 30 min
- **Clarity**: Easy to understand what to do next
- **Consistency**: Same approach across all projects
- **Adaptability**: Works for different project types and sizes

### Quality Outcomes
- **Documentation Coverage**: All key decisions and designs documented
- **Task Completion**: Clear progress tracking and completion criteria
- **Knowledge Retention**: Easy to resume work after breaks
- **Team Coordination**: Effective collaboration (team projects)

---

This methodology evolves based on practical usage and feedback. Document successful patterns and suggest improvements to keep it effective and current.

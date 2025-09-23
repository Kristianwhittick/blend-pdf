# Amazon Q Development Methodology

## ⚠️ CRITICAL RULE #1
**Amazon Q MUST follow the appropriate processes and checklists for EVERY prompt given during development work.**

- Check `.amazonq/rules/checklists.md` for the relevant checklist before responding
- Follow the process steps defined in `.amazonq/rules/process-map.md`
- Apply coding and documentation standards from `.amazonq/rules/documentation-standards.md`
- Use proper task formats from `.amazonq/rules/task-formats.md`

**This is not optional - it ensures consistency, quality, and proper documentation across all development activities.**

---

## Overview
This document defines the standardised approach to development using Amazon Q Developer. Choose between Light and Full methodologies based on your project needs.

**Master Files**: All detailed standards, checklists, and templates are maintained in `~/.aws/amazonq/rules/` and copied into projects during setup.

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
- **Backlog**: Prioritised list of work items (Epics, Stories, Tasks)

### Kanban Board Structure
- **📋 To Do**: Ready for work, prioritised
- **🔄 In Progress**: Currently being worked on
- **✅ Done**: Completed and validated
- **🗂️ Backlog**: Future work, not yet prioritised

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
2. **Initialise version control** with proper .gitignore
3. **Copy master standards files** from `~/.aws/amazonq/rules/` to project `.amazonq/rules/`:
   ```bash
   mkdir -p .amazonq/rules/
   cp ~/.aws/amazonq/rules/git-workflow.md .amazonq/rules/
   cp ~/.aws/amazonq/rules/documentation-standards.md .amazonq/rules/
   cp ~/.aws/amazonq/rules/process-map.md .amazonq/rules/
   cp ~/.aws/amazonq/rules/checklists.md .amazonq/rules/
   cp ~/.aws/amazonq/rules/project-templates.md .amazonq/rules/
   cp ~/.aws/amazonq/rules/task-formats.md .amazonq/rules/
   cp ~/.aws/amazonq/rules/suggestions.md .amazonq/rules/
   ```
   These provide:
   - `git-workflow.md` - Git branching, commits, and merge standards
   - `documentation-standards.md` - Code commenting and README requirements
   - `process-map.md` - Complete process tree and detailed cross-references
   - `checklists.md` - All operation checklists for consistent quality
   - `project-templates.md` - File structure templates for both methodologies
   - `task-formats.md` - Standard formats for tasks, stories, and epics
   - `suggestions.md` - Current methodology improvements and migration guides
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

## Project-Specific Standards

See `docs/project-standards.md` for this project's specific standards including:
- Go experiment organisation and API research process
- Experiment-driven development workflow
- Knowledge management approach
- Project-specific release process

---

## Reference Files

All detailed information is maintained in master files:

- **`.amazonq/rules/checklists.md`** - Operation checklists for all development activities
- **`.amazonq/rules/project-templates.md`** - File structure templates and setup commands
- **`.amazonq/rules/task-formats.md`** - Standard formats for tasks, stories, and epics
- **`.amazonq/rules/suggestions.md`** - Current methodology improvements and migration guides
- **`.amazonq/rules/git-workflow.md`** - Version control standards and practices
- **`.amazonq/rules/documentation-standards.md`** - Code commenting and documentation requirements
- **`.amazonq/rules/process-map.md`** - Complete process tree and cross-references

This methodology evolves based on practical usage and feedback. Document successful patterns and suggest improvements to keep it effective and current.

# Project Templates

## Light Methodology Structure
```
project-name/
├── README.md
├── AmazonQ.md
├── summary.txt                 # Daily progress summary
├── .amazonq/
│   └── rules/
│       ├── git-workflow.md
│       ├── documentation-standards.md
│       ├── process-map.md
│       ├── checklists.md
│       ├── project-templates.md
│       └── task-formats.md
├── docs/
│   ├── requirements.md         # Initial ideas and needs
│   ├── overview.md            # Project vision and goals
│   ├── stories.md             # User stories
│   ├── tasks.md               # Kanban task board
│   └── testing.md             # Basic test checklist
├── src/                       # Source code
└── tests/                     # Test files
```

## Full Methodology Structure
```
project-name/
├── README.md
├── AmazonQ.md
├── summary.txt                 # Daily progress summary
├── .amazonq/
│   └── rules/
│       ├── git-workflow.md
│       ├── documentation-standards.md
│       ├── process-map.md
│       ├── checklists.md
│       ├── project-templates.md
│       └── task-formats.md
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

## Setup Commands

### Copy Master Files
```bash
mkdir -p .amazonq/rules/
cp ~/.aws/amazonq/rules/git-workflow.md .amazonq/rules/
cp ~/.aws/amazonq/rules/documentation-standards.md .amazonq/rules/
cp ~/.aws/amazonq/rules/process-map.md .amazonq/rules/
cp ~/.aws/amazonq/rules/checklists.md .amazonq/rules/
cp ~/.aws/amazonq/rules/project-templates.md .amazonq/rules/
cp ~/.aws/amazonq/rules/task-formats.md .amazonq/rules/
```

## Quality Standards Templates

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

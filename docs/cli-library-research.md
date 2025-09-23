# Go CLI Library Research - Task 20

## Overview
Research of Go CLI libraries for enhancing BlendPDFGo's user experience with interactive menus, progress indicators, and advanced styling.

## Current State
BlendPDFGo uses basic `fmt.Printf` with ANSI color codes for output formatting and `bufio.Scanner` for input handling.

## Libraries Evaluated

### 1. Cobra + Viper ⭐ RECOMMENDED
**Repository**: https://github.com/spf13/cobra  
**Stars**: 37k+ | **Maintained**: Active

**Pros**:
- Industry standard (used by kubectl, docker, hugo)
- Excellent subcommand structure
- Built-in help generation
- Flag parsing and validation
- Viper integration for config files
- Completion generation (bash, zsh, fish)

**Cons**:
- Overkill for simple interactive menus
- More suited for command-based CLIs than interactive loops

**Use Case**: Better suited for expanding CLI commands rather than current interactive menu

### 2. Bubble Tea ⭐ RECOMMENDED
**Repository**: https://github.com/charmbracelet/bubbletea  
**Stars**: 27k+ | **Maintained**: Very Active

**Pros**:
- Modern TUI framework with Elm architecture
- Rich interactive components
- Excellent for real-time updates
- Great ecosystem (Lipgloss for styling, Bubbles for components)
- Perfect for interactive menus and progress displays

**Cons**:
- Learning curve for Elm-style architecture
- Might be complex for simple use cases
- Requires restructuring current menu logic

**Use Case**: Ideal for interactive file monitoring and menu systems

### 3. Survey ⭐ RECOMMENDED
**Repository**: https://github.com/AlecAivazis/survey  
**Stars**: 4k+ | **Maintained**: Active

**Pros**:
- Simple interactive prompts
- Built-in validation
- Multiple input types (select, multiselect, input, confirm)
- Easy to integrate with existing code
- Minimal learning curve

**Cons**:
- Limited to prompt-based interactions
- No real-time display updates
- Less suitable for continuous monitoring

**Use Case**: Perfect for replacing current menu system with minimal changes

### 4. Termui
**Repository**: https://github.com/gizak/termui  
**Stars**: 13k+ | **Maintained**: Limited

**Pros**:
- Dashboard-style layouts
- Charts and graphs
- Real-time updates

**Cons**:
- Maintenance concerns
- Overkill for file processing tool
- Complex for simple menus

**Use Case**: Not suitable for BlendPDFGo's needs

### 5. Progressbar
**Repository**: https://github.com/schollz/progressbar  
**Stars**: 4k+ | **Maintained**: Active

**Pros**:
- Simple progress bars
- Customizable styling
- Easy integration
- Minimal dependencies

**Cons**:
- Only progress bars, no other UI components
- Limited interaction capabilities

**Use Case**: Good for adding progress indicators to existing operations

### 6. Color + Fatih/Color
**Repository**: https://github.com/fatih/color  
**Stars**: 7k+ | **Maintained**: Active

**Pros**:
- Enhanced color support
- Cross-platform compatibility
- Simple API
- Better than current ANSI codes

**Cons**:
- Only colors, no interactive components
- Incremental improvement over current approach

**Use Case**: Easy upgrade for current color system

## Recommendations by Use Case

### Option A: Minimal Enhancement (Low Risk)
**Libraries**: Survey + Progressbar + Fatih/Color
- Replace current menu with Survey prompts
- Add progress bars for long operations
- Upgrade color system
- **Effort**: 1-2 days
- **Risk**: Low

### Option B: Modern Interactive (Medium Risk)
**Libraries**: Bubble Tea + Lipgloss
- Full TUI with real-time file monitoring
- Interactive file selection
- Live progress updates
- **Effort**: 1-2 weeks
- **Risk**: Medium

### Option C: Command Enhancement (Low Risk)
**Libraries**: Cobra + Viper + Progressbar
- Enhanced CLI with subcommands
- Configuration file support
- Better help system
- **Effort**: 3-5 days
- **Risk**: Low

## Specific UI Patterns Identified

### Interactive File Selection
```go
// Using Survey
prompt := &survey.Select{
    Message: "Choose files to merge:",
    Options: pdfFiles,
}
survey.AskOne(prompt, &selected)
```

### Real-time File Monitoring
```go
// Using Bubble Tea
type model struct {
    files []string
    cursor int
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
    // Handle file system changes
    // Update display in real-time
}
```

### Progress Indicators
```go
// Using progressbar
bar := progressbar.Default(100)
for i := 0; i < 100; i++ {
    bar.Add(1)
    // Process files
}
```

### Enhanced Colors
```go
// Using fatih/color
color.Green("Success: Files merged")
color.Red("Error: Invalid PDF")
color.Yellow("Warning: Large file detected")
```

## Implementation Priority

### Phase 1: Quick Wins (Recommended)
1. **Survey** for interactive menus
2. **Progressbar** for long operations
3. **Fatih/Color** for better colors

### Phase 2: Advanced Features (Optional)
1. **Bubble Tea** for full TUI experience
2. Real-time file monitoring
3. Interactive file selection

### Phase 3: CLI Enhancement (Future)
1. **Cobra** for subcommands
2. **Viper** for configuration
3. Shell completion

## Screen Takeover Analysis

### Full Terminal Control Requirements
Based on discussion, the goal is to implement **screen takeover** with terminal segmentation to replace the current scrolling input interface.

### Bubble Tea - Full Screen Capabilities
- **Complete terminal takeover**: Takes control of entire terminal window
- **Layout segmentation**: Draw borders and divide screen into sections
- **Real-time updates**: File counts, status, progress without scrolling
- **Interactive navigation**: Arrow keys, visual selection
- **Dynamic layouts**: Adapt to terminal size changes

### Example Full-Screen Layout:
```
┌─────────────────────────────────────────────────┐
│ BlendPDFGo v1.0.0                              │
├─────────────────────────────────────────────────┤
│ Files: Main(2) Archive(0) Output(0) Error(0)   │
├─────────────────────────────────────────────────┤
│ Available Files:                                │
│ • document1.pdf (2.3M)                         │
│ • document2.pdf (1.8M)                         │
├─────────────────────────────────────────────────┤
│ [S] Single File  [M] Merge  [Q] Quit           │
├─────────────────────────────────────────────────┤
│ Status: Ready for operation                     │
└─────────────────────────────────────────────────┘
```

## Windows Compatibility Concerns

### PowerShell 5 / Legacy Windows Issues
- **PowerShell 5**: Limited ANSI support, may show escape codes as text
- **CMD**: Very limited terminal capabilities, poor TUI rendering
- **Windows Console Host (legacy)**: No proper full-screen TUI support
- **Result**: Broken layouts, garbled output for Bubble Tea

### Cross-Platform Support Summary
- **Windows 10+ with Windows Terminal**: ✅ Full Bubble Tea support
- **Windows 10+ with PowerShell 7+**: ✅ Full support
- **Windows with PowerShell 5/CMD**: ❌ Bubble Tea unusable
- **Linux (all terminals)**: ✅ Full support
- **macOS (all terminals)**: ✅ Full support

## Updated Recommendations

### Option A: Screen Takeover Priority (Recommended for Goal)
**Primary**: Bubble Tea + Lipgloss for full terminal control
- Implement terminal capability detection
- Fall back to basic Survey interface on legacy Windows
- Provides desired screen takeover experience on modern terminals
- **Effort**: 1-2 weeks
- **Risk**: Medium (compatibility handling required)

### Option B: Universal Compatibility
**Primary**: Survey + Progressbar + Fatih/Color
- Works on all platforms including PowerShell 5
- No screen takeover but improved UX
- **Effort**: 1-2 days
- **Risk**: Low

### Option C: Hybrid Approach
**Implementation**: Detect terminal capabilities at runtime
- Modern terminals: Use Bubble Tea full-screen interface
- Legacy terminals: Fall back to Survey-based prompts
- Best of both worlds but more complex
- **Effort**: 2-3 weeks
- **Risk**: Medium-High

## Conclusion

**For Screen Takeover Goal**: Implement **Bubble Tea + Lipgloss** with terminal capability detection and graceful fallback to Survey for legacy Windows systems.

This achieves the desired full-screen interface while maintaining compatibility across all target platforms.

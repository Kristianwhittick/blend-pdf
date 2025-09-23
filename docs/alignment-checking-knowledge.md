# Terminal UI Alignment Checking Knowledge

## Problem: Visual Alignment Issues in Terminal Output

When working with formatted terminal output, especially bordered layouts, alignment issues can be difficult to detect and debug.

## Common Alignment Problems

### 1. Width Calculation Errors
- **Issue**: Miscounting total content width vs border width
- **Symptom**: Text extends beyond borders or appears inconsistently positioned
- **Example**: Using 77-character content in 80-character border without accounting for border characters

### 2. Format String Spacing
- **Issue**: Not accounting for implicit spacing in printf format strings
- **Symptom**: Numbers appear at different distances from border edges
- **Example**: `%-62s %6d` creates implicit spacing that must be counted in total width

### 3. Left vs Right Alignment Confusion
- **Issue**: Using wrong alignment type for consistent positioning
- **Symptom**: Variable-length content causes inconsistent number positioning
- **Example**: Left-aligned paths with different lengths push numbers to different positions

## Debugging Techniques

### 1. Character Counting Method
```bash
# Count visual width of border
echo "┌─────────────────────────────────────────────────────────────────────────────┐" | wc -m
# Result: 80 characters

# Count content width needed
# Border (80) - Border chars (2) = Content (78)
# But title uses 77, so match that for consistency
```

### 2. Manual Position Verification
- Copy actual terminal output to text editor
- Count character positions manually
- Compare alignment across different examples
- Look for consistent column positioning

### 3. Test with Extreme Cases
- Test with very short paths
- Test with very long paths  
- Test with large numbers (3+ digits)
- Test with single digits

## Solution Pattern for BlendPDF Header

### Problem Analysis
Original issue: Numbers appeared at inconsistent positions relative to right border when file counts varied.

### Root Cause
Width calculation error: Content was 3 characters too wide for the border, causing overflow and misalignment.

### Correct Calculation
```
Border width: 80 characters
Content width: 77 characters (to match title line)
Format breakdown:
- Label: 9 characters ("Archive: ")
- Path: 59 characters (left-aligned with %-59s)
- Space: 3 characters (implicit in format)
- Number: 6 characters (right-aligned with %6d)
Total: 9 + 59 + 3 + 6 = 77 characters ✓
```

### Final Format
```go
fmt.Printf("│ Archive: %-59s %6d │\n", archiveDir, count)
```

## Key Lessons Learned

### 1. Why Visual Checking Failed Initially
- **Terminal escape codes**: ANSI color codes made raw output harder to parse visually
- **Character counting errors**: Miscounted border vs content width repeatedly
- **Assumption errors**: Assumed calculations were correct without manual verification
- **Pattern recognition failure**: Didn't notice consistent 3-character overflow pattern

### 2. Effective Debugging Approach
1. **Measure border width accurately** using `wc -m`
2. **Account for all spacing** in format strings (implicit and explicit)
3. **Test with real examples** rather than theoretical calculations
4. **Verify alignment manually** by counting character positions
5. **Use extreme test cases** to expose edge case failures

### 3. Prevention Strategies
- **Document width calculations** in code comments
- **Test alignment immediately** after format changes
- **Use consistent measurement methods** (character counting tools)
- **Verify with multiple examples** before considering complete

## Tools for Alignment Verification

### Command Line Tools
```bash
# Measure visual width
echo "text" | wc -m

# Count characters including newlines  
echo "text" | wc -c

# Extract specific lines for analysis
program | head -10 | tail -1
```

### Manual Verification
1. Copy terminal output to text editor
2. Use editor's column/position indicators
3. Count characters manually for precision
4. Compare multiple examples side-by-side

## Application to Other UI Elements

This knowledge applies to any formatted terminal output:
- Table layouts
- Progress bars
- Status displays
- Menu systems
- Bordered content

Always verify alignment with actual output rather than trusting calculations alone.

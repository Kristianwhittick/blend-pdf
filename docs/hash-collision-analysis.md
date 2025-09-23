# Hash Collision Analysis for Directory-Specific Lock Files

## Overview
Analysis of hash collision risk for directory-specific lock file naming using truncated MD5 hashes.

## Hash Length Options Considered

### 8 Characters (32 bits) - **SELECTED**
- **Total combinations**: 2^32 = 4,294,967,296 (~4.3 billion)
- **Birthday paradox threshold**: √(2^32) ≈ 65,536 paths
- **50% collision probability**: At ~65,000 different watch directories
- **Practical risk**: Low for typical usage patterns
- **Benefits**: Shorter filenames, faster computation
- **Use case suitability**: Excellent for individual users and small teams

### 12 Characters (48 bits) - Alternative
- **Total combinations**: 2^48 = 281,474,976,710,656 (~281 trillion)
- **Birthday paradox threshold**: √(2^48) ≈ 16,777,216 paths
- **50% collision probability**: At ~16.7 million different directories
- **Practical risk**: Negligible for any realistic usage
- **Trade-offs**: Longer filenames, minimal performance impact

### 16 Characters (64 bits) - Overkill
- **Total combinations**: 2^64 = 18,446,744,073,709,551,616
- **Practical risk**: Virtually zero
- **Trade-offs**: Unnecessarily long filenames

## Decision Rationale

### Why 8 Characters Was Chosen:

1. **Sufficient Security**
   - 65,000 different watch folders before significant collision risk
   - Most users won't approach this threshold
   - Even power users rarely exceed 1,000 different project folders

2. **User Experience**
   - Shorter lock file names are easier to read and manage
   - Less visual clutter in temp directories
   - Easier to identify and debug if needed

3. **Performance**
   - Faster hash computation (minimal impact)
   - Faster string comparisons
   - Smaller memory footprint

4. **Real-World Usage Patterns**
   - Individual developers: 10-100 different project folders
   - Small teams: 100-1,000 different project folders
   - Large organizations: Would use centralized solutions

## Implementation Details

### Hash Generation Process
```go
func generateDirectoryHash(watchDir string) string {
    // 1. Normalize path for consistency
    absPath, _ := filepath.Abs(watchDir)
    cleanPath := filepath.Clean(absPath)
    normalizedPath := strings.ToLower(filepath.ToSlash(cleanPath))
    
    // 2. Generate MD5 hash
    hash := md5.Sum([]byte(normalizedPath))
    
    // 3. Return first 8 characters as hex
    return fmt.Sprintf("%x", hash)[:8]
}
```

### Lock File Naming Convention
- **Format**: `blendpdf-<8-char-hash>.lock`
- **Example**: `blendpdf-a1b2c3d4.lock`
- **Location**: 
  - Unix: `/tmp/blendpdf-a1b2c3d4.lock`
  - Windows: `<watch-dir>/blendpdf-a1b2c3d4.lock`

## Risk Mitigation

### Collision Handling
If a collision occurs (extremely unlikely):
1. **Detection**: Lock file exists but different PID
2. **Behavior**: Same as current - prevent multiple instances
3. **User Action**: Manual lock file removal if needed
4. **Impact**: Minimal - affects only two specific directories with same hash

### Monitoring
- No special monitoring needed due to low probability
- Standard lock file error messages remain unchanged
- Users can identify collisions by examining lock file contents (PID)

## Testing Strategy

### Collision Testing
- Generate hashes for common directory patterns
- Verify no collisions in typical development scenarios
- Test with various path formats and edge cases

### Cross-Platform Testing
- Verify consistent hashes across Windows/Linux/macOS
- Test path normalization with different filesystem types
- Validate lock file creation in different locations

## Future Considerations

### If Collision Rate Becomes Problematic
1. **Increase to 12 characters**: Simple configuration change
2. **Add timestamp suffix**: `blendpdf-<hash>-<timestamp>.lock`
3. **Use different hash algorithm**: SHA256 for better distribution

### Monitoring Recommendations
- Log hash generation in debug mode
- Track collision occurrences if they happen
- Consider telemetry for hash distribution analysis

## Conclusion

The 8-character MD5 hash provides an excellent balance of:
- **Security**: Adequate protection against collisions
- **Usability**: Short, manageable filenames
- **Performance**: Fast computation and comparison
- **Scalability**: Supports realistic usage patterns

This approach enables the desired functionality (multiple instances in different directories) while maintaining simplicity and reliability.

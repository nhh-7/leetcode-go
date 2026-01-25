# Agent Guidelines for LeetCode Go Repository

## Build, Lint, and Test Commands

### Build
```bash
go build ./...
```

### Run specific problem
```bash
cd hotQ100
go run path/to/problem.go
```

### Lint
```bash
golangci-lint run ./...
# or
go vet ./...
```

### Test
```bash
# Run all tests
go test ./...

# Run single test file
go test -v hotQ100/1.hash/

# Run specific test
go test -v hotQ100/1.hash/ -run TestTwoSum
```

## Code Style Guidelines

### File Organization
- Files named with Chinese problem title + problem number (e.g., `1. 两数之和.go`)
- Organized by algorithm category in subdirectories (`1.hash/`, `2.double_ptr/`, `3.sliding_window/`, etc.)
- Each problem file contains only solution function(s) and type definitions if needed
- Multiple solutions in same file use numeric suffixes (e.g., `addTwoNumbers0`, `addTwoNumbers1`)

### Package Declaration
- Use category-based package names: `hot`, `backtrack`, `dp`, `heap`, `stack`, `binarytree`, etc.
- Match package name to directory category

### Imports
- Use only required standard library imports
- Import third-party packages only when necessary (e.g., `container/heap`)
- Group imports logically; keep to minimum
- Common imports: `"slices"`, `"maps"`, `"container/heap"`, `"math/rand"`

### Formatting
- Follow standard `gofmt` conventions
- Use tabs for indentation (Go standard)
- Keep lines under 120 characters when possible
- Minimal whitespace between related functions

### Naming Conventions
- Functions: `camelCase` (e.g., `twoSum`, `lengthOfLongestSubstring`)
- Variables: `camelCase`, single letters for indices/short scope (i, j, k, n, m)
- Constants: `UPPER_SNAKE_CASE` for named constants
- Types: `PascalCase` (e.g., `ListNode`, `TreeNode`, `ColorAndNode`)
- Short, descriptive names preferred over abbreviations

### Types
- Use built-in types: `int`, `string`, `bool`, `[]int`, `[][]int`
- Struct fields: `PascalCase` (e.g., `Val`, `Left`, `Right`)
- Use empty slices: `[]int{}` or `[][]int{}` instead of nil returns
- Use `make()` for pre-allocated slices/maps when size is known
- Common types: `ListNode`, `TreeNode`, `minHeap`, `ColorAndNode`

### Function Signatures
- Return slice results directly: `func solve(nums []int) []int`
- Use named returns only for clarity: `func dfs(node *TreeNode) (ans []int)`
- Accept pointers for large structs: `*ListNode`, `*TreeNode`
- Use slices, not arrays, for dynamic collections

### Control Flow
- Prefer `for-range` loops for iteration
- Use `for` with explicit indices when needed
- Early returns for edge cases (nil checks, length 0)
- `if-else` chains preferred over nested conditionals
- Switch for simple multi-way branches

### Error Handling
- Since this is LeetCode solutions, errors are typically returned as sentinel values
- Return `[]int{0, 0}` or `[][]int{}` for no result cases
- Return nil for pointer types when appropriate
- Use `ok` idiom for map access: `val, ok := m[key]`

### Comments
- Comments only when necessary; code should be self-documenting
- Use `//` for single-line comments
- Comment complex algorithms (e.g., partition explanation)
- Include LeetCode problem structure comments for tree/list types
- Minimal inline comments for obvious operations

### Common Patterns
- **Two pointers**: `left, right := 0, n-1` while `left < right`
- **Sliding window**: `cnt` array or map, expand right, shrink left
- **DFS/Backtracking**: closure function `var dfs func()` with path tracking
- **DP**: `dp := make([]int, n+1)` with base cases
- **BFS**: use slice as queue: `q := []*TreeNode{}`
- **Memoization**: map or array to cache results
- **Binary search**: `left, right := 0, n-1` while `left <= right`

### Performance Considerations
- Use arrays `[128]int{}` for ASCII char counting instead of map
- Pre-allocate slices when size is known: `make([]int, n)`
- Use slices for stack: append/pop from end
- Use `max()` builtin function (Go 1.21+)
- Prefer integer operations over floating point
- Use `continue` and `break` for loop control flow

### Multiple Solutions
- When multiple approaches exist, number them: `solution1`, `solution2`
- Keep simpler solution as base name (no suffix)
- Document trade-offs in comments if needed

### Directory Structure
- `1.hash/` - Hash table solutions
- `2.double_ptr/` - Two pointers technique
- `3.sliding_window/` - Sliding window algorithms
- `4.substring/` - Substring manipulation
- `5.normal_array/` - Array problems
- `6.matrix/` - Matrix operations
- `7.linkedlist/` - Linked list problems
- `8.binary_tree/` - Binary tree algorithms
- `9.graph/` - Graph traversal and algorithms
- `10.backtrack/` - Backtracking and recursion
- `11.binary_search/` - Binary search problems
- `12.stack/` - Stack-based solutions
- `13.heap/` - Heap/priority queue
- `14.greedy/` - Greedy algorithms
- `15.dp/` - Dynamic programming
- `16.multi_dimension_dp/` - Multi-dimensional DP

### Type Definitions
- `ListNode`: singly linked list node with `Val` and `Next` fields
- `TreeNode`: binary tree node with `Val`, `Left`, and `Right` fields
- Custom types defined inline when problem-specific structure needed

## Testing Approach
- LeetCode solutions typically don't include test files in this repo
- Test manually by creating test cases or using LeetCode platform
- When adding tests, follow standard Go testing conventions
- Test function names: `Test{FunctionName}`

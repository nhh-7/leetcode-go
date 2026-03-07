package stack

func solve(height []int) int {
	stack := []int{}
	stack = append(stack, -1)
	height = append(height, 0)

	ans := 0
	for right, h := range height {
		for len(stack) > 1 && height[stack[len(stack)-1]] > h {
			i := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			left := stack[len(stack)-1]
			ans = max(height[i]*(right-left-1), ans)
		}
		stack = append(stack, right)
	}
	return ans
}

func maximalRectangle(matrix [][]byte) int {
	ans := 0
	height := make([]int, len(matrix[0]))
	for _, row := range matrix {
		for j, ch := range row {
			if ch == '1' {
				height[j]++
			} else {
				height[j] = 0
			}
		}
		s := solve(height)
		ans = max(ans, s)
	}
	return ans
}

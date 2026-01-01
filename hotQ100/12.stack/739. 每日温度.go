package stack

// 寻找下一次出现则使用单调栈
// 使用单调栈，遍历温度，找到高于栈顶温度的，则记录结果
// 单调栈中递减
func dailyTemperatures(temperatures []int) []int {
	stack := []int{}
	ans := make([]int, len(temperatures))

	for i, t := range temperatures {
		for len(stack) > 0 && t > temperatures[stack[len(stack)-1]] {
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			ans[top] = i - top
		}
		stack = append(stack, i)
	}
	return ans
}

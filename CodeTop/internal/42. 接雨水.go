package internal

func trap(height []int) int {
	n := len(height)
	leftMax, rightMax := height[0], height[n-1]

	i, j := 0, n-1
	ans := 0
	for i <= j {
		leftMax = max(leftMax, height[i])
		rightMax = max(rightMax, height[j])

		if leftMax < rightMax {
			ans += leftMax - height[i]
			i++
		} else {
			ans += rightMax - height[j]
			j--
		}
	}
	return ans
}

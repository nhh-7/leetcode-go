package hot

func trap(height []int) int {
	ans := 0
	n := len(height)
	left, right := 0, n-1
	leftMax, rightMax := height[left], height[right]

	for left < right {
		leftMax = max(leftMax, height[left])
		rightMax = max(rightMax, height[right])

		if leftMax < rightMax {
			ans += leftMax - height[left]
			left++
		} else {
			ans += rightMax - height[right]
			right--
		}
	}
	return ans
}

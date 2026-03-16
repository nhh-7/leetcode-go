package times1

func trap(height []int) int {
	left, right := 0, len(height)-1
	leftMax, rightMax := height[0], height[right]

	ans := 0
	for left <= right {
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

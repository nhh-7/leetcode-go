package hot

func maxArea(height []int) int {
	ans := 0

	left, right := 0, len(height)-1

	for left <= right {
		ans = max(ans, (right-left)*min(height[left], height[right]))

		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}

	return ans
}

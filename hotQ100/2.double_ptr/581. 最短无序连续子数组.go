package hot

func findUnsortedSubarray(nums []int) int {
	left, right := -1, -1
	n := len(nums)
	left_max, right_min := nums[0], nums[n-1]

	for i := 0; i < n; i++ {
		if nums[i] < left_max {
			right = i
		} else {
			left_max = nums[i]
		}
	}

	for i := n - 1; i >= 0; i-- {
		if nums[i] > right_min {
			left = i
		} else {
			right_min = nums[i]
		}
	}
	if right == -1 {
		return 0
	}
	return right - left + 1
}

package internal

import "slices"

func nextPermutation(nums []int) {
	n := len(nums)
	i := n - 2
	for i >= 0 && nums[i] >= nums[i+1] {
		i--
	}

	if i >= 0 {
		j := n - 1
		for j > i && nums[j] <= nums[i] {
			j--
		}
		nums[j], nums[i] = nums[i], nums[j]
	}
	slices.Reverse(nums[i+1:])
}

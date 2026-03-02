package skills

import "slices"

func nextPermutation1(nums []int) {
	n := len(nums)
	i := n - 2
	for i >= 0 && nums[i] >= nums[i+1] {
		i--
	}
	if i >= 0 {
		j := n - 1
		for j >= 0 && nums[j] <= nums[i] {
			j--
		}
		nums[i], nums[j] = nums[j], nums[i]
	}
	slices.Reverse(nums[i+1:])
}

// 1 3 5 4 2
func nextPermutation(nums []int) {
	n := len(nums)
	i := n - 2
	for i >= 0 && nums[i+1] <= nums[i] {
		i--
	} // i之后一定是降序的
	if i >= 0 {
		j := n - 1
		for j >= 0 && nums[j] <= nums[i] {
			j--
		} // 寻找到i右侧比nums[i]大的最小值
		nums[i], nums[j] = nums[j], nums[i]
	}
	slices.Reverse(nums[i+1:])
}

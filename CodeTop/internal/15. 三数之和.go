package internal

import "slices"

func threeSum(nums []int) [][]int {
	slices.Sort(nums)
	n := len(nums)
	ans := [][]int{}

	for i := 0; i < n-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		j, k := i+1, n-1
		for j < k {
			res := nums[i] + nums[j] + nums[k]
			if res < 0 {
				j++
			} else if res > 0 {
				k--
			} else {
				ans = append(ans, []int{nums[i], nums[j], nums[k]})
				for j < k && nums[j] == nums[j+1] {
					j++
				}
				j++
				for j < k && nums[k] == nums[k-1] {
					k--
				}
				k--
			}
		}
	}
	return ans
}

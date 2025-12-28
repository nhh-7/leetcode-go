package hot

import "slices"

func threeSum(nums []int) [][]int {
	slices.Sort(nums)

	n := len(nums)
	ans := [][]int{}

	for i, x := range nums[:n-2] {
		if i > 0 && x == nums[i-1] {
			continue
		}

		if x+nums[i+1]+nums[i+2] > 0 {
			continue
		}

		if x+nums[n-2]+nums[n-1] < 0 {
			continue
		}

		j, k := i+1, n-1

		for j < k {
			if x+nums[j]+nums[k] > 0 {
				k--
			} else if x+nums[j]+nums[k] < 0 {
				j++
			} else {
				ans = append(ans, []int{x, nums[j], nums[k]})

				for j < k && nums[j] == nums[j+1] {
					j++
				}
				for j < k && nums[k] == nums[k-1] {
					k--
				}
				j++
				k--
			}
		}
	}
	return ans
}

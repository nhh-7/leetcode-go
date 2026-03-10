package internal

import "math"

func maxSubArray(nums []int) int {
	ans := math.MinInt
	sum := 0
	for _, n := range nums {
		sum += n
		ans = max(sum, ans)
		if sum < 0 {
			sum = 0
		}
	}
	return ans
}

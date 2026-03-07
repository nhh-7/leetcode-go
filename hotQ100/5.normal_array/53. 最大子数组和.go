package hot

import (
	"math"
)

func maxSubArray1(nums []int) int {
	ans := math.MinInt

	sum := 0
	for _, v := range nums {
		sum += v
		ans = max(ans, sum)
		if sum < 0 {
			sum = 0
		}
	}

	return ans
}

func maxSubArray(nums []int) int {
	ans := math.MinInt
	sum := 0
	for _, v := range nums {
		sum += v
		ans = max(sum, ans)
		if sum < 0 {
			sum = 0
		}
	}
	return ans
}

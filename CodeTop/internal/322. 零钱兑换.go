package internal

import "math"

func coinChange(coins []int, amount int) int {
	dp := make([]int, amount+1)
	for i := range dp {
		dp[i] = math.MaxInt
	}
	dp[0] = 0

	for _, c := range coins {
		for j := c; j <= amount; j++ {
			if dp[j-c] != math.MaxInt {
				dp[j] = min(dp[j], dp[j-c]+1)
			}
		}
	}
	if dp[amount] == math.MaxInt {
		return -1
	}
	return dp[amount]
}

package dp

func findTargetSumWays(nums []int, target int) int {
	sum := 0
	for _, x := range nums {
		sum += x
	}
	if (target+sum)%2 == 1 || (sum+target) < 0 {
		return 0
	}
	cap := (target + sum) / 2
	dp := make([]int, cap+1)

	dp[0] = 1 // 凑出和为 0 的方案默认有 1 种（什么都不选）
	for _, num := range nums {
		for j := cap; j >= num; j-- {
			dp[j] = dp[j] + dp[j-num]
		}
	}
	return dp[cap]
}

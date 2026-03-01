package dp

func lengthOfLIS1(nums []int) int {
	n := len(nums)
	dp := make([]int, n)
	for i := range dp {
		dp[i] = 1
	}
	ans := 1
	for i := 1; i < n; i++ {
		for j := i - 1; j >= 0; j-- {
			if nums[j] < nums[i] {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
		ans = max(ans, dp[i])
	}
	return ans
}

func lengthOfLIS(nums []int) int {
	n := len(nums)
	dp := make([]int, n)
	for i := range dp {
		dp[i] = 1
	}
	ans := 1
	for i := 1; i < n; i++ {
		for j := i - 1; j >= 0; j-- {
			if nums[i] > nums[j] {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
		ans = max(ans, dp[i])
	}
	return ans
}

package dp

// 01背包
func canPartition1(nums []int) bool {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	if sum%2 != 0 {
		return false
	}
	target := sum / 2

	dp := make([]int, target+1)

	dp[0] = 0
	for _, num := range nums {
		for j := target; j >= num; j-- {
			dp[j] = max(dp[j], dp[j-num]+num)
		}
	}
	return dp[target] == target
}

func canPartition(nums []int) bool {
	sum := 0
	for _, x := range nums {
		sum += x
	}
	if sum%2 == 1 {
		return false
	}
	cap := sum / 2
	dp := make([]int, cap+1)
	dp[0] = 0

	for _, num := range nums {
		for j := cap; j >= num; j-- {
			dp[j] = max(dp[j], dp[j-num]+num)
		}
	}
	return dp[cap] == cap
}

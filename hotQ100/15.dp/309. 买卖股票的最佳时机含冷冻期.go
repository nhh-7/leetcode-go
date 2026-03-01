package dp

/*
dp[i][0]：持有股票状态
手上有票。可能是之前买的没卖，或者是今天刚买。

dp[i][1]：不持有股票，且没干啥（保持不持有）
指的是那种“我已经空仓很久了，今天也没买，且我不在冷冻期”的状态。

dp[i][2]：今天卖出股票
这是一个瞬时动作状态。特指今天卖了，所以明天必然进入冷冻期。

dp[i][3]：冷冻期
特指今天不能买入的状态。通常是因为昨天刚刚卖出了（即从 dp[i-1][2] 转移而来）
*/
func maxProfit(prices []int) int {
	n := len(prices)
	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, 4)
	}

	dp[0][0] = -prices[0] // 持有股票
	dp[0][1] = 0          // 不持有股票
	dp[0][2] = 0          // 卖出股票
	dp[0][3] = 0          // 冷冻期

	for i := 1; i < n; i++ {
		dp[i][0] = max(dp[i-1][0], dp[i-1][3]-prices[i], dp[i-1][1]-prices[i])
		dp[i][1] = max(dp[i-1][1], dp[i-1][3])
		dp[i][2] = dp[i-1][0] + prices[i]
		dp[i][3] = dp[i-1][2]
	}
	return max(dp[n-1][1], dp[n-1][2], dp[n-1][3])
}

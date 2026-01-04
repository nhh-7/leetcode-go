package greedy

func maxProfit(prices []int) int {
	in := prices[0] // 获取买入最少的价格
	ans := 0
	for _, p := range prices {
		in = min(in, p)
		ans = max(ans, p-in)
	}
	return ans
}

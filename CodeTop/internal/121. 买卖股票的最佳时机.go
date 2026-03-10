package internal

func maxProfit(prices []int) int {
	ans := 0
	minPrice := prices[0]
	for _, p := range prices {
		minPrice = min(minPrice, p)
		ans = max(ans, p-minPrice)
	}
	return ans
}

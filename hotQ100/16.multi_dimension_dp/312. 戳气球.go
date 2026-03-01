package multidimensiondp

func maxCoins(nums []int) int {
	n := len(nums)
	// 1. 预处理：加上左右边界 1
	temp := make([]int, n+2)
	temp[0], temp[n+1] = 1, 1
	for i := 0; i < n; i++ {
		temp[i+1] = nums[i]
	}

	// 2. 初始化 DP 数组
	// dp[i][j] 表示开区间 (i, j) 的最大收益
	dp := make([][]int, n+2)
	for i := range dp {
		dp[i] = make([]int, n+2)
	}

	// 3. 开始套路：区间长度 len 从 1 开始到 n
	for length := 1; length <= n; length++ {
		// i 是左边界
		for i := 0; i <= n-length; i++ {
			j := i + length + 1 // j 是右边界
			// k 是区间 (i, j) 中最后被戳破的气球
			for k := i + 1; k < j; k++ {
				/*
					为什么 $k$ 左右的气球是 $i$ 和 $j$？**
					因为我们定义 $k$ 是该区间**最后一个**被戳破的，意味着 $(i, k)$ 和 $(k, j)$ 里的气球都已经消失了，所以 $k$ 直接挨着 $i$ 和 $j$。
				*/
				score := temp[i] * temp[k] * temp[j]
				dp[i][j] = max(dp[i][j], dp[i][k]+dp[k][j]+score)
			}
		}
	}

	return dp[0][n+1]
}

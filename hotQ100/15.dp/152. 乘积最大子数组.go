package dp

func maxProduct(nums []int) int {
	n := len(nums)
	dp_max, dp_min := make([]int, n), make([]int, n)
	dp_max[0] = nums[0]
	dp_min[0] = nums[0]

	ans := nums[0]
	for i := 1; i < n; i++ {
		if nums[i] < 0 {
			dp_max[i-1], dp_min[i-1] = dp_min[i-1], dp_max[i-1]
		}
		dp_max[i] = max(dp_max[i-1]*nums[i], nums[i])
		dp_min[i] = min(dp_min[i-1]*nums[i], nums[i])

		ans = max(ans, dp_max[i])
	}
	return ans
}

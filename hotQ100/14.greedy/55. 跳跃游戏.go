package greedy

func canJump(nums []int) bool {
	max_cover := 0
	n := len(nums)
	for i := 0; i < n; i++ {
		if i > max_cover {
			return false
		}
		max_cover = max(max_cover, i+nums[i])
		if max_cover >= n-1 {
			return true
		}
	}
	return false
}

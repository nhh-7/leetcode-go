package greedy

func canJump(nums []int) bool {
	max_cover := 0
	for i := 0; i <= max_cover; i++ {
		max_cover = max(max_cover, i+nums[i])
		if max_cover >= len(nums)-1 {
			return true
		}
	}
	return false
}

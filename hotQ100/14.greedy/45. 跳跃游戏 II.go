package greedy

// 逼不得已才跳
// 如果要跳，那么要跳到能去到最远的
func jump(nums []int) int {
	cover := 0
	ans := 0
	nxt_cover := 0
	n := len(nums)
	if n == 1 {
		return ans
	}

	for i, num := range nums {
		nxt_cover = max(nxt_cover, num+i)
		if cover == i {
			ans++
			cover = nxt_cover
			nxt_cover = 0
			if cover >= n-1 {
				return ans
			}
		}
	}
	return ans
}

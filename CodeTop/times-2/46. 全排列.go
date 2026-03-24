package times2

func permute(nums []int) [][]int {
	ans := [][]int{}
	path := []int{}

	var backtrack func(map[int]bool)

	backtrack = func(m map[int]bool) {
		if len(path) == len(nums) {
			tmp := make([]int, len(nums))
			copy(tmp, path)
			ans = append(ans, tmp)
			return
		}

		for i := 0; i < len(nums); i++ {
			if m[nums[i]] {
				continue
			}
			m[nums[i]] = true
			path = append(path, nums[i])
			backtrack(m)
			path = path[:len(path)-1]
			m[nums[i]] = false
		}
	}
	backtrack(make(map[int]bool))
	return ans
}

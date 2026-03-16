package internal

func combinationSum(candidates []int, target int) [][]int {
	ans := [][]int{}
	path := []int{}

	var backtrack func(start, target int)
	backtrack = func(start, target int) {
		if target == 0 {
			tmp := make([]int, len(path))
			copy(tmp, path)
			ans = append(ans, tmp)
			return
		}
		if target < 0 {
			return
		}

		for i := start; i < len(candidates); i++ {
			path = append(path, candidates[i])
			backtrack(i, target-candidates[i])
			path = path[:len(path)-1]
		}
	}
	backtrack(0, target)
	return ans
}

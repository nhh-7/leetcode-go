package backtrack

func combinationSum(candidates []int, target int) [][]int {
	ans := [][]int{}
	path := []int{}
	n := len(candidates)

	var dfs func(int, int)
	dfs = func(startIdx, targetNum int) {
		if targetNum == 0 {
			tmp := make([]int, len(path))
			copy(tmp, path)
			ans = append(ans, tmp)
			return
		} else if targetNum < 0 {
			return
		}

		for i := startIdx; i < n; i++ {
			path = append(path, candidates[i])
			dfs(i, targetNum-candidates[i])
			path = path[:len(path)-1]
		}
	}
	dfs(0, target)
	return ans
}

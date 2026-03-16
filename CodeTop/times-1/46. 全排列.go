package times1

func permute(nums []int) [][]int {
	vis := make(map[int]bool)

	ans := [][]int{}
	path := []int{}
	var dfs func(map[int]bool)
	dfs = func(m map[int]bool) {
		if len(path) == len(nums) {
			tmp := make([]int, len(path))
			copy(tmp, path)
			ans = append(ans, tmp)
			return
		}

		for _, v := range nums {
			if m[v] {
				continue
			}
			path = append(path, v)
			m[v] = true
			dfs(m)
			m[v] = false
			path = path[:len(path)-1]
		}
	}
	dfs(vis)
	return ans
}

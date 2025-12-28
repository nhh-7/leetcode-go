package backtrack

func subsets(nums []int) [][]int {
	ans := [][]int{}
	n := len(nums)
	path := []int{}
	var dfs func(int)
	dfs = func(idx int) {
		tmp := make([]int, len(path))
		copy(tmp, path)
		ans = append(ans, tmp)

		for j := idx; j < n; j++ {
			path = append(path, nums[j])
			dfs(j + 1)
			path = path[:len(path)-1]
		}
	}
	dfs(0)
	return ans
}

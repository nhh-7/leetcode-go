package backtrack

func subsets1(nums []int) [][]int {
	ans := [][]int{}
	n := len(nums)
	path := []int{}
	var dfs func(int)
	dfs = func(start int) {
		tmp := make([]int, len(path))
		copy(tmp, path)
		ans = append(ans, tmp)

		// 从 start 开始遍历，避免产生 [1,2] 和 [2,1] 这种重复组合
		for j := start; j < n; j++ {
			path = append(path, nums[j])
			dfs(j + 1)
			path = path[:len(path)-1]
		}
	}
	dfs(0)
	return ans
}

func subsets(nums []int) [][]int {
	ans := [][]int{}
	path := []int{}

	var backtrack func(int)
	backtrack = func(start int) {
		tmp := make([]int, len(path))
		copy(tmp, path)
		ans = append(ans, tmp)

		for i := start; i < len(nums); i++ {
			path = append(path, nums[i])
			backtrack(i + 1)
			path = path[:len(path)-1]
		}
	}
	backtrack(0)
	return ans
}

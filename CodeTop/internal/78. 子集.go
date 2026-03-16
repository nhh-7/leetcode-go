package internal

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

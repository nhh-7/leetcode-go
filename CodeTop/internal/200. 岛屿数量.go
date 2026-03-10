package internal

func numIslands(grid [][]byte) int {
	DIRS := [4][2]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	m, n := len(grid), len(grid[0])

	var dfs func(i, j int)
	dfs = func(i, j int) {
		if grid[i][j] == '0' {
			return
		}
		grid[i][j] = '0'

		for _, d := range DIRS {
			ni, nj := i+d[0], j+d[1]
			if ni < 0 || ni >= m || nj < 0 || nj >= n {
				continue
			}
			dfs(ni, nj)
		}
	}
	ans := 0
	for i, row := range grid {
		for j, ch := range row {
			if ch == '1' {
				ans++
				dfs(i, j)
			}
		}
	}
	return ans
}

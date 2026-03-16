package internal

func maxAreaOfIsland(grid [][]int) int {
	DIRS := [4][2]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	m := len(grid)
	n := len(grid[0])

	area := 0
	var dfs func(i, j int)
	dfs = func(i, j int) {
		if grid[i][j] == 0 {
			return
		}
		grid[i][j] = 0
		area++
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
		for j, v := range row {
			if v == 1 {
				area = 0
				dfs(i, j)
				ans = max(ans, area)
			}
		}
	}
	return ans
}

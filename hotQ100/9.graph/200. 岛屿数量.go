package graph

func numIslands(grid [][]byte) int {
	dirs := [][]int{
		{0, 1},
		{1, 0},
		{0, -1},
		{-1, 0},
	}
	m, n := len(grid), len(grid[0])
	var dfs func(int, int)
	dfs = func(x, y int) {
		if grid[x][y] == '0' {
			return
		}
		grid[x][y] = '0'
		for _, d := range dirs {
			nx, ny := x+d[0], y+d[1]
			if nx >= m || nx < 0 || ny >= n || ny < 0 {
				continue
			}
			dfs(nx, ny)
		}
	}
	ans := 0
	for i, row := range grid {
		for j, c := range row {
			if c == '1' {
				dfs(i, j)
				ans++
			}
		}
	}
	return ans
}

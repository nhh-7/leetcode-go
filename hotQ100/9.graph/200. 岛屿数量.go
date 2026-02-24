package graph

func numIslands1(grid [][]byte) int {
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

func numIslands(grid [][]byte) int {
	DIRS := [4][2]int{
		{0, 1}, {1, 0}, {0, -1}, {-1, 0},
	}
	var dfs func(i, j int)
	m, n := len(grid), len(grid[0])
	dfs = func(i, j int) {
		if grid[i][j] != '1' {
			return
		}
		grid[i][j] = '0'
		for _, row := range DIRS {
			nx, ny := i+row[0], j+row[1]
			if nx < 0 || nx >= m || ny < 0 || ny >= n {
				continue
			}
			dfs(nx, ny)
		}
	}
	ans := 0
	for i, row := range grid {
		for j, ch := range row {
			if ch == '1' {
				dfs(i, j)
				ans++
			}
		}
	}
	return ans
}

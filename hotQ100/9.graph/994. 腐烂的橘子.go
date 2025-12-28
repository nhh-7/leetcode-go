package graph

type node struct {
	x, y int
}

func orangesRotting(grid [][]int) int {
	fresh := 0
	q := []node{}
	for i, row := range grid {
		for j, c := range row {
			switch c {
			case 1:
				fresh++
			case 2:
				q = append(q, node{i, j})
			}
		}
	}
	ans := 0

	// 开始bfs
	dirs := [][]int{
		{0, 1},
		{1, 0},
		{0, -1},
		{-1, 0},
	}
	m, n := len(grid), len(grid[0])
	for fresh > 0 && len(q) > 0 {
		k := len(q)
		ans++

		for i := 0; i < k; i++ {
			rotten := q[0]
			q = q[1:]
			for _, d := range dirs {
				nx, ny := rotten.x+d[0], rotten.y+d[1]

				if nx >= m || nx < 0 || ny >= n || ny < 0 {
					continue
				}

				if grid[nx][ny] == 1 {
					q = append(q, node{nx, ny})
					grid[nx][ny] = 2
					fresh--
				}

			}
		}
	}
	if fresh > 0 {
		return -1
	}
	return ans
}

package backtrack

import "slices"

func exist(board [][]byte, word string) bool {
	DIRs := [][]int{
		{0, 1},
		{1, 0},
		{0, -1},
		{-1, 0},
	}

	w := []byte(word)
	m, n := len(board), len(board[0])

	// 如果board中某字符的出现次数比word中的少，那么false
	cnt := make(map[byte]int)
	for _, row := range board {
		for _, c := range row {
			cnt[c]++
		}
	}
	wordCnt := make(map[byte]int)
	for _, c := range w {
		wordCnt[c]++
		if wordCnt[c] > cnt[c] {
			return false
		}
	}

	// 如果 word=abcd 但 board 中的 a 很多，d 很少（比如只有一个），那么从 d 开始搜索，能更快地找到答案。
	// 如果word首字符在board中出现次数比尾字符多，那么从word尾部开始寻找
	if cnt[w[0]] > cnt[w[len(w)-1]] {
		slices.Reverse(w)
	}

	var dfs func(int, int, int) bool
	dfs = func(i, j, k int) bool {
		if board[i][j] != w[k] {
			return false
		}
		if k == len(w)-1 {
			return true
		}

		tmp := board[i][j]
		board[i][j] = ' '

		for _, d := range DIRs {
			nx, ny := i+d[0], j+d[1]
			if nx >= m || nx < 0 || ny >= n || ny < 0 {
				continue
			}
			if dfs(nx, ny, k+1) {
				return true
			}
		}

		board[i][j] = tmp
		return false
	}

	for i, row := range board {
		for j, c := range row {
			if c == w[0] {
				if dfs(i, j, 0) {
					return true
				}
			}
		}
	}
	return false
}

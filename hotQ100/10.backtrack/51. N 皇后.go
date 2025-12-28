package backtrack

func isValid(board [][]byte, row, col int) bool {
	n := len(board)
	for i, j := row-1, col-1; i >= 0 && j >= 0; i, j = i-1, j-1 {
		if board[i][j] == 'Q' {
			return false
		}
	}
	for i, j := row-1, col+1; i >= 0 && j < n; i, j = i-1, j+1 {
		if board[i][j] == 'Q' {
			return false
		}
	}
	for i := 0; i < row; i++ {
		if board[i][col] == 'Q' {
			return false
		}
	}
	return true
}

func solveNQueens(n int) [][]string {
	ans := [][]string{}
	board := make([][]byte, n)

	for i := range n {
		row := make([]byte, n)
		for j := range n {
			row[j] = '.'
		}
		board[i] = row
	}

	var dfs func(int)
	dfs = func(row int) {
		if row >= n {
			tmp := make([]string, n)
			for i := range n {
				tmp[i] = string(board[i])
			}
			ans = append(ans, tmp)
		}

		for j := 0; j < n; j++ {
			if isValid(board, row, j) {
				board[row][j] = 'Q'
				dfs(row + 1)
				board[row][j] = '.'
			}
		}
	}
	dfs(0)
	return ans
}

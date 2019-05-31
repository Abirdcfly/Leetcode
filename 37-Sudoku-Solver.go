func solveSudoku(board [][]byte) {
	row, col, box := make([]map[byte]bool, 9), make([]map[byte]bool, 9), make([]map[byte]bool, 9)
	for i := 0; i < 9; i++ {
		row[i] = make(map[byte]bool)
		col[i] = make(map[byte]bool)
		box[i] = make(map[byte]bool)
	}
	getmap(board, row, col, box)
	dfs(0, 0, board, row, col, box)
	return
}
func getmap(board [][]byte, row, col, box []map[byte]bool) {
	for r := 0; r < 9; r++ {
		for c := 0; c < 9; c++ {
			if board[r][c] == '.' {
				continue
			}
			row[r][board[r][c]-byte('0')] = true
			col[c][board[r][c]-byte('0')] = true
			box[r/3*3+c/3][board[r][c]-byte('0')] = true
		}
	}
	return
}

func dfs(r, c int, board [][]byte, row, col, box []map[byte]bool) bool {
	if r == 9 {
		return true
	}

	nc := (c + 1) % 9
	nr := r
	if nc == 0 {
		nr = r + 1
	}

	if board[r][c] != '.' {
		return dfs(nr, nc, board, row, col, box)
	}

	for i := 1; i < 10; i++ {
		s1, _ := row[r][byte(i)]
		s2, _ := col[c][byte(i)]
		s3, _ := box[r/3*3+c/3][byte(i)]
		if !s1 && !s2 && !s3 {
			board[r][c] = byte(i) + byte('0')
			row[r][board[r][c]-byte('0')] = true
			col[c][board[r][c]-byte('0')] = true
			box[r/3*3+c/3][board[r][c]-byte('0')] = true
			if dfs(nr, nc, board, row, col, box) {
				return true
			}
			row[r][board[r][c]-byte('0')] = false
			col[c][board[r][c]-byte('0')] = false
			box[r/3*3+c/3][board[r][c]-byte('0')] = false
			board[r][c] = byte('.')
		}
	}
	return false
}

//自己理解写的
func solveSudoku(board [][]byte) {
	n := len(board)
	row, col, box := [9][9]bool{}, [9][9]bool{}, [9][9]bool{}
	check(n, board, &row, &col, &box)
	dfs(0, 0, n, board, &row, &col, &box)
	return
}

func dfs(r, c, n int, board [][]byte, row, col, box *[9][9]bool) bool {
	if r >= n {
		return true
	}
	nr := r
	nc := (c + 1) % n
	if nc == 0 {
		nr++
	}

	if board[r][c] != '.' {
		return dfs(nr, nc, n, board, row, col, box)
	}

	for i := 1; i <= 9; i++ {
		if row[r][i-1] || col[c][i-1] || box[r/3*3+c/3][i-1] {
			continue
		}
		board[r][c] = byte(i + int('0'))
		row[r][i-1], col[c][i-1], box[r/3*3+c/3][i-1] = true, true, true
		res := dfs(nr, nc, n, board, row, col, box)
		if !res {
			board[r][c] = '.'
			row[r][i-1], col[c][i-1], box[r/3*3+c/3][i-1] = false, false, false
		} else {
			return true
		}
	}
	return false
}

func check(n int, board [][]byte, row, col, box *[9][9]bool) {
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			v := board[i][j]
			if v == '.' {
				continue
			}
			vint := v - '0'
			row[i][vint-1] = true
			col[j][vint-1] = true
			box[i/3*3+j/3][vint-1] = true
		}
	}
	return
}

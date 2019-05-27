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

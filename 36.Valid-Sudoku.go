func isValidSudoku(board [][]byte) bool {
	n := len(board)
	m := make([]map[byte]bool, 27) //前9个行 中9个列 后9个正方形
	for i := 0; i < 27; i++ {
		m[i] = make(map[byte]bool)
	}
	for r := 0; r < n; r++ {
		for c := 0; c < n; c++ {
			val := board[r][c]
			if val == '.' {
				continue
			}
			if _, e := m[r][val]; e {
				return false
			}
			m[r][val] = true

			if _, e := m[c+9][val]; e {
				return false
			}
			m[c+9][val] = true

			if _, e := m[18+c/3*3+r/3][val]; e {
				return false
			}
			m[18+c/3*3+r/3][val] = true
		}
	}
	return true
}

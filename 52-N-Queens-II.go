func totalNQueens(n int) int {
	res := 0
	if n <= 0 {
		return res
	}
	row, pie, na := make([]bool, n), make([]bool, 2*n+1), make([]bool, 2*n+1)
	dfs(0, n, &res, row, pie, na, []int{})
	return res
}

func dfs(r, n int, res *int, row, pie, na []bool, resCur []int) {
	if r == n {
		(*res)++
		return
	}
	for c := 0; c < n; c++ {
		if row[c] || pie[c+r] || na[c-r+n] {
			continue
		}
		row[c], pie[c+r], na[c-r+n] = true, true, true
		dfs(r+1, n, res, row, pie, na, append(resCur, c))
		row[c], pie[c+r], na[c-r+n] = false, false, false
	}
}

//用位运算也可以做
func totalNQueens(n int) int {
	res := 0
	dfs(0, n, &res, 0, 0, 0)
	return res
}

func dfs(r, n int, res *int, col, pie, na int) {
	if r >= n {
		(*res)++
		return
	}
	bits := (^(col | pie | na)) & (1<<uint(n) - 1)
	for bits != 0 {
		p := bits & (-bits)
		dfs(r+1, n, res, col|p, (pie|p)<<1, (na|p)>>1)
		bits = bits & (bits - 1)
	}
	return
}

func solveNQueens(n int) [][]string {
	if n <= 0 {
		return nil
	}
	resAll := make([][]int, 0)
	resCur := make([]int, 0, n)
	col, pie, na := make([]bool, n), make([]bool, 2*n+1), make([]bool, 2*n+1)
	dfs(0, n, col, pie, na, resCur, &resAll)
	resNeed := make([][]string, len(resAll))
	for index, i := range resAll {
		resNeed[index] = make([]string, n)
		for j := 0; j < n; j++ {
			for t := 0; t < i[j]; t++ {
				resNeed[index][j] += "."
			}
			resNeed[index][j] += "Q"
			for t := 0; t < n-1-i[j]; t++ {
				resNeed[index][j] += "."
			}
		}
	}
	return resNeed
}

func dfs(r, n int, col, pie, na []bool, resCur []int, resAll *[][]int) {
	if r == n {
		//直接append resCur 会出错，因为resCur共用一个底层数组，反复试验等时候会把已经保存到resAll里的内容改掉。eg n =
		//7时
		t := make([]int, n)
		copy(t, resCur)
		*resAll = append(*resAll, t)
		return
	}
	for c := 0; c < n; c++ {
		if col[c] || pie[r+c] || na[c-r+n] {
			continue
		}
		col[c], pie[r+c], na[c-r+n] = true, true, true
		dfs(r+1, n, col, pie, na, append(resCur, c), resAll)
		col[c], pie[r+c], na[c-r+n] = false, false, false
	}
}

// 用数字代替数组
func solveNQueens(n int) [][]string {
	res := make([][]string, 0)
	dfs(0, n, 0, 0, 0, &res, make([]int, 0))
	return res
}

func dfs(i, n int, col, pie, na int, resAll *[][]string, resNow []int) {
	if i == n {
		r := make([]string, n)
		for index, j := range resNow {
			b := make([]rune, n)
			for x := 0; x < n; x++ {
				if j == 1<<uint(x) {
					b[x] = 'Q'
				} else {
					b[x] = '.'
				}
			}
			r[index] = string(b)
		}
		*resAll = append(*resAll, r)
		return
	}
	for choices := (1<<uint(n) - 1) & (^(col | pie | na)); choices != 0; choices = choices & (choices - 1) {
		c := choices & (-choices)
		dfs(i+1, n, col|c, (pie|c)<<1, (na|c)>>1, resAll, append(resNow, c))
	}
}


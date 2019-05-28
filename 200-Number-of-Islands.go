func numIslands(grid [][]byte) int {
	// set 是最笨的方法了吧？
	if len(grid) == 0 {
		return 0
	}
	m := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == '1' {
				bfs(i, j, grid)
				m++
				//p(grid)
			}
		}
	}
	return m
}

func bfs(r, c int, grid [][]byte) {
	q := make([][]int, 1)
	dr := [4]int{0, 1, 0, -1}
	dc := [4]int{1, 0, -1, 0}
	q[0] = []int{r, c}
	for len(q) != 0 {
		t := q[len(q)-1]
		q = q[:len(q)-1]
		r, c := t[0], t[1]
		grid[r][c] = '0'
		for i := 0; i < 4; i++ {
			nr, nc := r+dr[i%4], c+dc[i%4]
			if nr == r && nc == c {
				continue
			}
			if nr >= 0 && nr < len(grid) && nc >= 0 && nc < len(grid[0]) && grid[nr][nc] == '1' {
				q = append(q, []int{nr, nc})
			}
		}
	}
	return
}

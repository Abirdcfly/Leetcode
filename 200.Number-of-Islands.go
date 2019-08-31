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

//并查集
func numIslands(grid [][]byte) int {
	if len(grid) == 0 {
		return 0
	}
	m, n := len(grid), len(grid[0])
	UF := make([]int, m*n)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == '1' {
				UF[n*i+j] = n*i + j
			}
		}
	}
	dr := [4]int{0, 1, 0, -1}
	dc := [4]int{1, 0, -1, 0}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == '1' {
				for k := 0; k < 4; k++ {
					ni, nj := i+dr[k%4], j+dc[k%4]
					if ni >= 0 && ni < m && nj >= 0 && nj < n && grid[ni][nj] == '1' {
						union(UF, UF[n*i+j], UF[n*ni+nj])
					}
				}
			}
		}
	}
	count := 0
	for i := range UF {
		if UF[i] == i && grid[i/n][i%n] == '1' { // 防止[["0"]] 这种错误
			count++
		}
	}
	return count
}

func union(UF []int, i, j int) {
	iroot := find(UF, i)
	jroot := find(UF, j)
	if iroot != jroot {
		UF[jroot] = iroot
	}
}

func find(UF []int, i int) int {
	for UF[i] != i {
		i = UF[i]
	}
	return i
}

func findCircleNum(M [][]int) int {
	n := len(M)
	if n == 0 || n == 1 {
		return n
	}
	UF := make([]int, n*n)
	count := 0
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			UF[i*n+j] = i*n + j
		}
	}

	for i := 0; i < n; i++ {
		for j := i; j < n; j++ {
			if M[i][j] == 1 {
				for k := j; k < n; k++ {
					ni, nj := i, k
					if ni >= 0 && ni < n && nj >= 0 && nj < n && M[ni][nj] == 1 {
						union(UF, UF[i*n+j], UF[ni*n+nj])
					}
				}
				for k := i; k < n; k++ {
					ni, nj := k, j
					if ni >= 0 && ni < n && nj >= 0 && nj < n && M[ni][nj] == 1 {
						union(UF, UF[i*n+j], UF[ni*n+nj])
					}
				}
			}
		}
	}
	for i := 0; i < n; i++ {
		for j := i; j < n; j++ {
			if M[i][j] == 1 && UF[i*n+j] == i*n+j {
				count++
			}
		}
	}
	return count
}

func union(UF []int, i, j int) {
	iroot := find(UF, i)
	jroot := find(UF, j)
	if iroot < jroot {
		UF[jroot] = iroot
	} else {
		UF[iroot] = jroot
	}
}

func find(UF []int, i int) int {
	for UF[i] != i {
		i = UF[i]
	}
	return i
}

// 上面想复杂了。正确且简单的Union Find
func findCircleNum(M [][]int) int {
	n := len(M)
	if n == 0 || n == 1 {
		return n
	}
	UF := make([]int, n)
	count := 0
	for i := 0; i < n; i++ {
		UF[i] = i
	}
	for i := 0; i < n; i++ {
		for j := i; j < n; j++ {
			if M[i][j] == 1 {
				union(UF, i, j)
			}
		}
	}
	for i := 0; i < n; i++ {
		if UF[i] == i {
			count++
		}
	}
	return count
}

func union(UF []int, i, j int) {
	iroot := find(UF, i)
	jroot := find(UF, j)
	if iroot < jroot {
		UF[jroot] = iroot
	} else {
		UF[iroot] = jroot
	}
}

func find(UF []int, i int) int {
	for UF[i] != i {
		i = UF[i]
	}
	return i
}

//uf 写的行数少点
func findcirclenum(m [][]int) int {
    root := func(uf []int, i int) int {
        for uf[i] != i{
            i = uf[i]
        }
        return i
    }
    join := func(uf []int, i, j int){
        x := root(uf, i)
        y := root(uf, j)
        if x < y {
            uf[y] = x
        }else{
            uf[x] = y
        }
    }
    uf := make([]int, len(m))
    for i:= range uf{
        uf[i] = i
    }
    for i:= range uf{
        for j:=range uf{
            if m[i][j] == 1{
                join(uf,i,j)
            }
        }
    }
    n:=0
    for i:= range uf{
        if uf[i] == i{
            n++
        }
    }
    return n
}

//dfs
func findCircleNum(M [][]int) int {
	count := 0
	visited := make([]bool, len(M))
	for i := range M {
		if !visited[i] {
			dfs(i, visited, M)
			count++
		}
	}
	return count
}

func dfs(i int, visited []bool, M [][]int) {
	for j := range M {
		if M[i][j] == 1 && !visited[j] {
			visited[j] = true
			dfs(j, visited, M)
		}
	}
}

//dfs 差不多的
func findCircleNum(M [][]int) int {
    n := 0
    if len(M) == 0 {
        return n
    }
    visited := make([]bool, len(M))
    for i:= range M{
        if !visited[i]{
            n++
        }
        dfs(i, M, visited)
    }
    return n
}

func dfs(i int, M [][]int, visited []bool){
    if visited[i] {
        return
    }
    visited[i] = true
    for j:= range M{
        if j!= i && M[i][j] == 1 {
            dfs(j, M, visited)
        }
    }
}


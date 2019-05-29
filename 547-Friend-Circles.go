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


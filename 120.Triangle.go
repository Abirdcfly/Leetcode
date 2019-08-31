import (
	"math"
	"sort"
)

func minimumTotal(triangle [][]int) int {
	if l := len(triangle); l == 0 {
		return 0
	}
	m := make([][]int, len(triangle))
	for i := range m {
		m[i] = make([]int, i+1)
	}
	m[0][0] = triangle[0][0]
	for r := 1; r < len(triangle); r++ {
		t := triangle[r][0]
		m[r][0] = m[r-1][0] + t
		for c := 1; c < r; c++ {
			t := triangle[r][c]
			m[r][c] = int(math.Min(float64(t+m[r-1][c-1]), float64(t+m[r-1][c])))
		}
		t = triangle[r][r]
		m[r][r] = m[r-1][r-1] + t
	}
	sort.Ints(m[len(m)-1])
	return m[len(m)-1][0]
}

func minimumTotal(triangle [][]int) int {
	l := len(triangle)
	if l == 0 {
		return 0
	}
	m := make([]int, l)
	copy(m, triangle[l-1])
	for r := l - 2; r >= 0; r-- {
		for c := 0; c <= r; c++ {
			m[c] = triangle[r][c] + int(math.Min(float64(m[c]), float64(m[c+1])))
		}
	}
	return m[0]
}

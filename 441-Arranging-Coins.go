import "math"

func arrangeCoins(n int) int {
	if n < 2 {
		return n
	}
	// r*(r+1)/2<= n so maxr <= 1/2 *(-1+sqrt(8*n+1))
	// 或者暴力算
	maxr := (-1 + math.Sqrt(float64(8*n+1))) / 2
	for r := int(maxr) + 1; r > 0; r-- {
		if r*(r+1)/2 <= n {
			return r
		}
	}
	return -1
}

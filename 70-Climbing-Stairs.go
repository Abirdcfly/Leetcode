func climbStairs(n int) int {
	if n <= 0 {
		return 0
	}
	if n == 1 || n == 2 {
		return n
	}
	res := make([]int, n+1)
	res[1] = 1
	res[2] = 2
	for i := 3; i <= n; i++ {
		res[i] = res[i-1] + res[i-2]
	}
	return res[n]
}

func climbStairs(n int) int {
	if n <= 0 {
		return 0
	}
	if n == 1 || n == 2 {
		return n
	}
	rn := 0
	rn_2 := 1
	rn_1 := 2
	for i := 3; i <= n; i++ {
		rn = rn_2 + rn_1
		rn_2 = rn_1
		rn_1 = rn
	}
	return rn
}

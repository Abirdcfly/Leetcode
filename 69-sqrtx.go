func mySqrt(x int) int {
	res := x
	for res*res > x {
		res = (res + x/res) / 2 //牛顿迭代法求开方
	}
	return res
}

func mySqrt(x int) int {
	l, r := 0, x
	for l <= r {
		m := l + (r-l)>>1
		if m*m == x {
			return m
		} else if m*m > x {
			r = m - 1
		} else {
			l = m + 1
		}
	}
	return r
}

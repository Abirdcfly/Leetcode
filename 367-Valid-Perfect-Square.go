func isPerfectSquare(num int) bool {
	l, r, m := 1, num, 0
	for l <= r {
		m = l + (r-l)>>1
		if m*m == num {
			return true
		} else if m*m > num {
			r = m - 1
		} else {
			l = m + 1
		}
	}
	return false
}

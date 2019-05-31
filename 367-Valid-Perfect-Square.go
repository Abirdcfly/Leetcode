func isPerfectSquare(num int) bool {
	l, r, m := 1, num, 0
	for l <= r {
		m = l + (r-l)>>1
		if m*m == num {
			//一般是num/m == m 处理，为了所谓不溢出。M*M很可能太大溢出。但是这里用除法可能会遇到 5/2==2
			//的情况，显然，5不是某个数的平方
			return true
		} else if m*m > num {
			r = m - 1
		} else {
			l = m + 1
		}
	}
	return false
}

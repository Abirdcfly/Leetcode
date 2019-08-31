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

//用小于号要记得考虑1
func isPerfectSquare(num int) bool {
    if num == 1{
        return true
    }
    s,e,m := 1,num,0
    for s<e{
        m = (s+e)/2
        if m*m == num{
            return true
        }
        if m*m < num{
            s = m+1
        }else{
            e = m
        }
    }
    return false
}

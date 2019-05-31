func mySqrt(x int) int {
	res := x
	for res*res > x {
		res = (res + x/res) / 2 //牛顿迭代法求开方
	}
	return res
}

//https://www.zhihu.com/question/20690553/answer/146104283

func mySqrt(x int) int {
	s, e, mid := 1, x, 0
	for s <= e {
		mid = s + (e-s)>>1
		if x/mid == mid {
			return mid
		} else if x/mid > mid {
			s = mid + 1
		} else {
			e = mid - 1
		}
	}
	return e
}

//https://leetcode.com/problems/sqrtx/discuss/25047/A-Binary-Search-Solution/24042

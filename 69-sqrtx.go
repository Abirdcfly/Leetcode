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

//使用二分 求第一个大于等于
func mySqrt(x int) int {
    if x == 0{
        return 0
    }
    s,e,m :=1,x,0
    for s<e{
        m = (s+e)/2
        if m2:=m*m; m2 == x{
            return m
        }else if m2< x{
			//额外步骤，因为求的实际是最大小于
            if (m+1)*(m+1) > x{
                return m
            } 
            s = m+1
        }else{
            e = m
        }
    }
    return s
}

//另一种用大于等于第一个数的实现。
func mySqrt(x int) int {
    if x == 0{
        return 0
    }
    if x == 1 {
        return 1
    }
    s,e,m :=1,x,0
    for s<e{
        m = (s+e)/2
        if m2:=m*m; m2 == x{
            return m
        }else if m2 < x{
            s = m +1
        }else{
            e = m 
        }
    }
    return s-1 //除去1这种例外,没进循环体。其他得出的返回值或者是正好等于，或者是大于，大于的情况减1就好
}

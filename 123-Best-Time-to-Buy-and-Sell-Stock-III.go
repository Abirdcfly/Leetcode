import "math"

func Max(i ...int) (max int) {
	if len(i) == 0 {
		return
	}
	max = i[0]
	for _, j := range i {
		if j > max {
			max = j
		}
	}
	return
}

func maxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}
	mp := make([][]int, 0)
	// mp[i][j] 表示第i天时最大利润，j=0,1,2,3分别表示第i天时，买了1次，卖了1次，买了2次卖了1次，买卖2次
	//注意初始值，和条件比如买2次只能在第3天及以后发生。
	for i := 0; i < len(prices); i++ {
		mp = append(mp, []int{-prices[0], math.MinInt64, math.MinInt64, math.MinInt64})
	}
	res := 0
	for i := 1; i < len(prices); i++ {
		mp[i][0] = Max(-prices[i], mp[i-1][0])
		if i > 0 {
			mp[i][1] = Max(mp[i-1][0]+prices[i], mp[i-1][1])
		}
		if i > 1 {
			mp[i][2] = Max(mp[i-1][2], mp[i-1][1]-prices[i])
		}
		if i > 2 {
			mp[i][3] = Max(mp[i-1][3], mp[i-1][2]+prices[i])
		}
		res = Max(res, mp[i][1], mp[i][3], 0)
	}
	//fmt.Println(mp)
	return res
}

import "math"

func Max(i ...int) (res int) {
	res = i[0]
	for _, j := range i {
		if j > res {
			res = j
		}
	}
	return
}

func maxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}
	res := 0
	n := len(prices)
	mp := make([][][]int, n)
	//mp[i][j][k] 表示第i天，j{0:当天最后未持有股票 1:当天最后持有股票} k{0:当天没卖股票 1:当天卖了股票}
	for i := range mp {
		mp[i] = make([][]int, 2)
		for j := range mp[i] {
			mp[i][j] = make([]int, 2)
		}
	}
	inf := math.MinInt64
	mp[0][0][0] = 0
	mp[0][0][1] = inf
	mp[0][1][0] = -prices[0]
	mp[0][1][1] = inf
	for i := 1; i < n; i++ {
		mp[i][0][0] = Max(mp[i-1][0][0], mp[i-1][0][1])
		mp[i][0][1] = mp[i-1][1][0] + prices[i]
		mp[i][1][0] = Max(mp[i-1][1][0], mp[i-1][0][0]-prices[i])
		mp[i][1][1] = inf
	}
	res = Max(mp[n-1][0]...)
	return res
}

//https://www.youtube.com/watch?v=oL6mRyTn56M 这个很清晰

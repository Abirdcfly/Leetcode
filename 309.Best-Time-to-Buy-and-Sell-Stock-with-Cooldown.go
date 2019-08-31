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

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}
func maxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}
	dp := make([][]int, len(prices)) // dp[i][j] 代表当天状态是持有还是未持有
	for i := range dp {
		dp[i] = make([]int, 2)
	}
	dp[0][1] = -prices[0]
	for i := 1; i < len(prices); i++ {
		if i < 2 { // 只有第2天是没有冷却一说的。
			dp[i][1] = max(dp[i-1][1], dp[i-1][0]-prices[i])
		} else {
			dp[i][1] = max(dp[i-1][1], dp[i-2][0]-prices[i])
		}
		dp[i][0] = max(dp[i-1][0], dp[i-1][1]+prices[i])
	}
	return dp[len(prices)-1][0]
}

//上面方法的优化
func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}
func maxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}
	dp := [2][2]int{} //前一个是今天昨天，后一个是是否持有股票
	dp[0][1] = -prices[0]
	for i := 1; i < len(prices); i++ {
		t, y, yby := i%2, (i+1)%2, (i+2)%2 // 今天 昨天 前天
		dp[t][1] = max(dp[y][1], dp[yby][0]-prices[i])
		dp[t][0] = max(dp[y][0], dp[y][1]+prices[i])
	}
	return dp[(len(prices)-1)%2][0]
}

//自己另一种ac版本 状态定义的好写的会很顺
func maxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}
	max := func(i, j int) int {
		if i > j {
			return i
		}
		return j
	}
	dp := make([][][]int, 2) // 昨天和今天
	for i := range dp {
		dp[i] = make([][]int, 2) // 当天是否持有
		for j := range dp[i] {
			dp[i][j] = make([]int, 2) //当天是否卖出
		}
	}
	dp[0][1][0] = -prices[0]
	for i := 1; i < len(prices); i++ {
		t, y := i%2, (i-1)%2
		dp[t][0][0] = max(dp[y][0][0], dp[y][0][1])
		dp[t][0][1] = dp[y][1][0] + prices[i]
		dp[t][1][0] = max(dp[y][1][0], dp[y][0][0]-prices[i])
	}
	return max(dp[(len(prices)-1)%2][0][0], dp[(len(prices)-1)%2][0][1])
}

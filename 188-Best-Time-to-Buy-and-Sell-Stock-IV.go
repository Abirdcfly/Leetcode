import "math"

//状态要定义好，下面的定义会导致超时，但是结果是正确的
//不是状态问题，是数组太大，优化是搞2个数组就行，一个昨天一个今天，循环用，不必保留所有天的记录
//上次问题找差了，是那种k远大于价格数组长度的时候，k应该变成价格数组长度的一半而不是继续用k。
func Max(array ...int) int {
	if len(array) == 0 {
		return 0
	}
	max := array[0]
	for i := 1; i < len(array); i++ {
		if array[i] > max {
			max = array[i]
		}
	}
	return max
}

func maxProfit(k int, prices []int) int {
	if k < 2 || len(prices) == 0 {
		return 0
	}
	mp := make([][]int, 0)
	//mp[i][j]表示第i天，j是 买1次，卖1次，买2次。。。。。。
	lenr := k * 2
	if lenr > len(prices) {
		lenr = len(prices)
	}
	for i := 0; i < len(prices); i++ {
		r := make([]int, lenr)
		mp = append(mp, r)
	}
	res := 0
	mp[0][0] = -prices[0]
	for i := 1; i < len(prices); i++ {
		mp[i][0] = Max(mp[i-1][0], -prices[i])
		for j := 1; j < k*2 && j <= i; j++ {
			mp[i][j] = Max(mp[i-1][j], mp[i-1][j-1]+prices[i]*int(math.Pow((-1), float64(j+1))))
		}
		res = Max(Max(mp[i]...), res)
	}
	return res
}

//ac
func Max(array ...float64) float64 {
	if len(array) == 0 {
		return 0
	}
	max := array[0]
	for i := 1; i < len(array); i++ {
		if array[i] > max {
			max = array[i]
		}
	}
	return max
}

func maxProfit(k int, prices []int) int {
	if k < 1 || len(prices) == 0 {
		return 0
	}
	res := 0.0
	if k >= len(prices)/2 {
		//退化为不限制交易次数的情形
		res := 0
		for i := 1; i < len(prices); i++ {
			if prices[i] > prices[i-1] {
				res += prices[i] - prices[i-1]
			}
		}
		return res
	}
	mp := make([][][]float64, len(prices))
	//mp[i][j][k] 表示第i+1天，j为0表示第i天未持有股票，1为持有 k为到目前的成功交易次数取值[0，k+1]。
	//mp[0][1][0] 表示第1天，持有，已成功交易0次
	for i := 0; i < 2; i++ { //本来是len(prices)长度，但是leetcode给了一个1W数组1W交易的例子，这么搞内存占太大，通不过，只用昨天和今天来处理
		mp[i] = make([][]float64, 2)
		for j := range mp[i] {
			mp[i][j] = make([]float64, k+1)
		}
	}
	mp[0][0][0] = 0
	for t := 1; t <= k; t++ {
		mp[0][0][t] = math.Inf(-1)
	}
	mp[0][1][0] = float64(-prices[0]) // 次数，认为是卖出后次数加1
	for t := 1; t <= k; t++ {
		mp[0][1][t] = math.Inf(-1)
	}
	for i := 1; i < len(prices); i++ {
		x := i % 2
		y := (i - 1) % 2
		mp[x][0][0] = 0
		mp[x][1][0] = Max(mp[y][1][0], float64(-prices[i]))
		for k1 := 1; k1 <= k; k1++ {
			mp[x][0][k1] = Max(mp[y][0][k1], mp[y][1][k1-1]+float64(prices[i]))
			mp[x][1][k1] = Max(mp[y][1][k1], mp[y][0][k1]-float64(prices[i]))
		}
		res = Max(Max(mp[x][0]...), res)
	}
	return int(res)
}

//另一个ac版本。用float是因为出现了负无穷+一个数然后溢出为正的情况，无奈只能用-inf。实际是状态未定义好
func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}
func maxProfit(k int, prices []int) int {
	if len(prices) == 0 || k <= 0 {
		return 0
	}
	if k > len(prices)/2 {
		k = len(prices) / 2
	}
	dp := make([][][]int, 2) //当天和前一天
	for i := range dp {
		dp[i] = make([][]int, k+1) //交易次数0到k
		for j := range dp[i] {
			dp[i][j] = make([]int, 2) //是否持有股票
		}
	}
	//dp[i][j][k] 表示第i天最多j次交易时的最大收益，k==0是未持有股票 k==1是持有股票
	dp[0][0][1] = -prices[0]
	for i := 1; i <= k; i++ {
		dp[0][i][1] = -prices[0]
	}
	for i := 1; i < len(prices); i++ {
		ith, last := i%2, (i-1)%2
		dp[ith][0][1] = max(-prices[i], dp[last][0][1])
		for j := 1; j <= k; j++ {
			dp[ith][j][0] = max(dp[last][j-1][1]+prices[i], dp[last][j][0])
			dp[ith][j][1] = max(dp[last][j][0]-prices[i], dp[last][j][1])
		}
	}

	ith := (len(prices) - 1) % 2
	res := 0
	for i := 0; i <= k; i++ {
		res = max(res, dp[ith][i][0])
	}
	return res
}

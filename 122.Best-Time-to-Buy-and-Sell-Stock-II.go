import "math"

func maxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}
	// 状态定义是mp[i][0] 表示第i天空仓时收益，mp[i][1]是第i天持有股票时收益
	mp := make([][]int, len(prices))
	mp[0] = []int{0, -prices[0]}
	res := 0
	for i := 1; i < len(prices); i++ {
		mp[i] = make([]int, 2)
		mp[i][0] = int(math.Max(float64(mp[i-1][0]), float64(mp[i-1][1]+prices[i])))
		mp[i][1] = int(math.Max(float64(mp[i-1][1]), float64(mp[i-1][0]-prices[i])))
		res = int(math.Max(float64(mp[i][0]), float64(res)))
	}
	return res
}

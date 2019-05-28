func minDistance(word1 string, word2 string) int {
	n, m := len(word1), len(word2)
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, m+1)
	}
	// dp[i][j] 表示word1前i个字符和word2前j个字符相同时的最小编辑距离
	for i := range dp {
		dp[i][0] = i
	}
	for i := range dp[0] {
		dp[0][i] = i
	}
	for i := 1; i < n+1; i++ {
		for j := 1; j < m+1; j++ {
			if word1[i-1] == word2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				t1, t2, t3 := dp[i-1][j], dp[i][j-1], dp[i-1][j-1]
				min := 0
				if t1 < t2 && t1 < t3 {
					min = t1
				} else if t2 < t3 && t2 < t1 {
					min = t2
				} else {
					min = t3
				}
				dp[i][j] = min + 1
			}
		}
	}
	return dp[n][m]
}

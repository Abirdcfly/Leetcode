/*
 * @lc app=leetcode.cn id=5 lang=golang
 *
 * [5] 最长回文子串
 *
 * https://leetcode-cn.com/problems/longest-palindromic-substring/description/
 *
 * algorithms
 * Medium (26.76%)
 * Likes:    1201
 * Dislikes: 0
 * Total Accepted:    97.6K
 * Total Submissions: 363.6K
 * Testcase Example:  '"babad"'
 *
 * 给定一个字符串 s，找到 s 中最长的回文子串。你可以假设 s 的最大长度为 1000。
 *
 * 示例 1：
 *
 * 输入: "babad"
 * 输出: "bab"
 * 注意: "aba" 也是一个有效答案。
 *
 *
 * 示例 2：
 *
 * 输入: "cbbd"
 * 输出: "bb"
 *
 *
 */
func longestPalindrome(s string) string {
	// dp法
	//dp[i][j] 表示j到i是回文的
	if len(s) == 0 {
		return s
	}
	maxs, maxe := 0, 0
	dp := make([][]bool, len(s))
	for i:= range dp{
		dp[i] = make([]bool, i)
	}
	for i:= range s {
		for j:=0;j<i;j++{
			if i-j == 1 || i-1 == j+1 {
				// c b b     c b b
				// j i       j   i
				dp[i][j] = s[i]==s[j]
			}else{
				//中间不止有0或1个
				dp[i][j] = dp[i-1][j+1] && s[i]==s[j]
			}
			if dp[i][j] && i-j > maxe - maxs{
				maxs, maxe = j, i
			}
		}
	}
	return s[maxs:maxe+1]
}

func longestPalindrome(s string) string {
	if len(s) == 0 {
		return s
	}
	// 中间扩张法
	maxs, maxe := 0, 0
	for i := range s {
		// 奇数，单一中心
		for start, end := i-1, i+1; start >= 0 && end < len(s); start, end = start-1, end+1 {
			if s[start] != s[end] {
				break
			}
			if end-start > maxe-maxs {
				maxs, maxe = start, end
			}
		}
		// 偶数，双中心
		if i == 0 || s[i] != s[i-1] {
			continue
		}
		for start, end := i, i-1; start >= 0 && end < len(s); start, end = start-1, end+1 {
			if s[start] != s[end] {
				break
			}
			if end-start > maxe-maxs {
				maxs, maxe = start, end
			}
		}
	}
	return s[maxs : maxe+1]
}


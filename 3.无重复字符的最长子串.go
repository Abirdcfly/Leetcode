/*
 * @lc app=leetcode.cn id=3 lang=golang
 *
 * [3] 无重复字符的最长子串
 *
 * https://leetcode-cn.com/problems/longest-substring-without-repeating-characters/description/
 *
 * algorithms
 * Medium (30.99%)
 * Likes:    2283
 * Dislikes: 0
 * Total Accepted:    193.4K
 * Total Submissions: 623.6K
 * Testcase Example:  '"abcabcbb"'
 *
 * 给定一个字符串，请你找出其中不含有重复字符的 最长子串 的长度。
 *
 * 示例 1:
 *
 * 输入: "abcabcbb"
 * 输出: 3
 * 解释: 因为无重复字符的最长子串是 "abc"，所以其长度为 3。
 *
 *
 * 示例 2:
 *
 * 输入: "bbbbb"
 * 输出: 1
 * 解释: 因为无重复字符的最长子串是 "b"，所以其长度为 1。
 *
 *
 * 示例 3:
 *
 * 输入: "pwwkew"
 * 输出: 3
 * 解释: 因为无重复字符的最长子串是 "wke"，所以其长度为 3。
 * 请注意，你的答案必须是 子串 的长度，"pwke" 是一个子序列，不是子串。
 *
 *
 */
func lengthOfLongestSubstring(s string) int {
	//最初级的滑动窗口思想
	if len(s) == 0 {
		return 0
	}
	max := 1
	t := []byte{s[0]}
	for i:=1;i<len(s);i++{
		for j:=len(t)-1;j>=0;j--{
			if t[j] == s[i]{
				t = t[j+1:]
				break
			}
		}
		t = append(t, s[i])
		if len(t) > max{
			max = len(t)
		}
	}
	return max
}

func lengthOfLongestSubstring(s string) int {
	//改进滑动窗口
	w := [256]int{}
	for i:= range w{
		w[i] = -1
	}
	l, max := 0, 0
	for i:=0;i<len(s);i++{
		if w[s[i]] >= l {
			l = w[s[i]]+1
		}else if i+1-l > max{
			max = i+1-l
		}
		w[s[i]] = i
	}
	return max
}

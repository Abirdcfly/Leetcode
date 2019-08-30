/*
 * @lc app=leetcode.cn id=14 lang=golang
 *
 * [14] 最长公共前缀
 *
 * https://leetcode-cn.com/problems/longest-common-prefix/description/
 *
 * algorithms
 * Easy (34.53%)
 * Likes:    675
 * Dislikes: 0
 * Total Accepted:    113.9K
 * Total Submissions: 329.5K
 * Testcase Example:  '["flower","flow","flight"]'
 *
 * 编写一个函数来查找字符串数组中的最长公共前缀。
 *
 * 如果不存在公共前缀，返回空字符串 ""。
 *
 * 示例 1:
 *
 * 输入: ["flower","flow","flight"]
 * 输出: "fl"
 *
 *
 * 示例 2:
 *
 * 输入: ["dog","racecar","car"]
 * 输出: ""
 * 解释: 输入不存在公共前缀。
 *
 *
 * 说明:
 *
 * 所有输入只包含小写字母 a-z 。
 *
 */
func longestCommonPrefix(strs []string) string {
	if len(strs) == 0{
		return ""
	}
	r := 0
	s := make([]byte,0)
	var t byte
	for {
		for index, i:= range strs{
			if i == ""{
				return ""
			}
			if r == len(i) {
				return string(s)
			}
			if index == 0 {
				t = i[r]
			}
			if i[r] != t {
				return string(s)
			}
		}
		s = append(s, t)
		r+=1
	}
}


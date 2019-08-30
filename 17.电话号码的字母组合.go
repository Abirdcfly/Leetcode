/*
 * @lc app=leetcode.cn id=17 lang=golang
 *
 * [17] 电话号码的字母组合
 *
 * https://leetcode-cn.com/problems/letter-combinations-of-a-phone-number/description/
 *
 * algorithms
 * Medium (50.72%)
 * Likes:    431
 * Dislikes: 0
 * Total Accepted:    40.9K
 * Total Submissions: 80.5K
 * Testcase Example:  '"23"'
 *
 * 给定一个仅包含数字 2-9 的字符串，返回所有它能表示的字母组合。
 *
 * 给出数字到字母的映射如下（与电话按键相同）。注意 1 不对应任何字母。
 *
 *
 *
 * 示例:
 *
 * 输入："23"
 * 输出：["ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"].
 *
 *
 * 说明:
 * 尽管上面的答案是按字典序排列的，但是你可以任意选择答案输出的顺序。
 *
 */
func letterCombinations(digits string) []string {
	kv := map[string][]string{
		"2": []string{"a", "b", "c"}, "3": []string{"d", "e", "f"}, "4": []string{"g", "h", "i"}, "5": []string{"j", "k", "l"}, "6": []string{"m", "n", "o"}, "7": []string{"p", "q", "r", "s"}, "8": []string{"t", "u", "v"}, "9": []string{"w", "x", "y", "z"},
	}
	if len(digits) == 0 {
		return []string{}
	} else if len(digits) == 1 {
		return kv[digits]
	} else {
		res := []string{}
		last := letterCombinations(digits[1:])
		for _, i := range kv[string(digits[0])] {
			for _, j := range last {
				res = append(res, i+j)
			}
		}
		return res
	}
}



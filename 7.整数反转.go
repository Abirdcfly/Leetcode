/*
 * @lc app=leetcode.cn id=7 lang=golang
 *
 * [7] 整数反转
 *
 * https://leetcode-cn.com/problems/reverse-integer/description/
 *
 * algorithms
 * Easy (32.89%)
 * Likes:    1261
 * Dislikes: 0
 * Total Accepted:    172K
 * Total Submissions: 522.4K
 * Testcase Example:  '123'
 *
 * 给出一个 32 位的有符号整数，你需要将这个整数中每位上的数字进行反转。
 *
 * 示例 1:
 *
 * 输入: 123
 * 输出: 321
 *
 *
 * 示例 2:
 *
 * 输入: -123
 * 输出: -321
 *
 *
 * 示例 3:
 *
 * 输入: 120
 * 输出: 21
 *
 *
 * 注意:
 *
 * 假设我们的环境只能存储得下 32 位的有符号整数，则其数值范围为 [−2^31,  2^31 − 1]。请根据这个假设，如果反转后整数溢出那么就返回
 * 0。
 *
 */
func reverse(x int) int {
	// 转为str算是偷懒吧
	if x<0{
		return -reverse(-x)
	}
	a := strconv.Itoa(x)
	b := []rune(a)
	for left, right := 0, len(b)-1; left < right; left, right = left+1, right-1 {
		b[left], b[right] = b[right], b[left]
	}
	res, err := strconv.Atoi(string(b))
	if err != nil{
		return 0
	}
	if res > 1<<31 -1{
		return 0
	}
	return res
}

func reverse(x int) int {
	if x < 0 {
		return -reverse(-x)
	}
	res := 0
	for ; x != 0; x = x / 10 {
		t := x % 10
		res = res*10 + t
	}
	if res > (1<<31 - 1) {
		return 0
	}
	return res
}

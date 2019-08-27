/*
 * @lc app=leetcode.cn id=118 lang=golang
 *
 * [118] 杨辉三角
 *
 * https://leetcode-cn.com/problems/pascals-triangle/description/
 *
 * algorithms
 * Easy (63.67%)
 * Likes:    180
 * Dislikes: 0
 * Total Accepted:    34.2K
 * Total Submissions: 53.7K
 * Testcase Example:  '5'
 *
 * 给定一个非负整数 numRows，生成杨辉三角的前 numRows 行。
 *
 *
 *
 * 在杨辉三角中，每个数是它左上方和右上方的数的和。
 *
 * 示例:
 *
 * 输入: 5
 * 输出:
 * [
 * ⁠    [1],
 * ⁠   [1,1],
 * ⁠  [1,2,1],
 * ⁠ [1,3,3,1],
 * ⁠[1,4,6,4,1]
 * ]
 *
 */
func generate(numRows int) [][]int {
	if numRows == 0 {
		return nil
	}
	res := make([][]int, numRows)
	res[0] = []int{1}
	for i:=2;i<=numRows;i++{
		row := make([]int, i)
		row[0], row[i-1] = 1, 1
		if i>2 {
			for j:=1;j<i-1;j++{
				row[j] = res[i-2][j-1] + res[i-2][j]
			}
		}
		res[i-1] = row
	}
	return res

}


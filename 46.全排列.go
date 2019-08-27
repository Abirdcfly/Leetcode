/*
 * @lc app=leetcode.cn id=46 lang=golang
 *
 * [46] 全排列
 *
 * https://leetcode-cn.com/problems/permutations/description/
 *
 * algorithms
 * Medium (71.35%)
 * Likes:    358
 * Dislikes: 0
 * Total Accepted:    40.3K
 * Total Submissions: 56.5K
 * Testcase Example:  '[1,2,3]'
 *
 * 给定一个没有重复数字的序列，返回其所有可能的全排列。
 *
 * 示例:
 *
 * 输入: [1,2,3]
 * 输出:
 * [
 * ⁠ [1,2,3],
 * ⁠ [1,3,2],
 * ⁠ [2,1,3],
 * ⁠ [2,3,1],
 * ⁠ [3,1,2],
 * ⁠ [3,2,1]
 * ]
 *
 */
func permute(nums []int) [][]int {
	// 1 + permute([2, 3, 4])
	// 2 + permute([1, 3, 4])
	// 3 + permute([1, 2, 4])
	// 4 + permute([1, 2, 3])

   if len(nums) == 0 {
	   return nil
   }else if len(nums) == 1 {
	   t := make([][]int, 1)
	   t[0] = nums
	   return t
   }
   res := make([][]int, 0)
   for i, key:= range nums{
	   tmp:= make([]int, len(nums))
	   copy(tmp, nums)
	   sub := permute(append(tmp[:i],tmp[i+1:]...))
	   for _, sub1 := range sub{
		   t := append([]int{key}, sub1...)
		   res = append(res, t)
	   }
   }
   return res
}

func permute(nums []int) [][]int {
	// dfs 理解为N个相连的点不重不漏走一遍所有解法
}

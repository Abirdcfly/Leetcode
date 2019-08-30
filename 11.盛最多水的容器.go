/*
 * @lc app=leetcode.cn id=11 lang=golang
 *
 * [11] 盛最多水的容器
 *
 * https://leetcode-cn.com/problems/container-with-most-water/description/
 *
 * algorithms
 * Medium (57.29%)
 * Likes:    748
 * Dislikes: 0
 * Total Accepted:    70.9K
 * Total Submissions: 123.8K
 * Testcase Example:  '[1,8,6,2,5,4,8,3,7]'
 *
 * 给定 n 个非负整数 a1，a2，...，an，每个数代表坐标中的一个点 (i, ai) 。在坐标内画 n 条垂直线，垂直线 i 的两个端点分别为
 * (i, ai) 和 (i, 0)。找出其中的两条线，使得它们与 x 轴共同构成的容器可以容纳最多的水。
 *
 * 说明：你不能倾斜容器，且 n 的值至少为 2。
 *
 *
 *
 * 图中垂直线代表输入数组 [1,8,6,2,5,4,8,3,7]。在此情况下，容器能够容纳水（表示为蓝色部分）的最大值为 49。
 *
 *
 *
 * 示例:
 *
 * 输入: [1,8,6,2,5,4,8,3,7]
 * 输出: 49
 *
 */
// func maxArea(height []int) int {
// 	//暴力法
// 	if len(height) == 0 {
// 		return 0
// 	}
// 	res, t := 0, 0
// 	for i, first := range height {
// 		for j:=i+1;j<len(height);j++{
// 			if second:=height[j]; second>first {
// 				t = (j-i) * first
// 			}else {
// 				t = (j-i) * second
// 			}
// 			if t > res{
// 				res = t
// 			}
// 		}
// 	}
// 	return res
// }

func maxArea(height []int) int {
	//头尾双指针
	if len(height) == 0 {
		return 0
	}
	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}
	max, l, r := 0, 0, len(height)-1
	for l < r {
		if t := min(height[l], height[r]) * (r - l); t > max {
			max = t
		}
		if height[l] > height[r] {
			r--
		} else {
			l++
		}
	}
	return max
}



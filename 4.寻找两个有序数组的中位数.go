/*
 * @lc app=leetcode.cn id=4 lang=golang
 *
 * [4] 寻找两个有序数组的中位数
 *
 * https://leetcode-cn.com/problems/median-of-two-sorted-arrays/description/
 *
 * algorithms
 * Hard (35.85%)
 * Likes:    1451
 * Dislikes: 0
 * Total Accepted:    84.4K
 * Total Submissions: 235.5K
 * Testcase Example:  '[1,3]\n[2]'
 *
 * 给定两个大小为 m 和 n 的有序数组 nums1 和 nums2。
 *
 * 请你找出这两个有序数组的中位数，并且要求算法的时间复杂度为 O(log(m + n))。
 *
 * 你可以假设 nums1 和 nums2 不会同时为空。
 *
 * 示例 1:
 *
 * nums1 = [1, 3]
 * nums2 = [2]
 *
 * 则中位数是 2.0
 *
 *
 * 示例 2:
 *
 * nums1 = [1, 2]
 * nums2 = [3, 4]
 *
 * 则中位数是 (2 + 3)/2 = 2.5
 *
 *
 */
// func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
// 	//不符合时间复杂度为 O(log(m + n))的解法
// 	//合并2个有序数组到新数组
// 	m, n := len(nums1), len(nums2)
// 	t := make([]int, 0)
// 	i, j := 0, 0
// 	for i != m || j != n {
// 		if i == m && j != n {
// 			for k := j; k < n; k++ {
// 				t = append(t, nums2[k])
// 			}
// 			break
// 		} else if i != m && j == n {
// 			for k := i; k < m; k++ {
// 				t = append(t, nums1[k])
// 			}
// 			break
// 		}
// 		if nums1[i] < nums2[j] {
// 			t, i = append(t, nums1[i]), i+1
// 		} else {
// 			t, j = append(t, nums2[j]), j+1
// 		}
// 	}
// 	if sum := m + n; sum%2 == 1 {
// 		return float64(t[sum/2])
// 	} else {
// 		return float64((t[sum/2-1] + t[sum/2])) / 2
// 	}
// }

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	//https://www.youtube.com/watch?v=KB9IcSCDQ9k
	//https://www.youtube.com/watch?v=do7ibYtv5nkga
	//O(log min(m, n))解法
	//m1,m2 数值上等于切割长度，m1+m2 = k 故只需确定一个即可。log时间复杂度说明用了二分思想。
	//在nums1和nums2较短的那个上用二分，时间复杂度为O(log min(m, n))
	// 长度为偶数
	// 					  m1
	// 	nums1     1 3 5 | 7
	// 	nums2     2 4 | 6 8 9 10
	// 					m2
	// 		1 2 3 4 5 | 6 7 8 9 10 前半部分长度为k = (4+6+1)/2 = 11/2 = 5
	// 				cl  cr
	// 长度为奇数
	// 			          m1
	// 	nums1     1 3 5 | 7
	// 	nums2     2 4 6 | 8 9 10 11
	// 			          m2
	// 		1 2 3 4 5 6 | 7 8 9 10 11 前半部分长度为k = (4+7+1)/2 = 12/2 = 6
	// 				  cl
	n1, n2 := len(nums1), len(nums2)
	if n1 > n2 {
		//1. 因为下面的边界条件都是基于n1<=n2假设得出的，注释掉，可能会出现nums2[m2-1]越界问题。（改一下代码可以规避）
		//2. 假设n1<=n2，并在nums1上二分才能实现O(log min(m, n))
		return findMedianSortedArrays(nums2, nums1)
	}
	k := (n1 + n2 + 1) / 2 //前半部分长度
	m1, m2 := 0, 0
	l, r := 0, n1
	for l < r {
		m1 = l + (r-l)/2
		m2 = k - m1
		if nums1[m1] < nums2[m2-1] { //这里nums1[m1-1] < nums2[m2] 也成立。但是考虑到nums1,nums2长度关系。-1可能会越界
			l = m1 + 1
		} else {
			r = m1
		}
	}
	m1, m2 = l, k-l
	cl, cr := 0, 0
	// cl := max(nums1[m1-1], nums2[m2-1])
	// 第一个if实质是 m1==0 ? math.MinInt : nums1[m1-1]
	// 对应的特殊情况
	// nums1  | 1001 1002 1003               或者为空
	// nums2  100 101 102 ... | ... 1000
	// 第二个if同理
	// 实质是考虑nums1的几种情况，为空，|在0， |在中间， |在n
	// nums2考虑 |在0， |在中间， |在n。
	if m1 == 0 {
		cl = nums2[m2-1]
	} else if m2 == 0 {
		cl = nums1[m1-1]
	} else if nums1[m1-1] > nums2[m2-1] {
		cl = nums1[m1-1]
	} else {
		cl = nums2[m2-1]
	}

	if (n1+n2)%2 == 1 {
		return float64(cl)
	}

	// cr := min(nums1[m1], nums2[m2])
	// 第一个if实质是 m1==n1 ? math.MaxInt : nums1[m1]
	// 对应的特殊情况
	// nums1  1 2 3 |
	// nums2  100 101 102 ... | ... 1000
	// 第二个if同理
	if m1 == n1 {
		cr = nums2[m2]
	} else if m2 == n2 {
		cr = nums1[m1]
	} else if nums1[m1] > nums2[m2] {
		cr = nums2[m2]
	} else {
		cr = nums1[m1]
	}
	return float64(cl+cr) / 2
}

/*
 * @lc app=leetcode.cn id=78 lang=golang
 *
 * [78] 子集
 */
func subsets(nums []int) [][]int {
	n := len(nums)
	res := make([][]int,1<<uint(n))
	for i:=0;i<1<<uint(n);i++{
		tmp := make([]int, 0)
		for j:=0;j<n;j++{
			if i&(1<<uint(j))!=0{
				tmp = append(tmp, nums[j])
			}
		}
		res[i] = tmp
	}
	return res
}


import "math"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isValidBST(root *TreeNode) bool {
	res := in(root)
	if len(res) < 2 {
		return true
	}
	for i := 0; i < len(res)-1; i++ {
		if res[i] >= res[i+1] {
			return false
		}
	}
	return true
	/*
			上面判断无重复有序的片段也可借助标准库实现
			for i:= range res {
		        if i>0 && res[i] == res[i-1] {
		            return false
		        }
		    }
		    return sort.IntsAreSorted(res)
	*/
}
func in(root *TreeNode) []int {
	o := make([]int, 0)
	if root == nil {
		return o
	}
	if root.Left != nil {
		o = append(o, in(root.Left)...)
	}
	o = append(o, root.Val)
	if root.Right != nil {
		o = append(o, in(root.Right)...)
	}
	return o
}

// 递归
func isValidBST(root *TreeNode) bool {
	return test(root, math.MinInt64, math.MaxInt64)
}

func test(root *TreeNode, min, max int) bool {
	if nil == root {
		return true
	}
	if root.Val >= max || root.Val <= min {
		return false
	}
	lflag := test(root.Left, min, root.Val)
	rflag := test(root.Right, root.Val, max)
	return lflag && rflag
}

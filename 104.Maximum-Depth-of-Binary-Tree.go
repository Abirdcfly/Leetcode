/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	d1, d2 := 1, 1
	if nil != root.Left {
		d1 = maxDepth(root.Left) + 1
	}
	if nil != root.Right {
		d2 = maxDepth(root.Right) + 1
	}
	if d1 > d2 {
		return d1
	}
	return d2
}

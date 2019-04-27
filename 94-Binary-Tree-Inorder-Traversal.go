/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func inorderTraversal(root *TreeNode) []int {
	res := make([]int, 0)
	in(root, &res)
	return res
}

func in(root *TreeNode, res *[]int) {
	if root == nil {
		return
	}
	if root.Left != nil {
		in(root.Left, res)
	}
	*res = append(*res, root.Val)
	if root.Right != nil {
		in(root.Right, res)
	}
}

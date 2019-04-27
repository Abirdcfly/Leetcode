/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func postorderTraversal(root *TreeNode) []int {
	res := make([]int, 0)
	post(root, &res)
	return res
}

func post(root *TreeNode, res *[]int) {
	if root == nil {
		return
	}
	if root.Left != nil {
		post(root.Left, res)
	}
	if root.Right != nil {
		post(root.Right, res)
	}
	*res = append(*res, root.Val)
}

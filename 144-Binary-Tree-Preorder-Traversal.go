/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func preorderTraversal(root *TreeNode) []int {
	res := make([]int, 0)
	pre(root, &res)
	return res
}
func pre(root *TreeNode, res *[]int) {
	if root == nil {
		return
	}
	*res = append(*res, root.Val)
	if root.Left != nil {
		pre(root.Left, res)
	}
	if root.Right != nil {
		pre(root.Right, res)
	}
}
func preorderTraversal(root *TreeNode) []int {
	res := make([]int, 0)
	stack := make([]*TreeNode, 0)
	for root != nil || len(stack) != 0 {
		for root != nil {
			res = append(res, root.Val)
			stack = append(stack, root)
			root = root.Left
		}
		if len(stack) != 0 {
			root, stack = stack[len(stack)-1], stack[:len(stack)-1]
			root = root.Right
		}
	}
	return res
}

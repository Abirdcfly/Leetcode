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
func inorderTraversal(root *TreeNode) []int {
	//https://www.youtube.com/watch?v=fgEZMCrFrt4
	stack := make([]*TreeNode, 0)
	res := make([]int, 0)
	for root != nil || len(stack) != 0 {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		if len(stack) != 0 {
			root, stack = stack[len(stack)-1], stack[:len(stack)-1]
			res = append(res, root.Val)
			root = root.Right
		}
	}
	return res
}

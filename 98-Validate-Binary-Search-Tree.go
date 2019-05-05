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

/**
 * Definition for TreeNode.
 * type TreeNode struct {
 *     Val int
 *     Left *ListNode
 *     Right *ListNode
 * }
 */
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	return findPorQ(root, p, q)
}
func findPorQ(root, p, q *TreeNode) *TreeNode {
	if root == nil || root == p || root == q {
		return root
	}
	l := findPorQ(root.Left, p, q)
	r := findPorQ(root.Right, p, q)
	if l != nil && r != nil {
		return root
	} else if l == nil {
		return r
	}
	return l
}

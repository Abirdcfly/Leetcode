/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func hasPathSum(root *TreeNode, sum int) bool {
    if root == nil {
        return false
    }
    if root.Left == nil && root.Right == nil && sum - root.Val == 0 {
        return true
    }
    l,r := false,false
    if root.Left != nil {
        l  = hasPathSum(root.Left, sum - root.Val)
    }
    if root.Right != nil {
        r  = hasPathSum(root.Right, sum - root.Val)
    }
    return l || r
}

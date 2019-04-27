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

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func postorderTraversal(root *TreeNode) []int {
	/*简单的2-1-3二叉树
	  前序是123 后序是231
	  先右再左是132 正好是后序反过来。
	*/
	res := make([]int, 0)
	stack := make([]*TreeNode, 0)
	for root != nil || len(stack) != 0 {
		for root != nil {
			stack = append(stack, root)
			res = append(res, root.Val) // res = append(root.Val, res) 在前面放结果，就省去最后做倒序了。
			root = root.Right
		}
		if len(stack) != 0 {
			root, stack = stack[len(stack)-1], stack[:len(stack)-1]
			root = root.Left
		}
	}
	for i := 0; i < len(res)/2; i++ {
		res[i], res[len(res)-1-i] = res[len(res)-1-i], res[i]
	}
	return res
}

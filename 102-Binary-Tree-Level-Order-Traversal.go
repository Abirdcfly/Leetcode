/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
type newQ struct {
	d *TreeNode
	l int
}

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	Q := make([]newQ, 0)
	Q = append(Q, newQ{root, 1})
	res := make([][]int, 0)
	return lo(Q, res)
}

func lo(Q []newQ, res [][]int) [][]int {
	if len(Q) == 0 {
		return res
	}
	root, level := Q[0].d, Q[0].l
	//fmt.Printf("node:%d, level:%d, res:%v\n", root.Val, level, res)
	if level != len(res) {
		res = append(res, make([]int, 0))
	}
	res[level-1] = append(res[level-1], root.Val)
	Q = Q[1:]
	if root.Left != nil {
		Q = append(Q, newQ{root.Left, level + 1})
	}
	if root.Right != nil {
		Q = append(Q, newQ{root.Right, level + 1})
	}
	return lo(Q, res)
}

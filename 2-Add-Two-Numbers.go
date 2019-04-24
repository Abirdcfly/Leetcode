/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	flag := 0
	sums := 0
	head := &ListNode{}
	cur := head
	for l1 != nil || l2 != nil || flag == 1 {
		if l1 != nil {
			sums += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			sums += l2.Val
			l2 = l2.Next
		}
		sums += flag
		if sums >= 10 {
			sums = sums - 10
			flag = 1
		} else {
			flag = 0
		}
		n := ListNode{Val: sums}
		sums = 0
		cur.Next = &n
		cur = cur.Next

	}
	return head.Next
}

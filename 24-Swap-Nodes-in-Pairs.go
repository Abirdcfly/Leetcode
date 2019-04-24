/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	next := head.Next
	head.Next, next.Next = swapPairs(head.Next.Next), head
	return next
}
func swapPairs(head *ListNode) *ListNode {
	node := &ListNode{Next: head}
	pre := node
	for pre.Next != nil && pre.Next.Next != nil {
		cur := pre.Next
		next := pre.Next.Next
		cur.Next = next.Next
		pre.Next = next
		next.Next = cur
		pre = pre.Next.Next
	}
	return node.Next
}


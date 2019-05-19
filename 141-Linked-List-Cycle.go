/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func hasCycle(head *ListNode) bool {
	var slow, fast *ListNode = nil, nil
	slow = head
	fast = head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			return true
		}
	}
	return false
}

//除去快慢指针法和暴力法。实际上用map存节点地址也是很好的办法。而且这种办法更不容易出错，变形题找节点位置也更好做。

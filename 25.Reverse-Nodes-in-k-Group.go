/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseKGroup(head *ListNode, k int) *ListNode {
	nodes := make([]*ListNode, k)
	newNode := head
	for i := 0; i < k; i++ {
		if newNode != nil {
			nodes[i] = newNode
			newNode = newNode.Next
		} else {
			return head
		}
	}

	next := reverseKGroup(nodes[k-1].Next, k)
	for i := 1; i < k; i++ {
		nodes[i].Next = nodes[i-1]
	}
	nodes[0].Next = next
	return nodes[k-1]
}

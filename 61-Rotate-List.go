/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil {
		return head
	}
	t := head
	length := 0
	for t.Next != nil {
		t = t.Next
		length++
	}
	length++
	t.Next = head
	for i := 0; i < length-(k)%length+1; i++ {
		t = t.Next
	}
	j := t
	for j.Next != t {
		j = j.Next
	}
	j.Next = nil
	return t

}

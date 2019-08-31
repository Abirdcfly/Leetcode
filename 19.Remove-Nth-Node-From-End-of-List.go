//看一下https://leetcode.com/articles/remove-nth-node-from-end-of-list/
//还有一个法子是走一遍，记住长度L，然后走第二遍，删掉第L-n+1个节点
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	realhead := &ListNode{Val: 0, Next: head}
	slow := realhead
	fast := realhead
	t := head
	for fast.Next != nil {
		t = slow
		for i := 0; i < n; i++ {
			t = t.Next
		}
		fast = t
		if fast.Next != nil {
			slow = slow.Next
		}
	}
	slow.Next = slow.Next.Next
	return realhead.Next
}

//https://leetcode.com/problems/palindrome-linked-list/discuss/64500/11-lines-12-with-restore-O(n)-time-O(1)-space
/*
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func isPalindrome(head *ListNode) bool {
	var next, pre *ListNode = nil, nil
	fast := head
	slow := head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		next = slow.Next
		slow.Next = pre
		pre = slow
		slow = next
	}

	if fast != nil {
		slow = slow.Next
	}
	for pre != nil && pre.Val == slow.Val {
		pre = pre.Next
		slow = slow.Next
	}
	return pre == nil
}

# Definition for singly-linked list.
# class ListNode(object):
#     def __init__(self, x):
#         self.val = x
#         self.next = None

class Solution(object):
    def detectCycle(self, head):
        """
        :type head: ListNode
        :rtype: ListNode
        """
        try:
            slow = head
            fast = head.next
            while slow != fast :
                slow = slow.next
                fast = fast.next.next

            s1 = head
            while s1 != slow.next:
                s1 = s1.next
                slow = slow.next
            return s1
        except:
            return None

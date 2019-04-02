# Definition for singly-linked list.
# class ListNode(object):
#     def __init__(self, x):
#         self.val = x
#         self.next = None

class Solution(object):
    def mergeKLists(self, lists):
        """
        :type lists: List[ListNode]
        :rtype: ListNode
        """
        min = ListNode(None)
        if not lists:
            return None
        tail = min
        lists = [i for i in lists if i is not None]
        while len(lists) > 0:
            minval = float('inf')
            for index,i in enumerate(lists):
                if minval > i.val:
                    target = i
                    x = index
                    minval = i.val
            tail.next = target
            tail = tail.next
            if target.next != None:
                lists[x] = target.next
            else:
                del lists[x]
            tail.next = None
        return min.next

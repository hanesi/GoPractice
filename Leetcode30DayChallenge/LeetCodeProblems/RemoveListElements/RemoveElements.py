"""
Remove all elements from a linked list
of integers that have value val.

Example:

Input:  1->2->6->3->4->5->6, val = 6
Output: 1->2->3->4->5
"""

# Definition for singly-linked list.
class ListNode(object):
    def __init__(self, val=0, next=None):
        self.val = val
        self.next = next

class Solution(object):
    def removeElements(self, head, val):
        """
        :type head: ListNode
        :type val: int
        :rtype: ListNode
        """
        head, head.next = ListNode(0), head
        p = head
        while p.next:
            if p.next.val == val:
                p.next = p.next.next
            else: p = p.next
        return head.next

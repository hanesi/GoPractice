"""
Given a linked list, remove the n-th node
from the end of list and return its head.

Example:

Given linked list: 1->2->3->4->5, and n = 2.

After removing the second node from the end,
the linked list becomes 1->2->3->5.

Note:
Given n will always be valid.
"""

# Definition for singly-linked list.
class ListNode:
    def __init__(self, val=0, next=None):
        self.val = val
        self.next = next

class Solution:
    def removeNthFromEnd(self, head: ListNode, n: int) -> ListNode:
        length = 0
        lenRun = head
        while lenRun:
            length += 1
            lenRun = lenRun.next

        target = length - n

        tempNode = ListNode(0)
        tempNode.next = head
        lastRun = tempNode

        while target != 0:
            target -= 1
            lastRun = lastRun.next

        lastRun.next = lastRun.next.next
        return tempNode.next

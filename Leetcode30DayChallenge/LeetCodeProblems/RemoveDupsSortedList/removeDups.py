"""
Given a sorted linked list, delete all duplicates
such that each element appear only once.

Example 1:
Input: 1->1->2
Output: 1->2
"""

# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, val=0, next=None):
#         self.val = val
#         self.next = next


class Solution(object):
    def deleteDuplicates(self, head):
        curr = head

        while curr:
            next = curr.next
            while next and next.val == curr.val:
                next = next.next

            curr.next = next
            curr = curr.next

        return head

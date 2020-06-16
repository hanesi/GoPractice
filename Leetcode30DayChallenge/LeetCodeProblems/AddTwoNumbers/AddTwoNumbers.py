"""
You are given two non-empty linked lists representing
two non-negative integers. The digits are stored in
reverse order and each of their nodes contain a single
digit. Add the two numbers and return it as a linked list.

You may assume the two numbers do not contain any
leading zero, except the number 0 itself.

Example:
Input: (2 -> 4 -> 3) + (5 -> 6 -> 4)
Output: 7 -> 0 -> 8
Explanation: 342 + 465 = 807.
"""


class ListNode:
    def __init__(self, val=0, next=None):
        self.val = val
        self.next = next


class Solution:
    def addTwoNumbers(self, l1: ListNode, l2: ListNode) -> ListNode:
        string1 = ""
        string2 = ""

        while l1:
            string1 += str(l1.val)
            l1 = l1.next
        while l2 is not None:
            string2 += str(l2.val)
            l2 = l2.next

        int1 = int(string1[::-1])
        int2 = int(string2[::-1])

        retStr = str(int1+int2)[::-1]

        root = retLL = ListNode(retStr[0])
        for i in range(len(retStr)):
            if i == 0:
                continue
            retLL.next = ListNode(int(retStr[i]))
            retLL = retLL.next
        return root

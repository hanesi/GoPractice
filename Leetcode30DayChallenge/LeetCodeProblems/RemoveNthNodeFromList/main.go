/*
Given a linked list, remove the n-th node
from the end of list and return its head.

Example:

Given linked list: 1->2->3->4->5, and n = 2.

After removing the second node from the end,
the linked list becomes 1->2->3->5.

Note:
Given n will always be valid.
*/
package removeNthFromEnd

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	length := 0
	x := head
	for x != nil {
		length += 1
		x = x.Next
	}
	remIndex := length - n
	start := 0

	// Means we are removing the first element of the list.
	if remIndex == 0 {
		newHead := head.Next
		head = newHead
		return head
	}

	tempNode := &ListNode{}
	current := head

	for start != remIndex {
		start++
		tempNode = current
		current = current.Next
	}

	if current.Next == nil {
		tempNode.Next = nil
	} else if current.Next != nil {
		tempNode.Next = current.Next
		current.Next = nil
		current = nil
	}

	return head
}

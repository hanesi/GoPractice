/*
Given a non-empty, singly linked list with head node head,
return a middle node of linked list.

If there are two middle nodes, return the second middle node.

Example 1:

Input: [1,2,3,4,5]
Output: Node 3 from this list (Serialization: [3,4,5])
The returned node has value 3.
(The judge's serialization of this node is [3,4,5]).
Note that we returned a ListNode object ans, such that:
ans.val = 3, ans.next.val = 4, ans.next.next.val = 5, and ans.next.next.next = NULL.
*/

package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	fmt.Println(middleNode())
}

func middleNode(head *ListNode) *ListNode {
	fastPtr := head
	slowPtr := fastPtr

	for fastPtr.Next != nil {
		slowPtr = slowPtr.Next

		if fastPtr.Next.Next == nil {
			fastPtr = fastPtr.Next
		} else {
			fastPtr = fastPtr.Next.Next
		}
	}

	return slowPtr
}

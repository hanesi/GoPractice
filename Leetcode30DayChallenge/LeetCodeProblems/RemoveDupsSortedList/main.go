/*
Given a sorted linked list, delete all duplicates
such that each element appear only once.

Example 1:
Input: 1->1->2
Output: 1->2
*/
package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	one := &ListNode{1, nil}
	one1 := &ListNode{1, nil}
	two := &ListNode{2, nil}

	testList := one
	testList.Next = one1
	testList.Next = two
	fmt.Println(deleteDuplicates(testList))
}

func deleteDuplicates(head *ListNode) *ListNode {
	curr := head
	res := head
	for head != nil {
		if head.Val != curr.Val {
			curr.Next = head
			curr = head
		} else {
			head, head.Next = head.Next, nil
		}
	}
	return res
}

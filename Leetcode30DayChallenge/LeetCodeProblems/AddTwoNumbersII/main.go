/*
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
*/
package addTwoNumbers

import (
	"strconv"
	"strings"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	string1 := ""
	string2 := ""

	for l1 != nil {
		string1 += strconv.Itoa(l1.Val)
		l1 = l1.Next
	}

	for l2 != nil {
		string2 += strconv.Itoa(l2.Val)
		l2 = l2.Next
	}
	int1, _ := strconv.Atoi(string1)
	int2, _ := strconv.Atoi(string2)
	retIntStr := strconv.Itoa(int1 + int2)

	retSl := strings.Split(retIntStr, "")

	retList := []int{}
	for _, v := range retSl {
		val, _ := strconv.Atoi(v)
		retList = append(retList, val)
	}

	retLL := &ListNode{retList[0], nil}
	for i := 1; i < len(retList); i++ {
		retLL = addNodeEnd(&ListNode{retList[i], nil}, retLL)
	}
	return retLL
}

func addNodeEnd(newNode, NodeLs *ListNode) *ListNode {
	if NodeLs == nil {
		return NodeLs
	}
	for p := NodeLs; p != nil; p = p.Next {
		if p.Next == nil {
			p.Next = newNode
			return NodeLs
		}
	}
	return NodeLs
}

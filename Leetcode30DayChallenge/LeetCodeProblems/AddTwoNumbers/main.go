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
	"fmt"
	"strconv"
	"strings"
)

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	numStr1 := ""
	numStr2 := ""

	for p := l1; p != nil; p = p.Next {
		numStr1 += strconv.Itoa(p.Val)
	}
	for q := l2; q != nil; q = q.Next {
		numStr2 += strconv.Itoa(q.Val)
	}
	num1, _ := strconv.ParseInt(Reverse(numStr1), 10, 64)
	num2, _ := strconv.ParseInt(Reverse(numStr2), 10, 64)

	fmt.Println(numStr1, numStr2)
	fmt.Println(num1, num2)

	retStr := Reverse(strconv.Itoa(int(num1 + num2)))
	retLs := strings.Split(retStr, "")
	retList := []int{}
	for _, v := range retLs {
		val, _ := strconv.Atoi(v)
		retList = append(retList, val)
	}

	retLL := &ListNode{retList[0], nil}
	for i := 1; i < len(retList); i++ {
		retLL = addNodeEnd(&ListNode{retList[i], nil}, retLL)
	}
	return retLL
}

func Reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
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

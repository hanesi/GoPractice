/*
Given a non-empty array of digits representing a
non-negative integer, plus one to the integer.

The digits are stored such that the most significant
digit is at the head of the list, and each element
in the array contain a single digit.

You may assume the integer does not contain any
leading zero, except the number 0 itself.

Example 1:
Input: [1,2,3]
Output: [1,2,4]

Explanation: The array represents the integer 123.
*/
package main

import "fmt"

func main() {
	fmt.Println(plusOne([]int{1, 2, 3, 4}))
}

func plusOne(digits []int) []int {
	var n int = len(digits)
	for i := n - 1; i >= 0; i-- {
		if digits[i] < 9 {
			digits[i] += 1
			return digits
		} else {
			digits[i] = 0
		}
	}
	var a = make([]int, n+1)
	a[0] = 1
	return a
}

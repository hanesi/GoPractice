/*
Determine whether an integer is a palindrome.
An integer is a palindrome when it reads the same backward as forward.

Example 1:

Input: 121
Output: true
*/

package main

import "fmt"

func main() {
	fmt.Println(isPalindrome(123454321))
}

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	xRev := reverse_int(x)
	if xRev == x {
		return true
	}
	return false
}

func reverse_int(n int) int {
	new_int := 0
	for n > 0 {
		remainder := n % 10
		new_int *= 10
		new_int += remainder
		n /= 10
	}
	return new_int
}

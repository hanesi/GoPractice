/*
Write a function that reverses a string. The input
string is given as an array of characters char[].

Do not allocate extra space for another array, you
 must do this by modifying the input array in-place
 with O(1) extra memory.

You may assume all the characters consist of printable
ascii characters.

Example 1:
Input: ["h","e","l","l","o"]
Output: ["o","l","l","e","h"]
*/
package main

import "fmt"

func main() {
	testS := []string{"o", "l", "l", "e", "h"}
	reverseString(testS)
	fmt.Println(testS)
}

func reverseString(s []string) {
	if len(s) == 0 {
		return
	}
	l := len(s) - 1
	for i := 0; i <= l/2; i++ {
		s[i], s[l-i] = s[l-i], s[i]
	}
	return
}

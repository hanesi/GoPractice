/*
Given a string containing only three types of characters:
'(', ')' and '*', write a function to check whether this string is valid.
We define the validity of a string by these rules:

Any left parenthesis '(' must have a corresponding right parenthesis ')'.
Any right parenthesis ')' must have a corresponding left parenthesis '('.
Left parenthesis '(' must go before the corresponding right parenthesis ')'.
'*' could be treated as a single right parenthesis ')' or a single left parenthesis '(' or an empty string.
An empty string is also valid.

Example 1:
Input: "()"
Output: True

Example 2:
Input: "(*)"
Output: True

Example 3:
Input: "(*))"
Output: True

Note:
The string size will be in the range [1, 100].
*/
package main

import "fmt"

func main() {
	fmt.Println(checkValidString("(*))"))
}

func checkValidString(s string) bool {
	l, r := 0, 0
	n := len(s)
	for i := 0; i < n; i++ {
		j := n - i - 1
		if s[i] == ')' {
			l--
		} else {
			l++
		}
		if s[j] == '(' {
			r--
		} else {
			r++
		}
		if l < 0 || r < 0 {
			return false
		}
	}

	return true
}

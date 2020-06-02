/*
Given a string s consists of upper/lower-case alphabets
and empty space characters ' ', return the length of last
word (last word means the last appearing word if we loop
from left to right) in the string.

If the last word does not exist, return 0.

Note: A word is defined as a maximal substring consisting
of non-space characters only.

Example:
Input: "Hello World"
Output: 5
*/

package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(lengthOfLastWord("a "))
	fmt.Println(strings.Split("a ", " "))
}

func lengthOfLastWord(s string) int {
	s = strings.TrimSpace(s)
	wordSlice := strings.Split(s, " ")
	if strings.Contains(wordSlice[len(wordSlice)-1], " ") {
		return 0
	}
	return len(wordSlice[len(wordSlice)-1])
}

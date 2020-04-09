/*
Given two strings S and T,
return if they are equal when
both are typed into empty text editors.
# means a backspace character.

Example 1:
Input: S = "ab#c", T = "ad#c"
Output: true

Explanation: Both S and T become "ac".

Example 2:
Input: S = "a#c", T = "b"
Output: false

Explanation: S becomes "c" while T becomes "b".
*/
package main

import (
	"fmt"
	"strings"
)

func main() {
	S := "a##c"
	T := "a#c#"
	fmt.Println(backspaceCompare(S, T))
}

func backspaceCompare(S string, T string) bool {
	combStr := strings.ToLower(fmt.Sprintf(S + T))
	runeSlice := []rune(combStr)
	retSlice := []rune{}
	for i := 0; i < len(runeSlice)-1; i++ {
		if runeSlice[i] != 35 && runeSlice[i+1] != 35 {
			retSlice = append(retSlice, runeSlice[i])
		}
	}
	if runeSlice[len(runeSlice)-1] != 35 {
		retSlice = append(retSlice, runeSlice[len(runeSlice)-1])
	} else {
		retSlice = retSlice[:len(retSlice)-1]
	}
	retVal := compareSplit(string(retSlice))
	return retVal
}

func compareSplit(input string) bool {
	length := len(input)
	if length%2 != 0 {
		return false
	}
	if input[length/2:] == input[:length/2] {
		return true
	}
	return false
}

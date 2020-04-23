/*
Given a range [m, n] where 0 <= m <= n <= 2147483647,
return the bitwise AND of all numbers in this range, inclusive.

Example 1:
Input: [5,7]
Output: 4

Example 2:
Input: [0,1]
Output: 0
*/
package main

import "fmt"

func main() {
	fmt.Println(rangeBitwiseAnd(5, 7))
}

func rangeBitwiseAnd(m int, n int) int {
	diff := n - m
	bit := uint(0)
	for diff != 0 {
		diff = diff >> 1
		bit++
	}
	return (m >> bit << bit) & (n >> bit << bit)
}

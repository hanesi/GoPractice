/*
This function takes a number and says whether or not it's "happy"
A happy number is a number defined by the following process:
Starting with any positive integer, replace the number by the sum of the squares of its digits,
and repeat the process until the number equals 1 (where it will stay),
or it loops endlessly in a cycle which does not include 1.
Those numbers for which this process ends in 1 are happy numbers.

Example:
Input = 19
Output = true

Explanation:
1^2 + 9^2 = 82
8^2 + 2^2 = 68
6^2 + 8^2 = 100
1^2 + 0^2 + 0^2 = 1
*/
package main

import (
	"fmt"
	"strconv"
)

var counter int

func main() {
	fmt.Println(isHappy(2))
}

func isHappy(n int) bool {
	str := strconv.Itoa(n)
	result := 0
	if counter > 1000 {
		return false
	}
	for _, v := range str {
		digit := (int(v - '0'))
		result += (digit * digit)
	}
	fmt.Println(result)
	// fmt.Println(counter)
	switch {
	case result == 1:
		return true
	case result != 1:
		counter++
		return isHappy(result)
	}
	return false
}

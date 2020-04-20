/*
Go natively employs a binary search with the built in search
function, but this is for practice

Given a sorted list and a number to search, return its index
*/
package main

import "fmt"

func main() {
	testSlice := []int{0, 1, 2, 3, 4, 5, 6, 7, 8}
	fmt.Println(binSearch(7, testSlice))
}

func binSearch(num int, list []int) int {
	left := 0
	right := len(list) - 1
	for left <= right {
		mid := left + (right-left)/2
		if num > list[mid] {
			left++
		} else {
			right--
		}
		if list[mid] == num {
			return mid
		}
	}
	return -1
}

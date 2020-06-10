/*
Given an integer array of size n, find all elements
that appear more than ⌊ n/3 ⌋ times.

Note: The algorithm should run in linear time and in O(1) space.

Example 1:
Input: [3,2,3]
Output: [3]
*/
package main

import "fmt"

func main() {
	fmt.Println(majorityElement([]int{3, 2, 3}))
}

func majorityElement(nums []int) []int {
	ctDict := make(map[int]int)
	ctList := []int{}
	for _, v := range nums {
		ctDict[v]++
	}
	for k, v := range ctDict {
		if v > len(nums)/3 {
			ctList = append(ctList, k)
		}
	}
	return ctList
}

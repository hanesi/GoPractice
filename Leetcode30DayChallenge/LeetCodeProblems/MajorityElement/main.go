/*
Given an array of size n, find the majority element.
The majority element is the element that appears
more than ⌊ n/2 ⌋ times.

You may assume that the array is non-empty and
the majority element always exist in the array.

Example 1:
Input: [3,2,3]
Output: 3

Example 2:
Input: [2,2,1,1,1,2,2]
Output: 2
*/
package main

import "fmt"

func main() {
	fmt.Println(majorityElement([]int{2, 2, 1, 1, 1, 2, 2}))
}

func majorityElement(nums []int) int {
	ctDict := make(map[int]int)
	for _, v := range nums {
		ctDict[v]++
		if ctDict[v] > (len(nums) / 2) {
			return v
		}
	}
	return 0
}

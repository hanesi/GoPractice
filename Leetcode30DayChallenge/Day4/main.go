/*
Given an array nums, write a function to move all 0's to the end of it
while maintaining the relative order of the non-zero elements.

Example:

Input: [0,1,0,3,12]
Output: [1,3,12,0,0]

Note:
You must do this in-place without making a copy of the array.
Minimize the total number of operations.
*/
package main

import "fmt"

func main() {
	testSlice := []int{1, 0, 0, 1}
	fmt.Println(moveZeroes(testSlice))
}

func moveZeroes(nums []int) []int {
	// Main "gotcha" here is to iterate backwards through the list
	// If you iterate forwards and remove elements, your indexing gets screwy
	for i := len(nums) - 1; i >= 0; i-- {
		if nums[i] == 0 {
			nums = remove(nums, i)
			nums = append(nums, 0)
		}
	}
	return nums
}

func remove(slice []int, s int) []int {
	return append(slice[:s], slice[s+1:]...)
}

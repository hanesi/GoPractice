/*
Given a sorted array nums, remove the duplicates in-place such
that each element appear only once and return the new length.

Do not allocate extra space for another array, you must do
this by modifying the input array in-place with O(1) extra memory.

Example 1:

Given nums = [1,1,2],

Your function should return length = 2, with the first two
elements of nums being 1 and 2 respectively.

It doesn't matter what you leave beyond the returned length.
*/
package main

import "fmt"

func main() {
	numSlice := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	fmt.Println(removeDuplicates(numSlice))
}

func removeDuplicates(nums []int) int {
	for i := len(nums) - 1; i > 0; i-- {
		if nums[i] == nums[i-1] {
			nums = append(nums[0:i-1], nums[i:]...)
		}
	}
	return len(nums)
}

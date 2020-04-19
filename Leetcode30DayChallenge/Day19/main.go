/*
Suppose an array sorted in ascending order
is rotated at some pivot unknown to you beforehand.

(i.e., [0,1,2,4,5,6,7] might become [4,5,6,7,0,1,2]).

You are given a target value to search. If found
in the array return its index, otherwise return -1.

You may assume no duplicate exists in the array.

Example 1:
Input: nums = [4,5,6,7,0,1,2], target = 0
Output: 4

Example 2:
Input: nums = [4,5,6,7,0,1,2], target = 3
Output: -1
*/
package main

import "fmt"

func main() {
	testSlice := []int{4, 5, 6, 7, 0, 1, 2}
	targ := 3
	fmt.Println(search(testSlice, targ))
}

func search(nums []int, target int) int {
	left := 0
	right := len(nums) - 1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			return mid
		}
		if nums[mid] <= nums[right] {
			if target <= nums[right] && target >= nums[mid] {
				left = mid + 1
			} else {
				right = mid - 1
			}
		} else {
			if target >= nums[left] && target <= nums[mid] {
				right = mid - 1
			} else {
				left = mid + 1
			}
		}
	}
	return -1
}

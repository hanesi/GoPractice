/*
Given an array of non-negative integers, you are
initially positioned at the first index of the array.

Each element in the array represents your maximum
jump length at that position.

Determine if you are able to reach the last index.

Example 1:
Input: nums = [2,3,1,1,4]
Output: true
Explanation: Jump 1 step from index 0 to 1, then 3 steps to the last index.
*/
package main

import "fmt"

func main() {
	fmt.Println(canJump([]int{2, 3, 1, 1, 4}))
}

func canJump(nums []int) bool {
	lastPos := len(nums) - 1 // last position can reach the end index
	for i := len(nums) - 1; i >= 0; i-- {
		if i+nums[i] >= lastPos {
			lastPos = i
		}
	}
	return lastPos == 0
}

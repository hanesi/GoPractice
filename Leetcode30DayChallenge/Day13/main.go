/*
Given a binary array, find the maximum length of a
contiguous subarray with equal number of 0 and 1.

Example 1:
Input: [0,1]
Output: 2
Explanation: [0, 1] is the longest contiguous
subarray with equal number of 0 and 1.

Example 2:
Input: [0,1,0]
Output: 2
Explanation: [0, 1] (or [1, 0]) is a longest
contiguous subarray with equal number of 0 and 1.

Note: The length of the given binary array will not exceed 50,000.
*/
package main

import "fmt"

func main() {
	fmt.Println(findMaxLength([]int{1, 1, 1, 0, 0, 0}))
}

func findMaxLength(nums []int) int {
	l := len(nums)
	if l < 2 {
		return 0
	}

	for i := 0; i < l; i++ {
		if nums[i] == 0 {
			nums[i] = -1
		}
	}

	sum := 0
	max := 0

	sumToIndexMap := make(map[int]int)
	sumToIndexMap[0] = -1 //Initialize map; sum = 0 at index -1

	for i := 0; i < l; i++ {
		sum += nums[i]
		if _, v := sumToIndexMap[sum]; v {
			max = maxNo(max, i-sumToIndexMap[sum])
		} else {
			sumToIndexMap[sum] = i
		}
	}
	return max

}

func maxNo(x, y int) int {
	if x > y {
		return x
	}
	return y
}

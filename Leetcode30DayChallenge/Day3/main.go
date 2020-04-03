/*
This function ingests a slice of ints, returning the highest sum of a
contiguous sub array.

Example:

Input: [-2,1,-3,4,-1,2,1,-5,4],
Output: 6
Explanation: [4,-1,2,1] has the largest sum = 6.
*/
package main

import "fmt"

func main() {
	fmt.Println(maxSubArray([]int{1, 1, -2}))
}

func maxSubArray(nums []int) int {
	runningTotal := 0
	maxFound := 0
	fmt.Println(nums)
	if len(nums) == 1 {
		return nums[0]
	}
	for _, v := range nums {
		if runningTotal+v > 0 {
			runningTotal += v
			if runningTotal > maxFound {
				maxFound = runningTotal
			}
		} else {
			runningTotal = 0
		}
	}
	if maxFound == 0 {
		maxVal := nums[0]
		for _, v := range nums {
			if maxVal < v {
				maxVal = v
			}
		}
		return maxVal
	}
	return maxFound
}

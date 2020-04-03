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

	// If there's only one number, return it
	if len(nums) == 1 {
		return nums[0]
	}

	// Add current value to running total, updating maxFound when necessary
	for _, v := range nums {
		if runningTotal+v > 0 {
			runningTotal += v
			if runningTotal > maxFound {
				maxFound = runningTotal
			}
			// If runningTotal + v > 0, reset the running total
		} else {
			runningTotal = 0
		}
	}

	// If all numbers are negative, return the max negative value
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

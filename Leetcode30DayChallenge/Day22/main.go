/*
Given an array of integers and an integer k, you need to
find the total number of continuous subarrays whose sum equals to k.

Example 1:
Input:nums = [1,1,1], k = 2
Output: 2

Note:
The length of the array is in range [1, 20,000].
The range of numbers in the array is [-1000, 1000] and
the range of the integer k is [-1e7, 1e7].
*/
package main

import "fmt"

func main() {
	testSlice := []int{1, 1, 1}
	testSum := 2
	fmt.Println(subarraySum(testSlice, testSum))
}

func subarraySum(nums []int, k int) int {
	retVal, sum := 0, 0
	hashMap := make(map[int]int)
	hashMap[0] = 1
	for _, v := range nums {
		sum += v
		retVal += hashMap[sum-k]
		hashMap[sum]++
	}
	return retVal
}

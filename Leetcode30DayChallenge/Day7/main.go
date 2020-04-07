/*
Given an integer array arr, count element x such that x + 1 is also in arr.

If there're duplicates in arr, count them seperately.

Example 1:
Input: arr = [1,2,3]
Output: 2
Explanation: 1 and 2 are counted cause 2 and 3 are in arr.

Example 2:
Input: arr = [1,1,3,3,5,5,7,7]
Output: 0
Explanation: No numbers are counted, cause there's no 2, 4, 6, or 8 in arr.
*/
package main

import (
	"fmt"
)

func main() {
	testSlice := []int{1, 1, 2, 2}
	fmt.Println(countElements(testSlice))
}

func countElements(arr []int) int {
	mapCount := make(map[int]int)
	for _, v := range arr {
		mapCount[v]++
	}
	count := 0
	for k, v := range mapCount {
		if _, ok := mapCount[k+1]; ok {
			count += v
		}
	}
	return count
}

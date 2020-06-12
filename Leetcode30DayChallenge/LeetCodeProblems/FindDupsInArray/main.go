/*
Given an array of integers, 1 ≤ a[i] ≤ n (n = size of array),
some elements appear twice and others appear once.

Find all the elements that appear twice in this array.

Could you do it without extra space and in O(n) runtime?

Example:
Input:
[4,3,2,7,8,2,3,1]
Output:
[2,3]
*/
package main

import "fmt"

func main() {
	fmt.Println(findDuplicates([]int{4, 3, 2, 7, 8, 2, 3, 1}))
}

func findDuplicates(nums []int) []int {
	ctMap := make(map[int]int)
	retList := []int{}
	for _, v := range nums {
		ctMap[v]++
	}
	for k, v := range ctMap {
		if v > 1 {
			retList = append(retList, k)
		}
	}
	return retList
}

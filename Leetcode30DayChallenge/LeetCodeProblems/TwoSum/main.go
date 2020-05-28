/*
This function accepts a list of INTs and a target sum value.
It returns the indices of the two list values that combine
to result in the target sum. No repeated values and only
one possible target sum combination are possible.
*/
package main

import "fmt"

func main() {
	intSlice := []int{2, 7, 11, 15}
	target := 9

	fmt.Println(twoSum(intSlice, target))
}

func twoSum(nums []int, target int) []int {
	hashMap := make(map[int]int)
	for i, v := range nums {
		index, val := hashMap[target-v]
		if val {
			return []int{index, i}
		}
		hashMap[v] = i
	}
	return nil
}

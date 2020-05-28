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

package main

import "fmt"

func main() {
	testSlice := []int{2, 2, 1}
	fmt.Println(singleNumber(testSlice))
}

func singleNumber(nums []int) int {
	ctMap := make(map[int]int)
	for _, v := range nums {
		ctMap[v]++
	}

	for k, v := range ctMap {
		if v == 1 {
			return k
		}
	}
	return -1
}

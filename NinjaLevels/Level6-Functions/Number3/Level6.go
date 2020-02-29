package main

import "fmt"

func main() {
	ls := []int{1, 2, 3, 4, 5, 6, 7}
	ls2 := []int{6, 7, 8, 9}

	defer foo(ls...)
	bar(ls2)
}

func foo(x ...int) int {
	sum := 0
	for _, v := range x {
		sum += v
	}
	fmt.Println("foo sum:", sum)
	return sum
}

func bar(x2 []int) int {
	sum := 0
	for _, v := range x2 {
		sum += v
	}
	fmt.Println("bar sum:", sum)
	return sum
}

package main

import "fmt"

func main() {
	a := plusOne(1)
	fmt.Println(a)
	fmt.Println(a)
	fmt.Println(a)
}

func plusOne(x int) int {
	return x + 1
}

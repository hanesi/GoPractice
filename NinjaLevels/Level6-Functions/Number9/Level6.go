package main

import "fmt"

func main() {
	g := func(ls []int) int {
		if len(ls) == 0 {
			return 0
		}
		if len(ls) == 1 {
			return ls[0]
		}
		return ls[0] + ls[len(ls)-1]
	}
	x := foo(g, []int{1, 2, 3, 4, 5, 6})

	fmt.Println(x)
}

func foo(f func(ls []int) int, ii []int) int {
	n := f(ii)
	n++
	return n
}

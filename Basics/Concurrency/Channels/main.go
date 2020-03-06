package main

import (
	"fmt"
)

var c chan int

func main() {
	c := generate()
	rec(c)

	fmt.Println("almost done")
}

func generate() <-chan []int {
	c := make(chan []int)
	go func() {
		defer close(c)
		sl := []int{1, 2, 3, 4}
		for i := 0; i < len(sl); i++ {
			sl[i] += 5
			newSl := make([]int, len(sl))
			copy(newSl, sl)
			c <- newSl
		}
	}()

	return c
}

func rec(c <-chan []int) {
	for v := range c {
		fmt.Println(v)
	}
}

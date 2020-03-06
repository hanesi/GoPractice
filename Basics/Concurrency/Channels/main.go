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

func generate() <-chan []string] {
	c := make(chan []string)
	go func() {
		i := []string{"one","two","three"}
    c <- i
		close(c)
	}()

	return c
}

func rec(c <-chan []string) {
	for v := range c {
		fmt.Println(v)
	}
}

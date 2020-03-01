package main

import (
	"fmt"
)

func main() {
	// c := make(chan int) DOESNT WORK
	// c := make(chan int, 1) // Buffered channel works
	// Anonymous Func works too
	c := make(chan int)
	go func() {
		c <- 42
	}()

	fmt.Println(<-c)
}

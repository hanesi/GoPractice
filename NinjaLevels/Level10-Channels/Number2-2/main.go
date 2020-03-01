package main

import (
	"fmt"
)

func main() {
	// cr := make(<-chan int) errors because send to recieve only
	cr := make(chan int) //works because bidirectional channel

	go func() {
		cr <- 42
	}()
	fmt.Println(<-cr)

	fmt.Printf("------\n")
	fmt.Printf("cr\t%T\n", cr)
}

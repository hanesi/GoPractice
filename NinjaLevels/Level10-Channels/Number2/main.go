package main

import (
	"fmt"
)

func main() {
	// cs := make(chan<- int) THIS doesnt work, need to generalize
	cs := make(chan int) // this works

	go func() {
		cs <- 42
	}()
	fmt.Println(<-cs)

	fmt.Printf("------\n")
	fmt.Printf("cs\t%T\n", cs)
}

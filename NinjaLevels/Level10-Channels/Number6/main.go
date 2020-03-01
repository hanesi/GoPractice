package main

import "fmt"

func main() {
	c := make(chan int)

	go func() {
		for v := 0; v < 100; v++ {
			c <- v
		}
		close(c)
	}()
	for v := range c {
		fmt.Println(v)
	}
}

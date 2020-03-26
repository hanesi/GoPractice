package main

import "fmt"

func main() {
	c := make(chan int)
	done := make(chan bool)
	// One go routine sending to a channel
	// go func() {
	// 	for i := 0; i < 25; i++ {
	// 		c <- i
	// 	}
	// 	close(c)
	// }()
	//
	// Create 5 go routines that write to a channel
	// for i := 0; i < 5; i++ {
	// 	go func() {
	// 		for j := 0; j < 10; j++ {
	// 			c <- j
	// 		}
	// 		close(c)
	// 	}()
	// }

	// 5 go routines writing to a channel, use another channel to say
	// When it's done
	for i := 0; i < 5; i++ {
		go func() {
			for j := 0; j < 10; j++ {
				c <- j
			}
			done <- true
		}()
	}
	go func() {
		for i := 0; i < 5; i++ {
			<-done
		}
		close(c)
	}()
	for n := range c {
		fmt.Println(n)
	}
}

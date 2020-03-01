package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	wg.Add(2)
	go func() {
		fmt.Println("First Func")
		wg.Done()
	}()

	go func() {
		fmt.Println("Second Func")
		wg.Done()
	}()
	wg.Wait()
	fmt.Println("exiting...")
}

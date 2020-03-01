package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup
var mu sync.Mutex

func main() {
	wg.Add(100)
	fmt.Println("starting")
	go foo()
	wg.Wait()
	fmt.Println("done")
}

func foo() {
	inc := 0
	for i := 0; i < 100; i++ {
		mu.Lock()
		v := inc
		v++
		inc = v
		fmt.Println(inc)
		mu.Unlock()
		wg.Done()
	}
}

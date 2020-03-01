package main

import (
	"fmt"
	"sync"
	"sync/atomic"
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
	var inc int64
	for i := 0; i < 100; i++ {
		atomic.AddInt64(&inc, 1)
		r := atomic.LoadInt64(&inc)
		fmt.Println(r)
		wg.Done()
	}
}

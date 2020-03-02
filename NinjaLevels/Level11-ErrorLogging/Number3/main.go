package main

import (
	"fmt"
)

type customErr struct {
	info string
}

func (ce customErr) Error() string {
	return fmt.Sprintf("There was an error: %v", ce.info)
}

func main() {
	c1 := customErr{
		info: "An error happened",
	}

	foo(c1)
}

func foo(e error) {
	fmt.Println("foo ran -", e)
}

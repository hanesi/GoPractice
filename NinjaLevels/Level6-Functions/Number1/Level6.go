package main

import "fmt"

func main() {
	fmt.Println(foo())
	fmt.Println(bar())
}

func foo() int {
	return 42
}

func bar() (string, int) {
	return "Hello World, it's ", 2020
}

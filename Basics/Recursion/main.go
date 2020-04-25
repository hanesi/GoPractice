package main

import "fmt"

func main() {
	num := 10
	fmt.Println(factorial(num))
	fmt.Println(fibonacci(num))
}

func factorial(n int) int {
	if n <= 0 {
		return 1
	}
	return n * factorial(n-1)
}

func fibonacci(n int) int {
	if n <= 2 {
		return 1
	}
	return fibonacci(n-1) + fibonacci(n-2)
}

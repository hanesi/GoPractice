/*
This function takes an int and returns it reversed. If the int is bigger than
32 bits, it returns 0
*/
package main

import "fmt"

func main() {
	fmt.Println(reverse(123456789))
}

func reverse(x int) int {
	if x < 0 {
		revX := reverse_int(0 - x)
		if revX > 2147483647 {
			return 0
		}
		return 0 - revX
	}
	revX := reverse_int(x)
	if revX > 2147483647 {
		return 0
	}
	return revX
}

func reverse_int(n int) int {
	new_int := 0
	for n > 0 {
		remainder := n % 10
		new_int *= 10
		new_int += remainder
		n /= 10
	}
	return new_int
}

package main

import (
	"fmt"
	"strconv"
)

var counter int

func main() {
	fmt.Println(isHappy(2))
}

func isHappy(n int) bool {
	str := strconv.Itoa(n)
	result := 0
	if counter > 1000 {
		return false
	}
	for _, v := range str {
		digit := (int(v - '0'))
		result += (digit * digit)
	}
	fmt.Println(result)
	// fmt.Println(counter)
	switch {
	case result == 1:
		return true
	case result != 1:
		counter++
		return isHappy(result)
	}
	return false
}

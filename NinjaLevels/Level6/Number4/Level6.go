package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

func (p person) speak() {
	fmt.Println("I am", p.first, p.last)
}
func main() {
	p1 := person{
		first: "John",
		last:  "Doe",
		age:   32,
	}

	p1.speak()
}

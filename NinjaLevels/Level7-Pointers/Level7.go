package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

func changeMe(p *person) {
	p.first = "two"
	p.last = "two"
	p.age = 2
}

func main() {
	x := person{
		first: "one",
		last:  "one",
		age:   1,
	}
	fmt.Println(x)
	y := &x
	changeMe(y)
	fmt.Println(x)
}

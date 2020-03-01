package main

import "fmt"

type person struct {
	first string
	last  string
}

func (p *person) speak() {
	fmt.Println("hello:", p.first)
}

type human interface {
	speak()
}

func saySomething(h human) {
	h.speak()
}

func main() {
	p1 := person{
		first: "Ian",
		last:  "Flemming",
	}
	// saySomething(p1) -> Wont work!
	saySomething(&p1)
}

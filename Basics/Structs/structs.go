package main

import "fmt"

type person struct {
	first_name string
	last_name  string
}

func main() {
	p1 := person{
		first_name: "Ian",
		last_name:  "Hanes",
	}
	p2 := person{
		first_name: "John",
		last_name:  "Doe",
	}
	fmt.Println(p1.first_name, p2.first_name)
	fmt.Println(p1.last_name, p2.last_name)
}

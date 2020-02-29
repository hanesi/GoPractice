package main

import "fmt"

type person struct {
	first_name string
	last_name  string
}

type secretAgent struct {
	person
	boo bool
}

func main() {
	sa1 := secretAgent{
		person: person{
			first_name: "Ian",
			last_name:  "Hanes",
		},
		boo: true,
	}
	p2 := person{
		first_name: "John",
		last_name:  "Doe",
	}
	fmt.Println(sa1.first_name, sa1.last_name, sa1.boo)
	fmt.Println(p2.first_name, p2.last_name)
}

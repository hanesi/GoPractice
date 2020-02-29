package main

import "fmt"

type person struct {
	first string
	last  string
	likes []string
}

type vehicle struct {
	doors string
	color string
}

type truck struct {
	vehicle
	fourWheel bool
}

type sedan struct {
	vehicle
	luxury bool
}

func main() {
	// # 1
	p1 := person{
		first: "John",
		last:  "Doe1",
		likes: []string{"one, two"},
	}

	p2 := person{
		first: "Jane",
		last:  "Doe2",
		likes: []string{"three, four"},
	}

	for _, v := range p1.likes {
		fmt.Println(v)
	}

	for _, v := range p2.likes {
		fmt.Println(v)
	}

	// #2, store in map
	m := map[string]person{
		p1.last: p1,
		p2.last: p2,
	}

	fmt.Println(m)

	// #3
	t := truck{
		vehicle: vehicle{
			doors: "two",
			color: "green",
		},
		fourWheel: false,
	}

	s := sedan{
		vehicle: vehicle{
			doors: "four",
			color: "green",
		},
		luxury: true,
	}

	fmt.Println(t, t.color)
	fmt.Println(s, s.doors)

	// #4 anonymous struct

	as1 := struct {
		first   string
		friends map[string]int
		likes   []string
	}{
		first: "Rick",
		friends: map[string]int{
			"James": 123,
			"Mark":  234,
			"Dave":  345,
		},
		likes: []string{
			"one thing",
			"two things",
		},
	}

	fmt.Println(as1)
}

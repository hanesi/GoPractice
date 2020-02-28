package main

import "fmt"

func main() {
	x := []int{1, 2, 3, 4}
	fmt.Println(x)

	y := []int{5, 6, 7, 8, 9}
	x = append(x, y...)
	fmt.Println(x)

	x = append(x[:2], x[4:]...)
	fmt.Println(x)

	z := make([]int, 10, 100)

	fmt.Println(z)

	m := map[string]int{
		"Ian":   27,
		"Julia": 23,
	}

	fmt.Println(m)
	fmt.Println(m["Ian"])

	v, ok := m["ian"]
	fmt.Println(v)
	fmt.Println(ok)

	if v, ok := m["Ian"]; ok {
		fmt.Println(v)
	}

	m["Bob"] = 34

	fmt.Println("Added Bob")
	for k, v := range m {
		fmt.Println(k, v)
	}

	delete(m, "Bob")
	fmt.Println("Added Bob")
	for k, v := range m {
		fmt.Println(k, v)
	}

}

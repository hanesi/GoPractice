package main

import "fmt"

func main() {
	// # 1 array
	x := [5]int{1, 2, 3, 4, 5}

	for _, v := range x {
		fmt.Println(v)
	}

	fmt.Printf("%T", x)

	// #2
	y := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}

	for _, v := range y {
		fmt.Println(v)
	}

	fmt.Printf("%T\n", y)

	// #3
	z := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	fmt.Println(z[0:5])
	fmt.Println(z[5:])
	fmt.Println(z[2:7])
	fmt.Println(z[1:6])

	// #4
	a := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	a = append(a, 52)
	fmt.Println(a)
	a = append(a, 53, 54, 55)
	fmt.Println(a)
	b := []int{56, 57, 58, 59, 60}
	a = append(a, b...)
	fmt.Println(a)

	// #5
	num5 := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	num5 = append(num5[0:3], num5[6:]...)
	fmt.Println(num5)

	// #6
	states := make([]string, 50, 50)
	states = []string{"all", "the", "states"}

	for i := 0; i < len(states); i++ {
		fmt.Println(states[i])
	}

	// # 7
	xs1 := []string{"James", "Bond"}
	xs2 := []string{"Value", "Two"}
	fmt.Println(xs1)
	fmt.Println(xs2)

	xs3 := [][]string{xs1, xs2}
	fmt.Println(xs3)

	for _, v := range xs3 {
		fmt.Println(v)
		for _, j := range v {
			fmt.Println(j)
		}
	}

	// #8 - 10
	mp := map[string][]string{
		"bond_james": []string{"one", "two"},
	}
	// #9 is adding records like this
	mp["guy_bad"] = []string{"three", "four"}
	mp["guy_good"] = []string{"five", "six"}

	// #10 delete a record
	delete(mp, "guy_bad")

	for _, v := range mp {
		for j, v2 := range v {
			fmt.Println(j, v2)
		}
	}

	// 10

}

package main

import "fmt"

func main() {
	{
		x := 23
		fmt.Println(x)
	}
	fmt.Println(x) // this error left here intentionally to demonstrate enclosure

}

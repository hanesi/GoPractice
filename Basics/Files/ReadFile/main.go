package main

import (
	"fmt"
	"os"
)

func main() {
	// dat, err := ioutil.ReadFile("/Users/ianhanes/downloads/ESPN_NFL_TeamIDs.xlsx")
	// if err != nil {
	// 	fmt.Println(err)
	// }

	f, err := os.Open("/Users/ianhanes/downloads/ESPN_NFL_TeamIDs.xlsx")
	if err != nil {
		fmt.Println(err)
	}
	b1 := make([]byte, 50)
	n1, err := f.Read(b1)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%d bytes: %s\n", n1, string(b1[:n1]))
}

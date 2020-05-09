package main

import (
	"fmt"
	"os"
)

func main() {
	fp := "/Users/ianhanes/documents/github/gopractice/dataframestuff/nfl_data/NFL_data.csv"
	csv, err := os.Open(fp)
	if err != nil {
		fmt.Println(err)
	}

	df := qframe(csv)
	fmt.Println(df)
}

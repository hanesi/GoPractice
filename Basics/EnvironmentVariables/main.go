package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	// Store the PATH environment variable in a variable
	path, exists := os.LookupEnv("PATH")

	if exists {
		// Print the value of the environment variable
		fmt.Print(path)
	}
	fmt.Println()
	for _, e := range os.Environ() {
		pair := strings.SplitN(e, "=", 2)
		fmt.Println(pair)
	}
}

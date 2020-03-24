package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

func init() {
	// loads values from .env into the system
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}

func main() {
	// Store the PATH environment variable in a variable
	path, exists := os.LookupEnv("PATH")

	if exists {
		// Print the value of the environment variable
		fmt.Print(path)
	}

	fmt.Println()
	fmt.Println(os.LookupEnv("ATHENADB"))
	for _, e := range os.Environ() {
		pair := strings.SplitN(e, "=", 2)
		fmt.Println(pair)
	}
}

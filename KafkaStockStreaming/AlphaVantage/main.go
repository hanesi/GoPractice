package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

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
	AVkey, exists := os.LookupEnv("AVkey")

	if exists {
		// Print the value of the environment variable
		fmt.Print(AVkey)
	}
	AVkey = "OL54MJQGMZ7DX0N0"
	url := (buildQueryURL("GOOG", AVkey))

	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%s", body)
}

func buildQueryURL(s, av string) string {
	return fmt.Sprintf("https://www.alphavantage.co/query?function=GLOBAL_QUOTE&symbol=%s&apikey=%s", s, av)
}

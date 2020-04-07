package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

type StockData struct {
	GlobalQuote TickerData `json:"Global Quote"`
}

type TickerData struct {
	Symbol           string `json:"01. symbol"`
	Open             string `json:"02. open"`
	High             string `json:"03. high"`
	Low              string `json:"04. low"`
	Price            string `json:"05. price"`
	Volume           string `json:"06. volume"`
	LatestTradingDay string `json:"07. latest trading day"`
	PreviousClose    string `json:"08. previous close"`
	Change           string `json:"09. change"`
	ChangePercent    string `json:"10. change percent"`
}

func init() {
	// loads values from .env into the system
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
	return
}

func main() {
	// Store the PATH environment variable in a variable
	AVkey, _ := os.LookupEnv("AVkey")

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
	bodyString := fmt.Sprintf("%s", body)

	var data StockData
	decoder := json.NewDecoder(strings.NewReader(bodyString))
	err = decoder.Decode(&data)
	if err != nil {
		fmt.Println("twas an error")
		// return
	}
	fmt.Println(data)
}

func buildQueryURL(s, av string) string {
	return fmt.Sprintf("https://www.alphavantage.co/query?function=GLOBAL_QUOTE&symbol=%s&apikey=%s", s, av)
}

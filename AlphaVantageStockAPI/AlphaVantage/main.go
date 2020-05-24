package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
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

type StockInfo []struct {
	Ticker         string `json:"ticker"`
	BoughtPrice    string `json:"boughtPrice"`
	NumberOfShares string `json:"numberOfShares"`
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
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("us-east-1")},
	)
	if err != nil {
		panic(err)
	}

	s3Client := s3.New(sess)
	bucket := "ian-test-bucket-go-python"
	key := "StockInfo.json"

	AVkey, _ := os.LookupEnv("AVkey")

	requestInput := &s3.GetObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
	}

	result, err := s3Client.GetObject(requestInput)
	if err != nil {
		fmt.Println(err)
	}
	defer result.Body.Close()
	body1, err := ioutil.ReadAll(result.Body)
	if err != nil {
		fmt.Println(err)
	}
	bodyString1 := fmt.Sprintf("%s", body1)

	var s3data StockInfo
	decoder := json.NewDecoder(strings.NewReader(bodyString1))
	err = decoder.Decode(&s3data)
	if err != nil {
		fmt.Println("twas an error")
	}

	for _, v := range s3data {
		dataGrabber(v.Ticker, v.BoughtPrice, v.NumberOfShares, AVkey)
		// Alpha Vantage limits users to 1 request per 18 seconds
		fmt.Println("Sleeping for 18 seconds")
		time.Sleep(18 * time.Second)
	}

}

func buildQueryURL(s, av string) string {
	return fmt.Sprintf("https://www.alphavantage.co/query?function=GLOBAL_QUOTE&symbol=%s&apikey=%s", s, av)
}

func dataGrabber(ticker, boughtprice, numshares, av string) {
	queryString := buildQueryURL(ticker, av)
	numShares, _ := strconv.ParseFloat(numshares, 32)
	boughtPrice, _ := strconv.ParseFloat(boughtprice, 32)

	resp, err := http.Get(queryString)
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

	price, _ := strconv.ParseFloat(data.GlobalQuote.Price, 32)
	boughtEquity := numShares * boughtPrice
	currentEquity := numShares * price

	if currentEquity/boughtEquity > 1.5 {
		fmt.Println("SELLLLLLLLL")
	} else {
		fmt.Println("Stock", ticker, "ain't there yet")
	}
}

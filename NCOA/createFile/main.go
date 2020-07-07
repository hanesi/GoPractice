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

type Response struct {
	ID          string      `json:"Id"`
	Status      string      `json:"Status"`
	Caption     string      `json:"Caption"`
	Name        string      `json:"Name"`
	CreateDate  string      `json:"CreateDate"`
	RecordCount int         `json:"RecordCount"`
	URL         interface{} `json:"Url"`
}

func init() {
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}
func main() {
	ok := createFile("testFile")
	fmt.Println(ok)
}

func createFile(filename string) string {
	url := fmt.Sprintf("https://app.testing.truencoa.com/api/files/%s/index", filename)
	method := "POST"
	payload := strings.NewReader("caption=This%20should%20be%20the%20file%20caption")

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
	}

	login, _ := os.LookupEnv("NCOALogin")
	password, _ := os.LookupEnv("NCOAPassword")
	req.Header.Add("user_name", login)
	req.Header.Add("password", password)

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
	}

	var responseObject Response
	json.Unmarshal(body, &responseObject)
	return responseObject.ID
}

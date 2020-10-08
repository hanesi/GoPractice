package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {
	fileID := "f07a566d-5ffd-4bfc-a59c-7d5af7af8939"
	url := fmt.Sprintf("https://app.truencoa.com/api/files/%s/files", fileID)
	method := "GET"

	payload := strings.NewReader("")

	fmt.Println("defining client and request")
	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
	}

	req.Header.Add("user_name", "")
	req.Header.Add("password", "")
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	fmt.Println("executing request")
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(body))
}

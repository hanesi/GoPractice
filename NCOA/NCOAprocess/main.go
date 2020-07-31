package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {

	url := "https://app.testing.truencoa.com/api/files/989c1c39-4c4e-4789-8dde-ef9a04e47c88/index?status=submit"
	method := "PATCH"

	payload := strings.NewReader("")

	fmt.Println("defining client and request")
	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("adding headers")
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

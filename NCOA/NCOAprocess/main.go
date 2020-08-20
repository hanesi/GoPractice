package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {

	url := "https://app.testing.truencoa.com/api/files/9e050ede-a4de-412c-8298-540ca17380d8/index?status=submit"
	method := "PATCH"

	payload := strings.NewReader("")

	fmt.Println("defining client and request")
	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("adding headers")
	req.Header.Add("user_name", "ian@sharelocalmedia.com")
	req.Header.Add("password", "cokkyg-juczuF-8sasqi")
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

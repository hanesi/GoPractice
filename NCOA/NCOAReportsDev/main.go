package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/fatih/structs"
)

type Response struct {
	ID          string       `json:"Id"`
	Status      string       `json:"Status"`
	Name        string       `json:"Name"`
	CreateDate  string       `json:"CreateDate"`
	RecordCount int          `json:"RecordCount"`
	Items       []ItemReport `json:"Items"`
}

type ItemReport struct {
	Name         string `json:"Name"`
	Item         string `json:"Item"`
	Value        string `json:"Value"`
	OrderIndex   string `json:"OrderIndex"`
	IsPercentage bool   `json:"IsPercentage"`
	IsError      bool   `json:"IsError"`
}

func main() {
	reports := getReports("64ba1e6f-6f5c-4602-8358-671c9b790a85")
	// fmt.Println(reports[0].Items)
	// fmt.Println(reports[1].Items)
	s := make([]string, 0)

	for _, v := range reports[1].Items {
		for _, x := range structs.Values(v) {
			s = append(s, x.(string))
		}
	}

	fmt.Println(s)
	// w := csv.NewWriter(os.Stdout)
	// headers := []string{
	// 	"Name",
	// 	"Item",
	// 	"Value",
	// 	"OrderIndex",
	// 	"IsPercentage",
	// 	"IsError",
	// }
	// if err := w.Write(headers); err != nil {
	// 	//write failed do something
	// }
	// for _, v := range reports {
	// 	values := v.ToSlice()
	// 	if err := w.Write(values); err != nil {
	// 		//write failed do something
	// 	}
	// }
}

func getReports(fileID string) []Response {
	urlC := fmt.Sprintf("https://app.testing.truencoa.com/api/files/%s/reports/cass", fileID)
	urlN := fmt.Sprintf("https://app.testing.truencoa.com/api/files/%s/reports/ncoa", fileID)
	retList := []Response{}

	retList = append(retList, processRequest(urlC))
	retList = append(retList, processRequest(urlN))

	return retList

}

func processRequest(url string) Response {
	method := "GET"
	payload := strings.NewReader("")

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
	}
	req.Header.Add("user_name", "")
	req.Header.Add("password", "")
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

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

	return responseObject
}

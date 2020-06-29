package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {

	url := "https://app.testing.truencoa.com/api/files/d2f4c0b9-1627-472d-8d75-56712a8db242/records"
	method := "POST"

	payload := strings.NewReader(`[
    {
          "individual_id":"1",
          "individual_first_name":"Christopher",
          "individual_last_name":"Klimko",
          "address_line_1":"3000 Westminster Ave",
          "address_line_2":"",
          "address_city_name":"Dallas",
          "address_state_code":"TX",
          "address_postal_code":"75205",
          "address_country_code":""
       },
       {
          "individual_id":"2",
          "individual_first_name":"Edward",
          "individual_last_name":"Kay",
          "address_line_1":"9696 New Buffalo Rd",
          "address_line_2":"",
          "address_city_name":"Canfield",
          "address_state_code":"OH",
          "address_postal_code":"44406",
          "address_country_code":""
       },
       {
          "individual_id":"3",
          "individual_first_name":"John",
          "individual_last_name":"Dance",
          "address_line_1":"515 Brookpark Dr",
          "address_line_2":"",
          "address_city_name":"Canfield",
          "address_state_code":"OH",
          "address_postal_code":"44406",
          "address_country_code":""
       },
       {
          "individual_id":"4",
          "individual_first_name":"Patricia",
          "individual_last_name":"Guarnieri",
          "address_line_1":"114 Mill Creek Dr",
          "address_line_2":"",
          "address_city_name":"Youngstown",
          "address_state_code":"OH",
          "address_postal_code":"44512",
          "address_country_code":""
       },
       {
          "individual_id":"5",
          "individual_first_name":"Chris",
          "individual_last_name":"Meta",
          "address_line_1":"4120 Kirk Rd",
          "address_line_2":"",
          "address_city_name":"Columbiana",
          "address_state_code":"OH",
          "address_postal_code":"44408",
          "address_country_code":""
       },
       {
          "individual_id":"6",
          "individual_first_name":"Jason",
          "individual_last_name":"Peretti",
          "address_line_1":"28 Bob White Ct",
          "address_line_2":"",
          "address_city_name":"Youngstown",
          "address_state_code":"OH",
          "address_postal_code":"44511",
          "address_country_code":""
       },
       {
          "individual_id":"7",
          "individual_first_name":"Raymond",
          "individual_last_name":"Manofsky",
          "address_line_1":"1155 Paige Ave NE",
          "address_line_2":"",
          "address_city_name":"Warren",
          "address_state_code":"OH",
          "address_postal_code":"44483",
          "address_country_code":""
       },
       {
          "individual_id":"8",
          "individual_first_name":"Joshua",
          "individual_last_name":"Kollat",
          "address_line_1":"120 W Lytle Ave",
          "address_line_2":"",
          "address_city_name":"State College",
          "address_state_code":"PA",
          "address_postal_code":"16801",
          "address_country_code":""
       },
       {
          "individual_id":"9",
          "individual_first_name":"JoAnn",
          "individual_last_name":"Stock",
          "address_line_1":"190 Southwoods Ave",
          "address_line_2":"",
          "address_city_name":"Youngstown",
          "address_state_code":"OH",
          "address_postal_code":"44512",
          "address_country_code":""
       },
       {
          "individual_id":"10",
          "individual_first_name":"Judith",
          "individual_last_name":"Demay",
          "address_line_1":"6210 Warren Sharon Rd",
          "address_line_2":"",
          "address_city_name":"Brookfield",
          "address_state_code":"OH",
          "address_postal_code":"44403",
          "address_country_code":""
       },
       {
          "individual_id":"11",
          "individual_first_name":"Timothy",
          "individual_last_name":"Taggart",
          "address_line_1":"1714 Hamilton Pl",
          "address_line_2":"",
          "address_city_name":"Steubenville",
          "address_state_code":"OH",
          "address_postal_code":"43952",
          "address_country_code":""
       },
       {
          "individual_id":"12",
          "individual_first_name":"Matthew",
          "individual_last_name":"Stiffler",
          "address_line_1":"1751 Damos Way",
          "address_line_2":"",
          "address_city_name":"Marysville",
          "address_state_code":"OH",
          "address_postal_code":"43040",
          "address_country_code":""
       },
       {
          "individual_id":"13",
          "individual_first_name":"Nicole",
          "individual_last_name":"Hively",
          "address_line_1":"496 S Briarcliff Dr",
          "address_line_2":"",
          "address_city_name":"Canfield",
          "address_state_code":"OH",
          "address_postal_code":"44406",
          "address_country_code":""
       },
       {
          "individual_id":"14",
          "individual_first_name":"Mary",
          "individual_last_name":"Morrone",
          "address_line_1":"2679 S Hubbard Rd",
          "address_line_2":"",
          "address_city_name":"Lowellville",
          "address_state_code":"OH",
          "address_postal_code":"44436",
          "address_country_code":""
       },
       {
          "individual_id":"15",
          "individual_first_name":"Edward",
          "individual_last_name":"Hartwig",
          "address_line_1":"6662 NW 98th Dr.",
          "address_line_2":"",
          "address_city_name":"Parkland",
          "address_state_code":"FL",
          "address_postal_code":"33076",
          "address_country_code":""
       },
       {
          "individual_id":"16",
          "individual_first_name":"Lori",
          "individual_last_name":"Marshall",
          "address_line_1":"250 Parkview Dr",
          "address_line_2":"",
          "address_city_name":"Hubbard",
          "address_state_code":"OH",
          "address_postal_code":"44425",
          "address_country_code":""
       },
       {
          "individual_id":"17",
          "individual_first_name":"Michael",
          "individual_last_name":"Straniak",
          "address_line_1":"2818 Citadel Dr NE",
          "address_line_2":"",
          "address_city_name":"Warren",
          "address_state_code":"OH",
          "address_postal_code":"44483",
          "address_country_code":""
       },
       {
          "individual_id":"18",
          "individual_first_name":"Russell",
          "individual_last_name":"Hoover",
          "address_line_1":"1444 Robbins Ave",
          "address_line_2":"",
          "address_city_name":"Niles",
          "address_state_code":"OH",
          "address_postal_code":"44446",
          "address_country_code":""
       },
       {
          "individual_id":"19",
          "individual_first_name":"Eric",
          "individual_last_name":"Peterson",
          "address_line_1":"3418 Sandalwood Ln",
          "address_line_2":"",
          "address_city_name":"Youngstown",
          "address_state_code":"OH",
          "address_postal_code":"44511",
          "address_country_code":""
       },
       {
          "individual_id":"20",
          "individual_first_name":"Karen",
          "individual_last_name":"Campf",
          "address_line_1":"645 High St",
          "address_line_2":"",
          "address_city_name":"Washingtonville",
          "address_state_code":"OH",
          "address_postal_code":"44490",
          "address_country_code":""
       },
       {
          "individual_id":"21",
          "individual_first_name":"Connie",
          "individual_last_name":"Gorby",
          "address_line_1":"13324 McCormick Run Rd",
          "address_line_2":"",
          "address_city_name":"Lisbon",
          "address_state_code":"OH",
          "address_postal_code":"44432",
          "address_country_code":""
       },
       {
          "individual_id":"22",
          "individual_first_name":"Kenneth",
          "individual_last_name":"Johnson",
          "address_line_1":"5465 Chapel Rd. Frnt.",
          "address_line_2":"",
          "address_city_name":"Madison",
          "address_state_code":"OH",
          "address_postal_code":"44057",
          "address_country_code":""
       },
       {
          "individual_id":"23",
          "individual_first_name":"Michael",
          "individual_last_name":"Marshall",
          "address_line_1":"401 Moores River Dr",
          "address_line_2":"",
          "address_city_name":"Lansing",
          "address_state_code":"MI",
          "address_postal_code":"48910",
          "address_country_code":""
       },
       {
          "individual_id":"24",
          "individual_first_name":"Dori",
          "individual_last_name":"MacMillan",
          "address_line_1":"1237 Four Winds Ct",
          "address_line_2":"",
          "address_city_name":"Niles",
          "address_state_code":"OH",
          "address_postal_code":"44446",
          "address_country_code":""
       },
       {
          "individual_id":"25",
          "individual_first_name":"Judith",
          "individual_last_name":"Young",
          "address_line_1":"6373 Tara Dr",
          "address_line_2":"",
          "address_city_name":"Poland",
          "address_state_code":"OH",
          "address_postal_code":"44514",
          "address_country_code":""
       },
       {
          "individual_id":"26",
          "individual_first_name":"Kenneth",
          "individual_last_name":"Larson",
          "address_line_1":"2830 Spring Meadow Cir",
          "address_line_2":"",
          "address_city_name":"Youngstown",
          "address_state_code":"OH",
          "address_postal_code":"44515",
          "address_country_code":""
       },
       {
          "individual_id":"27",
          "individual_first_name":"John",
          "individual_last_name":"Bukovinsky",
          "address_line_1":"8031 Forest Lake Dr",
          "address_line_2":"",
          "address_city_name":"Youngstown",
          "address_state_code":"OH",
          "address_postal_code":"44512",
          "address_country_code":""
       },
       {
          "individual_id":"28",
          "individual_first_name":"Frank",
          "individual_last_name":"Saccomen",
          "address_line_1":"2314 Bell Wick Rd",
          "address_line_2":"",
          "address_city_name":"Hubbard",
          "address_state_code":"OH",
          "address_postal_code":"44425",
          "address_country_code":""
       },
       {
          "individual_id":"29",
          "individual_first_name":"Jody",
          "individual_last_name":"Cutrer",
          "address_line_1":"150 Talsman Dr Unit 1",
          "address_line_2":"",
          "address_city_name":"Canfield",
          "address_state_code":"OH",
          "address_postal_code":"44406",
          "address_country_code":""
       },
       {
          "individual_id":"30",
          "individual_first_name":"Mary",
          "individual_last_name":"Eicher",
          "address_line_1":"859 Pasadena Ave",
          "address_line_2":"",
          "address_city_name":"Youngstown",
          "address_state_code":"OH",
          "address_postal_code":"44502",
          "address_country_code":""
       }]`)
	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
	}
	req.Header.Add("user_name", "{{api_user_name}}")
	req.Header.Add("password", "{{api_password}}")
	req.Header.Add("Content-Type", "application/json")

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

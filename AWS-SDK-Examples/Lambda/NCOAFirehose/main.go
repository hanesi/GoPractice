package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/firehose"
)

type EventResponse struct {
	UTC time.Time `json:"utc"`
}

type Response struct {
	ID          string      `json:"Id"`
	Status      string      `json:"Status"`
	Caption     string      `json:"Caption"`
	Name        string      `json:"Name"`
	CreateDate  string      `json:"CreateDate"`
	RecordCount int         `json:"RecordCount"`
	URL         interface{} `json:"Url"`
}

type ProcessedResponse struct {
	ID          string             `json:"Id"`
	Status      string             `json:"Status"`
	Caption     string             `json:"Caption"`
	Name        string             `json:"Name"`
	CreateDate  string             `json:"CreateDate"`
	RecordCount int                `json:"RecordCount"`
	URL         interface{}        `json:"Url"`
	Records     []ProcessedRecords `json:"Records"`
}

type ProcessedRecords struct {
	SLMIndividualID                  string      `json:"individual_id"`
	HouseholdPosition                int         `json:"Household Position"`
	IndividualFirstName              string      `json:"individual_first_name"`
	NameID                           int         `json:"Name ID"`
	IndividualLastName               string      `json:"individual_last_name"`
	IndividualRecordID               int         `json:"Individual Record ID"`
	SLMAddressLine1                  string      `json:"address_line_1"`
	FirstName                        string      `json:"First Name"`
	SLMAddressLine2                  string      `json:"address_line_2"`
	LastName                         string      `json:"Last Name"`
	AddressCityName                  string      `json:"address_city_name"`
	CompanyName                      interface{} `json:"Company Name"`
	AddressStateCode                 string      `json:"address_state_code"`
	StreetNumber                     string      `json:"Street Number"`
	AddressPostalCode                string      `json:"address_postal_code"`
	StreetPreDirection               interface{} `json:"Street Pre Direction"`
	AddressCountryCode               string      `json:"address_country_code"`
	StreetName                       string      `json:"Street Name"`
	StreetPostDirection              interface{} `json:"Street Post Direction"`
	StreetSuffix                     string      `json:"Street Suffix"`
	UnitType                         interface{} `json:"Unit Type"`
	UnitNumber                       interface{} `json:"Unit Number"`
	BoxNumber                        interface{} `json:"Box Number"`
	CityName                         string      `json:"City Name"`
	StateCode                        string      `json:"State Code"`
	PostalCode                       string      `json:"Postal Code"`
	PostalCodeExtension              string      `json:"Postal Code Extension"`
	CarrierRoute                     string      `json:"Carrier Route"`
	AddressStatus                    string      `json:"Address Status"`
	ErrorNumber                      string      `json:"Error Number"`
	AddressType                      string      `json:"Address Type"`
	DeliveryPoint                    string      `json:"Delivery Point"`
	CheckDigit                       string      `json:"Check Digit"`
	DeliveryPointVerification        string      `json:"Delivery Point Verification"`
	DeliveryPointVerificationNotes   string      `json:"Delivery Point Verification Notes"`
	Vacant                           string      `json:"Vacant"`
	CongressionalDistrictCode        string      `json:"Congressional District Code"`
	AreaCode                         string      `json:"Area Code"`
	Latitude                         string      `json:"Latitude"`
	Longitude                        string      `json:"Longitude"`
	TimeZone                         string      `json:"Time Zone"`
	CountyName                       string      `json:"County Name"`
	CountyFIPS                       string      `json:"County FIPS"`
	StateFIPS                        string      `json:"State FIPS"`
	Barcode                          string      `json:"Barcode"`
	LocatableAddressConversionSystem interface{} `json:"Locatable Address Conversion System"`
	LineOfTravel                     string      `json:"Line of Travel"`
	AscendingDescending              string      `json:"Ascending/Descending"`
	MoveApplied                      interface{} `json:"Move Applied"`
	MoveType                         interface{} `json:"Move Type"`
	MoveDate                         interface{} `json:"Move Date"`
	MoveDistance                     interface{} `json:"Move Distance"`
	MatchFlag                        interface{} `json:"Match Flag"`
	NXI                              interface{} `json:"NXI"`
	ANK                              interface{} `json:"ANK"`
	ResidentialDeliveryIndicator     string      `json:"Residential Delivery Indicator"`
	RecordType                       string      `json:"Record Type"`
	CountryCode                      string      `json:"Country Code"`
	AddressLine1                     string      `json:"Address Line 1"`
	AddressLine2                     interface{} `json:"Address Line 2"`
	AddressID                        int         `json:"Address Id"`
	HouseholdID                      int         `json:"Household Id"`
	IndividualID                     int         `json:"Individual Id"`
}

type ProcessedRecordsReturn struct {
	SLMIndividualID                  string      `json:"SLMIndividualID"`
	HouseholdPosition                int         `json:"HouseholdPosition"`
	IndividualFirstName              string      `json:"IndividualFirstName"`
	NameID                           int         `json:"NameID"`
	IndividualLastName               string      `json:"IndividualLastName"`
	IndividualRecordID               int         `json:"IndividualRecordID"`
	SLMAddressLine1                  string      `json:"SLMAddressLine1"`
	FirstName                        string      `json:"FirstName"`
	SLMAddressLine2                  string      `json:"SLMAddressLine2"`
	LastName                         string      `json:"LastName"`
	AddressCityName                  string      `json:"AddressCityName"`
	CompanyName                      interface{} `json:"CompanyName"`
	AddressStateCode                 string      `json:"AddressStateCode"`
	StreetNumber                     string      `json:"StreetNumber"`
	AddressPostalCode                string      `json:"AddressPostalCode"`
	StreetPreDirection               interface{} `json:"StreetPreDirection"`
	AddressCountryCode               string      `json:"AddressCountryCode"`
	StreetName                       string      `json:"StreetName"`
	StreetPostDirection              interface{} `json:"StreetPostDirection"`
	StreetSuffix                     string      `json:"StreetSuffix"`
	UnitType                         interface{} `json:"UnitType"`
	UnitNumber                       interface{} `json:"UnitNumber"`
	BoxNumber                        interface{} `json:"BoxNumber"`
	CityName                         string      `json:"CityName"`
	StateCode                        string      `json:"StateCode"`
	PostalCode                       string      `json:"PostalCode"`
	PostalCodeExtension              string      `json:"PostalCodeExtension"`
	CarrierRoute                     string      `json:"CarrierRoute"`
	AddressStatus                    string      `json:"AddressStatus"`
	ErrorNumber                      string      `json:"ErrorNumber"`
	AddressType                      string      `json:"AddressType"`
	DeliveryPoint                    string      `json:"DeliveryPoint"`
	CheckDigit                       string      `json:"CheckDigit"`
	DeliveryPointVerification        string      `json:"DeliveryPointVerification"`
	DeliveryPointVerificationNotes   string      `json:"DeliveryPointVerificationNotes"`
	Vacant                           string      `json:"Vacant"`
	CongressionalDistrictCode        string      `json:"CongressionalDistrictCode"`
	AreaCode                         string      `json:"AreaCode"`
	Latitude                         string      `json:"Latitude"`
	Longitude                        string      `json:"Longitude"`
	TimeZone                         string      `json:"TimeZone"`
	CountyName                       string      `json:"CountyName"`
	CountyFIPS                       string      `json:"CountyFIPS"`
	StateFIPS                        string      `json:"StateFIPS"`
	Barcode                          string      `json:"Barcode"`
	LocatableAddressConversionSystem interface{} `json:"LocatableAddressConversionSystem"`
	LineOfTravel                     string      `json:"LineOfTravel"`
	AscendingDescending              string      `json:"AscendingDescending"`
	MoveApplied                      interface{} `json:"MoveApplied"`
	MoveType                         interface{} `json:"MoveType"`
	MoveDate                         interface{} `json:"MoveDate"`
	MoveDistance                     interface{} `json:"MoveDistance"`
	MatchFlag                        interface{} `json:"MatchFlag"`
	NXI                              interface{} `json:"NXI"`
	ANK                              interface{} `json:"ANK"`
	ResidentialDeliveryIndicator     string      `json:"ResidentialDeliveryIndicator"`
	RecordType                       string      `json:"RecordType"`
	CountryCode                      string      `json:"CountryCode"`
	AddressLine1                     string      `json:"AddressLine1"`
	AddressLine2                     interface{} `json:"AddressLine2"`
	AddressID                        int         `json:"AddressId"`
	HouseholdID                      int         `json:"HouseholdId"`
	IndividualID                     int         `json:"IndividualId"`
}

func handleRequest(ctx context.Context, request events.SQSEvent) (events.APIGatewayProxyResponse, error) {
	now := time.Now()
	resp := &EventResponse{
		UTC: now.UTC(),
	}
	bodyLambda, err := json.Marshal(resp)
	if err != nil {
		return events.APIGatewayProxyResponse{}, err
	}

	fmt.Println("Recieved event", request)
	for i := range request.Records {
		msgBody := request.Records[i].Body
		fmt.Println("Processing file", msgBody)
		exportid := strings.Split(msgBody, "___")[0]
		start, _ := strconv.Atoi(strings.Split(msgBody, "___")[1])
		end, _ := strconv.Atoi(strings.Split(msgBody, "___")[2])

		// url := fmt.Sprintf("https://app.testing.truencoa.com/api/files/%s/index?status=export&export_template=export_default", id)
		// method := "PATCH"
		//
		// payload := strings.NewReader("")
		//
		// client := &http.Client{}
		// req, err := http.NewRequest(method, url, payload)
		//
		// if err != nil {
		// 	fmt.Println(err)
		// }
		// login, _ := os.LookupEnv("NCOALogin")
		// password, _ := os.LookupEnv("NCOAPassword")
		// req.Header.Add("user_name", login)
		// req.Header.Add("password", password)
		// req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
		//
		// fmt.Println("Starting export...")
		// res, err := client.Do(req)
		// if err != nil {
		// 	fmt.Println(err)
		// }
		//
		// defer res.Body.Close()
		// body, err := ioutil.ReadAll(res.Body)
		// if err != nil {
		// 	fmt.Println(err)
		// }
		//
		// fmt.Println("Export Body Response: ", string(body))
		//
		// var responseObject Response
		// json.Unmarshal(body, &responseObject)
		// fmt.Println("Response Body:")
		// fmt.Println(responseObject)

		// exportid := responseObject.ID
		fmt.Println("Export ID:", exportid)
		// exportid := "6165c142-2900-4095-a62a-d9fcaca76c9c"

		fmt.Println("Starting Download...")
		recordList := download(start, end, exportid)

		fmt.Println("Submitting Records to Firehose...")
		submitToFirehose(recordList)
	}
	return events.APIGatewayProxyResponse{Body: string(bodyLambda), StatusCode: 200}, nil
}

func submitToFirehose(records []ProcessedRecords) {
	streamName := "slm-ncoaprocessed-stream"

	sess := session.Must(session.NewSession())
	// Create a Firehose client with additional configuration
	firehoseService := firehose.New(sess, aws.NewConfig().WithRegion("us-east-1"))

	recordsBatchInput := &firehose.PutRecordBatchInput{}
	recordsBatchInput = recordsBatchInput.SetDeliveryStreamName(streamName)

	recordsInput := []*firehose.Record{}
	for i := 0; i < len(records); i++ {
		if len(recordsInput) == 500 {
			recordsBatchInput = recordsBatchInput.SetRecords(recordsInput)
			resp, err := firehoseService.PutRecordBatch(recordsBatchInput)
			num_failures := *resp.FailedPutCount
			if err != nil {
				fmt.Printf("PutRecordBatch err: %v\n", err)
			} else {
				fmt.Printf("FailedPuts: %v\n", num_failures)
				if num_failures > 0 {
					rec_index := 0
					for _, v := range resp.RequestResponses {
						if v.ErrorCode != nil {
							recInput := &firehose.PutRecordInput{}
							recInput = recInput.SetDeliveryStreamName(streamName)
							recInput = recInput.SetRecord(recordsInput[rec_index])
							_, _ = firehoseService.PutRecord(recInput)
						}
						rec_index++
					}
				}
			}
			recordsInput = []*firehose.Record{}
		}

		out := ProcessedRecordsReturn(records[i])
		b, err := json.Marshal(out)

		if err != nil {
			log.Printf("Error: %v", err)
		}

		record := &firehose.Record{Data: b}
		recordsInput = append(recordsInput, record)
	}

	if len(recordsInput) > 0 {
		recordsBatchInput = recordsBatchInput.SetRecords(recordsInput)
		resp, err := firehoseService.PutRecordBatch(recordsBatchInput)
		num_failures := *resp.FailedPutCount
		if err != nil {
			fmt.Printf("PutRecordBatch err: %v\n", err)
		} else {
			fmt.Printf("FailedPuts: %v\n", num_failures)
			if num_failures > 0 {
				rec_index := 0
				for _, v := range resp.RequestResponses {
					if v.ErrorCode != nil {
						recInput := &firehose.PutRecordInput{}
						recInput = recInput.SetDeliveryStreamName(streamName)
						recInput = recInput.SetRecord(recordsInput[rec_index])
						_, _ = firehoseService.PutRecord(recInput)
					}
					rec_index++
				}
			}
		}
	}
}

func download(start, max int, id string) []ProcessedRecords {
	var interval int
	var retList []ProcessedRecords
	for interval < max {
		interval = start + 999
		if interval > max {
			interval = max
		}
		url := fmt.Sprintf("https://app.truencoa.com/api/files/%s/records?start=%d&end=%d", id, start, interval)
		method := "GET"
		fmt.Println(url)

		client := &http.Client{}
		req, err := http.NewRequest(method, url, nil)

		if err != nil {
			fmt.Println(err)
		}
		login, _ := os.LookupEnv("NCOALogin")
		password, _ := os.LookupEnv("NCOAPassword")
		req.Header.Add("user_name", login)
		req.Header.Add("password", password)
		req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

		fmt.Println("Executing Download Request...")
		res, err := client.Do(req)
		if err != nil {
			fmt.Println(err)
		}
		defer res.Body.Close()
		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			fmt.Println(err)
		}

		var responseObject ProcessedResponse
		json.Unmarshal(body, &responseObject)

		retList = append(retList, responseObject.Records...)

		start += 1000
	}
	return retList
}

func main() {
	lambda.Start(handleRequest)
}

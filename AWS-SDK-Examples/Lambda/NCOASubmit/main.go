package main

import (
	"context"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/sqs"
	"github.com/google/uuid"
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

func handleRequest(ctx context.Context, request events.SQSEvent) (events.APIGatewayProxyResponse, error) {
	now := time.Now()
	resp := &EventResponse{
		UTC: now.UTC(),
	}
	body, err := json.Marshal(resp)
	if err != nil {
		return events.APIGatewayProxyResponse{}, err
	}

	fmt.Println("Recieved event", request)

	for i := range request.Records {
		msgBody := request.Records[i].Body
		fmt.Println("Processing file", msgBody)
		fmt.Printf("%T\n", msgBody)
		jsonString := strings.ReplaceAll(msgBody, "'", "\"")

		var expectedStringArray []string
		json.Unmarshal([]byte(jsonString), &expectedStringArray)

		var transformedRecs []map[string]string
		var fileName string
		var slmIDs []string
		for _, v := range expectedStringArray {
			bucket := strings.Split(v, "___")[0]
			key := strings.Split(v, "___")[1]
			fileName = strings.Split(v, "___")[2]
			fileType := strings.Split(v, "___")[3]
			slm_file_id := strings.Split(v, "___")[4]

			slmIDs = append(slmIDs, slm_file_id)
			recs := getObjectReturnMaps(key, bucket)
			transformedRecs = append(transformedRecs, transformRecordsForProcessing(recs, fileType, slm_file_id)...)
		}

		fmt.Println(transformedRecs[0])

		fileID := createFile(fileName)
		fmt.Println("Starting Record Submission")
		submitRecords(transformedRecs, fileID)

		fmt.Println("Starting Record Validation")
		startFileValidation(fileID)

		bodySl := append([]string{fileID}, slmIDs...)
		bodyMsg := strings.Join(bodySl, "___")
		fmt.Println("Sending SQS Message")
		sendSQSMessage(bodyMsg)
	}
	return events.APIGatewayProxyResponse{Body: string(body), StatusCode: 200}, nil
}

func sendSQSMessage(body string) {
	queueName := "https://sqs.us-east-1.amazonaws.com/363807257486/NCOAValidationPoll.fifo"
	id := uuid.New().String()
	sess := session.Must(session.NewSession())
	// Create a Firehose client with additional configuration
	queueService := sqs.New(sess, aws.NewConfig().WithRegion("us-east-1"))
	_, err := queueService.SendMessage(&sqs.SendMessageInput{
		MessageAttributes: map[string]*sqs.MessageAttributeValue{
			"MsgBody": &sqs.MessageAttributeValue{
				DataType:    aws.String("String"),
				StringValue: aws.String(body),
			},
		},
		MessageBody:    aws.String(body),
		QueueUrl:       &queueName,
		MessageGroupId: &id,
	})

	if err != nil {
		fmt.Println(err)
	}
}

func startFileValidation(fileID string) {
	url := fmt.Sprintf("https://app.testing.truencoa.com/api/files/%s/index?status=submit", fileID)
	method := "PATCH"

	payload := strings.NewReader("")

	fmt.Println("defining client and request")
	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
	}

	login := os.Getenv("NCOALogin")
	password := os.Getenv("NCOAPassword")
	req.Header.Add("user_name", login)
	req.Header.Add("password", password)
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

func submitRecords(records []map[string]string, fileID string) {
	ch := make(chan []map[string]string)
	client := &http.Client{}
	method := "POST"
	url := fmt.Sprintf("https://app.testing.truencoa.com/api/files/%s/records", fileID)

	var wg sync.WaitGroup
	n := 4
	wg.Add(n)
	for i := 0; i < n; i++ {
		go func() {
			defer wg.Done()
			for b := range ch {
				str, _ := json.Marshal(b)

				payload := strings.NewReader(string(str))
				req, err := http.NewRequest(method, url, payload)
				if err != nil {
					fmt.Println(err)
					continue
				}

				login, _ := os.LookupEnv("NCOALogin")
				password, _ := os.LookupEnv("NCOAPassword")
				req.Header.Add("user_name", login)
				req.Header.Add("password", password)
				req.Header.Add("Content-Type", "application/json")

				res, err := client.Do(req)
				fmt.Println(res.Status)
				if err != nil {
					fmt.Println(err)
					continue
				}
				defer res.Body.Close()
				body, err := ioutil.ReadAll(res.Body)
				if err != nil {
					fmt.Println(err)
				}
				var responseObject Response
				json.Unmarshal(body, &responseObject)
				fmt.Println(responseObject)
			}
		}()
	}

	bLen := 10000
	batch := []map[string]string{}
	firstBatch := true
	for _, v := range records {
		batch = append(batch, v)
		if firstBatch && len(batch) == bLen {
			fmt.Println("Sending First Batch!")
			str, _ := json.Marshal(batch)

			payload := strings.NewReader(string(str))
			req, err := http.NewRequest(method, url, payload)
			if err != nil {
				fmt.Println(err)
				continue
			}

			login := os.Getenv("NCOALogin")
			password := os.Getenv("NCOAPassword")
			req.Header.Add("user_name", login)
			req.Header.Add("password", password)
			req.Header.Add("Content-Type", "application/json")

			res, err := client.Do(req)
			if err != nil {
				fmt.Println(err)
				continue
			}
			defer res.Body.Close()
			body, err := ioutil.ReadAll(res.Body)
			if err != nil {
				fmt.Println(err)
			}

			var responseObject Response
			json.Unmarshal(body, &responseObject)
			fmt.Println(responseObject)

			firstBatch = false
			batch = []map[string]string{}
		} else if len(batch) == bLen {
			ch <- batch
			batch = []map[string]string{}
		}
	}
	if len(batch) > 0 {
		ch <- batch
	}
	close(ch)

	wg.Wait()
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
	fmt.Println("File Created!")
	return responseObject.ID
}

func transformRecordsForProcessing(records []map[string]string, filetype, slm_id string) []map[string]string {
	transformedRecords := []map[string]string{}
	if filetype == "custom" {
		for _, v := range records {
			tempDict := make(map[string]string)

			tempDict["individual_id"] = v["individual_id"]
			var fn, ln string
			if val, ok := v["full_name"]; ok && len(val) > 0 {
				tempDict["individual_full_name"] = val
			} else {
				if val, ok := v["first_name"]; ok {
					fn = val
				} else {
					fn = ""
				}
				if val, ok := v["last_name"]; ok {
					ln = val
				} else {
					ln = ""
				}
				tempDict["individual_full_name"] = fmt.Sprintf("%s %s", fn, ln)
			}

			if val, ok := v["primaryAddress"]; ok {
				tempDict["address_line_1"] = val
			} else {
				tempDict["address_line_1"] = v["address1"]
			}
			if val, ok := v["address2"]; ok {
				tempDict["address_line_2"] = val
			} else if val, ok := v["secondaryAddress"]; ok {
				tempDict["address_line_2"] = val
			}
			tempDict["address_city_name"] = v["city"]
			tempDict["address_state_code"] = v["state"]
			tempDict["address_postal_code"] = v["zipcode"]
			tempDict["address_country_code"] = ""
			tempDict["slm_file_id"] = slm_id

			transformedRecords = append(transformedRecords, tempDict)
		}
		fmt.Println("Records Transformed!")
		return transformedRecords
	}
	for _, v := range records {
		tempDict := make(map[string]string)

		tempDict["individual_id"] = v["individual_id"]
		// tempDict["individual_first_name"] = v["firstName"]
		// tempDict["individual_last_name"] = v["lastName"]
		tempDict["individual_full_name"] = fmt.Sprintf("%s %s", v["firstName"], v["lastName"])
		tempDict["address_line_1"] = v["primaryAddress"]
		tempDict["address_line_2"] = ""
		tempDict["address_city_name"] = v["city"]
		tempDict["address_state_code"] = v["state"]
		tempDict["address_postal_code"] = v["zipcode"]
		tempDict["address_country_code"] = ""
		tempDict["slm_file_id"] = slm_id

		transformedRecords = append(transformedRecords, tempDict)
	}
	fmt.Println("Records Transformed!")
	return transformedRecords
}

func getObjectReturnMaps(key, bucket string) []map[string]string {
	// Initialize an AWS session with configured credentials
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("us-east-1")},
	)
	if err != nil {
		panic(err)
	}

	// Define client
	s3Client := s3.New(sess)

	// Using the key, get the object from the bucket
	obj, err := s3Client.GetObject(&s3.GetObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
	})
	if err != nil {
		fmt.Println(err)
	}

	// Read the CSV body, pass into the csvToSliceOfMaps function
	reader := csv.NewReader(obj.Body)
	record, err := reader.ReadAll()
	if err != nil {
		fmt.Println("Error", err)
	}
	fmt.Println("Records Parsed!")
	return csvToSliceOfMaps(record)
}

func csvToSliceOfMaps(records [][]string) (returnMap []map[string]string) {
	// Create a slice from the first row of the input for the headers
	headers := []string{}
	for _, v := range records[0] {
		headers = append(headers, v)
	}
	// Iterate through the remaining records, creating a map with keys
	// from the headers Slice and values from the rest of the slices in
	// the input record set
	for _, record := range records[1:] {
		line := map[string]string{}
		for i := 0; i < len(record); i++ {
			line[headers[i]] = record[i]
		}
		returnMap = append(returnMap, line)
	}
	return returnMap
}

func main() {
	lambda.Start(handleRequest)
}

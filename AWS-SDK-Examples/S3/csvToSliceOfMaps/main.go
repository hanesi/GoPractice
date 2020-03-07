package main

import (
	"encoding/csv"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/joho/godotenv"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}

func main() {
	// Initialize an AWS session with configured credentials
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("us-east-1")},
	)
	if err != nil {
		panic(err)
	}

	// Define client and bucket name (Use environment variables)
	s3Client := s3.New(sess)
	bucket := "ian-test-bucket-go-python"

	// List objects in bucket, retrieve the key from the returned result
	input := &s3.ListObjectsV2Input{
		Bucket:  aws.String(bucket),
		MaxKeys: aws.Int64(2),
	}
	result, err := s3Client.ListObjectsV2(input)
	if err != nil {
		fmt.Println(err)
	}
	key := result.Contents[0].Key

	// Using the key, get the object from the bucket
	obj, err := s3Client.GetObject(&s3.GetObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(*key),
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
	test := csvToSliceOfMaps(record)
	fmt.Println(test)
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
			line[header[i]] = record[i]
		}
		returnMap = append(returnMap, line)
	}
	return returnMap
}

package main

import (
	"bufio"
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/ianlopshire/go-fixedwidth"
)

type leRecord struct {
	lastName             string
	firstName            string
	middleName           string
	maturitySuffix       string
	titleOfRespect       string
	companyName          string
	careOfAddress        string
	secondaryAddress     string
	primaryAddress       string
	addressLine2         string
	city                 string
	state                string
	zipcode              string
	plus4                string
	deliveryPointCode    string
	CarrierRouteCode     string
	gender               string
	IBehaviorCampaign    string
	Segment              string
	IBehaviorID          string
	MailKey              string
	ClientCustomerNumber string
	Filler               string
	EOFChar              string
}

func main() {
	getObjectReturnMaps("a077f0_mail-c6a12f87-a130-4b29-9c0d-a1352b83b7a7.dat", "slm-addressfile-incoming")

}

func getObjectReturnMaps(key, bucket string) {
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
	var records []leRecord

	dec := fixedwidth.NewDecoder(bufio.NewReader(obj.Body))
	// reader := bufio.NewReader(obj.Body)
	// err = fixedwidth.Unmarshal(reader, &records)
	if err != nil {
		fmt.Println("Error", err)
	}
	fmt.Println(dec.Decode(records))
	fmt.Println(records)
	fmt.Println("Records Parsed!")
	// return csvToSliceOfMaps(record)
}

package main

import (
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
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("us-east-1")},
	)
	if err != nil {
		panic(err)
	}

	s3Client := s3.New(sess)
	bucket := "ian-test-bucket-go-python"

	//files := []string
	input := &s3.ListObjectsV2Input{
		Bucket:  aws.String(bucket),
		MaxKeys: aws.Int64(2),
	}
	result, err := s3Client.ListObjectsV2(input)
	if err != nil {
		fmt.Println(err)
	}
	key := result.Contents[0].Key
	fmt.Println(result)
	fmt.Println(*key)
}

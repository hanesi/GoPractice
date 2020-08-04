package main

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
	"github.com/google/uuid"
)

func main() {
	sendSQSMessage("9e050ede-a4de-412c-8298-540ca17380d8")
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

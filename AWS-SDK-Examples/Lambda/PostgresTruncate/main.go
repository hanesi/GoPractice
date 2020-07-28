package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/glue"
	_ "github.com/lib/pq"
)

type response struct {
	UTC time.Time `json:"utc"`
}

func handleRequest(ctx context.Context, request events.SQSEvent) (events.APIGatewayProxyResponse, error) {
	sess := aws.NewConfig()
	sess.Region = "us-east-1"

	glueClient := glue.New(*sess)

	now := time.Now()
	resp := &response{
		UTC: now.UTC(),
	}
	body, err := json.Marshal(resp)
	if err != nil {
		return events.APIGatewayProxyResponse{}, err
	}
	fmt.Println("Recieved event", request)

	host := os.Getenv("PGhost")
	port := 5432
	user := os.Getenv("PGuser")
	password := os.Getenv("PGpassword")
	dbname := os.Getenv("PGdbname")

	pgConString := fmt.Sprintf("port=%d host=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		port, host, user, password, dbname)

	db, err := sql.Open("postgres", pgConString)
	if err != nil {
		panic(err)
	}
	fmt.Println("Connected Successfully")
	defer db.Close()

	sqlStatement := "truncate table scans_staging;"

	fmt.Println(sqlStatement)
	_, err = db.Query(sqlStatement)
	if err != nil {
		fmt.Println("Failed to run query", err)
		return events.APIGatewayProxyResponse{Body: "Query Failed To Run", StatusCode: 400}, err
	}

	fmt.Println("Query executed!")

	jobName := "ScanAggregateETL"
	var params *glue.StartJobRunInput
	params.JobName = &jobName

	req := glueClient.StartJobRunRequest(params)
	respGlue, err := req.Send(context.TODO())
	if err == nil {
		fmt.Println(respGlue)
	}
	return events.APIGatewayProxyResponse{Body: string(body), StatusCode: 200}, nil
}

func main() {
	lambda.Start(handleRequest)
}

package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	_ "github.com/lib/pq"
)

type response struct {
	UTC time.Time `json:"utc"`
}

func handleRequest(ctx context.Context, request events.SQSEvent) (events.APIGatewayProxyResponse, error) {
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

	for i := range request.Records {
		msgBody := request.Records[i].Body
		fmt.Println("Processing file", msgBody)
		bucket := strings.Split(msgBody, ",")[0]
		key := strings.Split(msgBody, ",")[1]
		table := strings.Split(key, "/")[0]

		db, err := sql.Open("postgres", pgConString)
		if err != nil {
			panic(err)
		}
		fmt.Println("Connected Successfully")
		defer db.Close()

		sqlStatement := `
							select aws_s3.table_import_from_s3(
							'%s',
							'',
							'(format csv)',
							'%s',
							'%s',
							'us-east-1'
							)
							;`
		switch {
		case table == "MailFiles":
			sqlStatement = fmt.Sprintf(sqlStatement, "printer_mailings", bucket, key)
		case table == "TX-Files":
			sqlStatement = fmt.Sprintf(sqlStatement, "orders", bucket, key)
		case table == "Holdouts":
			sqlStatement = fmt.Sprintf(sqlStatement, "holdouts", bucket, key)
		}

		_, err = db.Query(sqlStatement)
		if err != nil {
			fmt.Println("Failed to run query", err)
			return events.APIGatewayProxyResponse{Body: "Query Failed To Run", StatusCode: 400}, err
		}
		fmt.Println("Query executed!")
		return events.APIGatewayProxyResponse{Body: string(body), StatusCode: 200}, nil
	}

}

func main() {
	lambda.Start(handleRequest)
}

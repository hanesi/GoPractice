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

type message struct {
	Filename     string `json:"filename"`
	PrinterName  string `json:"printer_name"`
	CreatedAt    string `json:"created_at"`
	UpdatedAt    string `json:"updated_at"`
	Status       string `json:"status"`
	UploadMethod string `json:"upload_method"`
	FileType     string `json:"file_type"`
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
		queryMethod := strings.Split(msgBody, "__")[0]
		fieldDict := strings.Split(msgBody, "__")[1]

		db, err := sql.Open("postgres", pgConString)
		if err != nil {
			panic(err)
		}
		fmt.Println("Connected Successfully")
		defer db.Close()

		body := message{}
		json.Unmarshal([]byte(fieldDict), &body)
		fmt.Println(queryMethod)
		fmt.Println(body)

		var sqlStatement string
		switch {
		case queryMethod == "insert":
			sqlStatement = fmt.Sprintf(`INSERT INTO files VALUES ('%s', '%s', '%s', '%s', '%s', '%s', '%s');`,
				body.Filename,
				body.PrinterName,
				body.CreatedAt,
				body.UpdatedAt,
				body.Status,
				body.UploadMethod,
				body.FileType,
			)

		}
		fmt.Println(sqlStatement)
		_, err = db.Query(sqlStatement)
		if err != nil {
			fmt.Println("Failed to run query", err)
			return events.APIGatewayProxyResponse{Body: "Query Failed To Run", StatusCode: 400}, err
		}
	}
	fmt.Println("Query executed!")
	return events.APIGatewayProxyResponse{Body: string(body), StatusCode: 200}, nil
}

func main() {
	lambda.Start(handleRequest)
}

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
	ID             string `json:"id"`
	Filename       string `json:"filename"`
	Status         string `json:"status"`
	UploadMethod   string `json:"upload_method"`
	FileType       string `json:"file_type"`
	PrinterName    string `json:"printer_name"`
	PrintJobNumber string `json:"printer_job_number"`
	ClientSlug     string `json:"client_slug"`
	CreatedAt      string `json:"created_at"`
	UpdatedAt      string `json:"updated_at"`
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
		switch queryMethod {
		case "insert":
			sqlStatement = fmt.Sprintf(`INSERT INTO slm_files VALUES ('%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s');`,
				body.ID,
				body.Filename,
				body.Status,
				body.UploadMethod,
				body.FileType,
				body.PrinterName,
				body.PrintJobNumber,
				body.ClientSlug,
				body.CreatedAt,
				body.UpdatedAt,
			)
		case "update", "update error":
			sqlStatement = fmt.Sprintf(`UPDATE slm_files SET status = '%s', updated_at = '%s' where filename = '%s';`,
				body.Status,
				body.UpdatedAt,
				body.Filename,
			)
		case "update_copper":
			sqlStatement = fmt.Sprintf(`UPDATE slm_files SET status = '%s', updated_at = '%s', print_job_number = '%s', client_slug = '%s' where filename = '%s';`,
				body.Status,
				body.UpdatedAt,
				body.PrintJobNumber,
				body.ClientSlug,
				body.Filename,
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

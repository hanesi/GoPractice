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
	msgBody := request.Records[0].Body
	fmt.Println(msgBody)

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

	// sqlStatement := `
	//           select aws_s3.table_import_from_s3(
	//           'holdouts',
	//           '',
	//           '(format csv)',
	//           'slm-mail-planner-csvs',
	//           'Holdouts/run-1583957358818-part-r-00001',
	//           'us-east-1'
	//           )
	//           ;`
	//
	// _, err = db.Query(sqlStatement)
	// if err != nil {
	// 	fmt.Println("Failed to run query", err)
	// 	return events.APIGatewayProxyResponse{Body: "Query Failed To Run", StatusCode: 400}, err
	// }
	fmt.Println("Query executed!")
	//
	// cols, err := rows.Columns()
	// if err != nil {
	// 	fmt.Println("Failed to get columns", err)
	// 	return events.APIGatewayProxyResponse{Body: "Couldnt Get Columns", StatusCode: 400}, err
	// }
	// // Result is your slice string.
	// rawResult := make([][]byte, len(cols))
	// result := make([]string, len(cols))
	//
	// dest := make([]interface{}, len(cols)) // A temporary interface{} slice
	// for i := range rawResult {
	// 	dest[i] = &rawResult[i] // Put pointers to each string in the interface slice
	// }
	//
	// for rows.Next() {
	// 	err = rows.Scan(dest...)
	// 	if err != nil {
	// 		fmt.Println("Failed to scan row", err)
	// 		return events.APIGatewayProxyResponse{Body: "Couldnt Scan Row", StatusCode: 400}, err
	// 	}
	//
	// 	for i, raw := range rawResult {
	// 		if raw == nil {
	// 			result[i] = "\\N"
	// 		} else {
	// 			result[i] = string(raw)
	// 		}
	// 	}
	//
	// 	fmt.Printf("%#v\n", result[3])
	// }
	return events.APIGatewayProxyResponse{Body: string(body), StatusCode: 200}, nil
}

func main() {
	lambda.Start(handleRequest)
}

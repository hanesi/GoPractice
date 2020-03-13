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

func handleRequest(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	now := time.Now()
	resp := &response{
		UTC: now.UTC(),
	}
	body, err := json.Marshal(resp)
	if err != nil {
		return events.APIGatewayProxyResponse{}, err
	}
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
	defer db.Close()

	sqlStatement := "select * from orders limit 15"

	rows, err := db.Query(sqlStatement)
	if err != nil {
		fmt.Println("Failed to run query", err)
		return events.APIGatewayProxyResponse{Body: string(body), StatusCode: 400}, err
	}

	fmt.Println(rows)
	return events.APIGatewayProxyResponse{Body: string(body), StatusCode: 200}, nil
}

func main() {
	lambda.Start(handleRequest)
}

package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/athena"
	"github.com/joho/godotenv"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}

func main() {
	db, _ := os.LookupEnv("ATHENADB")
	table, _ := os.LookupEnv("ATHENATABLE")
	query := fmt.Sprintf("select * from %s limit 10;", table)
	// bucket, _ := os.LookupEnv("bucket")
	// outLoc := fmt.Sprintf("s3://%s/", bucket)

	queryResult, err := executeAthenaQuery(db, table, query)
	fmt.Println(queryResult, err)
	// fmt.Printf("%T\n", queryId)
}

func executeAthenaQuery(db, table, query string) ([][]interface{}, error) {
	var s athena.StartQueryExecutionInput
	var q athena.QueryExecutionContext
	var r athena.ResultConfiguration
	var qri athena.GetQueryExecutionInput
	var ip athena.GetQueryResultsInput
	var qrop *athena.GetQueryExecutionOutput
	var rc [][]interface{}

	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("us-east-1")},
	)
	if err != nil {
		panic(err)
	}

	svc := athena.New(sess, aws.NewConfig().WithRegion("us-east-1"))

	s.SetQueryString(query)
	q.SetDatabase(db)
	s.SetQueryExecutionContext(&q)
	r.SetOutputLocation("s3://slm-ian-test")
	s.SetResultConfiguration(&r)

	result, err := svc.StartQueryExecution(&s)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	qri.SetQueryExecutionId(*result.QueryExecutionId)

	duration := time.Duration(2) * time.Second // Pause for 2 seconds

	for {
		qrop, err = svc.GetQueryExecution(&qri)
		if err != nil {
			fmt.Println(err)
			return nil, err
		}
		if *qrop.QueryExecution.Status.State != "RUNNING" {
			break
		}
		fmt.Println("waiting.")
		time.Sleep(duration)

	}
	if *qrop.QueryExecution.Status.State == "SUCCEEDED" {
		ip.SetQueryExecutionId(*result.QueryExecutionId)

		op, err := svc.GetQueryResults(&ip)
		if err != nil {
			fmt.Println(err)
			return nil, err
		}
		fmt.Printf("%+v", op)
		for i := range op.ResultSet.Rows {
			if i == 0 {
				continue
			}
			var temp []interface{}
			for j := range op.ResultSet.Rows[i].Data {
				temp = append(temp, *op.ResultSet.Rows[i].Data[j].VarCharValue)
			}
			rc = append(rc, temp)
		}
	} else {
		fmt.Println(*qrop.QueryExecution.Status.State)

	}
	return rc, nil
}

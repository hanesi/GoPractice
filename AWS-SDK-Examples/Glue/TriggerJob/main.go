package main

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/glue"
)

func main() {
	sess := session.Must(session.NewSession())
	// Create a Firehose client with additional configuration
	glueService := glue.New(sess, aws.NewConfig().WithRegion("us-east-1"))

	startInput := &glue.StartJobRunInput{}
	job := "TX-HoldoutToRedshift"
	args := map[string]*string{}

	ft := "holdout"
	pdp := "run-1613160339868-part-r-00000"
	args["--filetype"] = &ft
	args["--pdp"] = &pdp

	startInput.JobName = &job
	startInput.Arguments = args

	glueService.StartJobRun(startInput)
}

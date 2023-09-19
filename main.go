package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/devJGuerrero/gambitnaxbeuser/functions"
)

func main() {
	lambda.Start(functions.RunLambdaUser)
}

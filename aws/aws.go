package aws

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
)

var Ctx context.Context
var Config aws.Config
var Failure error

func StartAWSConnection() {
	Ctx = context.TODO()
	Config, Failure = config.LoadDefaultConfig(Ctx, config.WithDefaultRegion("us-east-1"))
	if Failure != nil {
		panic("Error. Communication with AWS services failed. " + Failure.Error())
	}
}

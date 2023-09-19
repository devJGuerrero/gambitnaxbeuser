package functions

import (
	"context"
	"errors"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/devJGuerrero/gambitnaxbeuser/aws"
	"github.com/devJGuerrero/gambitnaxbeuser/database"
	"github.com/devJGuerrero/gambitnaxbeuser/models"
	"os"
	"strings"
)

func envExists(env string) bool {
	_, exists := os.LookupEnv(env)
	return exists
}

func RunLambdaUser(ctx context.Context, event events.CognitoEventUserPoolsPostConfirmation) (events.CognitoEventUserPoolsPostConfirmation, error) {
	aws.StartAWSConnection()
	if !envExists("secret-manager-name") {
		message := "Error. The environment configuration for the RDS secret manager was not set."
		fmt.Println(message)
		failure := errors.New(strings.ToLower(message))
		return event, failure
	}
	var data models.SignUp
	for row, att := range event.Request.UserAttributes {
		switch row {
		case "email":
			data.Email = att
			fmt.Println("Email -> " + data.Email)
		case "sub":
			data.UUID = att
			fmt.Println("UUId -> " + data.UUID)
		}
	}
	if failure := database.ReadSecretManager(); failure != nil {
		fmt.Println("Error. Unable to read the secret manager from the database. " + failure.Error())
		return event, failure
	}
	status := database.SignUp(data)
	return event, status
}

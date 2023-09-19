package secret_manager

import (
	"encoding/json"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/secretsmanager"
	_aws "github.com/devJGuerrero/gambitnaxbeuser/aws"
	"github.com/devJGuerrero/gambitnaxbeuser/models"
)

func GetSecret(name string) (models.SecretManagerRDSJson, error) {
	var secretManager models.SecretManagerRDSJson
	fmt.Println("-> Request secret manager aws " + name)
	svc := secretsmanager.NewFromConfig(_aws.Config)
	key, failure := svc.GetSecretValue(_aws.Ctx, &secretsmanager.GetSecretValueInput{
		SecretId: aws.String(name),
	})
	if failure != nil || json.Unmarshal([]byte(*key.SecretString), &secretManager) != nil {
		fmt.Println(failure.Error())
		return secretManager, failure
	}
	fmt.Println("-> Response secret manager aws " + name)
	return secretManager, nil
}

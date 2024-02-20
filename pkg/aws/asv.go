package aws

import (
	"encoding/json"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/secretsmanager"
)

type ApiKey struct {
	API_KEY string
}

func GetApikey(secret string) (string, error) {
	sess := NewSession()
	svc := secretsmanager.New(sess)

	input := &secretsmanager.GetSecretValueInput{
		SecretId:     aws.String(secret),
		VersionStage: aws.String("AWSCURRENT"),
	}

	result, err := svc.GetSecretValue(input)

	if err != nil {
		return "", err
	}

	var secretString ApiKey
	err = json.Unmarshal([]byte(*result.SecretString), &secretString)

	if err != nil {
		return "", err
	}

	return secretString.API_KEY, nil
}

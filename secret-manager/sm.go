package secretmanager

import (
	awsgo "aws/aws-go"
	"aws/models"
	"encoding/json"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/secretsmanager"
)

func GetSecret(secretName string) (models.Secret, error) {
	var dataSecret models.Secret
	fmt.Println("> Buscando secreto" + secretName)
	svc := secretsmanager.NewFromConfig(awsgo.Cfg)

	key, err := svc.GetSecretValue(awsgo.Ctx, &secretsmanager.GetSecretValueInput{
		SecretId: aws.String(secretName),
	})
	if err != nil {
		fmt.Println("Error al buscar el secreto" + err.Error())
		return dataSecret, err
	}
	json.Unmarshal([]byte(*key.SecretString), &dataSecret)

	fmt.Println("> Se encontro el secreto:" + secretName)
	return dataSecret, nil
}

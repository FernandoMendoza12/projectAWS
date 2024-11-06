package main

import (
	awsgo "aws/aws-go"
	"aws/db"
	"aws/handlers"
	"aws/models"
	secretmanager "aws/secret-manager"
	"context"
	"os"
	"strings"

	"github.com/aws/aws-lambda-go/events"
	lambda "github.com/aws/aws-lambda-go/lambda"
)

func main() {
	lambda.Start(execLambda)
}

func execLambda(ctx context.Context, request events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	var resp *events.APIGatewayProxyResponse

	awsgo.InitAWS()

	if !validateParameters() {
		resp = &events.APIGatewayProxyResponse{
			StatusCode: 400,
			Body:       "Error en las variables de entorno 'SecretName','UrlPrefix','BucketName'",
			Headers: map[string]string{
				"Content/Type": "application/json",
			},
		}
		return resp, nil
	}
	SecretModel, err := secretmanager.GetSecret(os.Getenv("SecretName"))
	if err != nil {
		resp = &events.APIGatewayProxyResponse{
			StatusCode: 400,
			Body:       "Error al buscar la variable de entorno",
			Headers: map[string]string{
				"Content/Type": "application/json",
			},
		}
		return resp, nil
	}

	path := strings.Replace(request.PathParameters["project"], os.Getenv("UrlPrefix"), "", -1)

	awsgo.Ctx = context.WithValue(awsgo.Ctx, models.Key("path"), path)
	awsgo.Ctx = context.WithValue(awsgo.Ctx, models.Key("method"), request.HTTPMethod)
	awsgo.Ctx = context.WithValue(awsgo.Ctx, models.Key("user"), SecretModel.Username)
	awsgo.Ctx = context.WithValue(awsgo.Ctx, models.Key("password"), SecretModel.Password)
	awsgo.Ctx = context.WithValue(awsgo.Ctx, models.Key("host"), SecretModel.Host)
	awsgo.Ctx = context.WithValue(awsgo.Ctx, models.Key("database"), SecretModel.Database)
	awsgo.Ctx = context.WithValue(awsgo.Ctx, models.Key("jwtSign"), SecretModel.JWTSign)
	awsgo.Ctx = context.WithValue(awsgo.Ctx, models.Key("body"), request.Body)
	awsgo.Ctx = context.WithValue(awsgo.Ctx, models.Key("bucketName"), os.Getenv("BucketName"))

	err = db.ConnectDB(awsgo.Ctx)
	if err != nil {
		resp = &events.APIGatewayProxyResponse{
			StatusCode: 500,
			Body:       "Error al llamar la conexion de mongo desde Main" + err.Error(),
			Headers: map[string]string{
				"Content/Type": "application/json",
			},
		}
		return resp, nil
	}
	respApi := handlers.Handler(awsgo.Ctx, request)
	if respApi.CustomResp == nil {
		resp = &events.APIGatewayProxyResponse{
			StatusCode: respApi.Status,
			Body:       string(respApi.Message),
			Headers: map[string]string{
				"Content-Type": "application/json",
			},
		}
		return resp, nil
	} else {
		return respApi.CustomResp, nil
	}
}

func validateParameters() bool {
	_, correct := os.LookupEnv("SecretName")
	if !correct {
		return false
	}

	_, correct = os.LookupEnv("BucketName")
	if !correct {
		return false
	}

	_, correct = os.LookupEnv("UrlPrefix")
	if !correct {
		return false
	}

	return correct
}

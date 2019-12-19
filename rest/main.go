package main

import (
	"context"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	echoLambda "github.com/awslabs/aws-lambda-go-api-proxy/echo"

	"go-serverless-demo/internal/api"
	"go-serverless-demo/internal/db"
	"go-serverless-demo/internal/echo"
)

// Handler ...
func lambdaHandler(
	ctx context.Context,
	req events.APIGatewayProxyRequest,
) (events.APIGatewayProxyResponse, error) {
	d := db.New(&aws.Config{
		Region: aws.String(os.Getenv("DYNAMO_REGION")),
	})
	d.SetTable("go-todo")
	a := api.NewAPI(d)
	e := echo.NewEcho(a)
	h := echoLambda.New(e)

	return h.ProxyWithContext(ctx, req)
}

func main() {
	lambda.Start(lambdaHandler)
}

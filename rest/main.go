package main

import (
	"context"
	"go-serverless-demo/internal/api"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	echoLambda "github.com/awslabs/aws-lambda-go-api-proxy/echo"
)

// Handler ...
func lambdaHandler(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	e := api.New()

	h := echoLambda.New(e.Echo)

	return h.ProxyWithContext(ctx, req)
}

func main() {
	lambda.Start(lambdaHandler)
}

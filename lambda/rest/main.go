package rest

import (
	"context"
	"fmt"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/awslabs/aws-lambda-go-api-proxy/handlerfunc"
)

// Handler ...
func Handler(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	f := handlerfunc.New(func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(w, "Hello world")
	})
	return f.ProxyWithContext(ctx, req)
}

func main() {
	lambda.Start(Handler)
}

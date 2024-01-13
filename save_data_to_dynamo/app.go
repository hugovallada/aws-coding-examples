package main

import (
	"context"
	"errors"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/hugovallada/save_data_to_dynamo/src/controller"
)

func handler(ctx context.Context, event events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	switch event.HTTPMethod {
	case "POST":
		return controller.CreatePersonController(ctx, event)
	default:
		return controller.ErrorResponse(405, "Method not allowed"), errors.ErrUnsupported
	}
}

func main() {
	lambda.Start(handler)
}

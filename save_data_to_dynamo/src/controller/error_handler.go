package controller

import "github.com/aws/aws-lambda-go/events"

func ErrorResponse(status int, body string) events.APIGatewayProxyResponse {
	return events.APIGatewayProxyResponse{
		StatusCode: status,
		Body:       body,
	}
}

func SuccessResponse(status int, body string) events.APIGatewayProxyResponse {
	return events.APIGatewayProxyResponse{
		StatusCode: status,
		Body:       body,
	}
}

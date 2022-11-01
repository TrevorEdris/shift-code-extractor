package services

import "github.com/aws/aws-lambda-go/events"

type (
	Login struct{}
)

func LoginHandler(request events.APIGatewayV2HTTPRequest) (events.APIGatewayProxyResponse, error) {
	// TODO: Redirect the user to the Cognito login page
	return events.APIGatewayProxyResponse{}, errNotImplemented
}

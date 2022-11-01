package main

import (
	"github.com/aws/aws-lambda-go/lambda"

	"github.com/TrevorEdris/shift-code-extractor/services"
)

func main() {
	lambda.Start(services.SubscribeUIHandler)
}

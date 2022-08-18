package main

import (
	"github.com/apotox/go-encrynote/functions/create_note"
	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	lambda.Start(create_note.Handler)
}

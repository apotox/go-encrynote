package main

import (
	"github.com/apotox/go-encrynote/functions/delete_note"
	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	lambda.Start(delete_note.Handler)
}

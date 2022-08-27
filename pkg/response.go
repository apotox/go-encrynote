package pkg

import (
	"encoding/json"

	"github.com/aws/aws-lambda-go/events"
)

// MakeResponse returns a response with the given body and status code.
func MakeResponse(obj interface{}, StatusCode int) (events.APIGatewayProxyResponse, error) {

	jsonItem, err := json.Marshal(obj)
	stringItem := string(jsonItem) + "\n"
	return events.APIGatewayProxyResponse{
		Body:       stringItem,
		StatusCode: StatusCode,
		Headers: map[string]string{
			"Access-Control-Allow-Origin":      "*",
			"Access-Control-Allow-Methods":     "POST, PUT, GET, DELETE, PATCH, OPTIONS",
			"Access-Control-Allow-Headers":     "*",
			"Access-Control-Allow-Credentials": "true",
			"Content-Type":                     "application/json",
		},
	}, err
}

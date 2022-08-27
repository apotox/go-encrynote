package pkg

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/apotox/go-encrynote/config"
	"github.com/apotox/go-encrynote/database"
	"github.com/aws/aws-lambda-go/events"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func ParseBody[T any](body string) (*T, error) {
	var event T
	err := json.Unmarshal([]byte(body), &event)

	if err != nil {
		log.Printf("Error parsing body: %s", err)
		return nil, err
	}
	return &event, nil
}

func GetClientContext() (*mongo.Client, context.Context, context.CancelFunc) {
	client := database.GetDatabaseClient()
	ctx, cancel := context.WithTimeout(context.TODO(), config.CONTEXT_TIMEOUT)
	return client, ctx, cancel
}

func ParseResponseBody(resp events.APIGatewayProxyResponse) (map[string]interface{}, error) {
	var b ResponseBodyType
	err := json.Unmarshal([]byte(resp.Body), &b)
	if err != nil {
		fmt.Printf("Error Unmarshal body: %v \n", err)
		return nil, err
	}
	return b, nil
}

func Marshal(entity interface{}) string {
	encoded, _ := json.Marshal(entity)
	return string(encoded)
}

// GetDateTime returns the current date time in the format YYYY-MM-DD HH:MM:SS
func GetDateTimeNow() Datetime {
	return Datetime{
		primitive.NewDateTimeFromTime(time.Now()),
	}
}

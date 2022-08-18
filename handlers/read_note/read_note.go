package main

import (
	"context"

	"github.com/apotox/go-encrynote/functions/read_note"
	"github.com/apotox/go-encrynote/pkg"
	"github.com/apotox/go-encrynote/queue"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	lambda.Start(func(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
		return read_note.Handler(ctx, request, pkg.Services{
			QueueDeleteNote: queue.GetQueueDeleteNote(),
		})
	})
}

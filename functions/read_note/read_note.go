package read_note

import (
	"context"
	"log"

	"github.com/apotox/go-encrynote/pkg"
	"github.com/aws/aws-lambda-go/events"
)

func Handler(ctx context.Context, request events.APIGatewayProxyRequest, services pkg.Services) (events.APIGatewayProxyResponse, error) {

	idAndKey := request.QueryStringParameters["id"]
	if idAndKey == "" || len(idAndKey) != 56 {
		return pkg.MakeResponse(map[string]interface{}{
			"error": "id is required",
		}, 400)
	}

	noteId, privateKey := idAndKey[:24], idAndKey[24:]

	note, err := Read(noteId)

	if err != nil {
		log.Printf("Error reading noteid:%s - %s", noteId, err)
		return pkg.MakeResponse(map[string]interface{}{
			"error": err.Error(),
		}, 400)
	}

	message := note.DecryptMessage(privateKey)

	defer func() {
		log.Printf("Deleting message: %s", noteId)
		services.QueueDeleteNote.Publish(noteId)
	}()

	return pkg.MakeResponse(map[string]interface{}{
		"itemId":  note.Id,
		"message": message,
	}, 200)
}

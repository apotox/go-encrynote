package read_note

import (
	"context"
	"log"

	"github.com/apotox/go-encrynote/pkg"
	"github.com/aws/aws-lambda-go/events"
)

func Handler(ctx context.Context, request events.APIGatewayProxyRequest, services pkg.Services) (events.APIGatewayProxyResponse, error) {

	embadedId := request.QueryStringParameters["id"]
	if embadedId == "" || len(embadedId) != 56 {
		return pkg.GetResponse(map[string]interface{}{
			"error": "id is required",
		}, 400)
	}

	noteId, privateKey := embadedId[:24], embadedId[24:]

	note, err := Read(noteId)

	if err != nil {
		log.Printf("Error reading noteid:%s - %s", noteId, err)
		return pkg.GetResponse(map[string]interface{}{
			"error": err.Error(),
		}, 400)
	}

	message := note.DecryptMessage(privateKey)

	defer func() {
		log.Printf("Deleting message: %s", noteId)
		services.QueueDeleteNote.Publish(noteId)
	}()

	return pkg.GetResponse(map[string]interface{}{
		"itemId":  note.Id,
		"message": message,
	}, 200)
}

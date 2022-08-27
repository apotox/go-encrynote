package create_note

import (
	"context"
	"log"

	"github.com/apotox/go-encrynote/pkg"
	"github.com/aws/aws-lambda-go/events"
	"gopkg.in/validator.v2"
)

func Handler(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	body := request.Body

	note, err := pkg.ParseBody[pkg.Note](body)

	if err != nil {
		return pkg.MakeResponse(map[string]interface{}{
			"error": err.Error(),
		}, 500)
	}

	if note == nil {
		return pkg.MakeResponse(map[string]interface{}{
			"error": "note parsed body is nil",
		}, 500)
	}

	if err := validator.Validate(note); err != nil {
		return pkg.MakeResponse(map[string]interface{}{
			"error": err.Error(),
		}, 400)
	}

	usedKey := note.EncryptMessage()

	err = Insert(note)

	if err != nil {
		log.Printf("Error inserting note: %s", err)
		return pkg.MakeResponse(map[string]interface{}{
			"error": err.Error(),
		}, 500)
	}

	return pkg.MakeResponse(map[string]interface{}{
		"itemId":  note.Id,
		"usedKey": usedKey,
	}, 200)
}

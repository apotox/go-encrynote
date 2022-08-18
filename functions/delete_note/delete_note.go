package delete_note

import (
	"context"
	"errors"
	"log"

	"github.com/aws/aws-lambda-go/events"
)

func Handler(ctx context.Context, request events.SQSEvent) error {

	noteId := request.Records[0].Body

	log.Printf("deleting note with id: %s", noteId)

	if noteId == "" {
		log.Printf("No note id found in request")
		return errors.New("No note id found in request")
	}

	_, err := ReadNote(noteId)

	if err != nil {
		log.Printf("Error reading note: %s", err)
		return err
	}

	err = Delete(noteId)

	if err != nil {
		log.Printf("Error deleting note: %s", err)
		return err
	}

	log.Printf("Note %s deleted", noteId)
	return nil
}

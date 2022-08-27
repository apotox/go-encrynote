package create_note

import (
	database "github.com/apotox/go-encrynote/database"
	"github.com/apotox/go-encrynote/pkg"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// insert note into database
func Insert(note *pkg.Note) error {

	client, ctx, cancel := pkg.GetClientContext()
	defer cancel()

	note.Id = primitive.NewObjectID()
	note.CreatedAt = pkg.GetDateTimeNow()

	_, err := database.GetCollection(client, note.GetCollectionName()).InsertOne(ctx, note)
	if err != nil {
		return err
	}

	return err
}

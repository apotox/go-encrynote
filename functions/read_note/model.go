package read_note

import (
	"github.com/apotox/go-encrynote/database"
	"github.com/apotox/go-encrynote/pkg"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func Read(noteId string) (*pkg.Note, error) {
	client, ctx, cancel := pkg.GetClientContext()
	defer cancel()

	id, _ := primitive.ObjectIDFromHex(noteId)
	result := database.GetCollection(client, string(pkg.COL_NOTES)).FindOne(ctx, bson.M{"_id": id})
	note := &pkg.Note{}
	err := result.Decode(note)

	if err != nil {
		return nil, err
	}

	return note, nil
}

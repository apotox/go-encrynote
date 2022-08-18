package tools

import (
	"github.com/apotox/go-encrynote/database"
	"github.com/apotox/go-encrynote/pkg"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CleanCollection(collectionName pkg.COLLECTION_NAME) {
	client, ctx, cancel := pkg.GetClientContext()
	defer cancel()
	database.GetCollection(client, string(collectionName)).Drop(ctx)
}

func Read[T any](collectionName pkg.COLLECTION_NAME, id string) (T, error) {
	client, ctx, cancel := pkg.GetClientContext()
	defer cancel()
	_id, err := primitive.ObjectIDFromHex(id)
	var x T
	if err != nil {
		return x, err
	}

	database.GetCollection(client, string(collectionName)).FindOne(ctx, bson.M{
		"_id": _id,
	}).Decode(&x)

	return x, nil
}

package migrations

import (
	"context"
	"fmt"

	migrate "github.com/xakep666/mongo-migrate"
	"go.mongodb.org/mongo-driver/mongo"
)

func init() {
	migrate.Register(func(db *mongo.Database) error {
		fmt.Printf("Adding %s collection\n", "notes")
		err := db.CreateCollection(context.Background(), "notes", nil)

		if err != nil {
			return err
		}

		return nil
	}, func(db *mongo.Database) error {
		fmt.Printf("Dropping %s collection\n", "notes")
		err := db.Collection("notes").Drop(context.TODO())
		if err != nil {
			return err
		}
		return nil
	})
}

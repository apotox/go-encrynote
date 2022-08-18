package database

import (
	"context"
	"log"
	"time"

	"github.com/apotox/go-encrynote/config"
	"github.com/apotox/go-encrynote/migrations"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var _client *mongo.Client

func GetDatabaseClient() *mongo.Client {

	if _client != nil {
		return _client
	}

	config := config.GetConfig()

	opts := options.Client().ApplyURI(config.MONGO_URL)

	client, err := mongo.NewClient(opts)
	if err != nil {
		log.Fatal(err)
	}

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	//ping the database
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	migrations.StartMigrations(client.Database(config.DATABASE_NAME))
	_client = client
	return client
}

func GetCollection(client *mongo.Client, collectionName string) *mongo.Collection {
	config := config.GetConfig()
	collection := client.Database(config.DATABASE_NAME).Collection(collectionName)
	return collection
}

package migrations

import (
	"log"

	migrate "github.com/xakep666/mongo-migrate"
	"go.mongodb.org/mongo-driver/mongo"
)

func StartMigrations(db *mongo.Database) {
	log.Printf("Starting migrations. DB: %s", db.Name())
	migrate.SetDatabase(db)

	if err := migrate.Up(migrate.AllAvailable); err != nil {
		log.Fatalf("Error migrate: %s", err)
	}
}

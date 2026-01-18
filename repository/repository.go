package repository

import (
	"context"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"log"
)

var db *mongo.Database

func Init(database *mongo.Database) {
	db = database

	ctx := context.Background()

	if err := NewDossierRepository(db).InitIndexes(ctx); err != nil {
		log.Printf("Dossier index creation failed: %v\n", err)
	}

	if err := NewProfileRepository(db).InitIndexes(ctx); err != nil {
		log.Printf("Profile index creation failed: %v\n", err)
	}

	if err := NewAccessRepository(db).InitIndexes(ctx); err != nil {
		log.Printf("Access index creation failed: %v\n", err)
	}
}

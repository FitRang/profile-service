package repository

import (
	"context"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type DossierRepository struct {
	Col *mongo.Collection
}

func NewDossierRepository(db *mongo.Database) *DossierRepository {
	return &DossierRepository{
		Col: db.Collection("dossiers"),
	}
}

func (r *DossierRepository) InitIndexes(ctx context.Context) error {
	models := []mongo.IndexModel{
		{
			Keys:    bson.M{"username": 1},
			Options: options.Index().SetUnique(true),
		},
	}

	_, err := r.Col.Indexes().CreateMany(ctx, models)
	return err
}

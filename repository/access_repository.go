package repository

import (
	"context"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type AccessRepository struct {
    Col *mongo.Collection
}

func NewAccessRepository(db *mongo.Database) *AccessRepository {
    return &AccessRepository{
        Col: db.Collection("access-request"),
    }
}

func (r *AccessRepository) InitIndexes(ctx context.Context) error {
    models := []mongo.IndexModel{
        {
            Keys: bson.D{
                {Key: "username", Value: 1},
                {Key: "requester", Value: 1},
            },
            Options: options.Index().SetUnique(true),
        },
    }

    _, err := r.Col.Indexes().CreateMany(ctx, models)
    return err
}

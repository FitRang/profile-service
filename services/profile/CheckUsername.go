package profile

import (
	"context"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

func (p *ProfileService) CheckUsername(ctx context.Context, username string) (bool, error) {
	filter := bson.M{"username": username}

	err := p.ProfileRepo.Col.FindOne(ctx, filter).Err()
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return false, nil
		}
		return false, err
	}

	return true, nil
}

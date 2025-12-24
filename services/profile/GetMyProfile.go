package profile

import (
	"context"
	"errors"
	"fmt"

	"github.com/Foxtrot-14/FitRang/profile-service/graph/model"
	"github.com/Foxtrot-14/FitRang/profile-service/middleware"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

func(p *ProfileService) GetMyProfile(ctx context.Context) (*model.Profile, error) {
	emailID := ctx.Value(middleware.EmailKey).(string)
	if emailID == "" {
		return nil, errors.New("unauthenticated: email-id missing in headers")
	}

	filter := bson.M{"email": emailID}
	res:= p.Repo.Col.FindOne(ctx, filter)
	if res.Err() != nil {
		if res.Err() == mongo.ErrNoDocuments {
			return nil, fmt.Errorf("profile not found")
		}
		return nil, res.Err()
	}

	var profile model.Profile
	if err := res.Decode(&profile); err != nil {
		return nil, errors.New("failed to decode profile")
	}
	return &profile, nil
}

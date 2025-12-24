package profile

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/Foxtrot-14/FitRang/profile-service/graph/model"
	"github.com/Foxtrot-14/FitRang/profile-service/middleware"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

func (p *ProfileService) UpdateProfile(ctx context.Context, input model.ProfileUpdateInput) (*model.Profile, error) {
	emailID := ctx.Value(middleware.EmailKey).(string)
	if emailID == "" {
		return nil, errors.New("unauthenticated: email-id missing in headers")
	}

	if emailID != input.Email {
		return nil, errors.New("forbidden: you can only update your own profile")
	}

	update := bson.M{}
	if input.FullName != nil {
		update["fullName"] = *input.FullName
	}

	if len(update) == 0 {
		return nil, errors.New("no fields provided for update")
	}

	now := time.Now().UTC().Format(time.RFC3339)
	update["updatedAt"] = now

	filter := bson.M{"username": input.Username}
	opts := options.FindOneAndUpdate().SetReturnDocument(options.After)

	res := p.Repo.Col.FindOneAndUpdate(
		ctx,
		filter,
		bson.M{"$set": update},
		opts,
	)

	if res.Err() != nil {
		if res.Err() == mongo.ErrNoDocuments {
			return nil, fmt.Errorf("profile not found")
		}
		return nil, res.Err()
	}

	var profile model.Profile
	if err := res.Decode(&profile); err != nil {
		return nil, err
	}

	p.sendProfileToIndex(profile)

	return &profile, nil
}

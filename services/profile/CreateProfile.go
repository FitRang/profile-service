package profile

import (
	"context"
	"time"

	"github.com/Foxtrot-14/FitRang/profile-service/apperror"
	"github.com/Foxtrot-14/FitRang/profile-service/db"
	"github.com/Foxtrot-14/FitRang/profile-service/graph/model"
	"github.com/Foxtrot-14/FitRang/profile-service/middleware"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

func (p *ProfileService) CreateProfile(ctx context.Context, input model.ProfileCreateInput) (*model.MyProfile, error) {
	emailID := ctx.Value(middleware.EmailKey).(string)
	if emailID == "" {
		return nil, apperror.New(
			apperror.Unauthenticated,
			"Authentication required",
		)
	}

	now := time.Now().UTC().Format(time.RFC3339)

	bsonProfile := db.Profile{
		Username:  input.Username,
		FullName:  input.FullName,
		Email:     emailID,
		CreatedAt: now,
		UpdatedAt: now,
	}

	resp, err := p.Repo.Col.InsertOne(ctx, bsonProfile)
	if err != nil {
		if mongo.IsDuplicateKeyError(err) {
			return nil, apperror.New(
				apperror.Conflict,
				"Profile already exists",
			)
		}
	}

	oid, ok := resp.InsertedID.(bson.ObjectID)
	if !ok {
		return nil, apperror.Wrap(
			apperror.Internal,
			"Failed to create profile",
			err,
		)
	}

	gqlProfile := &model.Profile{
		ID:         oid.Hex(),
		Username:   bsonProfile.Username,
		FullName:   bsonProfile.FullName,
		ProfileURL: bsonProfile.ProfileURL,
		CreatedAt:  bsonProfile.CreatedAt,
		UpdatedAt:  bsonProfile.UpdatedAt,
	}

	p.sendProfileToIndex(*gqlProfile)
	return &model.MyProfile{
		ID:         oid.Hex(),
		Username:   bsonProfile.Username,
		FullName:   bsonProfile.FullName,
		Email:      bsonProfile.Email,
		ProfileURL: bsonProfile.ProfileURL,
		CreatedAt:  bsonProfile.CreatedAt,
		UpdatedAt:  bsonProfile.UpdatedAt,
	}, nil
}

package profile

import (
	"context"

	"github.com/Foxtrot-14/FitRang/profile-service/apperror"
	"github.com/Foxtrot-14/FitRang/profile-service/graph/model"
	"github.com/Foxtrot-14/FitRang/profile-service/middleware"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

func (p *ProfileService) GetProfile(
	ctx context.Context,
	username string,
) (*model.Profile, error) {
	emailID, ok := ctx.Value(middleware.EmailKey).(string)
	if !ok || emailID == "" {
		return nil, apperror.New(
			apperror.Unauthenticated,
			"Authentication required",
		)
	}

	filter := bson.M{"username": username}
	res := p.Repo.Col.FindOne(ctx, filter)

	if err := res.Err(); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, apperror.New(
				apperror.NotFound,
				"Profile not found",
			)
		}

		return nil, apperror.Wrap(
			apperror.Internal,
			"Failed to fetch profile",
			err,
		)
	}

	var profile model.Profile
	if err := res.Decode(&profile); err != nil {
		return nil, apperror.Wrap(
			apperror.Internal,
			"Failed to decode profile",
			err,
		)
	}

	// TODO: check if has access
	profile.AccessStatus = model.AccessStatusNo
	return &profile, nil
}

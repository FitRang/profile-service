package profile

import (
	"context"

	"github.com/Foxtrot-14/FitRang/profile-service/apperror"
	"github.com/Foxtrot-14/FitRang/profile-service/graph/model"
	"github.com/Foxtrot-14/FitRang/profile-service/middleware"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

func (p *ProfileService) GetMyProfile(
	ctx context.Context,
) (*model.MyProfile, error) {

	emailID := ctx.Value(middleware.EmailKey).(string)
	if emailID == "" {
		return nil, apperror.New(
			apperror.Unauthenticated,
			"Authentication required",
		)
	}

	filter := bson.M{"email": emailID}
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

	var profile model.MyProfile
	if err := res.Decode(&profile); err != nil {
		return nil, apperror.Wrap(
			apperror.Internal,
			"Failed to decode profile",
			err,
		)
	}

	return &profile, nil
}

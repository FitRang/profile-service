package profile

import (
	"context"

	"github.com/Foxtrot-14/FitRang/profile-service/apperror"
	"github.com/Foxtrot-14/FitRang/profile-service/db"
	"github.com/Foxtrot-14/FitRang/profile-service/middleware"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

func (p *ProfileService) UpdateAvatar(ctx context.Context) (string, error) {
	emailID := ctx.Value(middleware.EmailKey).(string)
	if emailID == "" {
		return "", apperror.New(
			apperror.Unauthenticated,
			"Authentication required",
		)
	}
	filter := bson.M{"email": emailID}
	res := p.ProfileRepo.Col.FindOne(ctx, filter)
	if err := res.Err(); err != nil {
		if err == mongo.ErrNoDocuments {
			return "", apperror.New(
				apperror.NotFound,
				"Profile not found",
			)
		}

		return "", apperror.Wrap(
			apperror.Internal,
			"Failed to fetch profile",
			err,
		)
	}

	var profile db.Profile
	if err := res.Decode(&profile); err != nil {
		return "", apperror.Wrap(
			apperror.Internal,
			"Failed to decode profile",
			err,
		)
	}

	profileJSON := db.ToGraphQLProfile(&profile)

	url, err := getURL(ctx, profileJSON.ID)
	if err != nil {
		return "", apperror.Wrap(
			apperror.Internal,
			"Failed to generate upload link",
			err,
		)
	}

	return url, nil
}

package profile

import (
	"context"

	"github.com/Foxtrot-14/FitRang/profile-service/apperror"
	"github.com/Foxtrot-14/FitRang/profile-service/db"
	"github.com/Foxtrot-14/FitRang/profile-service/graph/model"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

func (p *ProfileService) getRequesterProfile(ctx context.Context, emailID string) (*model.MyProfile, error) {
	filter := bson.M{"email": emailID}
	res := p.ProfileRepo.Col.FindOne(ctx, filter)
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

	var dbProfile db.Profile
	if err := res.Decode(&dbProfile); err != nil {
		return nil, apperror.Wrap(
			apperror.Internal,
			"Failed to decode profile",
			err,
		)
	}
	profile := db.ToGraphQLProfile(&dbProfile)
	return profile, nil
}

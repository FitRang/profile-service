package profile

import (
	"context"
	"slices"

	"github.com/Foxtrot-14/FitRang/profile-service/apperror"
	"github.com/Foxtrot-14/FitRang/profile-service/graph/model"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

func (p *ProfileService) GetProfile(
	ctx context.Context,
	username string,
) (*model.Profile, error) {
	emailID, err := getEmailFromContext(ctx)
	if err != nil {
		return nil, err
	}
	reqProfile, err := p.getRequesterProfile(ctx, emailID)
	if err != nil {
		return nil, err
	}

	filter := bson.M{"username": username}
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

	var profile model.Profile
	if err := res.Decode(&profile); err != nil {
		return nil, apperror.Wrap(
			apperror.Internal,
			"Failed to decode profile",
			err,
		)
	}

	filter = bson.M{"username": username}
	res = p.DossierRepo.Col.FindOne(ctx, filter)
	if err := res.Err(); err != nil {
		if err == mongo.ErrNoDocuments {
			profile.AccessStatus = model.AccessStatusNo
		}
		return nil, apperror.Wrap(
			apperror.Internal,
			"Failed to fetch profile",
			err,
		)
	}

	var dossier model.Dossier
	if err := res.Decode(&dossier); err != nil {
		return nil, apperror.Wrap(
			apperror.Internal,
			"Failed to decode profile",
			err,
		)
	}

	filter = bson.M{
		"owner":     profile.Username,
		"requester": reqProfile,
	}
	res = p.AccessRepo.Col.FindOne(ctx, filter)
	if err := res.Err(); err != nil {
		if err == mongo.ErrNoDocuments {
			profile.AccessStatus = model.AccessStatusRequested
		}
		return nil, apperror.Wrap(
			apperror.Internal,
			"Failed to fetch record",
			err,
		)
	}
	if slices.Contains(dossier.Viewers, reqProfile.Username) {
		profile.AccessStatus = model.AccessStatusYes
	}
	profile.AccessStatus = model.AccessStatusNo
	return &profile, nil
}

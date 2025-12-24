package dossier

import (
	"context"
	"errors"
	"time"

	"github.com/Foxtrot-14/FitRang/profile-service/db"
	"github.com/Foxtrot-14/FitRang/profile-service/graph/model"
	"github.com/Foxtrot-14/FitRang/profile-service/middleware"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

func (d *DossierService) CreateDossier(ctx context.Context, input model.CreateDossier) (*model.Dossier, error) {
	emailID := ctx.Value(middleware.EmailKey).(string)
	if emailID == "" {
		return nil, errors.New("unauthenticated: email-id missing in headers")
	}

	var profile model.Profile
	err := d.ProfileRepo.Col.FindOne(
		ctx,
		bson.M{"email": emailID},
	).Decode(&profile)

	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, errors.New("profile not found for this user")
		}
		return nil, err
	}

	now := time.Now().UTC().Format(time.RFC3339)
	dossier := &model.Dossier{
		Email:           emailID,
		Username:        profile.Username,
		FaceType:        input.FaceType,
		SkinTone:        input.SkinTone,
		BodyType:        input.BodyType,
		Gender:          input.Gender,
		PreferredColors: input.PreferredColors,
		DislikedColors:  input.DislikedColors,
		Height:          input.Height,
		Weight:          input.Weight,
		CreatedAt:       now,
		UpdatedAt:       now,
	}

	dossierDB, err := db.ToBsonDossier(dossier)
	if err != nil {
		return nil, err
	}

	_, err = d.DossierRepo.Col.InsertOne(ctx, dossierDB)
	if err != nil {
		if mongo.IsDuplicateKeyError(err) {
			return nil, errors.New("dossier already exists for this user")
		}

		return nil, err
	}
	return dossier, nil
}

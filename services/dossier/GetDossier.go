package dossier

import (
	"context"
	"errors"
	"fmt"
	"slices"

	"github.com/Foxtrot-14/FitRang/profile-service/graph/model"
	"github.com/Foxtrot-14/FitRang/profile-service/middleware"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

func (d *DossierService) GetDossier(ctx context.Context, username string) (*model.Dossier, error) {
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

	filter := bson.M{"username": username}
	res := d.DossierRepo.Col.FindOne(ctx, filter)
	if res.Err() != nil {
		if res.Err() == mongo.ErrNoDocuments {
			return nil, fmt.Errorf("this user does not have a dossier")
		}
		return nil, res.Err()
	}

	var dossier model.Dossier
	if err := res.Decode(&dossier); err != nil {
		return nil, errors.New("failed to decode dossier")
	}

	if slices.Contains(dossier.Viewers, profile.Username) && dossier.Email != emailID {
		return &dossier, nil
    } else {
        return nil, fmt.Errorf("you are not allowed to view this dossier")
	}
}

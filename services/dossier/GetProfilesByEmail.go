package dossier

import (
	"context"
	"errors"

	"github.com/Foxtrot-14/FitRang/profile-service/db"
	"github.com/Foxtrot-14/FitRang/profile-service/graph/model"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

func (d *DossierService) GetProfilesByEmail(
	ctx context.Context,
	email string,
) ([]db.Profile, []db.Dossier, error) {
	var requesterProfile model.Profile
	if err := d.ProfileRepo.Col.FindOne(
		ctx,
		bson.M{"email": email},
	).Decode(&requesterProfile); err != nil {

		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, nil, errors.New("profile not found for this user")
		}
		return nil, nil, err
	}

	pipeline := mongo.Pipeline{
		{{Key: "$match", Value: bson.M{
			"viewers": requesterProfile.Username,
		}}},
		{{Key: "$lookup", Value: bson.M{
			"from":         "profiles",
			"localField":   "email",
			"foreignField": "email",
			"as":           "profile",
		}}},
		{{Key: "$unwind", Value: "$profile"}},
	}

	cursor, err := d.DossierRepo.Col.Aggregate(ctx, pipeline)
	if err != nil {
		return nil, nil, err
	}
	defer cursor.Close(ctx)

	type aggResult struct {
		db.Dossier `bson:",inline"`
		Profile    db.Profile `bson:"profile"`
	}

	var aggResults []aggResult
	if err := cursor.All(ctx, &aggResults); err != nil {
		return nil, nil, err
	}

	profiles := make([]db.Profile, 0, len(aggResults))
	dossiers := make([]db.Dossier, 0, len(aggResults))

	for _, r := range aggResults {
		profiles = append(profiles, r.Profile)
		dossiers = append(dossiers, r.Dossier)
	}

	return profiles, dossiers, nil
}

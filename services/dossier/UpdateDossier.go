package dossier

import (
	"context"
	"errors"
	"time"

	"github.com/Foxtrot-14/FitRang/profile-service/graph/model"
	"github.com/Foxtrot-14/FitRang/profile-service/middleware"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

func (d *DossierService) UpdateDossier(
	ctx context.Context,
	input model.UpdateDossier,
) (*model.Dossier, error) {

	emailID, ok := ctx.Value(middleware.EmailKey).(string)
	if !ok || emailID == "" {
		return nil, errors.New("unauthenticated: email-id missing in headers")
	}

	update := bson.M{}
	set := bson.M{}

	if input.FaceType != nil {
		set["faceType"] = *input.FaceType
	}
	if input.SkinTone != nil {
		set["skinTone"] = input.SkinTone
	}
	if input.BodyType != nil {
		set["bodyType"] = *input.BodyType
	}
	if input.Gender != nil {
		set["gender"] = input.Gender
	}
	if input.PreferredColors != nil {
		set["preferredColors"] = input.PreferredColors
	}
	if input.DislikedColors != nil {
		set["dislikedColors"] = input.DislikedColors
	}
	if input.Height != nil {
		set["height"] = input.Height
	}
	if input.Weight != nil {
		set["weight"] = input.Weight
	}

	if len(set) == 0 {
		return nil, errors.New("no fields provided to update")
	}

	set["updatedAt"] = time.Now().UTC().Format(time.RFC3339)
	update["$set"] = set

	res := d.DossierRepo.Col.FindOneAndUpdate(
		ctx,
		bson.M{"email": emailID},
		update,
		options.FindOneAndUpdate().SetReturnDocument(options.After),
	)
	if res.Err() != nil {
		if errors.Is(res.Err(), mongo.ErrNoDocuments) {
			return nil, errors.New("dossier not found")
		}
		return nil, res.Err()
	}

	var updated model.Dossier
	if err := res.Decode(&updated); err != nil {
		return nil, err
	}

	return &updated, nil
}

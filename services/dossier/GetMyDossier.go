package dossier

import (
	"context"
	"errors"
	"fmt"

	"github.com/Foxtrot-14/FitRang/profile-service/graph/model"
	"github.com/Foxtrot-14/FitRang/profile-service/middleware"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

func (d *DossierService) GetMyDossier(ctx context.Context) (*model.MyDossier, error) {
	emailID := ctx.Value(middleware.EmailKey).(string)
	if emailID == "" {
		return nil, errors.New("unauthenticated: email-id missing in headers")
	}

	filter := bson.M{"email": emailID}

	res := d.DossierRepo.Col.FindOne(ctx, filter)
	if res.Err() != nil {
		if res.Err() == mongo.ErrNoDocuments {
			return nil, fmt.Errorf("you don't have a Dossier")
		}
		return nil, res.Err()
	}

	var dossier model.MyDossier
	if err := res.Decode(&dossier); err != nil {
		return nil, errors.New("failed to decode dossier")
	}

	return &dossier, nil
}

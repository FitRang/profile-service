package access

import (
	"context"
	"errors"
	"log"

	"github.com/Foxtrot-14/FitRang/profile-service/graph/model"
	"github.com/Foxtrot-14/FitRang/profile-service/middleware"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

func (a *AccessService) GrantAccess(ctx context.Context, username string) (bool, error) {
	emailID, ok := ctx.Value(middleware.EmailKey).(string)
	if !ok || emailID == "" {
		return false, errors.New("unauthenticated: email-id missing in headers")
	}

	filter := bson.M{"email": emailID}
	update := bson.M{
		"$addToSet": bson.M{
			"viewers": username,
		},
	}

	opts := options.FindOneAndUpdate().
		SetReturnDocument(options.After)

	var dossier model.Dossier
	err := a.DossierRepo.Col.
		FindOneAndUpdate(ctx, filter, update, opts).
		Decode(&dossier)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return false, errors.New("this user does not have a dossier")
		}
		return false, err
	}

	ownerUsername := dossier.Username

	payload := map[string]any{
		"granter":  ownerUsername,
		"receiver": username,
	}
	b, err := jsonToBytes(payload)
	if err != nil {
		log.Printf("failed to marshal message: %v", err)
	}

	err = a.Bus.Publish(
		"notification.granted",
		stringToByte(username),
		b,
	)
	if err != nil {
		log.Printf("failed to send message to kafka: %v", err)
	}

	return true, nil
}

package access

import (
	"context"
	"errors"
	"log"

	"github.com/Foxtrot-14/FitRang/profile-service/db"
	"github.com/Foxtrot-14/FitRang/profile-service/graph/model"
	"github.com/Foxtrot-14/FitRang/profile-service/middleware"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

func (a *AccessService) RequestAccess(ctx context.Context, username string) (bool, error) {
	emailID, ok := ctx.Value(middleware.EmailKey).(string)
	if !ok || emailID == "" {
		return false, errors.New("unauthenticated: email-id missing in headers")
	}

	var profile model.Profile
	err := a.ProfileRepo.Col.FindOne(
		ctx,
		bson.M{"email": emailID},
	).Decode(&profile)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return false, errors.New("profile not found for this user")
		}
		return false, err
	}

	access := db.ToBsonAccess(username, profile.Username)
	_, err = a.AccessRepo.Col.InsertOne(ctx, access)
	if err != nil {
		return false, err
	}

	payload := map[string]any{
		"requester": profile.Username,
		"owner":     username,
	}
	b, err := jsonToBytes(payload)
	if err != nil {
		log.Printf("failed to marshal message: %v", err)
	}

	err = a.Bus.Publish(
		"notification",
		stringToByte("Requested: "+username),
		b,
	)
	if err != nil {
		log.Printf("failed to send message to kafka: %v", err)
	}

	return true, nil
}

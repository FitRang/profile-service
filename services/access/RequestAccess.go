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

	var myProfile model.MyProfile
	err := a.ProfileRepo.Col.FindOne(
		ctx,
		bson.M{"email": emailID},
	).Decode(&myProfile)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return false, errors.New("profile not found for this user")
		}
		return false, err
	}

	var targetProfile model.MyProfile
	err = a.ProfileRepo.Col.FindOne(
		ctx,
		bson.M{"email": emailID},
	).Decode(&targetProfile)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return false, errors.New("profile not found for this user")
		}
		return false, err
	}

	access := db.ToBsonAccess(username, myProfile.Username)
	_, err = a.AccessRepo.Col.InsertOne(ctx, access)
	if err != nil {
		return false, err
	}

	b, err := messageToBytes(
		myProfile.Username,
		myProfile.Email,
		targetProfile.Username,
		targetProfile.Email,
		"has requested access",
	)
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

package db

import (
	"github.com/Foxtrot-14/FitRang/profile-service/graph/model"
	"go.mongodb.org/mongo-driver/v2/bson"
)

type Profile struct {
	ID         bson.ObjectID `bson:"_id,omitempty"`
	Username   string        `bson:"username"`
	FullName   string        `bson:"fullName"`
	Email      string        `bson:"email"`
	ProfileURL *string       `bson:"profileUrl,omitempty"`
	CreatedAt  string        `bson:"createdAt"`
	UpdatedAt  string        `bson:"updatedAt"`
}

func ToGraphQLProfile(p *Profile) *model.MyProfile {
	if p == nil {
		return nil
	}

	return &model.MyProfile{
		ID:         p.ID.Hex(),
		Email:      p.Email,
		Username:   p.Username,
		FullName:   p.FullName,
		ProfileURL: p.ProfileURL,
		CreatedAt:  p.CreatedAt,
		UpdatedAt:  p.UpdatedAt,
	}
}

func ToBsonProfile(p *model.Profile) (Profile, error) {
	if p == nil {
		return Profile{}, nil
	}

	var oid bson.ObjectID
	var err error

	if p.ID != "" {
		oid, err = bson.ObjectIDFromHex(p.ID)
		if err != nil {
			return Profile{}, err
		}
	}

	return Profile{
		ID:         oid,
		Username:   p.Username,
		FullName:   p.FullName,
		ProfileURL: p.ProfileURL,
		CreatedAt:  p.CreatedAt,
		UpdatedAt:  p.UpdatedAt,
	}, nil
}

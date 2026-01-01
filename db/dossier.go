package db

import (
	"github.com/Foxtrot-14/FitRang/profile-service/graph/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Dossier struct {
	ID              primitive.ObjectID `bson:"_id,omitempty"`
	Username        string             `bson:"username"`
	FaceType        string             `bson:"faceType"`
	SkinTone        string             `bson:"skinTone"`
	BodyType        string             `bson:"bodyType"`
	Gender          string             `bson:"gender"`
	PreferredColors []string           `bson:"preferredColors"`
	DislikedColors  []string           `bson:"dislikedColors"`
	Viewers			[]string		   `bson:"viewers"`
	Height          *string            `bson:"height,omitempty"`
	Weight          *string            `bson:"weight,omitempty"`
	CreatedAt       string             `bson:"createdAt"`
	UpdatedAt       string             `bson:"updatedAt"`
}

type MyDossier struct {
	ID              primitive.ObjectID `bson:"_id,omitempty"`
	Email           string             `bson:"email"`
	Username        string             `bson:"username"`
	FaceType        string             `bson:"faceType"`
	SkinTone        string             `bson:"skinTone"`
	BodyType        string             `bson:"bodyType"`
	Gender          string             `bson:"gender"`
	PreferredColors []string           `bson:"preferredColors"`
	DislikedColors  []string           `bson:"dislikedColors"`
	Viewers			[]string		   `bson:"viewers"`
	Height          *string            `bson:"height,omitempty"`
	Weight          *string            `bson:"weight,omitempty"`
	CreatedAt       string             `bson:"createdAt"`
	UpdatedAt       string             `bson:"updatedAt"`
}

func ToBsonDossier(m *model.Dossier) (Dossier, error) {
	if m == nil {
		return Dossier{}, nil
	}

	var oid primitive.ObjectID
	var err error

	if m.ID != "" {
		oid, err = primitive.ObjectIDFromHex(m.ID)
		if err != nil {
			return Dossier{}, err
		}
	}

	viewers := m.Viewers
    if viewers == nil {
        viewers = []string{}
    }

	return Dossier{
		ID:              oid,
		Username:		 m.Username,
		FaceType:        m.FaceType,
		SkinTone:        string(m.SkinTone),
		BodyType:        m.BodyType,
		Gender:          string(m.Gender),
		PreferredColors: m.PreferredColors,
		DislikedColors:  m.DislikedColors,
		Viewers:		 viewers,
		Height:          m.Height,
		Weight:          m.Weight,
		CreatedAt:       m.CreatedAt,
		UpdatedAt:       m.UpdatedAt,
	}, nil
}

func ToBsonMyDossier(m *model.MyDossier) (MyDossier, error) {
	if m == nil {
		return MyDossier{}, nil
	}

	var oid primitive.ObjectID
	var err error

	if m.ID != "" {
		oid, err = primitive.ObjectIDFromHex(m.ID)
		if err != nil {
			return MyDossier{}, err
		}
	}

	viewers := m.Viewers
    if viewers == nil {
        viewers = []string{}
    }

	return MyDossier{
		ID:              oid,
		Email:			 m.Email,
		Username:		 m.Username,
		FaceType:        m.FaceType,
		SkinTone:        string(m.SkinTone),
		BodyType:        m.BodyType,
		Gender:          string(m.Gender),
		PreferredColors: m.PreferredColors,
		DislikedColors:  m.DislikedColors,
		Viewers:		 viewers,
		Height:          m.Height,
		Weight:          m.Weight,
		CreatedAt:       m.CreatedAt,
		UpdatedAt:       m.UpdatedAt,
	}, nil
}

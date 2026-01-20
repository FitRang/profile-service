package db

import (
	"github.com/Foxtrot-14/FitRang/profile-service/graph/model"
	"go.mongodb.org/mongo-driver/v2/bson"
)

type Dossier struct {
	ID              bson.ObjectID `bson:"_id,omitempty"`
	Username        string        `bson:"username"`
	FaceType        string        `bson:"faceType"`
	SkinTone        string        `bson:"skinTone"`
	BodyType        string        `bson:"bodyType"`
	Gender          string        `bson:"gender"`
	PreferredColors []string      `bson:"preferredColors"`
	DislikedColors  []string      `bson:"dislikedColors"`
	Viewers         []string      `bson:"viewers"`
	Height          *string       `bson:"height,omitempty"`
	Weight          *string       `bson:"weight,omitempty"`
	CreatedAt       string        `bson:"createdAt"`
	UpdatedAt       string        `bson:"updatedAt"`
}

type MyDossier struct {
	ID              bson.ObjectID `bson:"_id,omitempty"`
	Email           string        `bson:"email"`
	Username        string        `bson:"username"`
	FaceType        string        `bson:"faceType"`
	SkinTone        string        `bson:"skinTone"`
	BodyType        string        `bson:"bodyType"`
	Gender          string        `bson:"gender"`
	PreferredColors []string      `bson:"preferredColors"`
	DislikedColors  []string      `bson:"dislikedColors"`
	Viewers         []string      `bson:"viewers"`
	Height          *string       `bson:"height,omitempty"`
	Weight          *string       `bson:"weight,omitempty"`
	CreatedAt       string        `bson:"createdAt"`
	UpdatedAt       string        `bson:"updatedAt"`
}

func ToBsonDossier(m *model.Dossier) (Dossier, error) {
	if m == nil {
		return Dossier{}, nil
	}

	var oid bson.ObjectID
	var err error

	if m.ID != "" {
		oid, err = bson.ObjectIDFromHex(m.ID)
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
		Username:        m.Username,
		FaceType:        m.FaceType,
		SkinTone:        string(m.SkinTone),
		BodyType:        m.BodyType,
		Gender:          string(m.Gender),
		PreferredColors: m.PreferredColors,
		DislikedColors:  m.DislikedColors,
		Viewers:         viewers,
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

	var oid bson.ObjectID
	var err error

	if m.ID != "" {
		oid, err = bson.ObjectIDFromHex(m.ID)
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
		Email:           m.Email,
		Username:        m.Username,
		FaceType:        m.FaceType,
		SkinTone:        string(m.SkinTone),
		BodyType:        m.BodyType,
		Gender:          string(m.Gender),
		PreferredColors: m.PreferredColors,
		DislikedColors:  m.DislikedColors,
		Viewers:         viewers,
		Height:          m.Height,
		Weight:          m.Weight,
		CreatedAt:       m.CreatedAt,
		UpdatedAt:       m.UpdatedAt,
	}, nil
}

func ToGraphQLDossier(d *Dossier) *model.Dossier {
	if d == nil {
		return nil
	}

	var id string
	if d.ID != bson.NilObjectID {
		id = d.ID.Hex()
	}

	viewers := d.Viewers
	if viewers == nil {
		viewers = []string{}
	}

	return &model.Dossier{
		ID:              id,
		Username:        d.Username,
		FaceType:        d.FaceType,
		SkinTone:        model.SkinTone(d.SkinTone),
		BodyType:        d.BodyType,
		Gender:          model.Gender(d.Gender),
		PreferredColors: d.PreferredColors,
		DislikedColors:  d.DislikedColors,
		Viewers:         viewers,
		Height:          d.Height,
		Weight:          d.Weight,
		CreatedAt:       d.CreatedAt,
		UpdatedAt:       d.UpdatedAt,
	}
}

func ToGraphQLMyDossier(d *MyDossier) *model.MyDossier {
	if d == nil {
		return nil
	}

	var id string
	if d.ID != bson.NilObjectID {
		id = d.ID.Hex()
	}

	viewers := d.Viewers
	if viewers == nil {
		viewers = []string{}
	}

	return &model.MyDossier{
		ID:              id,
		Email:           d.Email,
		Username:        d.Username,
		FaceType:        d.FaceType,
		SkinTone:        model.SkinTone(d.SkinTone),
		BodyType:        d.BodyType,
		Gender:          model.Gender(d.Gender),
		PreferredColors: d.PreferredColors,
		DislikedColors:  d.DislikedColors,
		Viewers:         viewers,
		Height:          d.Height,
		Weight:          d.Weight,
		CreatedAt:       d.CreatedAt,
		UpdatedAt:       d.UpdatedAt,
	}
}

package model

import (
	"fmt"
	"slices"
)

type Gender    string
type SkinTone  string

const (
	GenderMale   Gender = "male"
	GenderFemale Gender = "female"
)

const (
	Pale  SkinTone = "pale"
	Light SkinTone = "light"
	Olive SkinTone = "olive"
	Dark  SkinTone = "dark"
)

type CreateStyleDossierRequest struct {
	FaceType         string		`json:"face_type" binding:"required"`
	SkinTone         SkinTone	`json:"skin_tone" binding:"required"`
	BodyType         string		`json:"body_type" binding:"required"`
	Gender           Gender		`json:"gender" binding:"required"`
	PreferredColors  *[]string	`json:"preferred_colors,omitempty"`
	DislikedColors   *[]string  `json:"disliked_colors,omitempty"`
	Height           *string    `json:"height,omitempty"`
	Weight           *string    `json:"weight,omitempty"`
}

func (r *CreateStyleDossierRequest) Validate() error {
	if r.Gender != GenderFemale && r.Gender != GenderMale {
		return fmt.Errorf("invalid gender: %s", r.Gender)
	}
	validBodyTypes := map[Gender][] string{
		GenderFemale : {"apple","pear","rectangular","hourglass"},
		GenderMale : {"rectangular","triangle","trapezoid","oval","invert-triangle"},
	}
	bodyTypes := validBodyTypes[r.Gender]
	if slices.Contains(bodyTypes, r.BodyType) {
		return fmt.Errorf("invalid body type '%s' for gender '%s'", r.BodyType, r.Gender)
	}
	validFaceTypes := map[Gender][] string{
		GenderFemale : {"oval","heart","diamond","square","round","oblong"},
		GenderMale : {"rectangular","round","oblong","triangular","heart"},
	}
	faceTypes := validFaceTypes[r.Gender]
	if slices.Contains(faceTypes, r.FaceType) {
		return fmt.Errorf("invalid body type '%s' for gender '%s'", r.FaceType, r.Gender)
	}
	return nil
}

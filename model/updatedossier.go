package model

import (
	"fmt"
	"slices"
)

type UpdateDossierRequest struct {
	DossierID		 string   `json:"id" binding:"required, uuid"`
	FaceType         string	  `json:"face_type" binding:"omitempty"`
	SkinTone         Skintone `json:"skin_tone" binding:"omitempty"`
	BodyType         string	  `json:"body_type" binding:"omitempty"`
	Gender           Gender	  `json:"gender" binding:"omitempty"`
	PreferredColors  []string `json:"preferred_colors,omitempty"`
	DislikedColors   []string `json:"disliked_colors,omitempty"`
	Height           string   `json:"height,omitempty"`
	Weight           string   `json:"weight,omitempty"`
}

type UpdateDossierResponse struct {
	Dossier Dossier `json:"dossier"`
}

func (r *UpdateDossierRequest) Validate() error {
	if r.Gender != GenderFemale && r.Gender != GenderMale {
		return fmt.Errorf("invalid gender: %s", r.Gender)
	}
	validBodyTypes := map[Gender][] string{
		GenderFemale : {"apple","pear","rectangular","hourglass"},
		GenderMale : {"rectangular","triangle","trapezoid","oval","invert-triangle"},
	}
	bodyTypes := validBodyTypes[r.Gender]
	if !slices.Contains(bodyTypes, r.BodyType) {
		return fmt.Errorf("invalid body type '%s' for gender '%s'", r.BodyType, r.Gender)
	}
	validFaceTypes := map[Gender][] string{
		GenderFemale : {"oval","heart","diamond","square","round","oblong"},
		GenderMale : {"rectangular","round","oblong","triangular","heart"},
	}
	faceTypes := validFaceTypes[r.Gender]
	if !slices.Contains(faceTypes, r.FaceType) {
		return fmt.Errorf("invalid face type '%s' for gender '%s'", r.FaceType, r.Gender)
	}
	return nil
}

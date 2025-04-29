package model

type Gender string

const (
	GenderMale Gender = "male"
	GenderFemale Gender = "female"
)

type CreateStyleDossierRequest struct {
	FaceType         string		`json:"face_type" binding:"required"`
	SkinTone         string		`json:"skin_tone" binding:"required"`
	BodyType         string		`json:"body_type" binding:"required"`
	Gender           Gender		`json:"gender" binding:"required"`
	PreferredColors  *[]string	`json:"preferred_colors,omitempty"`
	DislikedColors   *[]string  `json:"disliked_colors,omitempty"`
	Height           *string    `json:"height,omitempty"`
	Weight           *string    `json:"weight,omitempty"`
}

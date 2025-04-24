package domain

import (
	"database/sql"
	"github.com/FitRang/profile-service/model"
)

type ProfileServiceInterface interface {
	CreateProfile(profile *model.ProfileCreateRequest) error
	GetProfile(profileID string) (*model.ProfileGetResponse, error)
	UpdateProfile(profile *model.ProfileUpdateRequest) (*model.ProfileUpdateResponse, error)
}

type ProfileService struct {
	db *sql.DB
}

func NewProfileService(db *sql.DB) *ProfileService{
	return &ProfileService{
		db : db,
	}
}

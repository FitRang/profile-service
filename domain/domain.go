package domain

import (
	"database/sql"
	"github.com/FitRang/profile-service/model"
)

type ProfileServiceInterface interface {
	CreateProfile(profile *model.ProfileCreateRequest) error
	GetProfile(profileID string) (*model.ProfileGetResponse, error)
}

type ProfileService struct {
	db *sql.DB
}

func NewProfileService(db *sql.DB) *ProfileService{
	return &ProfileService{
		db : db,
	}
}

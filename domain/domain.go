package domain

import (
	"database/sql"
	"github.com/FitRang/profile-service/model"
)

type ProfileServiceInterface interface {
	CreateProfile(profile *model.ProfileCreateRequest) error
	GetProfile(profileID string) (*model.ProfileGetResponse, error)
	UpdateProfile(profile *model.ProfileUpdateRequest) (*model.ProfileUpdateResponse, error)
	CreateDossier(dossier *model.CreateStyleDossierRequest) error
	GetDossier(dossierID string) (*model.Dossier, error)
	GetAllDossiers(userID string) (*model.GetAllDossierRequest, error)
	UpdateDossier(req *model.UpdateDossierRequest) (*model.Dossier, error)
	UpdateDossierAccess(update *model.GrantAccessRequest) error
	IsDossierOwner(dossierID, profileID string) (bool, error)
}

type ProfileService struct {
	db *sql.DB
}

func NewProfileService(db *sql.DB) *ProfileService{
	return &ProfileService{
		db : db,
	}
}

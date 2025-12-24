package dossier

import (
	"github.com/Foxtrot-14/FitRang/profile-service/repository"
)

type DossierService struct {
	DossierRepo *repository.DossierRepository
	ProfileRepo *repository.ProfileRepository
}

func NewDossierService(d *repository.DossierRepository, p *repository.ProfileRepository) *DossierService {
	return &DossierService{
		DossierRepo: d,
		ProfileRepo: p,
	}
}

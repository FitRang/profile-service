package profile

import (
	"github.com/Foxtrot-14/FitRang/profile-service/eventbus"
	"github.com/Foxtrot-14/FitRang/profile-service/repository"
)

type ProfileService struct {
	ProfileRepo *repository.ProfileRepository
	DossierRepo *repository.DossierRepository
	AccessRepo  *repository.AccessRepository
	Bus         *eventbus.Producer
}

func NewProfileService(
	r *repository.ProfileRepository,
	b *eventbus.Producer,
	d *repository.DossierRepository,
	a *repository.AccessRepository,
) *ProfileService {
	return &ProfileService{
		ProfileRepo: r,
		Bus:         b,
		DossierRepo: d,
		AccessRepo:  a,
	}
}

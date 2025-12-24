package access

import (
	"github.com/Foxtrot-14/FitRang/profile-service/eventbus"
	"github.com/Foxtrot-14/FitRang/profile-service/repository"
)

type AccessService struct {
	AccessRepo  *repository.AccessRepository
	DossierRepo *repository.DossierRepository
	ProfileRepo *repository.ProfileRepository
	Bus         *eventbus.Producer
}

func NewAccessService(
	r *repository.AccessRepository,
	d *repository.DossierRepository,
	p *repository.ProfileRepository,
	b *eventbus.Producer,
) *AccessService {
	return &AccessService{
		AccessRepo:  r,
		DossierRepo: d,
		ProfileRepo: p,
		Bus:         b,
	}
}

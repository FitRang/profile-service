package profile

import (
	"github.com/Foxtrot-14/FitRang/profile-service/eventbus"
	"github.com/Foxtrot-14/FitRang/profile-service/repository"
)

type ProfileService struct {
	Repo *repository.ProfileRepository
	Bus  *eventbus.Producer
}

func NewProfileService(r *repository.ProfileRepository, b *eventbus.Producer) *ProfileService {
	return &ProfileService{
		Repo: r,
		Bus:  b,
	}
}

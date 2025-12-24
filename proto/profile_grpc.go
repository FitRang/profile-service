package proto

import (
	"github.com/Foxtrot-14/FitRang/profile-service/cache"
	"github.com/Foxtrot-14/FitRang/profile-service/services/dossier"
	"github.com/Foxtrot-14/FitRang/profile-service/services/profile"
)

type ProfileGRPCService struct {
	UnimplementedProfileServiceServer
	profileService *profile.ProfileService
	dossierService *dossier.DossierService
	rdb            *cache.RedisClient
}

func NewProfileGRPCService(p *profile.ProfileService, d *dossier.DossierService, r *cache.RedisClient) *ProfileGRPCService {
	return &ProfileGRPCService{
		profileService: p,
		dossierService: d,
		rdb:            r,
	}
}

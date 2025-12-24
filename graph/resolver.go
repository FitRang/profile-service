package graph

//go:generate go tool gqlgen generate
import (
	"github.com/Foxtrot-14/FitRang/profile-service/services/access"
	"github.com/Foxtrot-14/FitRang/profile-service/services/dossier"
	"github.com/Foxtrot-14/FitRang/profile-service/services/profile"
)

type Resolver struct{
	ProfileService  *profile.ProfileService
	DossierService  *dossier.DossierService
	AccessService   *access.AccessService
}

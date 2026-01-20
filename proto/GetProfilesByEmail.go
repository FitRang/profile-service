package proto

import (
	"context"
	"log"

	"github.com/Foxtrot-14/FitRang/profile-service/db"
	"github.com/Foxtrot-14/FitRang/profile-service/graph/model"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *ProfileGRPCService) GetProfilesByEmail(
	ctx context.Context,
	req *GetByEmailRequest,
) (*GetProfilesResponse, error) {
	email := req.GetEmail()
	if email == "" {
		return nil, status.Error(codes.InvalidArgument, "email is required")
	}

	profiles, dossiers, err := s.dossierService.GetProfilesByEmail(ctx, email)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	for _, dossier := range dossiers {
		dossierJSON := db.ToGraphQLDossier(&dossier)
		if err := s.rdb.StoreDossier(ctx, dossierJSON.Username, *dossierJSON); err != nil {
			log.Printf(
				"[WARN] failed to cache dossier %s: %v",
				dossier.Username,
				err,
			)
		}
	}

	var profilesJSON []model.MyProfile
	for _, profile := range profiles {
		profileJSON := db.ToGraphQLProfile(&profile)
		profilesJSON = append(profilesJSON, *profileJSON)
	}

	return &GetProfilesResponse{
		Profile: toProtoProfiles(profilesJSON),
	}, nil
}

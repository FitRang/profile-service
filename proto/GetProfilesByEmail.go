package proto

import (
	"context"
	"log"

	"github.com/Foxtrot-14/FitRang/profile-service/db"
	"github.com/Foxtrot-14/FitRang/profile-service/graph/model"
	"go.mongodb.org/mongo-driver/v2/bson"
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

	res := s.profileService.ProfileRepo.Col.FindOne(
		ctx,
		bson.M{"email": email},
	)

	var profile db.Profile
	if err := res.Decode(&profile); err != nil {
		log.Printf("[Error] while returning user profile: %v", err)
	}

	myProfile := db.ToGraphQLProfile(&profile)

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
	profilesJSON = append(profilesJSON, *myProfile)
	return &GetProfilesResponse{
		Profile: toProtoProfiles(profilesJSON),
	}, nil
}

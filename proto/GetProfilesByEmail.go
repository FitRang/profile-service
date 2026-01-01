package proto

import (
	"context"
	"log"

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
		if err := s.rdb.StoreDossier(ctx, dossier.Username, dossier); err != nil {
			log.Printf(
				"[WARN] failed to cache dossier %s: %v",
				dossier.Username,
				err,
			)
		}
	}

	return &GetProfilesResponse{
		Profile: toProtoProfiles(profiles),
	}, nil
}

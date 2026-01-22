package proto

import (
	"context"

	"github.com/Foxtrot-14/FitRang/profile-service/graph/model"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *ProfileGRPCService) GetDossierByUsername(
	ctx context.Context,
	req *GetByUsernameRequest,
) (*GetDossierResponse, error) {
	username := req.GetUsername()
	if username == "" {
		return nil, status.Error(codes.InvalidArgument, "username is required")
	}

	if cached, err := s.rdb.GetDossier(ctx, username); err == nil && cached != nil {
		return &GetDossierResponse{
			Dossier: toProtoDossier(*cached),
		}, nil
	}

	var dossier model.Dossier
	err := s.dossierService.DossierRepo.Col.FindOne(
		ctx,
		bson.M{"username": username},
	).Decode(&dossier)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, status.Error(codes.NotFound, "dossier not found")
		}
		return nil, status.Error(codes.Internal, err.Error())
	}

	_ = s.rdb.StoreDossier(ctx, username, dossier)

	return &GetDossierResponse{
		Dossier: toProtoDossier(dossier),
	}, nil
}

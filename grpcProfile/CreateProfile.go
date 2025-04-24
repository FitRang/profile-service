package grpcProfile

import (
	"context"

	"github.com/FitRang/profile-service/apperror"
	"github.com/FitRang/profile-service/pb"
) 

func (s *GrpcServer) CreateProfile(ctx context.Context, in *pb.CreateRequest) (*pb.CreateResponse, error)  {
	profile,validationErr := apperror.ValidateGrpcCreateRequest(in)
	if validationErr != nil {
		s.logger.Info().Str("method", "CreateProfile").Str("error",validationErr.Error()).Msg("Failed to create profile")
		return nil, validationErr
	}
	if err := s.domain.CreateProfile(profile); err != nil {
		s.logger.Info().Str("method", "CreateProfile").Str("error",validationErr.Error()).Msg("Failed to create profile")
		return nil, err;
	}
	s.logger.Info().Str("method", "CreateProfile").Str("user_id", profile.ID).Msg("Created Profile Successfully")
	return &pb.CreateResponse{
		Id: profile.ID,
		Status: "SUCCESS",
	}, nil
}

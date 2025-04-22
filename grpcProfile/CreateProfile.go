package grpcProfile

import (
	"context"

	"github.com/FitRang/profile-service/apperror"
	"github.com/FitRang/profile-service/pb"
) 

func (s *GrpcServer) CreateProfile(ctx context.Context, in *pb.CreateRequest) (*pb.CreateResponse, error)  {
	profile,validationErr := apperror.ValidateGrpcCreateRequest(in)
	if validationErr != nil {
		return nil, validationErr
	}
	if err := s.domain.CreateProfile(profile); err != nil {
		return nil, err;
	}
	return &pb.CreateResponse{
		Id: profile.ID,
		Status: "SUCCESS",
	}, nil
}

package grpcProfile

import (
	"context"
	"errors"

	"github.com/google/uuid"

	"github.com/FitRang/profile-service/domain"
	"github.com/FitRang/profile-service/pb"
)

func (s *GrpcServer) GetProfile(ctx context.Context, in *pb.GetRequest) (*pb.GetResponse, error) {
	if _, err := uuid.Parse(in.Id); err != nil {
		s.logger.Info().Str("method", "GetProfile").Str("user_id", in.Id).Msg("Invalid Id provided")
		return &pb.GetResponse{
			Result: &pb.GetResponse_Error{
				Error: &pb.Error{
					Code: "Invalid Id",
				},
			},
		}, nil
	}
	profile, err := s.domain.GetProfile(in.Id)
	if err != nil {
		if errors.Is(err, domain.ErrProfileNotFound) {
			s.logger.Info().Str("method", "GetProfile").Str("user_id", in.Id).Msg("Profile Not Found")
			return &pb.GetResponse{
				Result: &pb.GetResponse_Error{
					Error: &pb.Error{
						Code:    "NOT_FOUND",
						Message: "Profile does not exist",
					},
				},
			}, nil
		}
	}
	s.logger.Info().Str("method", "GetProfile").Str("user_id", in.Id).Msg("Profile found")
	return &pb.GetResponse{
		Result: &pb.GetResponse_Profile{
			Profile: &pb.Profile{
				Id:    profile.ID,
				Name:  profile.FullName,
				Email: profile.Email,
				PhoneNumber: profile.PhoneNumber,
				CreatedAt: profile.CreatedAT,
				UpdatedAt: profile.UpdatedAT,
			},
		},
	}, nil
}

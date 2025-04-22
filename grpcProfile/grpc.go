package grpcProfile

import (
	"database/sql"

	"github.com/FitRang/profile-service/domain"
	"github.com/FitRang/profile-service/pb"

	"google.golang.org/grpc"
)

type ProfileService interface {
	pb.ProfileServiceServer
}

type GrpcServer struct {
	pb.UnimplementedProfileServiceServer
	domain domain.ProfileService
}

func NewGrpcService(s *grpc.Server, db *sql.DB) {
	service := &GrpcServer{
		domain: *domain.NewProfileService(db),
	}
	pb.RegisterProfileServiceServer(s,service)
}

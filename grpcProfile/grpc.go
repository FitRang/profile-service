package grpcProfile

import (
	"database/sql"
	"os"

	"github.com/FitRang/profile-service/domain"
	"github.com/FitRang/profile-service/pb"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	"google.golang.org/grpc"
)

type ProfileService interface {
	pb.ProfileServiceServer
}

type GrpcServer struct {
	pb.UnimplementedProfileServiceServer
	domain domain.ProfileService
	logger zerolog.Logger
}

func NewGrpcService(s *grpc.Server, db *sql.DB) {
	logger := log.Output(zerolog.ConsoleWriter{Out: os.Stdout})

	service := &GrpcServer{
		domain: *domain.NewProfileService(db),
		logger: logger,
	}

	pb.RegisterProfileServiceServer(s, service)
}

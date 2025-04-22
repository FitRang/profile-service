package apperror

import (
	"errors"
	"regexp"

	"github.com/FitRang/profile-service/model"
	"github.com/FitRang/profile-service/pb"
	"github.com/google/uuid"
)

func ValidateGrpcCreateRequest(req *pb.CreateRequest) (*model.ProfileCreateRequest, error) {
	
	// Validate uuid
	if _, err := uuid.Parse(req.Id); err != nil {
		return nil, errors.New("invalid UUID format for ID")
	}		

	//Validate Email
	if !regexp.MustCompile(`^[^@]+@[^@]+\.[^@]+$`).MatchString(req.Email) {
		return nil, errors.New("invalid email format")
	}

	// Validate FullName (non-empty)
	if req.Name == "" {
		return nil, errors.New("full_name is required")
	}

	// Validate PhoneNumber (E.164)
	if !regexp.MustCompile(`^\+[1-9]\d{1,14}$`).MatchString(req.PhoneNumber) {
		return nil, errors.New("phone_number must be in E.164 format (e.g., +1234567890)")
	}

	return &model.ProfileCreateRequest{
		ID: req.Id,
		Email: req.Email,
		FullName: req.Name,
		PhoneNumber: req.PhoneNumber,
	}, nil
}


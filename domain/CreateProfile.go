package domain

import (
	"errors"
	"log"

	"github.com/FitRang/profile-service/model"
	"github.com/lib/pq"
)

const (
	// postgresUniqueConstraintViolationCode error code for unique constraint violation.
	postgresUniqueConstraintViolationCode = "23505"
)

var (
	ErrIDAlreadyExists = errors.New("a profile with ID already exists")
	ErrEmailAlreadyExists = errors.New("a profile with Email already exists")
	ErrPhoneNumberAlreadyExists = errors.New("a profile with Phone Number already exists")
	ErrProfileNotFound      = errors.New("requested profile not found")
	ErrProfileCreation      = errors.New("failed to create profile")
	ErrProfileGet           = errors.New("failed to get profile")
)

func (ps *ProfileService) CreateProfile(profile *model.ProfileCreateRequest) error {
	sqlStatement := `INSERT INTO "profile" VALUES ($1)`

	if _, err := ps.db.Exec(
		sqlStatement,
		profile.ID,
		profile.Email,
		profile.FullName,
		profile.PhoneNumber,
	); err != nil {
		pqErr, ok := err.(*pq.Error)
		if ok {
			if pqErr.Code == postgresUniqueConstraintViolationCode {
				switch pqErr.Constraint {
				case "profile_id_key":
					return ErrIDAlreadyExists
				case "profile_email_key":
					return ErrEmailAlreadyExists
				case "profile_phone_number":
					return ErrPhoneNumberAlreadyExists
				default:
					log.Println("[ERROR:CreateProfile]:", pqErr.Message)
					return ErrProfileCreation
				}
			}
			log.Println("[ERROR:CreateProfile]:", pqErr.Message)
			return ErrProfileCreation
		}
		log.Println("[ERROR:CreateProfile]:", err.Error())
		return ErrProfileCreation
	}
	return nil
}

package domain

import (
	"errors"
	"fmt"
	"log"

	"github.com/FitRang/profile-service/model"
	"github.com/lib/pq"
)

var (
	ErrDossierAlreadyExists = errors.New("the user already has a dossier")
	ErrDossierCreation      = errors.New("failed to create dossier")
	ErrDossierNotFound      = errors.New("dossier does not exist")
	ErrDossiersNotFound     = errors.New("no shared dossiers")
	ErrGetDossier			= errors.New("failed to get dossier")
)

func (ps *ProfileService) CreateDossier(dossier *model.CreateStyleDossierRequest) error {
	const insertDossierSQL = `
		INSERT INTO dossier (
			owner_id,
			face_type,
			skin_tone,
			body_type,
			gender,
			preferred_colors,
			disliked_colors,
			height,
			weight
		) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
	`
	_, err := ps.db.Exec(
		insertDossierSQL,
		dossier.OwnerID,
		dossier.FaceType,
		dossier.SkinTone,
		dossier.BodyType,
		dossier.Gender,
		dossier.PreferredColors,
		dossier.DislikedColors,
		dossier.Height,
		dossier.Weight,
	)
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok {
			if pqErr.Code == postgresUniqueConstraintViolationCode {
				return ErrDossierAlreadyExists
			}
			log.Printf("[ERROR:CreateDossier] Postgres error: %s (Code: %s)\n", pqErr.Message, pqErr.Code)
			return ErrDossierCreation
		}
		log.Printf("[ERROR:CreateDossier] Unexpected error: %v\n", err)
		return ErrDossierCreation
	}
	return nil
}


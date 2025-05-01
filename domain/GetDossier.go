package domain

import (
	"database/sql"
	"errors"
	"log"

	"github.com/FitRang/profile-service/model"
)

func (ps *ProfileService) GetDossier (dossierID string) (*model.Dossier, error)  {
	dossier := model.Dossier{}
	sqlStatement := `SELECT * FROM TABLE "dossier" WHERE id = $1`
	if err := ps.db.QueryRow(sqlStatement,dossierID).Scan(
		&dossier.Id,
		&dossier.OwnerID,
		&dossier.BodyType,
		&dossier.FaceType,
		&dossier.SkinTone,
		&dossier.PreferredColors,
		&dossier.DislikedColors,
		&dossier.Gender,
		&dossier.Height,
		&dossier.Weight,
	); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrDossierNotFound
		}
		log.Println("[ERROR:GETDossier]:",err.Error())
		return nil, ErrGetDossier
	}
	return &dossier, nil
}

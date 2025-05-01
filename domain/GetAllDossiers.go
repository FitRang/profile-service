package domain

import (
	"database/sql"
	"errors"
	"log"

	"github.com/FitRang/profile-service/model"
	"github.com/lib/pq"
)

func (ps *ProfileService) GetAllDossiers (userID string) (*model.GetAllDossierResponse, error) {
	dossiers := model.GetAllDossierResponse{}
	sqlStatement := `
		SELECT d.*
		FROM dossier_access da
		JOIN dossier d ON da.dossier_id = d.id
		WHERE da.profile_id = $1;
	`
	rows, err := ps.db.Query(sqlStatement, userID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrDossiersNotFound
		}
		log.Println("[ERRORGETDossiers]", err.Error())
		return nil, ErrGetDossier
	}

	defer rows.Close()
	for rows.Next() {
		var d model.Dossier
		var preferredColors, dislikedColors pq.StringArray

		err := rows.Scan(
			&d.Id,
			&d.OwnerID,
			&d.FaceType,
			&d.SkinTone,
			&d.BodyType,
			&d.Gender,
			&preferredColors,
			&dislikedColors,
			&d.Height,
			&d.Weight,
		)

		if err != nil {
			log.Println("[ERROR] scanning row:", err)
			return nil, ErrGetDossier
		}

		d.PreferredColors = preferredColors
		d.DislikedColors = dislikedColors
		dossiers.Dossiers = append(dossiers.Dossiers, d)

	}
	
	if err = rows.Err(); err != nil {
		log.Println("[ERROR] iterating rows:", err)
		return nil, ErrGetDossier
	}

	return &dossiers, nil
}

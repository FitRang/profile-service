package domain

import (
	"database/sql"

	"github.com/FitRang/profile-service/model"
)

func (ps *ProfileService) UpdateDossierAccess(update *model.GrantAccessRequest) error {
	sqlStatement := `INSERT INTO dossier_access (dossier_id, profile_id) VALUES ($1, $2)`
	_, err := ps.db.Exec(sqlStatement, update.DossierID, update.UserId)
	if err != nil {
		return err
	}
	return nil
}

func (ps *ProfileService) IsDossierOwner(dossierID, profileID string) (bool, error) {
	var ownerID string
	err := ps.db.QueryRow(`SELECT owner_id FROM dossier WHERE id = $1`, dossierID).Scan(&ownerID)
	if err != nil {
		if err == sql.ErrNoRows {
			return false, nil
		}
		return false, err
	}
	return ownerID == profileID, nil
}

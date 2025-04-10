package domain

import (
	"database/sql"
	"log"

	"github.com/FitRang/profile-service/model"
	"errors"
)

func (ps *ProfileService) GetProfile(profileID string) (*model.ProfileGetResponse, error) {
	profile := model.ProfileGetResponse{}
	sqlStatement := `
	SELECT * FROM "profile"
	WHERE id = $1
	`
	if err := ps.db.QueryRow(sqlStatement, profileID).Scan(
		&profile.ID,
		&profile.Email,
		&profile.FullName,
		&profile.PhoneNumber,
		&profile.CreatedAt,
		); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrProfileNotFound
		}
		log.Println("[ERROR:GETProfile]:",err.Error())
		return nil, ErrProfileGet
	}
	return &profile, nil
}

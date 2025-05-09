package domain

import (
	"errors"
	"fmt"
	"strings"

	"github.com/FitRang/profile-service/model"
)

func (ps *ProfileService) UpdateProfile(profile *model.ProfileUpdateRequest) (*model.ProfileUpdateResponse, error) {
	setClauses := []string{}
	args := []any{}
	argIdx := 1

	if profile.Email != "" {
		setClauses = append(setClauses, fmt.Sprintf("email = $%d", argIdx))
		args = append(args, profile.Email)
		argIdx++
	}
	if profile.FullName != "" {
		setClauses = append(setClauses, fmt.Sprintf("full_name = $%d", argIdx))
		args = append(args, profile.FullName)
		argIdx++
	}
	if profile.PhoneNumber != "" {
		setClauses = append(setClauses, fmt.Sprintf("phone_number = $%d", argIdx))
		args = append(args, profile.PhoneNumber)
		argIdx++
	}

	if len(setClauses) == 0 {
		return nil, errors.New("no fields to update")
	}

	setClauses = append(setClauses, "updated_at = NOW()")

	args = append(args, profile.ID)

	sqlStatement := fmt.Sprintf(`
		UPDATE profile 
		SET %s 
		WHERE id = $%d 
		RETURNING id, email, full_name, phone_number, created_at, updated_at
		`, strings.Join(setClauses, ", "), argIdx)

	resProfile := model.ProfileUpdateResponse{}
	err := ps.db.QueryRow(sqlStatement, args...).Scan(
		&resProfile.ID,
		&resProfile.Email,
		&resProfile.FullName,
		&resProfile.PhoneNumber,
		&resProfile.CreatedAT,
		&resProfile.UpdatedAT,
	)
	if err != nil {
		return nil, err
	}

	return &resProfile, nil
}

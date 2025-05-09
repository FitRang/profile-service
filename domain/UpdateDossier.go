package domain

import (
	"errors"
	"fmt"
	"log"
	"strings"

	"github.com/FitRang/profile-service/model"
	"github.com/lib/pq"
)


func (ps *ProfileService) UpdateDossier(req *model.UpdateDossierRequest) (*model.Dossier, error) {
	setClauses := []string{}
	args	   := []any{}
	argIndex   := 1

	if req.FaceType != "" {
		setClauses = append(setClauses, fmt.Sprintf(`"face_type" = $%d`, argIndex))
		args = append(args, req.FaceType)
		argIndex++
	}
	if req.SkinTone != "" {
		setClauses = append(setClauses, fmt.Sprintf(`"skin_tone" = $%d`, argIndex))
		args = append(args, req.SkinTone)
		argIndex++
	}
	if req.BodyType != "" {
		setClauses = append(setClauses, fmt.Sprintf(`"body_type" = $%d`, argIndex))
		args = append(args, req.BodyType)
		argIndex++
	}
	if req.Gender != "" {
		setClauses = append(setClauses, fmt.Sprintf(`"gender" = $%d`, argIndex))
		args = append(args, req.Gender)
		argIndex++
	}
	if req.PreferredColors != nil {
		setClauses = append(setClauses, fmt.Sprintf(`"preferred_colors" = $%d`, argIndex))
		args = append(args, pq.Array(req.PreferredColors))
		argIndex++
	}
	if req.DislikedColors != nil {
		setClauses = append(setClauses, fmt.Sprintf(`"disliked_colors" = $%d`, argIndex))
		args = append(args, pq.Array(req.DislikedColors))
		argIndex++
	}
	if req.Height != "" {
		setClauses = append(setClauses, fmt.Sprintf(`"height" = $%d`, argIndex))
		args = append(args, req.Height)
		argIndex++
	}
	if req.Weight != "" {
		setClauses = append(setClauses, fmt.Sprintf(`"weight" = $%d`, argIndex))
		args = append(args, req.Weight)
		argIndex++
	}

	if len(setClauses) == 0 {
		return nil, errors.New("no fields to update")
	}

	args = append(args, req.DossierID)
	query := fmt.Sprintf(`
		UPDATE dossier
		SET %s
		WHERE id = $%d
		RETURNING id, owner_id, face_type, skin_tone, body_type, gender, preferred_colors, disliked_colors, height, weight
	`, strings.Join(setClauses, ", "), argIndex)

	var dossier model.Dossier
	err := ps.db.QueryRow(query, args...).Scan(
		&dossier.Id,
		&dossier.OwnerID,
		&dossier.FaceType,
		&dossier.SkinTone,
		&dossier.BodyType,
		&dossier.Gender,
		pq.Array(&dossier.PreferredColors),
		pq.Array(&dossier.DislikedColors),
		&dossier.Height,
		&dossier.Weight,
	)
	if err != nil {
		log.Println("[ERROR:UpdateDossier]:", err)
		return nil, fmt.Errorf("failed to update dossier: %w", err)
	}

	return &dossier, nil
}


package model

// Sent from a different user
type RequestAccessRequest struct {
	DossierID string `json:"dossier_id" binding:"required, uuid"`
	UserId    string `json:"user_id" binding:"required, uuid"`
}

// Granted by the owner
type GrantAccessRequest struct {
	DossierID string `json:"dossier_id" binding:"required, uuid"`
	UserId    string `json:"user_id" binding:"required, uuid"`
}

// Removed by the owner
type RemoveAccessRequest struct {
	DossierID string `json:"dossier_id" binding:"required, uuid"`
	UserId    string `json:"user_id" binding:"required, uuid"`
}

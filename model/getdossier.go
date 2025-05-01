package model

type Dossier struct {
	Id				 string
	OwnerID			 string
	FaceType         string
	SkinTone         string
	BodyType         string		
	Gender           string		
	PreferredColors  []string
	DislikedColors   []string
	Height           *string
	Weight           *string
}

type GetDossierRequest struct {
	ID  string `json:"id" binding:"required, uuid"`
}

type GetAllDossierRequest struct {
	UserID string `json:"user_id" binding:"required, uuid"`
}

type GetAllDossierResponse struct {
	Dossiers []Dossier `json:"dossiers"`
}

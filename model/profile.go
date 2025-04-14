package model

type ProfileCreateRequest struct {
	ID          string `json:"id" binding:"required,uuid"`
	Email       string `json:"email" binding:"required,email"`
	FullName    string `json:"full_name" binding:"required"`
	PhoneNumber string `json:"phone_number" binding:"required,e164"`
}

type ProfileGetResponse struct {
	ID          string `json:"id"`
	Email       string `json:"email"`
	FullName    string `json:"full_name"`
	PhoneNumber string `json:"phone_number"`
}

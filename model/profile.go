package model

type ProfileCreateRequest struct {
	ID          string `json:"id" validate:"required,uuid4"`
	Email       string `json:"email" validate:"required,email"`
	FullName    string `json:"full_name" validate:"required"`
	PhoneNumber string `json:"phone_number" validate:"required,e164"`
}

type ProfileGetResponse struct {
	ID          string `json:"id"`
	Email       string `json:"email"`
	FullName    string `json:"full_name"`
	PhoneNumber string `json:"phone_number"`
	CreatedAt   string `json:"created_at"`
}

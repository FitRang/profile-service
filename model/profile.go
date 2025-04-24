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
	CreatedAT   string `json:"create_at"`
}

type ProfileUpdateRequest struct {
	ID			 string `json:"id" binding:"required,uuid"`
	Email		*string `json:"email" binding:"omitempty,email"`
	FullName    *string `json:"full_name" binding:"omitempty"`
	PhoneNumber *string `json:"phone_number" binding:"omitempty,e164"`
}

type ProfileUpdateResponse struct {
	ID          string `json:"id"`
	Email       string `json:"email"`
	FullName    string `json:"full_name"`
	PhoneNumber string `json:"phone_number"`
	CreatedAT   string `json:"created_at"`
	UpdatedAT	string `json:"updated_at"`
}

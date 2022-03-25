package dto

import "github.com/google/uuid"

type CreateUserRequest struct {
	PersonalNumber string `json:"personal_number" `
	Password       string `json:"password" `
	Email          string `json:"email" `
	Name           string `json:"name"`
}

type GetUserDetail struct {
	ID uuid.UUID `json:"user_id"`
}

package dto

import (
	"github.com/google/uuid"
)

type CreateUserRequest struct {
	PersonalNumber string `json:"personal_number" `
	Password       string `json:"password" `
	Email          string `json:"email" `
	Name           string `json:"name"`
}

type UpdateUserRequest struct {
	PersonalNumber string `json:"personal_number" `
	Password       string `json:"password" `
	Email          string `json:"email" `
	Name           string `json:"name"`
	Role           RoleID `json:"role"`
	Active         bool   `json:"active"`
}

type RoleID struct {
	ID uuid.UUID `json:"id"`
}

type CreateProductRequest struct {
	Name        string `json:"name"`
	Description string `json:"description" `
}

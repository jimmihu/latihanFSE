package dto

import (
	"latihanFSE/models/entity"

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
	Name        string            `json:"name"`
	Description string            `json:"description" `
	MakerID     uuid.UUID         `json:"-"`
	Maker       entity.UserResult `gorm:"foreignKey:MakerID"`
}

type UpdateProductRequest struct {
	Name        string `json:"name"`
	Description string `json:"description" `
}

type CheckProductRequest struct {
	Name        string            `json:"name"`
	Description string            `json:"description" `
	CheckerID   uuid.UUID         `json:"-"`
	Checker     entity.UserResult `gorm:"foreignKey:CheckerID"`
	Status      string            `json:"status"`
}

type PublishProductRequest struct {
	Name        string            `json:"name"`
	Description string            `json:"description" `
	SignerID    uuid.UUID         `json:"-"`
	Signer      entity.UserResult `gorm:"foreignKey:SignerID"`
	Status      string            `json:"status"`
}
type LoginRequest struct {
	PersonalNumber string `json:"personal_number" `
	Password       string `json:"password" `
}

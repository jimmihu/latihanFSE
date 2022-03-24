package dto

type CreateUserRequest struct {
	PersonalNumber string `json:"personalNumber" `
	Password       string `json:"password" `
	Email          string `json:"email" `
	Name           string `json:"name"`
}

package dto

type CreateUserRequest struct {
	PersonalNumber string `json:"personal_number" `
	Password       string `json:"password" `
	Email          string `json:"email" `
	Name           string `json:"name"`
}
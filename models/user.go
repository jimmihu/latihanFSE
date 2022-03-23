package models

type User struct {
	ID              string
	Personal_number string
	Password        string
	Email           string
	Name            string
	RoleID          Role
	Active          bool
}

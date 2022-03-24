package entity

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID             uuid.UUID `gorm:"PrimaryKey"`
	PersonalNumber string
	Password       string
	Email          string
	Name           string
	RoleID         uuid.UUID
	Role           Role `gorm:"ForeignKey:RoleId"`
	Active         bool
}

func (user *User) BeforeCreate(tx *gorm.DB) (err error) {
	user.ID = uuid.New()
	return
}

package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID              uuid.UUID `gorm:"PrimaryKey"`
	Personal_number string
	Password        string
	Email           string
	Name            string
	RoleID          Role `gorm:"ForeignKey:RoleId"`
	Active          bool
}

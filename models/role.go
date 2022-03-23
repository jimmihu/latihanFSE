package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Role struct {
	gorm.Model
	ID     uuid.UUID `gorm:"PrimaryKey"`
	Title  string    `gorm:"type:ENUM('admin','maker','checker','signer','viewer')"`
	Active bool
}

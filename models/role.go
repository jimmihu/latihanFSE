package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Role struct {
	ID     uuid.UUID `gorm:"PrimaryKey"`
	Title  string    `gorm:"type:ENUM('admin','maker','checker','signer','viewer')"`
	Active bool
}

func (role *Role) BeforeCreate(tx *gorm.DB) (err error) {
	role.ID = uuid.New()
	return
}

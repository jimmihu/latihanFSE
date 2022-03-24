package entity

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Role struct {
	ID     uuid.UUID `gorm:"primaryKey;column:id;type:varchar(55)" json:"role_id"`
	Title  string    `gorm:"type:enum('admin','maker','checker','signer','viewer');column:title" json:"title"`
	Active bool      `gorm:"column:active;type:bool" json:"active"`
}

func (role *Role) BeforeCreate(tx *gorm.DB) (err error) {
	role.ID = uuid.New()
	return
}

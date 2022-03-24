package entity

import (
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
)

type Role struct {
	ID     uuid.UUID `gorm:"primaryKey;column:id;type:varchar(55)"`
	Title  string    `gorm:"type:enum('admin','maker','checker','signer','viewer');column:title"`
	Active bool      `gorm:"column:active;type:bool"`
}

func (role *Role) BeforeCreate(tx *gorm.DB) (err error) {
	role.ID = uuid.New()
	return
}

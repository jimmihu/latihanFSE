package entity

import (
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
)

type User struct {
	ID             uuid.UUID `gorm:"primaryKey;column:id;type:varchar(55)"`
	PersonalNumber string    `gorm:"column:personal_number;type:varchar(65)"`
	Password       string    `gorm:"column:password;type:varchar(65)"`
	Email          string    `gorm:"column:emai;type:varchar(255)"`
	Name           string    `gorm:"column:name;type:varchar(255)"`
	RoleID         uuid.UUID `gorm:"column:role_id;type:varchar(55)"`
	Role           Role      `gorm:"ForeignKey:RoleID"`
	Active         bool      `gorm:"column:active;type:bool"`
}

func (user *User) BeforeCreate(tx *gorm.DB) (err error) {
	user.ID = uuid.New()
	return
}

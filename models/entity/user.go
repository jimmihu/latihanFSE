package entity

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID             uuid.UUID      `gorm:"primaryKey;column:id;type:varchar(55)" json:"user_id"`
	PersonalNumber string         `gorm:"column:personal_number;type:varchar(65)" json:"personal_number"`
	Password       string         `gorm:"column:password;type:varchar(65)" json:"password"`
	Email          string         `gorm:"column:email;type:varchar(255)" json:"email"`
	Name           string         `gorm:"column:name;type:varchar(255)" json:"name"`
	RoleID         uuid.UUID      `gorm:"column:role_id;type:varchar(55)" json:"-"`
	Role           Role           `gorm:"foreignKey:RoleID" json:"role"`
	Active         bool           `gorm:"column:active;type:bool" json:"active"`
	CreatedAt      time.Time      `gorm:"column:created_at" json:"created_at"`
	UpdatedAt      time.Time      `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt      gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}

func (user *User) BeforeCreate(tx *gorm.DB) (err error) {
	user.ID = uuid.New()
	return
}

type UserList struct {
	ID     uuid.UUID `json:"id"`
	Name   string    `json:"name"`
	RoleID uuid.UUID `json:"-"`
	Role   Role      `json:"role" gorm:"foreignKey:RoleID"`
	Active bool      `json:"active"`
}

type UserDetail struct {
	ID             uuid.UUID `json:"id"`
	Name           string    `json:"name"`
	PersonalNumber string    `json:"personal_number"`
	Email          string    `json:"email"`
	RoleID         uuid.UUID `json:"-"`
	Role           Role      `json:"role" gorm:"foreignKey:RoleID"`
	Active         bool      `json:"active"`
}

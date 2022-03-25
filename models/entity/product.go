package entity

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Product struct {
	ID          uuid.UUID      `gorm:"primaryKey;column:id;type:varchar(55)" json:"product_id"`
	Description string         `gorm:"column:description;type:text" json:"description"`
	Name        string         `gorm:"column:name;type:varchar(255)" json:"name"`
	Status      string         `gorm:"type:enum('inactive','approved','active');column:status" json:"status"`
	MakerID     uuid.UUID      `gorm:"column:maker_id;type:varchar(55)" json:"-"`
	Maker       User           `gorm:"foreignKey:MakerID" json:"maker"`
	CheckerID   uuid.UUID      `gorm:"column:checker_id;type:varchar(55)" json:"-"`
	Checker     User           `gorm:"foreignKey:CheckerID" json:"checker"`
	SignerID    uuid.UUID      `gorm:"column:signer_id;type:varchar(55)" json:"-"`
	Signer      User           `gorm:"foreignKey:SignerID" json:"signer"`
	CreatedAt   time.Time      `gorm:"column:created_at" json:"created_at"`
	UpdatedAt   time.Time      `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}

func (product *Product) BeforeCreate(tx *gorm.DB) (err error) {
	product.ID = uuid.New()
	return
}

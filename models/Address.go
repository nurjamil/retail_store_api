package models

import (
	"time"

	"gorm.io/gorm"
)

type Address struct {
	ID        uint           `gorm:"primaryKey" json:"address_id" form:"address_id"`
	CreatedAt time.Time      `json:"created_at" form:"created_at"`
	UpdatedAt time.Time      `json:"updated_at" form:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at" form:"deleted_at"`
	UserID    uint           `gorm:"foreignKey" json:"user_id" form:"user_id"`
	Name      string         `gorm:"type:varchar(255);not null" json:"name" form:"name"`
	Address   string         `gorm:"type:varchar(255);not null" json:"address" form:"address"`
}

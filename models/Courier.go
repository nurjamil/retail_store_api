package models

import (
	"time"

	"gorm.io/gorm"
)

type Courier struct {
	ID          uint           `gorm:"primaryKey" json:"courier_id" form:"courier_id"`
	CreatedAt   time.Time      `json:"created_at" form:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at" form:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"deleted_at" form:"deleted_at"`
	CompanyName string         `gorm:"type:varchar(100);unique;not null" json:"company_name" form:"company_name"`
}

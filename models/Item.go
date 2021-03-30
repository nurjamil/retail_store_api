package models

import (
	"time"

	"gorm.io/gorm"
)

type Item struct {
	ID             uint           `gorm:"primaryKey" json:"item_id" form:"item_id"`
	CreatedAt      time.Time      `json:"created_at" form:"created_at"`
	UpdatedAt      time.Time      `json:"updated_at" form:"updated_at"`
	DeletedAt      gorm.DeletedAt `gorm:"index" json:"deleted_at" form:"deleted_at"`
	Name           string         `gorm:"type:varchar(100);unique;not null" json:"name" form:"name"`
	Description    string         `gorm:"type:varchar(500)" json:"description" form:"description"`
	Stock          uint           `json:"stock" form:"stock"`
	Price          uint           `json:"price" form:"price"`
	ItemCategoryID uint           `gorm:"not null" json:"item_category_id" form:"item_category_id"`
	ItemCategory   ItemCategory   `json:"item_category" form:"item_category"`
}
type ItemAPI struct {
	ID           uint
	Name         string
	Description  string
	Stock        uint
	Price        uint
	ItemCategory ItemCategory
}

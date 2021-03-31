package models

import (
	"time"

	"gorm.io/gorm"
)

type OrderItem struct {
	ID        uint           `gorm:"primaryKey" json:"order_item_id" form:"order_item_id"`
	CreatedAt time.Time      `json:"created_at" form:"created_at"`
	UpdatedAt time.Time      `json:"updated_at" form:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at" form:"deleted_at"`
	OrderID   uint           `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"order_id" form:"order_id"`
	ItemID    uint           `json:"item_id" form:"item_id"`
	Item      Item           `json:"item" form:"item"`
	Quantity  uint           `gorm:"default:1" json:"quantity" form:"quantity"`
}

package models

import (
	"time"

	"gorm.io/gorm"
)

type Payment struct {
	ID              uint           `gorm:"primaryKey" json:"payment_id" form:"payment_id"`
	CreatedAt       time.Time      `json:"created_at" form:"created_at"`
	UpdatedAt       time.Time      `json:"updated_at" form:"updated_at"`
	DeletedAt       gorm.DeletedAt `gorm:"index" json:"deleted_at" form:"deleted_at"`
	OrderID         uint           `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"order_id" form:"order_id"`
	UserID          uint           `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"user_id" form:"user_id"`
	TransactionCode string         `gorm:"type:varchar(50);unique;not null" json:"transaction_code" form:"transaction_code"`
	Status          string         `gorm:"type:varchar(20);not null" json:"status" form:"status"`
	TotalAmount     uint           `json:"total_amount" form:"total_amount"`
}

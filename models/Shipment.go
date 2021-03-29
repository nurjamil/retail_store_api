package models

import (
	"time"

	"gorm.io/gorm"
)

type Shipment struct {
	ID         uint           `gorm:"primaryKey" json:"shipment_id" form:"shipment_id"`
	CreatedAt  time.Time      `json:"created_at" form:"created_at"`
	UpdatedAt  time.Time      `json:"updated_at" form:"updated_at"`
	DeletedAt  gorm.DeletedAt `gorm:"index" json:"deleted_at" form:"deleted_at"`
	PaymentID  uint           `json:"payment_id" form:"payment_id"`
	Payment    Payment        `json:"payment" form:"payment"`
	OrderID    uint           `json:"order_id" form:"order_id"`
	Order      Order          `json:"order" form:"order"`
	AddressID  uint           `json:"address_id" form:"address_id"`
	Address    Address        `json:"address" form:"address"`
	ResiNumber string         `gorm:"type:varchar(100);unique;not null" json:"resi_number" form:"resi_number"`
}

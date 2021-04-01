package models

import (
	"time"

	"gorm.io/gorm"
)

type Order struct {
	ID               uint           `gorm:"primaryKey" json:"order_id" form:"order_id"`
	CreatedAt        time.Time      `json:"created_at" form:"created_at"`
	UpdatedAt        time.Time      `json:"updated_at" form:"updated_at"`
	DeletedAt        gorm.DeletedAt `gorm:"index" json:"-" form:"deleted_at"`
	UserID           uint           `json:"user_id" form:"user_id"`
	TotalAmount      uint           `json:"total_amount" form:"total_amount"`
	CourierID        uint           `json:"courier_id" form:"courier_id"`
	Courier          Courier        `json:"courier" form:"courier"`
	AddressID        uint           `json:"address_Id" form:"address_Id"`
	Address          Address        `json:"address" form:"address"`
	PaymentServiceID uint           `json:"payment_service_id" form:"payment_service_id"`
	PaymentService   PaymentService `json:"payment_service" form:"payment_service"`
	OrderItem        []OrderItem    `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"order_item" form:"order_item"`
	Payment          Payment        `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"payment" form:"payment"`
}

type ResponseOrder struct {
	Code    uint
	Status  string
	Message string
	Data    Order
}

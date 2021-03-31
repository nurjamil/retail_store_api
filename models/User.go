package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID           uint           `gorm:"primaryKey" json:"user_id" form:"user_id"`
	CreatedAt    time.Time      `json:"created_at" form:"created_at"`
	UpdatedAt    time.Time      `json:"updated_at" form:"updated_at"`
	DeletedAt    gorm.DeletedAt `gorm:"index" json:"deleted_at" form:"deleted_at"`
	Username     string         `gorm:"type:varchar(50);unique;not null" json:"username" form:"username"`
	Name         string         `gorm:"type:varchar(50);not null" json:"name" form:"name"`
	Email        string         `gorm:"type:varchar(100);unique;not null" json:"email" form:"email"`
	Password     string         `gorm:"type:varchar(100);not null" json:"password" form:"password"`
	PhoneNumber  string         `gorm:"type:varchar(50);unique;not null" json:"phone_number" form:"phone_number"`
	Role  string         `gorm:"type:varchar(50);not null" json:"role" form:"role"`
	Token        string         `gorm:"type:varchar(255);not null" json:"token" form:"token"`
	Address      []Address      `json:"address" form:"address"`
	ShoppingCart ShoppingCart   `json:"shopping_cart" form:"shopping_cart"`
}
type APIUser struct {
	Name  string
	Email string
	Token string
}

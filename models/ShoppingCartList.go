package models

import (
	"time"

	"github.com/labstack/echo"
	"gorm.io/gorm"
)

type ShoppingCartList struct {
	ID             uint           `gorm:"primaryKey" json:"shopping_cart_list_id" form:"shopping_cart_list_id"`
	CreatedAt      time.Time      `json:"created_at" form:"created_at"`
	UpdatedAt      time.Time      `json:"updated_at" form:"updated_at"`
	DeletedAt      gorm.DeletedAt `gorm:"index" json:"deleted_at" form:"deleted_at"`
	ShoppingCartID uint           `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"shopping_cart_id" form:"shopping_cart_id"`
	ItemID         uint           `json:"item_id" form:"item_id"`
	Item           Item           `json:"item" form:"item"`
	Quantity       uint           `json:"quantity" form:"quantity"`
}

func (s *ShoppingCartList) Find(c echo.Context, DB *gorm.DB) error {
	return DB.Where(s).First(s).Error
}

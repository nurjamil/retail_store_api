package models

import (
	"time"

	"github.com/labstack/echo"
	"gorm.io/gorm"
)

type ShoppingCart struct {
	ID               uint               `gorm:"primaryKey" json:"shopping_cart_id" form:"shopping_cart_id"`
	CreatedAt        time.Time          `json:"created_at" form:"created_at"`
	UpdatedAt        time.Time          `json:"updated_at" form:"updated_at"`
	DeletedAt        gorm.DeletedAt     `gorm:"index" json:"deleted_at" form:"deleted_at"`
	ShoppingCartList []ShoppingCartList `json:"shopping_cart_list" form:"shopping_cart_list"`
	UserID           uint               `json:"user_id" form:"user_id"`
}

type ShoppingCartAPI struct {
	ID     uint `gorm:"primaryKey" json:"shopping_cart_id" form:"shopping_cart_id"`
	UserID uint `json:"user_id" form:"user_id"`
}

func (s *ShoppingCart) GetShoppingCart(c echo.Context, DB *gorm.DB, model *ShoppingCart) error {
	c.Bind(s)
	if err := DB.Preload("ShoppingCartList.Item.ItemCategory").Where(model).First(s).Error; err != nil {
		return err
	}

	return nil
}

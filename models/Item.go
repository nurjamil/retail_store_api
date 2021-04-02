package models

import (
	"time"

	"github.com/labstack/echo"
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

//create in database
func (i *Item) Create(c echo.Context, DB *gorm.DB) error {
	c.Bind(i)
	return DB.Save(&i).Error
}

//find in the database
func (i *Item) Find(c echo.Context, DB *gorm.DB) error {
	c.Bind(i)
	return DB.Where(i).First(i).Error
}

type ItemResponse struct {
	Code    uint   `json:"code" form:"code"`
	Status  string `json:"status" form:"status"`
	Message string `json:"message" form:"message"`
	Data    Item   `json:"data,omitempty" form:"data"`
}

type ItemResponseArr struct {
	Code    uint   `json:"code" form:"code"`
	Status  string `json:"status" form:"status"`
	Message string `json:"message" form:"message"`
	Data    []Item `json:"data,omitempty" form:"data"`
}

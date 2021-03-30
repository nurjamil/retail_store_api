package config

import (
	"fmt"
	"retailStore/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {

	config := map[string]string{
		"DB_Username": "root",
		"DB_Password": "",
		"DB_Port":     "3306",
		"DB_Host":     "localhost",
		"DB_Name":     "retail_store_db",
	}

	connectionString :=
		fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
			config["DB_Username"],
			config["DB_Password"],
			config["DB_Host"],
			config["DB_Port"],
			config["DB_Name"],
		)
	var err error
	DB, err = gorm.Open(mysql.Open(connectionString), &gorm.Config{})

	if err != nil {
		panic(err)
	}
	InitialMigration()
}

func InitialMigration() {
	DB.AutoMigrate(&models.User{}, &models.Address{}, &models.Courier{}, &models.Item{}, &models.ItemCategory{}, &models.Order{}, &models.OrderItem{}, &models.Payment{}, &models.PaymentService{}, &models.Shipment{}, &models.ShoppingCart{}, &models.ShoppingCartList{})
}

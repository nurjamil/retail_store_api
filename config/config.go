package config

import (
	"fmt"
	"os"
	"retailStore/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {

	connectionString :=
		fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
			os.Getenv("DB_Username"),
			os.Getenv("DB_Password"),
			os.Getenv("DB_Host"),
			os.Getenv("DB_Port"),
			os.Getenv("DB_Name"),
		)
	var err error
	DB, err = gorm.Open(mysql.Open(connectionString), &gorm.Config{})

	if err != nil {
		panic(err)
	}
}

func InitDBTest() {

	connectionString :=
		fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
			os.Getenv("DB_Username_Test"),
			os.Getenv("DB_Password_Test"),
			os.Getenv("DB_Host_Test"),
			os.Getenv("DB_Port_Test"),
			os.Getenv("DB_Name_Test"),
		)
	var err error
	DB, err = gorm.Open(mysql.Open(connectionString), &gorm.Config{})

	if err != nil {
		panic(err)
	}
	DropTable()
}

func InitialMigration() {
	DB.AutoMigrate(&models.User{}, &models.Address{}, &models.Courier{}, &models.Item{}, &models.ItemCategory{}, &models.Order{}, &models.OrderItem{}, &models.Payment{}, &models.PaymentService{}, &models.Shipment{}, &models.ShoppingCart{}, &models.ShoppingCartList{})
}

func DropTable() {
	DB.Migrator().DropTable(&models.User{}, &models.Address{}, &models.Courier{}, &models.Item{}, &models.ItemCategory{}, &models.Order{}, &models.OrderItem{}, &models.Payment{}, &models.PaymentService{}, &models.Shipment{}, &models.ShoppingCart{}, &models.ShoppingCartList{})
}

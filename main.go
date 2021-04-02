package main

import (
	"fmt"
	"retailStore/config"
	"retailStore/lib/seeders"
	"retailStore/routes"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}

	config.InitDB()
	//config.DropTable() //reset tables
	config.InitialMigration()

	seeders.Seed()
	seeders.ItemSeed() // seeders for insert categories, paymentservices, and couries. for dev purposes

	e := routes.New()

	e.Logger.Fatal(e.Start(":3000"))

}

package main

import (
	"fmt"
	"retailStore/config"
	"retailStore/routes"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}
	config.InitDB()

	e := routes.New()

	e.Logger.Fatal(e.Start(":3000"))

}

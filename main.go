package main

import (
	"retailStore/config"
	"retailStore/routes"
)

func main() {
	config.InitDB()

	e := routes.New()

	e.Logger.Fatal(e.Start(":3000"))

	
}




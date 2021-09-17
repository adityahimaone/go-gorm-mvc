package main

import (
	"orm-crud/config"
	"orm-crud/routes"
)

func main() {
	config.InitDB()
	config.InitalMigration()
	e := routes.NewRouters()
	e.Start(":8080")
}

package main

import (
	"orm-crud/config"
	m "orm-crud/middleware"
	"orm-crud/routes"
)

func main() {
	config.InitDB()
	config.InitalMigration()
	e := routes.NewRouters()
	m.LogMiddleware(e)
	e.Logger.Fatal(e.Start(":8080"))
}

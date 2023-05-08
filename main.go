package main

import (
	"cm-api/db"
	"cm-api/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
	db.InitDb()
	app := fiber.New()

	routes.Init(app)

	app.Listen(":8080")
}
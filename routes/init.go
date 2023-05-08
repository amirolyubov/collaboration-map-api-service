package routes

import (
	"cm-api/categories"
	"cm-api/points"

	"github.com/gofiber/fiber/v2"
)

func Init(app *fiber.App) {
	app.Get("/points", points.GetPoints)
	app.Get("/points/:pointId", points.GetPoint)
	app.Patch("/points/:pointId", points.EditPoint)
	app.Delete("/points/:pointId", points.DeletePoint)
	app.Get("/categories", categories.GetCategories)
}
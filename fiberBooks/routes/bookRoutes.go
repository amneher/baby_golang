package routes

import (
	"github.com/amneher/fiberBooks/views"
	"github.com/gofiber/fiber/v3"
)

func BookRoutes(app fiber.Router) {
	r := app.Group("/books")

	r.Get("/all", views.GetAllBooks)
	r.Post("/create", views.CreateBook)
}

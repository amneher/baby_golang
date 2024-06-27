package routes

import (
	"github.com/amneher/fiberBooks/views"
	"github.com/gofiber/fiber/v3"
)

func AuthorRoutes(app fiber.Router) {
	r := app.Group("/authors")

	r.Get("/all", views.GetAllAuthors)
	r.Post("/create", views.CreateAuthor)
}

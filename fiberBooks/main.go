package main

import (
	"github.com/amneher/fiberBooks/initializers"
	"github.com/amneher/fiberBooks/models"
	"github.com/amneher/fiberBooks/routes"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/helmet"
	"github.com/gofiber/fiber/v3/middleware/logger"
)

func init() {
	initializers.InitDB("fiberBooks")

	initializers.DB.AutoMigrate(&models.Book{}, &models.Author{})
}

func main() {
	app := fiber.New()

	app.Use(helmet.New())
	app.Use(logger.New())

	routes.BookRoutes(app)
	routes.AuthorRoutes(app)

	app.Listen(":8080")
}

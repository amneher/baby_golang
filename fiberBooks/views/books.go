package views

import (
	"github.com/amneher/fiberBooks/models"
	"github.com/gofiber/fiber/v3"
)

func CreateBook(c fiber.Ctx) error {
	var book models.Book
	if err := c.Bind().JSON(book); err != nil {
		return err
	}
	if err := models.CreateBook(&book).Error; err != nil {
		return err
	}
	return c.JSON(book)
}

func GetAllBooks(c fiber.Ctx) error {
	var books []models.Book
	if err := models.GetAllBooks(books).Error; err != nil {
		return err
	}
	return c.JSON(books)
}

package views

import (
	"github.com/amneher/fiberBooks/models"
	"github.com/gofiber/fiber/v3"
)

func GetAllAuthors(c fiber.Ctx) error {
	var authors []models.Author
	if err := models.GetAllAuthors(authors).Error; err != nil {
		return err
	}
	return c.JSON(authors)
}

func CreateAuthor(c fiber.Ctx) error {
	var author models.Author
	if err := c.Bind().JSON(author); err != nil {
		return err
	}
	if err := models.CreateAuthor(&author).Error; err != nil {
		return err
	}
	return c.JSON(author)
}

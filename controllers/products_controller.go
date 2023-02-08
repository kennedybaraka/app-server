package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kennedybaraka/app-server/database"
	"gorm.io/gorm"
)

type products struct {
	db *gorm.DB
}

type Products interface {
	Products(c *fiber.Ctx) error
}

func NewProductsController() *products {
	return &products{
		db: database.Data,
	}
}

func (p *products) Products(c *fiber.Ctx) error {

	return c.JSON(
		fiber.Map{
			"message": "These are all paginated products.",
		},
	)
}

func (p *products) Create(c *fiber.Ctx) error {

	return c.JSON(
		fiber.Map{
			"message": "These are all paginated products.",
		},
	)
}

func (p *products) One(c *fiber.Ctx) error {

	return c.JSON(
		fiber.Map{
			"message": "These are all paginated products.",
		},
	)
}

func (p *products) Delete(c *fiber.Ctx) error {

	return c.JSON(
		fiber.Map{
			"message": "These are all paginated products.",
		},
	)
}

func (p *products) Update(c *fiber.Ctx) error {

	return c.JSON(
		fiber.Map{
			"message": "These are all paginated products.",
		},
	)
}

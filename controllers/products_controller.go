package controllers

import (
	"github.com/go-redis/redis/v8"
	"github.com/gofiber/fiber/v2"
	"github.com/kennedybaraka/app-server/database"
	"gorm.io/gorm"
)

type products struct {
	db     *gorm.DB
	client *redis.Client
}

type Products interface {
	ReadProducts(c *fiber.Ctx) error
	AdminReadProducts(c *fiber.Ctx) error
	CreateProduct(c *fiber.Ctx) error
	ReadProduct(c *fiber.Ctx) error
	AdminReadProduct(c *fiber.Ctx) error
	UpdateProduct(c *fiber.Ctx) error
	DeleteProduct(c *fiber.Ctx) error
	SearchProducts(c *fiber.Ctx) error
	AdminSearchProducts(c *fiber.Ctx) error
	FilterProducts(c *fiber.Ctx) error
	AdminFilterProducts(c *fiber.Ctx) error
}

func NewProductsController() *products {
	return &products{
		db:     database.Data,
		client: database.Client,
	}
}

func (p *products) ReadProducts(c *fiber.Ctx) error {

	return c.JSON(
		fiber.Map{
			"message": "These are all paginated products.",
		},
	)
}

func (p *products) AdminReadProducts(c *fiber.Ctx) error {

	return c.JSON(
		fiber.Map{
			"message": "These are all paginated products.",
		},
	)
}
func (p *products) CreateProduct(c *fiber.Ctx) error {

	return c.JSON(
		fiber.Map{
			"message": "These are all paginated products.",
		},
	)
}

func (p *products) ReadProduct(c *fiber.Ctx) error {

	return c.JSON(
		fiber.Map{
			"message": "These are all paginated products.",
		},
	)
}

func (p *products) AdminReadProduct(c *fiber.Ctx) error {

	return c.JSON(
		fiber.Map{
			"message": "These are all paginated products.",
		},
	)
}

func (p *products) DeleteProduct(c *fiber.Ctx) error {

	return c.JSON(
		fiber.Map{
			"message": "These are all paginated products.",
		},
	)
}

func (p *products) UpdateProduct(c *fiber.Ctx) error {

	return c.JSON(
		fiber.Map{
			"message": "These are all paginated products.",
		},
	)
}

func (p *products) SearchProducts(c *fiber.Ctx) error {

	return c.JSON(
		fiber.Map{
			"message": "These are all paginated products.",
		},
	)
}
func (p *products) AdminSearchProducts(c *fiber.Ctx) error {

	return c.JSON(
		fiber.Map{
			"message": "These are all paginated products.",
		},
	)
}
func (p *products) FilterProducts(c *fiber.Ctx) error {

	return c.JSON(
		fiber.Map{
			"message": "These are all paginated products.",
		},
	)
}

func (p *products) AdminFilterProducts(c *fiber.Ctx) error {

	return c.JSON(
		fiber.Map{
			"message": "These are all paginated products.",
		},
	)
}

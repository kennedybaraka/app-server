package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kennedybaraka/app-server/database"
	"gorm.io/gorm"
)

type shop struct {
	db *gorm.DB
}

type Shop interface {
	CreateShop(c *fiber.Ctx) error
	DeleteShop(c *fiber.Ctx) error
	UpdateShop(c *fiber.Ctx) error
}

func NewShopController() *shop {
	return &shop{
		db: database.Data,
	}
}

func (a *shop) CreateShop(c *fiber.Ctx) error {
	return c.JSON(
		fiber.Map{
			"message": "These are all paginated users.",
		},
	)
}

func (a *shop) UpdateShop(c *fiber.Ctx) error {
	return c.JSON(
		fiber.Map{
			"message": "These are all paginated users.",
		},
	)
}
func (a *shop) DeleteShop(c *fiber.Ctx) error {
	return c.JSON(
		fiber.Map{
			"message": "These are all paginated users.",
		},
	)
}

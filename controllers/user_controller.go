package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kennedybaraka/app-server/database"
	"gorm.io/gorm"
)

type auth struct {
	db *gorm.DB
}

type Auth interface {
	Users(c *fiber.Ctx) error
}

func NewAuthController() *auth {
	return &auth{
		db: database.Data,
	}
}

func (a *auth) Users(c *fiber.Ctx) error {

	return c.JSON(
		fiber.Map{
			"message": "These are all paginated users.",
		},
	)
}

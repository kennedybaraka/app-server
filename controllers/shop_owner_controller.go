package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kennedybaraka/app-server/database"
	"github.com/kennedybaraka/app-server/models"
	"github.com/kennedybaraka/app-server/utils/email"
	"github.com/kennedybaraka/app-server/utils/encryption"
	"github.com/kennedybaraka/app-server/utils/validation"
	"gorm.io/gorm"
)

var (
	argon encryption.Argon = encryption.NewArgon2()
)

type owner struct {
	db *gorm.DB
}

type ShopOwner interface {
	CreateOwner(c *fiber.Ctx) error
	AuthenticateShopOwner(c *fiber.Ctx) error
	UpdateOwner(c *fiber.Ctx) error
	DeleteOwner(c *fiber.Ctx) error
	RefreshToken(c *fiber.Ctx) error
}

func NewShopOwnerController() *owner {
	return &owner{
		db: database.Data,
	}
}

func (o *owner) CreateOwner(c *fiber.Ctx) error {
	var owner models.Owner

	if err := c.BodyParser(&owner); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": fiber.Map{
				"type":    "Input",
				"message": "bad formatting of input data",
			},
		})
	}
	// validation
	if err := validation.ValidateCreateOwnerStruct(owner); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": fiber.Map{
				"type":    "Input",
				"message": err.Error(),
			},
		})
	}
	// password hashing
	hash, err := argon.HashPassword(owner.OwnerPassword)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": fiber.Map{
				"type":    "Server",
				"message": "failed to hash your password",
			},
		})
	}
	// save to db
	owner.OwnerPassword = hash
	result := o.db.Omit("store_id").Create(&owner)

	// send email
	// go func() {
	email.SendVerificationEmail(owner.OwnerEmail)
	// }()

	if err := result.Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": fiber.Map{
				"type":    "Database",
				"message": "email already exists",
			},
		})
	}

	return c.JSON(
		fiber.Map{
			"message": "an account has been created successfully",
			"owner":   owner,
		},
	)
}

func (o *owner) AuthenticateShopOwner(c *fiber.Ctx) error {
	var owner models.Owner
	var input models.Owner

	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": fiber.Map{
				"type":    "Input",
				"message": "bad formatting of input data",
			},
		})
	}
	// if user exists
	result := o.db.First(&owner, "owner_email = ?", input.OwnerEmail)
	if result.Error != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": fiber.Map{
				"type":    "Database",
				"message": "your email is not registered yet",
			},
		})
	}
	// compare passwords
	err := argon.VerifyPassword(owner.OwnerPassword, input.OwnerPassword)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": fiber.Map{
				"type":    "Authentication",
				"message": "passwords don't match",
			},
		})
	}
	// jwt creation

	return c.JSON(
		fiber.Map{
			"message": "These are all paginated users.",
		},
	)
}
func (a *owner) UpdateOwner(c *fiber.Ctx) error {
	return c.JSON(
		fiber.Map{
			"message": "These are all paginated users.",
		},
	)
}

func (a *owner) DeleteOwner(c *fiber.Ctx) error {
	return c.JSON(
		fiber.Map{
			"message": "These are all paginated users.",
		},
	)
}

func (a *owner) RefreshToken(c *fiber.Ctx) error {
	return c.JSON(
		fiber.Map{
			"message": "These are all paginated users.",
		},
	)
}

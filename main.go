package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	"github.com/kennedybaraka/app-server/config"
	"github.com/kennedybaraka/app-server/database"
	"github.com/kennedybaraka/app-server/routes"
	"gorm.io/gorm"
)

var (
	PORT     = config.Env("PORT")
	Database *gorm.DB
)

func init() {
	// database stuff
	database.ConnectToPostgreSql()
	database.ConnectToRedis()
}

func main() {
	// fiber initialization
	app := fiber.New()

	app.Use(logger.New())
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))
	app.Use(compress.New(compress.Config{
		Level: compress.LevelBestSpeed, // 1
	}))

	app.Get("/metrics", monitor.New(monitor.Config{Title: "My Service Metrics Page"}))

	// controllers

	// route handler
	routes.Routes(app)

	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(
			fiber.Map{
				"message": "Server is up and running!",
			},
		)
	})

	// 404
	app.Use(func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Sorry you are lost",
			"error": fiber.Map{
				"code":    fiber.StatusNotFound,
				"message": c.Path() + " does not exist",
			},
		})
	})

	app.Listen(":" + PORT)

}

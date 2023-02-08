package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kennedybaraka/app-server/controllers"
)

var (
	products controllers.Products = controllers.NewProductsController()
	auth     controllers.Auth     = controllers.NewAuthController()
)

func Routes(r fiber.Router) {
	api := r.Group("/api")

	// users
	a := api.Group("/auth")
	a.Get("/", auth.Users)

	// products
	p := api.Group("/products")
	p.Get("/", products.Products)
}

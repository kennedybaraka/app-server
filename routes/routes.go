package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kennedybaraka/app-server/controllers"
)

var (
	products controllers.Products  = controllers.NewProductsController()
	owners   controllers.ShopOwner = controllers.NewShopOwnerController()
	shops    controllers.Shop      = controllers.NewShopController()
)

func Routes(r fiber.Router) {
	api := r.Group("/api")

	// owners
	a := api.Group("/owner")
	a.Post("/register", owners.CreateOwner)
	a.Post("/login", owners.AuthenticateShopOwner)
	a.Put("/update/:id", owners.UpdateOwner)
	a.Delete("/delete/:id", owners.DeleteOwner)
	a.Post("/refresh", owners.RefreshToken)

	// shops
	s := api.Group("/shop")
	s.Post("/create", shops.CreateShop)
	s.Put("/update/:id", shops.CreateShop)
	s.Delete("/delete/:id", shops.CreateShop)

	// products
	p := api.Group("/products")

	// admins
	pa := p.Group("/admin")
	pa.Post("/create", products.CreateProduct)
	pa.Post("/filter", products.AdminFilterProducts)
	pa.Post("/search", products.AdminSearchProducts)
	pa.Delete("/delete/:id", products.DeleteProduct)
	pa.Put("/update/:id", products.UpdateProduct)
	pa.Get("/all", products.AdminReadProducts)
	pa.Get("/one/:id", products.AdminReadProduct)

	// client
	pc := p.Group("/client")
	pc.Get("/all", products.ReadProducts)
	pc.Get("/one/:id", products.ReadProduct)
	pa.Post("/filter", products.FilterProducts)
	pa.Post("/search", products.SearchProducts)

}

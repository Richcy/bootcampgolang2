package main

import (
	"rapidtech/fiberrestapi/controllers"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	// controllers
	productApiController := controllers.InitProductAPIController()

	p := app.Group("/api")
	p.Get("/hello", productApiController.Greeting)
	p.Get("/products", productApiController.GetAllProducts)
	p.Post("/products", productApiController.CreateProduct)
	p.Get("/products/productdetail", productApiController.GetDetailProduct)
	p.Get("/products/detail/:id", productApiController.GetDetailProduct2)
	p.Put("/products/:id", productApiController.EditProduct)
	p.Delete("/products/:id", productApiController.DeleteProduct)

	// prod := app.Group("/products")
	// prod.Get("/", prodController.IndexProduct)
	// prod.Get("/create", prodController.AddProduct)
	// prod.Post("/create", prodController.AddPostedProduct)
	// prod.Get("/productdetail", prodController.GetDetailProduct)
	// prod.Get("/detail/:id", prodController.GetDetailProduct2)
	// prod.Get("/editproduct/:id", prodController.EditlProduct)
	// prod.Post("/editproduct/:id", prodController.EditlPostedProduct)
	// prod.Get("/deleteproduct/:id", prodController.DeleteProduct)

	// app.Get("/login",authController.Login)
	// app.Post("/login",authController.LoginPosted)
	// app.Get("/logout",authController.Logout)

	// app.Get("/profile",
	// 	authController.CheckLogin,
	// 	authController.Profile)

	app.Listen(":3000")
}

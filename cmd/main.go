package main

import (
	"fmt"
	"log"

	"github.com/go-project/internal/handlers"
	"github.com/gofiber/fiber/v2"
)

func main() {

	app := fiber.New()
	app.Get("/healthcheck", func(c *fiber.Ctx) error {
		return c.SendString("OK")
	})

	app.Post("/api/products", handlers.CreateProduct)
	app.Get("api/products", handlers.GetAllProducts)

	log.Fatal(app.Listen(":3000"))
	fmt.Println("Hello World")
}
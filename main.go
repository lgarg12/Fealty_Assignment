package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/lgarg12/Fealty_Assignment/routes"
)

func main() {
	app := fiber.New()

	app.Use(cors.New())

	routes.Routers(app)

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	app.Listen(":8000")
}

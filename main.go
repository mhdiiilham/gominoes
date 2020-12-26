package main

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := setupApp()
	app.Listen(":3000")
}

func setupApp() *fiber.App {
	app := fiber.New()
	app.Get("/", func(c *fiber.Ctx) error {
		msg := map[string]string{
			"message": "GO-MINOES",
		}
		return c.Status(http.StatusOK).JSON(msg)
	})
	return app
}

package app

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
)

// SetupApp fiber route
func SetupApp() *fiber.App {
	app := fiber.New(fiber.Config{
		ErrorHandler: customErrHandler,
	})
	app.Get("/", func(c *fiber.Ctx) error {
		msg := map[string]string{
			"message": "GO-MINOES",
		}
		return c.Status(http.StatusOK).JSON(msg)
	})
	return app
}

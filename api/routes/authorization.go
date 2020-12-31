package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mhdiiilham/gominoes/api/controllers"
)

// AuthRoutes handlers
func AuthRoutes(r fiber.Router, c *controllers.UserController) {
	auth := r.Group("/auth")
	auth.Post("/registrations", c.Register)
}

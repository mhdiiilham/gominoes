package app

import (
	"net/http"

	ut "github.com/go-playground/universal-translator"
	"github.com/gofiber/fiber/v2"
	"github.com/mhdiiilham/gominoes/api/controllers"
	"github.com/mhdiiilham/gominoes/api/routes"
	"github.com/mhdiiilham/gominoes/entity/user"
	"github.com/mhdiiilham/gominoes/pkg/jwt"
	"gopkg.in/go-playground/validator.v9"
)

// Managers struct
type Managers struct {
	UserManager user.Manager
}

// SetupApp fiber route
func SetupApp(managers *Managers, jwtService jwt.TokenService, v *validator.Validate, trans ut.Translator) *fiber.App {
	userHandler := controllers.NewUserController(managers.UserManager, jwtService, v, trans)

	app := fiber.New(fiber.Config{
		ErrorHandler: customErrHandler,
	})
	app.Get("/", func(c *fiber.Ctx) error {
		msg := map[string]string{
			"message": "GO-MINOES",
		}
		return c.Status(http.StatusOK).JSON(msg)
	})
	api := app.Group("/api")
	routes.AuthRoutes(api, userHandler)
	return app
}

package app

import (
	"net/http"
	"os"

	ut "github.com/go-playground/universal-translator"
	"github.com/gofiber/fiber/v2"
	"github.com/mhdiiilham/gominoes/api/controllers"
	"github.com/mhdiiilham/gominoes/api/routes"
	db "github.com/mhdiiilham/gominoes/db"
	"github.com/mhdiiilham/gominoes/entity/user"
	"github.com/mhdiiilham/gominoes/pkg/jwt"
	"gopkg.in/go-playground/validator.v9"
)

// SetupApp fiber route
func SetupApp(v *validator.Validate, trans ut.Translator) *fiber.App {
	client, err := db.NewMongoDBConnection(os.Getenv("MONGO_DB_USER"), os.Getenv("MONGO_DB_PASS"), os.Getenv("MONGO_DB"))
	if err != nil {
		panic(err)
	}
	jwtService := jwt.NewJWTService(os.Getenv("JWT_SECRET"), os.Getenv("APP_NAME"))
	userCollection := client.Database(os.Getenv("MONGO_DB")).Collection("users")
	userRepo := user.NewMongoDBRepository(userCollection)
	userManager := user.NewManager(userRepo)
	userHandler := controllers.NewUserController(userManager, jwtService, v, trans)

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

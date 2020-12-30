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

func customErrHandler(ctx *fiber.Ctx, err error) error {
	code := http.StatusInternalServerError
	msg := "Internal Server Error"

	if e, ok := err.(*fiber.Error); ok {
		code = e.Code
		msg = e.Message
	}

	err = ctx.Status(code).JSON(struct {
		Code    int
		Message string
	}{
		Code:    code,
		Message: msg,
	})

	if err != nil {
		return ctx.Status(500).JSON(struct {
			Code int
			Msg  string
		}{
			Code: 500,
			Msg:  "Fuck",
		})
	}

	return nil
}

package main

import (
	"net/http"

	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/gofiber/fiber/v2"
	"gopkg.in/go-playground/validator.v9"
	en_translations "gopkg.in/go-playground/validator.v9/translations/en"
)

func main() {
	translator := en.New()
	uni := ut.New(translator, translator)

	trans, found := uni.GetTranslator("en")
	if !found {
		panic("Translator not found!")
	}
	v := validator.New()
	if err := en_translations.RegisterDefaultTranslations(v, trans); err != nil {
		panic(err)
	}
	_ = v.RegisterTranslation("required", trans, func(ut ut.Translator) error {
		return ut.Add("required", "{0} is required field", true)
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("required", fe.Field())
		return t
	})
	_ = v.RegisterTranslation("email", trans, func(ut ut.Translator) error {
		return ut.Add("email", "{0} must be a valid email", true)
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("email", fe.Field())
		return t
	})

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

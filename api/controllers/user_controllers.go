package controllers

import (
	ut "github.com/go-playground/universal-translator"
	"github.com/gofiber/fiber/v2"
	"github.com/mhdiiilham/gominoes/entity/user"
	"github.com/mhdiiilham/gominoes/pkg/jwt"
	"gopkg.in/go-playground/validator.v9"
)

// UserController struct
type UserController struct {
	m        user.Manager
	Token    jwt.TokenService
	Validate *validator.Validate
	Trans    ut.Translator
}

type registerInput struct {
	Fullname string `json:"fullname" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

// NewUserController function
func NewUserController(r fiber.Router, m user.Manager, ts jwt.TokenService, v *validator.Validate, t ut.Translator) {
	controller := &UserController{
		m:        m,
		Token:    ts,
		Validate: v,
		Trans:    t,
	}
	r.Get("/auth/registrations", controller.register)
}

func (c *UserController) register(ctx *fiber.Ctx) error {
	msg := map[string]string{
		"message": "Hello World",
	}
	return ctx.Status(200).JSON(msg)
}

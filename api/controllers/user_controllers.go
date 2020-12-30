package controllers

import (
	"net/http"

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
	r.Post("/auth/registrations", controller.register)
}

func (c *UserController) register(ctx *fiber.Ctx) error {
	i := registerInput{}

	if err := ctx.BodyParser(&i); err != nil {
		return fiber.ErrInternalServerError
	}

	user := user.User{
		Fullname: i.Fullname,
		Email:    i.Email,
		Password: i.Password,
	}
	id := c.m.Register(user)

	if id == "" {
		return fiber.ErrInternalServerError
	}

	token := c.Token.Generate(&user)

	return ctx.Status(200).JSON(struct {
		Code        int    `json:"code"`
		Message     string `json:"message"`
		AccessToken string `json:"access_token"`
	}{
		Code:        http.StatusCreated,
		Message:     "Success Create User",
		AccessToken: token,
	})
}

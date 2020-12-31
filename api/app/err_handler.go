package app

import (
	"encoding/json"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

// ErrDetail struct
type ErrDetail struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

// ErrResponse struct
type ErrResponse struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Errors  []ErrDetail `json:"errors"`
}

func customErrHandler(ctx *fiber.Ctx, err error) error {
	code := http.StatusInternalServerError
	msg := "Internal Server Error"
	details := []ErrDetail{}

	if e, ok := err.(*fiber.Error); ok {
		code = e.Code
		msg = e.Message

		if code == http.StatusNotAcceptable {
			msg = "Not Acceptable"
			parseJSON := ErrResponse{}
			if err := json.Unmarshal([]byte(e.Message), &parseJSON); err != nil {
				return ctx.Status(http.StatusInternalServerError).JSON(ErrResponse{
					Code:    http.StatusInternalServerError,
					Message: "Internal Server Error",
				})
			}
			code = parseJSON.Code
			msg = parseJSON.Message
			details = parseJSON.Errors
		}
	}

	err = ctx.Status(code).JSON(ErrResponse{
		Code:    code,
		Message: msg,
		Errors:  details,
	})

	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(ErrResponse{
			Code:    http.StatusInternalServerError,
			Message: "Internal Server Error",
		})
	}

	return nil
}

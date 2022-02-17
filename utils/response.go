package utils

import (
	"github.com/gofiber/fiber/v2"
)

type Response struct {
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
	Errors  interface{} `json:"errors"`
	Status  bool        `json:"status"`
}

//SuccessJSON response of success json.
func SuccessJSON(c *fiber.Ctx, code int, message string, data interface{}) error {
	res := Response{
		Message: message,
		Errors:  "",
		Status:  true,
		Data:    data,
	}
	return c.Status(code).JSON(res)
}

//ErrorJSON response of error is json.
func ErrorJSON(c *fiber.Ctx, code int, message string, errors interface{}) error {
	res := Response{
		Message: message,
		Errors:  errors,
		Status:  false,
		Data:    "",
	}
	return c.Status(code).JSON(res)
}

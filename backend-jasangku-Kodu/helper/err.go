package helper

import (
	"github.com/faridanang/jasangku-kodu/model"
	"github.com/gofiber/fiber/v2"
)

func ErrorResponse(c *fiber.Ctx, err, status string, code int) error {
	return c.Status(code).JSON(model.Response{
		Code:   code,
		Status: status,
		Data:   err,
	})
}

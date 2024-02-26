package middleware

import (
	"os"

	"github.com/dgrijalva/jwt-go"
	"github.com/faridanang/jasangku-kodu/model"
	"github.com/gofiber/fiber/v2"
)

func Authorization() func(*fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		if c.Method() == "POST" && (c.Path() == "/api/user/signup" || c.Path() == "/api/user/signin") {
			return c.Next()
		} else {
			tokenAuth := c.Get("Authorization")
			if tokenAuth == "" {
				return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
					"code":   fiber.StatusUnauthorized,
					"status": "Unauthorized",
				})
			}
			jwtSecretKet := []byte(os.Getenv("JWT_SECRET_KEY"))
			claims := &model.ClaimToken{}
			token, err := jwt.ParseWithClaims(tokenAuth, claims, func(t *jwt.Token) (interface{}, error) {
				return jwtSecretKet, nil
			})
			if err != nil {
				return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
					"code":   fiber.StatusUnauthorized,
					"status": "Unauthorized",
					"error":  err.Error(),
				})
			}
			if !token.Valid {
				return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
					"code":   fiber.StatusUnauthorized,
					"status": "Unauthorized",
					"error":  err.Error(),
				})
			}
			return c.Next()

		}
	}
}

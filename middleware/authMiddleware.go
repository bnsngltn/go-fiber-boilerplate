package middleware

import (
	"github.com/bnsngltn/go-fiber-boilerplate/config"
	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v2"
)

// Protected middleware, all routes that uses this handler will require a JWT with a valid signature
func Protected() fiber.Handler {
	return jwtware.New(jwtware.Config{
		SigningKey:    []byte(config.AuthSecret),
		SigningMethod: "HS512",
		ErrorHandler:  jwtError,
	})
}

func jwtError(c *fiber.Ctx, err error) error {
	return c.Status(401).JSON(fiber.Map{
		"error": "Missing or malformed JWT",
	})
}

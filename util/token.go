package util

import (
	"time"

	"github.com/bnsngltn/go-fiber-boilerplate/config"
	"github.com/bnsngltn/go-fiber-boilerplate/database"
	"github.com/bnsngltn/go-fiber-boilerplate/model"
	"github.com/form3tech-oss/jwt-go"
	"github.com/gofiber/fiber/v2"
)

// TokenFromResource func
func TokenFromResource(generic *model.User) (string, error) {
	token := jwt.New(jwt.SigningMethodHS512)

	claims := token.Claims.(jwt.MapClaims)
	claims["username"] = generic.Username

	claims["exp"] = time.Now().Add(time.Hour * config.TokenExpiration).Unix()
	t, err := token.SignedString([]byte(config.AuthSecret))

	return t, err
}

// ResourceFromToken func
func ResourceFromToken(c *fiber.Ctx) (*model.User, error) {
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	username := claims["username"].(string)

	u, err := database.GetUserByUsername(username)
	return u, err
}

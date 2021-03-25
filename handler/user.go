package handler

import (
	"github.com/bnsngltn/go-fiber-boilerplate/database"
	"github.com/bnsngltn/go-fiber-boilerplate/model"
	"github.com/bnsngltn/go-fiber-boilerplate/util"
	"github.com/gofiber/fiber/v2"
)

// Register godoc
// @Tags users
// @Summary Register user
// @Description Register a brand new user to the database.
// @Accept json
// @Produce json
// @Param payload body model.User true "Details needed for registration"
// @Success 201 {object} model.User
// @Router /users/register [POST]
func Register(c *fiber.Ctx) error {
	user := new(model.User)

	if err := c.BodyParser(user); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Malformed request",
		})
	}

	// Hash the password
	var err error
	user.Password, err = util.HashPassword(user.Password)
	if err != nil {
		return err
	}

	if err := database.CreateUser(user); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.Status(201).JSON(user)
}

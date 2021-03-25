package handler

import "github.com/gofiber/fiber/v2"

// PlaceBet godoc
// @Security BearerToken
// @Tags generic
// @Summary Need token
// @Description This route needs token
// @Accept json
// @Produce json
// @Router /generic [GET]
func ThisHandlerNeedsAToken(c *fiber.Ctx) error {
	return c.Status(200).JSON(fiber.Map{
		"message": "Welcome to the protected route",
	})
}

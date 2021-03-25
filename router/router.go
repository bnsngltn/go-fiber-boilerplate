package router

import (
	swagger "github.com/arsmn/fiber-swagger/v2"
	_ "github.com/bnsngltn/go-fiber-boilerplate/docs" //For swagger
	"github.com/bnsngltn/go-fiber-boilerplate/handler"
	"github.com/bnsngltn/go-fiber-boilerplate/middleware"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api")

	// For Swagger Docs
	api.Get("/docs/*", swagger.Handler)

	api.Get("/generic", middleware.Protected(), handler.ThisHandlerNeedsAToken)

	api.Get("/users/register", handler.Register)
}

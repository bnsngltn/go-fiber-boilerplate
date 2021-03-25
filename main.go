package main

import (
	"log"

	"github.com/bnsngltn/go-fiber-boilerplate/database"
	"github.com/bnsngltn/go-fiber-boilerplate/router"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

// @title Go Fiber Boilerplate
// @version 1.0
// @description Boilerplate
// @securityDefinitions.apikey BearerToken
// @in header
// @name Authorization
// @BasePath /api
func main() {
	log.Println("Connecting to the database...")
	if err := database.Connect(); err != nil {
		log.Fatal(err)
	}

	log.Println("Creating new fiber app...")
	app := fiber.New()

	app.Use(logger.New())

	log.Println("Setting up routes...")
	router.SetupRoutes(app)

	app.Listen(":3000")

}

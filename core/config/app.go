package config

import (
	"Find-Backend/core/middleware"

	"github.com/gofiber/fiber/v2"
)

func SetupApp() *fiber.App {
	app := fiber.New(fiber.Config{
			Prefork:       false,
			CaseSensitive: true,
			StrictRouting: true,
			ServerHeader:  "Archilst",
			AppName:       "Find Backend",
	})

	app.Use(middleware.JWTMiddleware())

	return app
}
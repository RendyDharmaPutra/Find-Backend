package routes

import (
	"Find-Backend/features/auth"
	"Find-Backend/features/features_init"
	"Find-Backend/features/user"

	"github.com/gofiber/fiber/v2"
)


func SetupRoutes(app *fiber.App, module *features_init.Module) {

	authGroup :=  app.Group("/auth")
	auth.AuthController(authGroup, module.AuthService)

	userGroup := app.Group("/users")
	user.UserController(userGroup, module.UserService)
}
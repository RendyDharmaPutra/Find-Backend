package main

import (
	"Find-Backend/core/config"
	"Find-Backend/core/routes"
	"Find-Backend/features/features_init"
	"log"
	"os"
)

func main() {
    config.LoadEnv()

    app := config.SetupApp()

    db := config.InitDB()

    module := features_init.InitializeModules(db)

    routes.SetupRoutes(app, module)

    log.Fatal(app.Listen(os.Getenv("APP_PORT")))
}
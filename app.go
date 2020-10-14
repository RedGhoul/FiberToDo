package main

import (
	"log"
	"os"
	"strconv"

	"github.com/RedGhoul/fibertodo/database"
	"github.com/RedGhoul/fibertodo/middleware"
	"github.com/RedGhoul/fibertodo/providers"
	"github.com/RedGhoul/fibertodo/routes"
	"github.com/gofiber/fiber"
	"github.com/gofiber/session"
	"github.com/gofiber/template/django"
	"github.com/joho/godotenv"
)

func StartUp() {
	database.InitDb()
	providers.SetUpSessionProvider(session.New())
	providers.SetUpHashProvider()
	app := setupViewEngine()
	middleware.SetupMiddleware(app)
	routes.SetupRoutes(app)
	setupAppListenPort(app)
}

func setupViewEngine() *fiber.App {
	engine := django.New("./views", ".django")

	configErr := godotenv.Load()
	if configErr != nil {
		log.Fatal("Error loading .env file")
	}

	DebugFlag, _ := strconv.ParseBool(os.Getenv("DEBUG"))
	if DebugFlag {
		engine.Reload(true)
		engine.Debug(true)
	}

	return fiber.New(&fiber.Settings{
		Views: engine,
	})
}

func setupAppListenPort(app *fiber.App) {
	app.Listen(os.Getenv("BASEPORT") + ":" + os.Getenv("SERVERPORT"))
}

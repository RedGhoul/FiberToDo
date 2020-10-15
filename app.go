package main

import (
	"log"
	"os"
	"strconv"

	"fibertodo/database"
	"fibertodo/middleware"
	"fibertodo/providers"
	"fibertodo/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/session/v2"
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

	return fiber.New(fiber.Config{
		Views: engine,
	})
}

func setupAppListenPort(app *fiber.App) {
	app.Listen(os.Getenv("BASEPORT") + ":" + os.Getenv("SERVERPORT"))
}

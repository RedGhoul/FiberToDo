package middleware

import (
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func SetupMiddleware(app *fiber.App) {
	app.Use(recover.New())
	app.Use(requestid.New(requestid.Config{
		Header: "X-Custom-Header",
		Generator: func() string {
			return "static-id"
	}))
	app.Use(helmet.New())
	app.Use(compress.New(compress.Config{
		Level: compress.LevelBestSpeed, // 1
	}))
	app.Use(logger.New(logger.Config{
		Format:     "${pid} ${status} - ${method} ${path}\n",
		TimeFormat: "02-Jan-2006",
		TimeZone:   "America/New_York",
		Output:     os.Stdout,
	}))
	app.Static("/asset", "./public")
}

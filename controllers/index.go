package controllers

import (
	"fibertodo/utils"

	"github.com/gofiber/fiber/v2"
)

func Show_Index(c *fiber.Ctx) error {
	return utils.Render(c, "Home/index", "layouts/main", fiber.Map{
		"Title": "Welcome to Fiber Snips"})
}

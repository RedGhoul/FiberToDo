package controllers

import (
	"github.com/RedGhoul/fibertodo/utils"
	"github.com/gofiber/fiber"
)

func Show_Index(c *fiber.Ctx) {
	utils.Render(c, "Home/index", "layouts/main", fiber.Map{
		"Title": "Welcome to Fiber Snips"})
}

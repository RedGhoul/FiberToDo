package controllers

import (
	"github.com/RedGhoul/fibertodo/repos"
	"github.com/RedGhoul/fibertodo/utils"
	"github.com/gofiber/fiber"
)

func Show_Index(c *fiber.Ctx) {
	users := repos.GetAllUsers()
	utils.Render(c, "Home/index", "layouts/main", fiber.Map{"users": users,
		"Title": "Hello, World!"})
}

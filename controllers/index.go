package controllers

import (
	"github.com/RedGhoul/fibertodo/repos"
	"github.com/RedGhoul/fibertodo/utils"
	"github.com/gofiber/fiber"
)

func ShowIndex(c *fiber.Ctx) {
	users := repos.GetAllUsers()
	utils.Render(c, "Home/index", "layouts/main", fiber.Map{"users": users,
		"Title": "Hello, World!"})
}

func ShowSecrect(c *fiber.Ctx) {
	utils.Render(c, "Home/secrect", "layouts/main", fiber.Map{"msg": "I like dogs"})

}

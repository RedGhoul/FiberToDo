package controllers

import (
	"github.com/RedGhoul/fibertodo/repos"
	"github.com/gofiber/fiber"
)

func ShowIndex(c *fiber.Ctx) {
	users := repos.GetAllUsers()
	if err := c.Render("Home/index", fiber.Map{"users": users,
		"Title": "Hello, World!"}, "layouts/main"); err != nil {
		c.Status(500).Send(err.Error())
	}
}

func ShowSecrect(c *fiber.Ctx) {
	if err := c.Render("Home/secrect", fiber.Map{"msg": "I like dogs"}, "layouts/main"); err != nil {
		c.Status(500).Send(err.Error())
	}
}

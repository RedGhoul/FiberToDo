package controllers

import (
	"github.com/RedGhoul/fibertodo/repos"
	"github.com/RedGhoul/fibertodo/utils"
	"github.com/gofiber/fiber/v2"
)

func Show_Register_Form(c *fiber.Ctx) error {
	utils.Render(c, "Auth/register", "layouts/main", fiber.Map{})
}

func Post_Register_Form(c *fiber.Ctx) error {
	username := c.FormValue("username")
	password1 := c.FormValue("password")
	password2 := c.FormValue("password2")
	if password1 != password2 {
		c.SendString("Your passwords didn't match")
	}

	if !repos.CheckIfUserExists(username) {
		if repos.CreateUser(username, username, password1) {
			c.Redirect("/Login")
		}
	}
	return nil
}

func Show_Login_Form(c *fiber.Ctx) error {
	utils.Render(c, "Auth/login", "layouts/main", fiber.Map{})
}

func Post_Login_Form(c *fiber.Ctx) error {
	username := c.FormValue("username")
	password := c.FormValue("password")
	didmatch, curuser := utils.MatchPasswords(username, password)
	if didmatch {
		utils.SetAuthCookie(curuser, c)
		c.Redirect("/TodoLists")
	} else {
		c.SendString("The entered details do not match our records.")
	}
}

func Logout(c *fiber.Ctx) error {
	if utils.RemoveCookie(c) {
		c.Redirect("/")
	}
	return nil
}

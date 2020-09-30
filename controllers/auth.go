package controllers

import (
	"github.com/RedGhoul/fibertodo/repos"
	"github.com/RedGhoul/fibertodo/utils"
	"github.com/gofiber/fiber"
)

func ShowRegisterForm(c *fiber.Ctx) {
	utils.Render(c, "Auth/register", "layouts/main", fiber.Map{})
}

func PostRegisterForm(c *fiber.Ctx) {
	username := c.FormValue("username")
	password1 := c.FormValue("password")
	password2 := c.FormValue("password2")
	if password1 != password2 {
		c.Send("Your passwords didn't match")
	}

	if !repos.CheckIfUserExists(username) {
		if repos.CreateUser(username, username, password1) {
			c.Redirect("/Login")
		}
	}
	c.Send("Could not register")
}

func ShowLoginForm(c *fiber.Ctx) {
	utils.Render(c, "Auth/login", "layouts/main", fiber.Map{})
}

func PostLoginForm(c *fiber.Ctx) {
	username := c.FormValue("username")
	password := c.FormValue("password")
	didmatch, curuser := utils.MatchPasswords(username, password)
	if didmatch {
		utils.SetAuthCookie(curuser, c)
		c.Redirect("/TodoLists")
	} else {
		c.Send("The entered details do not match our records.")
	}
}

func Logout(c *fiber.Ctx) {
	if utils.RemoveCookie(c) {
		c.Redirect("/")
	}
	c.Redirect(string(c.Fasthttp.Request.Header.Referer()))
}

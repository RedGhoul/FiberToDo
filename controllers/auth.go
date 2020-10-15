package controllers

import (
	"fibertodo/literals"
	"fibertodo/repos"
	"fibertodo/utils"
	"log"

	"github.com/gofiber/fiber/v2"
)

func Show_Register_Form(c *fiber.Ctx) error {
	return utils.Render(c, "Auth/register", "layouts/main", fiber.Map{})
}

func Post_Register_Form(c *fiber.Ctx) error {
	username := c.FormValue("username")
	password1 := c.FormValue("password")
	password2 := c.FormValue("password2")
	if password1 != password2 {
		return c.Redirect(literals.SysRoutes.Register)
	}

	if !repos.CheckIfUserExists(username) {
		if repos.CreateUser(username, username, password1) {
			return c.Redirect(literals.SysRoutes.Login)
		}
	}
	return c.Redirect(literals.SysRoutes.Home)
}

func Show_Login_Form(c *fiber.Ctx) error {
	return utils.Render(c, "Auth/login", "layouts/main", fiber.Map{})
}

func Post_Login_Form(c *fiber.Ctx) error {
	username := c.FormValue("username")
	password := c.FormValue("password")
	didmatch, curuser := utils.MatchPasswords(username, password)
	if didmatch {
		utils.SetAuthCookie(curuser, c)
		return c.Redirect(literals.SysRoutes.Todolists)
	} else {
		//"The entered details do not match our records."
		log.Fatal("The entered details do not match our records.")
		return c.Redirect(literals.SysRoutes.Login)
	}
}

func Logout(c *fiber.Ctx) error {
	if utils.RemoveCookie(c) {
		return c.Redirect(literals.SysRoutes.Home)
	}
	return c.Redirect(c.Route().Path)
}

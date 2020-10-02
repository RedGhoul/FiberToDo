package utils

import (
	"fmt"
	"log"

	"github.com/RedGhoul/fibertodo/models"
	"github.com/RedGhoul/fibertodo/providers"
	"github.com/RedGhoul/fibertodo/repos"
	"github.com/gofiber/fiber"
)

func MatchPasswords(username string, password string) (bool, *models.User) {
	curuser := repos.GetUserByUsername(username)
	match, err := providers.HashProvider().MatchHash(password, curuser.Password)
	if err != nil || match == false {
		if err != nil {
			log.Fatalf("Error when matching hash for password: %v", err)
		}
		return false, nil
	}
	return true, &curuser
}

func SetAuthCookie(curuser *models.User, c *fiber.Ctx) {
	store := providers.SessionProvider().Get(c)
	defer store.Save()
	store.Set("userid", curuser.ID)
}

func RemoveCookie(c *fiber.Ctx) bool {
	if providers.IsAuthenticated(c) {
		store := providers.SessionProvider().Get(c)
		store.Delete("userid")
		store.Save()
		c.ClearCookie()
		return true
	}
	return false
}

func GetUserId(c *fiber.Ctx) uint {
	if providers.IsAuthenticated(c) {
		store := providers.SessionProvider().Get(c)
		return uint(store.Get("userid").(int64))
	}
	return 0
}

func CheckAuth() fiber.Handler {
	return func(c *fiber.Ctx) {
		// Filter request to skip middleware
		if providers.IsAuthenticated(c) {
			c.Next()
			return
		}
		c.Redirect("/Login")
	}
}

func Render(c *fiber.Ctx, pageName string, layoutName string, data fiber.Map) {
	data["signedIn"] = providers.IsAuthenticated(c)
	fmt.Println(data)
	if err := c.Render(pageName, data, layoutName); err != nil {
		c.Status(500).Send(err.Error())
	}
}

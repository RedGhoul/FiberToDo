package utils

import (
	"log"

	"fibertodo/literals"
	"fibertodo/models"
	"fibertodo/providers"
	"fibertodo/repos"

	"github.com/gofiber/fiber/v2"
)

func MatchPasswords(username string, password string) (bool, *models.User) {
	curuser := repos.GetUserByUsername(username)
	if curuser.ID == 0 {
		return false, nil
	}
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
	return func(c *fiber.Ctx) error {
		// Filter request to skip middleware
		if providers.IsAuthenticated(c) {
			return c.Next()

		}
		return c.Redirect("/Login")
	}
}

func Render(c *fiber.Ctx, pageName string, layoutName string, data fiber.Map) error {
	data["signedIn"] = providers.IsAuthenticated(c)
	data["sys"] = literals.SysRoutes
	return c.Render(pageName, data, layoutName)
}

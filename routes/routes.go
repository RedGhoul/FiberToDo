package routes

import (
	"github.com/RedGhoul/fibertodo/controllers"
	"github.com/RedGhoul/fibertodo/utils"
	"github.com/gofiber/fiber"
)

func SetupRoutes(app *fiber.App) {
	setupAuthRoutes(app)
	setupBasicRoutes(app)
	app.Use(utils.CheckAuth())
	setupHiddenRoutes(app)
}

func setupAuthRoutes(app *fiber.App) {
	app.Get("/Login", controllers.ShowLoginForm)
	app.Post("/Login", controllers.PostLoginForm)
	app.Get("/Logout", controllers.PostLogoutForm)
	app.Get("/Register", controllers.ShowRegisterForm)
	app.Post("/Register", controllers.PostRegisterForm)
}

func setupBasicRoutes(app *fiber.App) {
	app.Get("/", controllers.ShowIndex)
}

func setupHiddenRoutes(app *fiber.App) {
	app.Get("/Secrect", controllers.ShowSecrect)
	setupTodoListRoutes(app)
}

func setupTodoListRoutes(app *fiber.App) {
	app.Get("/ListTodoLists", controllers.ShowListOfTodoLists)
	app.Get("/CreateNewToDoList", controllers.GetTodoListForm)
	app.Post("/CreateNewToDoList", controllers.CreateNewToDoList)
	app.Get("/TodoLists/:Id", controllers.ShowTodoList)
	app.Put("/TodoLists/:Id", controllers.UpdateNewToDoList)
	app.Delete("/TodoLists/:Id", controllers.DeleteTodoList)
}

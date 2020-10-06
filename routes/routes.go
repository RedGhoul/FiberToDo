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
	app.Get("/Logout", controllers.Logout)
	app.Get("/Register", controllers.ShowRegisterForm)
	app.Post("/Register", controllers.PostRegisterForm)
}

func setupBasicRoutes(app *fiber.App) {
	app.Get("/", controllers.ShowIndex)
}

func setupHiddenRoutes(app *fiber.App) {
	app.Get("/Secrect", controllers.ShowSecrect)
	setupTodoListRoutes(app)
	setupTodoListItemRoutes(app)
}

func setupTodoListRoutes(app *fiber.App) {
	app.Get("/TodoLists", controllers.ShowListOfTodoLists)
	app.Get("/CreateNewToDoList", controllers.ViewCreateTodoListForm)
	app.Post("/CreateNewToDoList", controllers.CreateNewToDoList)
	app.Get("/TodoList/:Id", controllers.ViewToDoList)
	app.Get("/UpdateTodoList/:Id", controllers.ViewUpdateToDoListForm)
	app.Post("/UpdateTodoList/:Id", controllers.UpdateToDoList)
	app.Delete("/DeleteTodoList/:Id", controllers.DeleteTodoList)
}

func setupTodoListItemRoutes(app *fiber.App) {
	app.Post("/TodoListItem/:Id", controllers.ToDoListItem_Create)
	app.Delete("/TodoListItem/:Id", controllers.ToDoListItem_Delete)
	app.Put("/TodoListItem/:Id", controllers.ToDoListItem_Update)
	app.Get("/TodoListItem/MarkDone/:Id", controllers.ToDoListItem_MarkDone)
}

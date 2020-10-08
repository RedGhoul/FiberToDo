package routes

import (
	"github.com/RedGhoul/fibertodo/controllers"
	"github.com/RedGhoul/fibertodo/literals"
	"github.com/RedGhoul/fibertodo/utils"
	"github.com/gofiber/fiber"
)

func SetupRoutes(app *fiber.App) {
	setup_Auth(app)
	setup_Basic(app)
	setup_Hidden(app)
}

func setup_Auth(app *fiber.App) {
	app.Get(literals.SysRoutes.Login, controllers.Show_Login_Form)
	app.Post(literals.SysRoutes.Login, controllers.Post_Login_Form)
	app.Get(literals.SysRoutes.Logout, controllers.Logout)
	app.Get(literals.SysRoutes.Register, controllers.Show_Register_Form)
	app.Post(literals.SysRoutes.Register, controllers.Post_Register_Form)
}

func setup_Basic(app *fiber.App) {
	app.Get(literals.SysRoutes.Home, controllers.Show_Index)
}

func setup_Hidden(app *fiber.App) {
	app.Use(utils.CheckAuth())
	setup_TodoList(app)
	setup_TodoListItem(app)
}

func setup_TodoList(app *fiber.App) {
	app.Get(literals.SysRoutes.Todolists, controllers.Show_List_Of_TodoLists)
	app.Get(literals.SysRoutes.Createtodolists, controllers.Show_Create_TodoList_Form)
	app.Post(literals.SysRoutes.Createtodolists, controllers.Create_ToDoList)
	app.Get(literals.SysRoutes.Todolistid, controllers.Show_ToDoList)
	app.Get(literals.SysRoutes.Updatetodolist, controllers.Show_Update_ToDoList_Form)
	app.Post(literals.SysRoutes.Updatetodolist, controllers.Update_ToDoList)
	app.Delete(literals.SysRoutes.Deletetodolist, controllers.Delete_TodoList)
}

func setup_TodoListItem(app *fiber.App) {
	app.Post(literals.SysRoutes.Todolistitem, controllers.Create_ToDoList_Item)
	app.Delete(literals.SysRoutes.Todolistitem, controllers.Delete_ToDoList_Item)
	app.Put(literals.SysRoutes.Todolistitem, controllers.Update_ToDoList_Item)
	app.Get(literals.SysRoutes.Todolistitemdone, controllers.Mark_Done_ToDoList_Item)
}

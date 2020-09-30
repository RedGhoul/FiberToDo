package controllers

import (
	"github.com/RedGhoul/fibertodo/repos"
	"github.com/RedGhoul/fibertodo/utils"
	"github.com/gofiber/fiber"
)

var HomePage = "/TodoLists"

/*
	Show a list of todo lists you have
*/
func ShowListOfTodoLists(c *fiber.Ctx) {

	if userID := utils.GetUserId(c); userID != 0 {
		todolists := repos.GetAllTodoListsByUserId(userID)
		utils.Render(c, "ToDoList/index", "layouts/main", fiber.Map{"todolists": todolists})
	} else {
		c.Redirect(HomePage)
	}

}

/*
	Show a todo list
*/
func ShowTodoList(c *fiber.Ctx) {
	utils.Render(c, "ToDoList/todolist", "layouts/main", fiber.Map{"ToDoListId": c.Params("Id")})
}

/*
	Delete a todo list
*/
func DeleteTodoList(c *fiber.Ctx) {
	//todoListId := c.Params("Id")
	c.JSON("OK")
}

/*
	Create a todo list
*/
func CreateNewToDoList(c *fiber.Ctx) {
	todolistname := c.FormValue("todolistname")
	if len(todolistname) == 0 {
		c.Redirect("/CreateNewToDoList")
	}
	if userID := utils.GetUserId(c); userID != 0 {
		repos.CreateTodoList(todolistname, userID)
		c.Redirect(HomePage)
	} else {
		c.Redirect("/CreateNewToDoList")
	}
}

/*
	Get a todo list form
*/
func GetTodoListForm(c *fiber.Ctx) {
	utils.Render(c, "ToDoList/createToDoList", "layouts/main", fiber.Map{})
}

/*
	Update a todo list
*/
func UpdateToDoList(c *fiber.Ctx) {
	c.JSON("OK")
}

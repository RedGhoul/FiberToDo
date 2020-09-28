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
		if err := c.Render("ToDoList/index", fiber.Map{"todolists": todolists}, "layouts/main"); err != nil {
			c.Status(500).Send(err.Error())
		}
	} else {
		c.Redirect(HomePage)
	}

}

/*
	Show a todo list
*/
func ShowTodoList(c *fiber.Ctx) {
	if err := c.Render("ToDoList/todolist",
		fiber.Map{"ToDoListId": c.Params("Id")},
		"layouts/main"); err != nil {
		c.Status(500).Send(err.Error())
	}
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
	if err := c.Render("ToDoList/createToDoList",
		fiber.Map{},
		"layouts/main"); err != nil {
		c.Status(500).Send(err.Error())
	}
}

/*
	Update a todo list
*/
func UpdateToDoList(c *fiber.Ctx) {
	c.JSON("OK")
}

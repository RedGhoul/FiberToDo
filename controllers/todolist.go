package controllers

import (
	"log"
	"strconv"

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
	View a todo list
*/
func ViewToDoList(c *fiber.Ctx) {
	if userID := utils.GetUserId(c); userID != 0 {
		value, _ := strconv.Atoi(c.Params("Id"))
		todolist := repos.GetTodolistByIDAndUserId(uint(value), userID)
		utils.Render(c, "ToDoList/view", "layouts/main", fiber.Map{"todolist": todolist})
	} else {
		c.Redirect(HomePage)
	}
}

/*
	View update form for todo list
*/
func ViewUpdateToDoListForm(c *fiber.Ctx) {
	if userID := utils.GetUserId(c); userID != 0 {
		value, _ := strconv.Atoi(c.Params("Id"))
		todolist := repos.GetTodolistByID(uint(value))
		log.Println(todolist)
		utils.Render(c, "ToDoList/index", "layouts/main", fiber.Map{"todolist": todolist})
	} else {
		c.Redirect(HomePage)
	}
	c.JSON("OK")
}

/*
	Update a todo list
*/
func UpdateToDoList(c *fiber.Ctx) {
	if userID := utils.GetUserId(c); userID != 0 {
		value, _ := strconv.Atoi(c.Params("Id"))
		todolist := repos.GetTodolistByID(uint(value))
		log.Println(todolist)
		utils.Render(c, "ToDoList/index", "layouts/main", fiber.Map{"todolist": todolist})
	} else {
		c.Redirect(HomePage)
	}
	c.JSON("OK")
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
func ViewCreateTodoListForm(c *fiber.Ctx) {
	utils.Render(c, "ToDoList/create", "layouts/main", fiber.Map{})
}

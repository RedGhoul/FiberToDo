package controllers

import (
	"strconv"

	"fibertodo/repos"
	"fibertodo/utils"

	"github.com/gofiber/fiber/v2"
)

var HomePage = "/TodoLists"

/*
	Show a list of todo lists you have
*/
func Show_List_Of_TodoLists(c *fiber.Ctx) error {

	if userID := utils.GetUserId(c); userID != 0 {
		todolists := repos.GetAllTodoListsByUserId(userID)
		return utils.Render(c, "ToDoList/index", "layouts/main", fiber.Map{"todolists": todolists})
	} else {
		c.Redirect(HomePage)
	}
	return nil
}

/*
	Show a todo list
*/
func Show_ToDoList(c *fiber.Ctx) error {
	if userID := utils.GetUserId(c); userID != 0 {
		value, _ := strconv.Atoi(c.Params("Id"))
		todolist := repos.GetTodolistByIDAndUserId(uint(value), userID)

		return utils.Render(c, "ToDoList/view", "layouts/main", fiber.Map{"todolist": todolist})
	} else {
		c.Redirect(HomePage)
	}
	return nil
}

/*
	Show update form for todo list
*/
func Show_Update_ToDoList_Form(c *fiber.Ctx) error {
	if userID := utils.GetUserId(c); userID != 0 {
		value, _ := strconv.Atoi(c.Params("Id"))
		todolist := repos.GetTodolistByID(uint(value))
		return utils.Render(c, "ToDoList/update", "layouts/main", fiber.Map{"todolist": todolist})

	} else {
		c.Redirect(HomePage)
	}
	return c.JSON("OK")
}

/*
	Update a todo list
*/
func Update_ToDoList(c *fiber.Ctx) error {
	if userID := utils.GetUserId(c); userID != 0 {
		value, _ := strconv.Atoi(c.Params("Id"))
		todolistname := c.FormValue("todolistname")
		if len(todolistname) > 0 {
			repos.UpdateTodolistByID(uint(value), todolistname)
			return c.Redirect(HomePage)
		}
		todolist := repos.GetTodolistByID(uint(value))
		return utils.Render(c, "ToDoList/index", "layouts/main", fiber.Map{"todolist": todolist})

	}
	return c.Redirect(HomePage)
}

/*
	Delete a todo list
*/
func Delete_TodoList(c *fiber.Ctx) error {
	if userID := utils.GetUserId(c); userID != 0 {
		value, _ := strconv.Atoi(c.Params("Id"))
		repos.DeleteTodoList(uint(value), userID)

	}
	return c.Redirect(HomePage)
}

/*
	Create a todo list
*/
func Create_ToDoList(c *fiber.Ctx) error {
	todolistname := c.FormValue("todolistname")
	if len(todolistname) == 0 {
		return c.Redirect("/CreateNewToDoList")
	}
	if userID := utils.GetUserId(c); userID != 0 {
		repos.CreateTodoList(todolistname, userID)
		return c.Redirect(HomePage)
	} else {
		return c.Redirect("/CreateNewToDoList")
	}
}

/*
	Show create todolist form
*/
func Show_Create_TodoList_Form(c *fiber.Ctx) error {
	return utils.Render(c, "ToDoList/create", "layouts/main", fiber.Map{})
}

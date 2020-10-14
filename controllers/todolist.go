package controllers

import (
	"strconv"

	"github.com/RedGhoul/fibertodo/repos"
	"github.com/RedGhoul/fibertodo/utils"
	"github.com/gofiber/fiber"
)

var HomePage = "/TodoLists"

/*
	Show a list of todo lists you have
*/
func Show_List_Of_TodoLists(c *fiber.Ctx) {

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
func Show_ToDoList(c *fiber.Ctx) {
	if userID := utils.GetUserId(c); userID != 0 {
		value, _ := strconv.Atoi(c.Params("Id"))
		todolist := repos.GetTodolistByIDAndUserId(uint(value), userID)
		// for _, val := range todolist.ToDoListItems {
		// 	val.CreatedAt = val.CreatedAt.Format("2006-01-02")
		// }

		utils.Render(c, "ToDoList/view", "layouts/main", fiber.Map{"todolist": todolist})
	} else {
		c.Redirect(HomePage)
	}
}

/*
	Show update form for todo list
*/
func Show_Update_ToDoList_Form(c *fiber.Ctx) {
	if userID := utils.GetUserId(c); userID != 0 {
		value, _ := strconv.Atoi(c.Params("Id"))
		todolist := repos.GetTodolistByID(uint(value))
		utils.Render(c, "ToDoList/update", "layouts/main", fiber.Map{"todolist": todolist})
		return
	} else {
		c.Redirect(HomePage)
	}
	c.JSON("OK")
}

/*
	Update a todo list
*/
func Update_ToDoList(c *fiber.Ctx) {
	if userID := utils.GetUserId(c); userID != 0 {
		value, _ := strconv.Atoi(c.Params("Id"))
		todolistname := c.FormValue("todolistname")
		if len(todolistname) > 0 {
			repos.UpdateTodolistByID(uint(value), todolistname)
			c.Redirect(HomePage)
			return
		}
		todolist := repos.GetTodolistByID(uint(value))
		utils.Render(c, "ToDoList/index", "layouts/main", fiber.Map{"todolist": todolist})
		return
	}
	c.Redirect(HomePage)
}

/*
	Delete a todo list
*/
func Delete_TodoList(c *fiber.Ctx) {
	if userID := utils.GetUserId(c); userID != 0 {
		value, _ := strconv.Atoi(c.Params("Id"))
		repos.DeleteTodoList(uint(value), userID)

	}
	c.Redirect(HomePage)
}

/*
	Create a todo list
*/
func Create_ToDoList(c *fiber.Ctx) {
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
	Show create todolist form
*/
func Show_Create_TodoList_Form(c *fiber.Ctx) {
	utils.Render(c, "ToDoList/create", "layouts/main", fiber.Map{})
}

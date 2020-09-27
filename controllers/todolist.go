package controllers

import (
	"github.com/gofiber/fiber"
)

/*
	Show a list of todo lists you have
*/
func ShowListOfTodoListsPage(c *fiber.Ctx) {
	if err := c.Render("ToDo/index", fiber.Map{}, "layouts/main"); err != nil {
		c.Status(500).Send(err.Error())
	}
}

/*
	Show a todo list
*/
func ShowTodoListPage(c *fiber.Ctx) {
	c.Send("ShowTodoListPage")
}

/*
	Delete a todo list
*/
func DeleteTodoList(c *fiber.Ctx) {
	if err := c.Render("ToDo/todolist", fiber.Map{}, "layouts/main"); err != nil {
		c.Status(500).Send(err.Error())
	}
}

/*
	Create a todo list
*/
func CreateNewToDoList(c *fiber.Ctx) {
	c.Send("CreateNewToDoList")
}

/*
	Update a todo list
*/
func UpdateNewToDoList(c *fiber.Ctx) {
	c.Send("CreateNewToDoList")
}

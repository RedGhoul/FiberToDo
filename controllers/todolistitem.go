package controllers

import (
	"github.com/gofiber/fiber"
)

// ToDoListItem

/*
	Create a todo list item
*/
func CreateNewToDoListItem(c *fiber.Ctx) {
	c.Send("CreateNewToDoListItem")
}

/*
	Delete a todo list item
*/
func DeleteNewToDoListItem(c *fiber.Ctx) {
	c.Send("DeleteNewToDoListItem")
}

/*
	Update a todo list item
*/
func UpdateNewToDoListItem(c *fiber.Ctx) {
	c.Send("UpdateNewToDoListItem")
}

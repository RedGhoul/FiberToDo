package controllers

import (
	"log"
	"strconv"

	"github.com/RedGhoul/fibertodo/models"
	"github.com/RedGhoul/fibertodo/repos"
	"github.com/RedGhoul/fibertodo/utils"
	"github.com/gofiber/fiber"
)

// ToDoListItem

/*
	Create a todo list item
*/
func ToDoListItem_Create(c *fiber.Ctx) {
	if userID := utils.GetUserId(c); userID != 0 {
		// Instantiate new Product struct
		p := new(models.TodoListItemDto)

		//  Parse body into product struct
		if err := c.BodyParser(p); err != nil {
			log.Println(err)
			c.Status(400).JSON(&fiber.Map{
				"success": false,
				"message": err,
			})
			return
		}
		repos.Create_TodoListItem(p.Title, p.ListId)
		c.Status(200).JSON(&fiber.Map{
			"success": true,
		})
		return
	} else {
		c.Redirect(HomePage)
	}
}

/*
	Delete a todo list item
*/
func ToDoListItem_Delete(c *fiber.Ctx) {
	if userID := utils.GetUserId(c); userID != 0 {
		value, _ := strconv.Atoi(c.Params("Id"))
		repos.TodoListItem_Delete(uint(value))
		c.Status(200).JSON(&fiber.Map{
			"success": true,
		})
		return
	} else {
		c.Redirect(HomePage)
	}
}

/*
	Update a todo list item
*/
func ToDoListItem_Update(c *fiber.Ctx) {
	c.Send("UpdateNewToDoListItem")
}

/*
	Update a todo list item
*/
func ToDoListItem_MarkDone(c *fiber.Ctx) {
	if userID := utils.GetUserId(c); userID != 0 {
		value, _ := strconv.Atoi(c.Params("Id"))
		repos.TodoListItem_ToggleDone(uint(value))
		c.Status(200).JSON(&fiber.Map{
			"success": true,
		})
		return
	} else {
		c.Redirect(HomePage)
	}
}

package controllers

import (
	"log"

	"github.com/RedGhoul/fibertodo/models"
	"github.com/RedGhoul/fibertodo/repos"
	"github.com/RedGhoul/fibertodo/utils"
	"github.com/gofiber/fiber"
)

// ToDoListItem

/*
	Create a todo list item
*/
func CreateNewToDoListItem(c *fiber.Ctx) {
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
func DeleteNewToDoListItem(c *fiber.Ctx) {
	c.Send("DeleteNewToDoListItem")
}

/*
	Update a todo list item
*/
func UpdateNewToDoListItem(c *fiber.Ctx) {
	c.Send("UpdateNewToDoListItem")
}

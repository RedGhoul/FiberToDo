package controllers

import (
	"log"
	"strconv"

	"github.com/RedGhoul/fibertodo/models"
	"github.com/RedGhoul/fibertodo/repos"
	"github.com/RedGhoul/fibertodo/utils"
	"github.com/gofiber/fiber/v2"
)

// ToDoListItem

/*
	Create a todo list item
*/
func Create_ToDoList_Item(c *fiber.Ctx) error {
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
		createdItem := repos.Create_TodoListItem(p.Title, p.ListId)
		things := createdItem.CreatedAt.Format("2006-01-02")

		c.Status(200).JSON(&fiber.Map{
			"success":   true,
			"id":        createdItem.ID,
			"create_at": things,
		})
		return
	} else {
		c.Redirect(HomePage)
	}
}

/*
	Delete a todo list item
*/
func Delete_ToDoList_Item(c *fiber.Ctx) error {
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
func Update_ToDoList_Item(c *fiber.Ctx) error {
	c.Send("UpdateNewToDoListItem")
}

/*
	Update a todo list item
*/
func Mark_Done_ToDoList_Item(c *fiber.Ctx) error {
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

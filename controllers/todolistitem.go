package controllers

import (
	"log"
	"strconv"

	"fibertodo/models"
	"fibertodo/repos"
	"fibertodo/utils"

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
			return c.Status(400).JSON(&fiber.Map{
				"success": false,
				"message": err,
			})

		}
		createdItem := repos.Create_TodoListItem(p.Title, p.ListId, userID)
		things := createdItem.CreatedAt.Format("2006-01-02")

		return c.Status(200).JSON(&fiber.Map{
			"success":   true,
			"id":        createdItem.ID,
			"create_at": things,
		})
	} else {
		return c.Redirect(HomePage)
	}
}

/*
	Delete a todo list item
*/
func Delete_ToDoList_Item(c *fiber.Ctx) error {
	if userID := utils.GetUserId(c); userID != 0 {
		value, _ := strconv.Atoi(c.Params("Id"))
		repos.TodoListItem_Delete(uint(value))
		return c.Status(200).JSON(&fiber.Map{
			"success": true,
		})

	} else {
		return c.Redirect(HomePage)
	}
}

/*
	Update a todo list item
*/
func Update_ToDoList_Item(c *fiber.Ctx) error {
	if userID := utils.GetUserId(c); userID != 0 {
		value, _ := strconv.Atoi(c.Params("Id"))
		todoItem := repos.Get_Todolist_Item_By_ID_And_UserId(uint(value), userID)
		return utils.Render(c, "TodoItem/update", "layouts/main", fiber.Map{"todoItem": todoItem})

	} else {
		return c.Redirect(HomePage)
	}
}

/*
	Update a todo list item
*/
func Mark_Done_ToDoList_Item(c *fiber.Ctx) error {
	if userID := utils.GetUserId(c); userID != 0 {
		value, _ := strconv.Atoi(c.Params("Id"))
		repos.TodoListItem_ToggleDone(uint(value))
		return c.Status(200).JSON(&fiber.Map{
			"success": true,
		})

	} else {
		return c.Redirect(HomePage)
	}
}

package repos

import (
	"github.com/RedGhoul/fibertodo/database"
	"github.com/RedGhoul/fibertodo/models"
)

func GetAll_TodoListItems_ByListId(listId int) []models.TodoListItem {
	db := database.DBConn
	var items []models.TodoListItem
	//"&" generates a pointer
	//Find fill in that book array
	db.Where("todo_list_refer = ?", listId).Find(&items)
	return items
}

func Create_TodoListItem(title string, listId uint) {
	db := database.DBConn
	var newTodoListItem models.TodoListItem
	newTodoListItem.Title = title
	newTodoListItem.TodoListRefer = listId
	db.Create(&newTodoListItem)
}

func MarkDone_TodoListItem(itemId uint) bool {
	db := database.DBConn

	var item models.TodoListItem
	db.First(&item, itemId)
	if item.ID == 0 {
		return false
	}
	if item.Done == true {
		item.Done = false
	} else {
		item.Done = true
	}
	db.Save(&item)
	return true
}

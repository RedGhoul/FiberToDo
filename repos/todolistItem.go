package repos

import (
	"fibertodo/database"
	"fibertodo/models"
)

func GetAll_TodoListItems_ByListId(listId int) []models.TodoListItem {
	db := database.DBConn
	var items []models.TodoListItem
	//"&" generates a pointer
	//Find fill in that book array
	db.Where("todo_list_refer = ?", listId).Find(&items)
	return items
}

func Create_TodoListItem(title string, listId uint, userId uint) models.TodoListItem {
	db := database.DBConn
	var newTodoListItem models.TodoListItem
	newTodoListItem.Title = title
	newTodoListItem.TodoListRefer = listId
	newTodoListItem.UserRefer = userId
	db.Create(&newTodoListItem)
	return newTodoListItem
}

func TodoListItem_ToggleDone(itemId uint) bool {
	db := database.DBConn

	var item models.TodoListItem
	db.First(&item, itemId)
	if item.ID == 0 {
		return false
	}
	item.Done = !item.Done
	db.Save(&item)
	return true
}

func TodoListItem_Delete(itemId uint) {
	db := database.DBConn

	var item models.TodoListItem
	db.First(&item, itemId)
	db.Delete(&item, itemId)
}

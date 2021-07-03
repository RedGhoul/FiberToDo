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

func Get_Todolist_Item_By_ID_And_UserId(todoItemId uint, userId uint) models.TodoListItem {
	db := database.DBConn
	var todolistItem models.TodoListItem
	db.Where("id = ? AND user_refer = ?",
		todoItemId, userId).Find(&todolistItem)

	return todolistItem
}

func Update_Todolist_Item(todoItemId uint, userId uint, title string) models.TodoListItem {
	db := database.DBConn
	var todolistItem models.TodoListItem
	db.Where("id = ? AND user_refer = ?",
		todoItemId, userId).Find(&todolistItem)
	todolistItem.Title = title
	db.Model(&todolistItem).Update("title", title)
	return todolistItem
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

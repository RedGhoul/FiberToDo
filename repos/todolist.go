package repos

import (
	"github.com/RedGhoul/fibertodo/database"
	"github.com/RedGhoul/fibertodo/models"
)

func GetAllTodoLists() []models.TodoList {
	db := database.DBConn
	var todolists []models.TodoList
	//"&" generates a pointer
	//Find fill in that book array
	db.Find(&todolists)
	return todolists
}

func GetAllTodoListsByUserId(userId uint) []models.TodoList {
	db := database.DBConn
	var todolists []models.TodoList
	//"&" generates a pointer
	//Find fill in that book array
	db.Where("user_refer = ?", userId).Find(&todolists)
	return todolists
}
func GetTodolistByIDAndUserId(todoId uint, userId uint) models.TodoList {
	db := database.DBConn
	var todolist models.TodoList
	db.Where("id = ? AND user_refer = ?", todoId, userId).Find(&todolist)
	var items []models.TodoListItem
	db.Model(&todolist).Association("ToDoListItems").Find(&items)
	todolist.ToDoListItems = items
	return todolist
}

func GetTodolistByID(todoId uint) models.TodoList {
	db := database.DBConn
	var todolist models.TodoList
	var items []models.TodoListItem
	db.Find(&todolist, todoId)
	db.Model(&todolist).Association("ToDoListItems").Find(&items)
	todolist.ToDoListItems = items
	return todolist
}

func CreateTodoList(title string, userId uint) {
	db := database.DBConn
	var newTodoList models.TodoList
	newTodoList.Title = title
	newTodoList.UserRefer = userId
	db.Create(&newTodoList)
}

func DeleteTodoList(todolistId uint, userId uint) bool {
	db := database.DBConn

	var curTodoList models.TodoList
	db.First(&curTodoList, todolistId).Where("userRefer = ?", userId)
	if curTodoList.ID == 0 {
		return false
	}

	db.Delete(&curTodoList)
	return true
}

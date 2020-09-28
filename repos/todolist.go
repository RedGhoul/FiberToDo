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

func GetTodolistByID(todoId int) models.TodoList {
	db := database.DBConn
	var todolist models.TodoList
	var todolistitem []models.TodoListItem
	db.Find(&todolist, todoId)
	db.Model(&todolist).Association("TodoListRefer").Find(&todolistitem)
	todolist.ToDoListItems = todolistitem
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

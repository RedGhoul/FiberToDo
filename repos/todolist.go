package repos

import (
	"log"

	"github.com/RedGhoul/fibertodo/database"
	"github.com/RedGhoul/fibertodo/models"
	"github.com/RedGhoul/fibertodo/providers"
)

func GetAllTodoLists() []models.TodoList {
	db := database.DBConn
	var todolists []models.TodoList
	//"&" generates a pointer
	//Find fill in that book array
	db.Find(&todolists)
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

func CreateUser(username string, email string, password string) bool {
	db := database.DBConn
	newHash, err := providers.HashProvider().CreateHash(password)
	if err != nil {
		log.Println("failed to create user")
		return false
	}
	var newUser models.User
	newUser.Email = username
	newUser.Username = username
	newUser.Password = newHash
	db.Create(&newUser)
	return true
}

func DeleteUser(userId int) bool {
	db := database.DBConn

	var curUser models.User

	db.First(&curUser, userId)
	if curUser.ID == uint(userId) {
		return false
	}

	db.Delete(&curUser)
	return true
}

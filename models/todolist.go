package models

import (
	"gorm.io/gorm"
)

type TodoList struct {
	gorm.Model
	Title         string         `josn:"title"`
	IsDone        bool           `json:"isDone"`
	ToDoListItems []TodoListItem `gorm:"ForeignKey:TodoListRefer"`
	UserRefer     uint           `json:"userRefer"`
}

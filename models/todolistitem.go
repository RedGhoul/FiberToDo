package models

import (
	"gorm.io/gorm"
)

type TodoListItem struct {
	gorm.Model
	Title         string `josn:"title"`
	Author        string `json:"author"`
	Rating        int    `json:"rating"`
	TodoListRefer uint   `json:"todoListRefer"`
}

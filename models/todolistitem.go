package models

import (
	"gorm.io/gorm"
)

type TodoListItem struct {
	gorm.Model
	Title         string `josn:"title"`
	Done          bool   `json:"done"`
	Author        string `json:"author"`
	TodoListRefer uint   `json:"TodoListRefer"`
}

type TodoListItemDto struct {
	ListId uint   `josn:"listId"`
	Title  string `josn:"title"`
	Done   bool   `json:"done"`
}

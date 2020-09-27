package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username  string     `gorm:"unique"`
	Email     string     `json:"email"`
	Password  string     `json:"password"`
	TodoLists []TodoList `gorm:"ForeignKey:UserRefer"`
}

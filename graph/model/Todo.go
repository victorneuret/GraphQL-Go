package model

import "github.com/jinzhu/gorm"

type Todo struct {
	gorm.Model
	Text     string `json:"text"`
	Done     bool   `json:"done"`
	UserName string `json:"user"`
}

type NewTodo struct {
	Text     string `json:"text"`
	Done     bool   `json:"done"`
	UserName string `json:"userName"`
}

type UpdateTodo struct {
	ID       int    `json:"id"`
	Done     bool   `json:"done"`
	UserName string `json:"userName"`
}

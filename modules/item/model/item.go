package model

import "RESTful_API_with_Go_and_Gin/common"

const (
	StatusDone    Status = "Done"
	StatusDoing   Status = "Doing"
	StatusDeleted Status = "Deleted"
)

type Status string

type TodoItem struct {
	common.SQLModel
	Title       string  `json:"title" gorm:"column:title;"`
	Description string  `json:"description" gorm:"column:description;"`
	Status      *Status `json:"status" gorm:"column:status;"`
}

func (TodoItem) TableName() string {
	return "todo_items"
}

type TodoItemCreation struct {
	Id          int     `json:"-" gorm:"column:id;"`
	Title       string  `json:"title" gorm:"column:title"`
	Description string  `json:"description" gorm:"column:description"`
	Status      *Status `json:"status" gorm:"column:status;"`
}

func (TodoItemCreation) TableName() string {
	return TodoItem{}.TableName()
}

type TodoItemUpdate struct {
	Id          int     `json:"-" gorm:"column:id;"`
	Title       *string `json:"title" gorm:"column:title"`
	Description *string `json:"description" gorm:"column:description"`
	Status      *string `json:"status" gorm:"column:status"`
}

func (TodoItemUpdate) TableName() string {
	return TodoItem{}.TableName()
}

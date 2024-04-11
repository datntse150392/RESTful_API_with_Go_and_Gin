package main

import (
	"encoding/json"
	"fmt"
	"time"
)

type TodoItem struct {
	Id          int        `json:"id"`
	Title       string     `json:"title"`
	Description string     `json:"description"`
	Status      string     `json:"status"`
	CreatedAt   *time.Time `json:"created_at"`
	UpdatedAt   *time.Time `json:"updated_at"`
}

func main() {
	now := time.Now().UTC()

	item := TodoItem{
		Id:          1,
		Title:       "This is taks 1",
		Description: "This is taks 2",
		Status:      "Done",
		CreatedAt:   &now,
		UpdatedAt:   &now,
	}

	jsonData, err := json.Marshal(item)

	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(jsonData))
}

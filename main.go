package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
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

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": item,
		})
	})
	r.Run(":3000") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

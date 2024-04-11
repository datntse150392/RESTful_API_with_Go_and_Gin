package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"net/http"
	"os"
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
	dsn := os.Getenv("DB_CONN_STR")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("Connected to DB", db)
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

	//CRUD: Create, Read, Update, Delete
	// POST /v1/items (create a new item)
	// GET /v1/items (get list item) or /v1/items?page=1
	// GET /v1/items/:id (get an item detail by id
	// (PUT/PATCH) /v1/items/:id (update an item by id)
	// DELETE /v1/items/:id (Delete an item by id)

	v1 := r.Group("v1")
	{
		items := v1.Group("/items")
		{
			items.POST("")
			items.GET("")
			items.GET("/:id")
			items.PUT("/:id")
			items.DELETE("/:id")
		}
	}
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": item,
		})
	})
	r.Run(":3000") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

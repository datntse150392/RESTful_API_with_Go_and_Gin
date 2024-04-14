package main

import (
	"RESTful_API_with_Go_and_Gin/common"
	"RESTful_API_with_Go_and_Gin/modules/item/model"
	"RESTful_API_with_Go_and_Gin/services/item"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"net/http"
	"os"
	"strconv"
)

func main() {
	dsn := os.Getenv("DB_CONN_STR")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("Connected to DB", db)
	r := gin.Default()

	//CRUD: Create, Read, Update, Delete
	// POST /v1/items (create a new item)
	// GET /v1/items (get list item) or /v1/items?page=1
	// GET /v1/items/:id (get an item detail by id
	// (PUT/PATCH) /v1/items/:id (update an item by id)
	// DELETE /v1/items/:id (Delete an item by id)S

	v1 := r.Group("v1")
	{
		items := v1.Group("/items")
		{
			items.POST("", item.CreateItem(db))
			items.GET("", GetListItems(db))
			items.GET("/:id", GetDetailItem(db))
			items.PUT("/:id", UpdateItem(db))
			items.DELETE("/:id", DeleteITem(db))
		}
	}
	r.Run(":3000") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

func CreateItem(db *gorm.DB) func(*gin.Context) {
	return func(c *gin.Context) {
		var data model.TodoItemCreation

		if err := c.ShouldBind(&data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"err": err.Error(),
			})
			return
		}

		if err := db.Create(&data).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"err": err.Error(),
			})

			return
		}

		c.JSON(http.StatusCreated, common.SimpleSuccessRes(data.Id))
	}
}

func GetDetailItem(db *gorm.DB) func(ctx *gin.Context) {
	return func(c *gin.Context) {
		var data model.TodoItem

		id, err := strconv.Atoi(c.Param("id"))

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"err": err.Error(),
			})
		}

		data.Id = id
		db.First(&data)

		if err := db.First(&data).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{
				"err": err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, common.SimpleSuccessRes(data))
	}
}

func UpdateItem(db *gorm.DB) func(*gin.Context) {
	return func(c *gin.Context) {
		var data model.TodoItemUpdate

		id, err := strconv.Atoi(c.Param("id"))

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"err": err.Error(),
			})
			return
		}

		if err := c.ShouldBind(&data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"err": err.Error(),
			})
			return
		}

		data.Id = id
		db.Updates(data)

		c.JSON(http.StatusOK, common.SimpleSuccessRes("Updated"))
	}
}

// Soft Delete
func DeleteITem(db *gorm.DB) func(*gin.Context) {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"err": err.Error(),
			})
			return
		}

		// Ví dụ ở đây không có truyền tham chiếu đến struct data thì làm sao nó biết được table name của mình ?
		if err := db.Table(model.TodoItem{}.TableName()).Where("id = ?", id).Updates(map[string]interface{}{
			"status": "Deleted",
		}).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{
				"err": err.Error(),
			})
		}

		c.JSON(http.StatusOK, common.SimpleSuccessRes("Deleted"))
	}
}

func GetListItems(db *gorm.DB) func(*gin.Context) {
	return func(c *gin.Context) {
		var paging common.Paging

		if err := c.ShouldBind(&paging); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"err": err.Error(),
			})
			return
		}

		paging.Process()

		var result []model.TodoItem

		db = db.Where("status <> ?", "Deleted")

		if err := db.Table(model.TodoItem{}.TableName()).Count(&paging.Total).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"err": err.Error(),
			})
			return
		}

		if err := db.Table(model.TodoItem{}.TableName()).Order("id desc").Offset((paging.Page - 1) * paging.Limit).
			Limit(paging.Limit).
			Find(&result).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"err": err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"data":  result,
			"total": paging.Total,
		})
	}
}

package item

import (
	"RESTful_API_with_Go_and_Gin/common"
	"RESTful_API_with_Go_and_Gin/modules/item/model"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

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

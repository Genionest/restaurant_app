package controller

import (
	"net/http"

	"example.com/m/v2/global"
	"github.com/gin-gonic/gin"
)

type Dish struct {
	ID       uint `gorm:"primaryKey`
	Name     string
	Price    int
	category string
}

func AddDish(ctx *gin.Context) {
	var dish Dish
	if err := ctx.ShouldBindJSON(&dish); err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if err := global.DB.AutoMigrate(&dish); err != nil {
		ctx.IndentedJSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	if err := global.DB.Create(&dish).Error; err != nil {
		ctx.IndentedJSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.IndentedJSON(http.StatusOK, dish)
}

func GetAllMenu(ctx *gin.Context) []string {
	// global.DB.Find()
}

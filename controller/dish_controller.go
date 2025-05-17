package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// AddDish 函数用于处理添加菜品的请求
//
// 参数:
//
//	ctx *gin.Context: Gin框架的上下文对象，用于获取请求参数和返回响应
//
// 返回值:
//
//	无
func AddDish(ctx *gin.Context) {
	var dish Dish
	if err := CreateData(ctx, &dish); err != nil {
		return
	}

	ctx.IndentedJSON(http.StatusOK, dish)
}

// GetDish 是处理HTTP GET请求的处理器函数，用于根据给定的ID获取菜品信息
// 参数：
//
//	ctx *gin.Context: Gin框架的上下文对象，包含了请求和响应信息
//
// 返回值：
//
//	无
func GetDish(ctx *gin.Context) {
	id := ctx.Param("id")
	var dish Dish
	if err := GetData(ctx, &dish, map[string]interface{}{"id": id}); err != nil {
		return
	}
	ctx.IndentedJSON(http.StatusOK, dish)
}

// UpdateDish 更新菜品信息
//
// 参数:
//
//	ctx: gin.Context上下文对象
//	dish: Dish结构体类型的菜品信息
//
// 返回值:
//
//	无
func UpdateDish(ctx *gin.Context) {
	var dish Dish
	if err := UpdateData(ctx, &dish); err != nil {
		return
	}
	// var new_dish Dish
	// GetData(ctx, &new_dish, DataQuery{"id": dish.ID})
	// ctx.IndentedJSON(http.StatusOK, new_dish)
	ctx.IndentedJSON(http.StatusOK, dish)
}

// DeleteDish 函数用于删除菜品
// 参数:
//
//	ctx: *gin.Context - gin框架的上下文对象
//
// 返回值:
//
//	无返回值
func DeleteDish(ctx *gin.Context) {
	id := ctx.Param("id")
	iid, _ := strconv.Atoi(id)
	dish := Dish{ID: uint(iid)}
	if err := DeleteData(ctx, &dish); err != nil {
		return
	}
	ctx.IndentedJSON(http.StatusOK, gin.H{
		"id":  id,
		"msg": "删除成功",
	})
}

// GetAllDishes 函数用于获取所有菜品信息
//
// 参数：
//   - ctx *gin.Context：Gin框架的上下文对象
//
// 返回值：
//
//	无返回值
func GetAllDishes(ctx *gin.Context) {
	var dishes []Dish
	if err := GetAllData(ctx, &dishes, nil); err != nil {
		return
	}
	ctx.IndentedJSON(http.StatusOK, dishes)
}

// GetDishesByCategory 函数根据传入的分类获取对应的菜品列表
// 参数：
//
//	ctx: *gin.Context - gin框架的上下文对象
//
// 返回值：
//
//	无
func GetDishesByCategory(ctx *gin.Context) {
	category := ctx.Param("category")
	var dishes []Dish
	if err := GetAllData(ctx, &dishes, map[string]interface{}{"category": category}); err != nil {
		return
	}
	ctx.IndentedJSON(http.StatusOK, dishes)
}

// GetAllRecords 函数用于获取所有记录信息
//
// 参数：
//   - ctx: *gin.Context - gin 框架的上下文对象，用于获取请求参数和返回响应
//
// 返回值：
//   - 无
func GetAllRecords(ctx *gin.Context) {
	var records []Record
	if err := GetAllData(ctx, &records, nil); err != nil {
		return
	}
	ctx.IndentedJSON(http.StatusOK, records)
}

func GetRecentRecords(ctx *gin.Context) {
	var records []Record
	if err := GetAllData(ctx, &records, nil); err != nil {
		return
	}

}

func GetTotalPrice(ctx *gin.Context) {
	totalPrice := 0
	var bills []Bill
	BindJSON(ctx, &bills)
	for _, bill := range bills {
		var dish Dish
		GetData(ctx, &dish, map[string]interface{}{"id": bill.DishID})
		totalPrice += dish.Price * bill.Count
	}
	ctx.IndentedJSON(http.StatusOK, gin.H{
		"total_price": totalPrice,
	})
}

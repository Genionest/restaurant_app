package controller

import (
	"errors"
	"net/http"

	"example.com/m/v2/global"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type DataQuery map[string]interface{}

// CreateData 函数用于创建一个数据对象，并将其保存到数据库中
//
// 参数：
//
//	ctx *gin.Context: Gin框架的上下文对象，用于处理HTTP请求和响应
//	data *T: 指向任意类型数据的指针，用于存储请求体中解析出的数据
//
// 返回值：
//
//	error: 如果函数执行过程中出现错误，将返回一个非空error对象；否则返回nil
func CreateData[T any](ctx *gin.Context, data *T) error {
	if err := ctx.ShouldBindJSON(data); err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{
			"error":  err.Error(),
			"msg":    "ShouldBindJSON error",
			"struct": *data,
		})
		return err
	}

	if err := global.DB.AutoMigrate(data); err != nil {
		ctx.IndentedJSON(http.StatusInternalServerError, gin.H{
			"error":  err.Error(),
			"msg":    "AutoMigrate error",
			"struct": *data,
		})
		return err
	}

	if err := global.DB.Create(data).Error; err != nil {
		ctx.IndentedJSON(http.StatusInternalServerError, gin.H{
			"error":  err.Error(),
			"msg":    "Create error",
			"struct": *data,
		})
		return err
	}
	return nil
}

// GetData 是一个泛型函数，用于从数据库中获取数据
//
// 参数：
//
//	ctx *gin.Context: gin框架的上下文对象，用于处理HTTP请求和响应
//	data *T: 一个指向泛型类型的指针，用于存储从数据库中获取的数据
//	query DataQuery: 一个DataQuery类型的变量，用于指定查询条件
//
// 返回值：
//
//	error: 如果查询过程中出现错误，则返回错误信息；否则返回nil
//
// 备注：
//
//	如果查询到的记录不存在，函数会以HTTP状态码404返回错误信息和查询条件；
//	如果查询过程中出现其他错误，函数会以HTTP状态码500返回错误信息和查询条件。
func GetData[T any](ctx *gin.Context, data *T, query DataQuery) error {
	if err := global.DB.Where(query).First(data).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			ctx.IndentedJSON(http.StatusNotFound, gin.H{
				"error":  err.Error(),
				"msg":    "Record not found",
				"query":  query,
				"struct": *data,
			})
			return err
		} else {
			ctx.IndentedJSON(http.StatusInternalServerError, gin.H{
				"error":  err.Error(),
				"msg":    "Query error",
				"query":  query,
				"struct": *data,
			})
			return err
		}
	}
	return nil
}

// GetAllData 函数用于根据给定的查询条件从数据库中获取所有数据，并将结果存储到给定的切片中。
//
// 参数：
// ctx: *gin.Context - gin框架的上下文对象，用于处理HTTP请求和响应。
// datas: *[]T - 指向切片类型的指针，用于存储查询结果。T为泛型类型，表示切片中元素的类型。
// query: DataQuery - 查询条件，用于筛选数据库中的数据。
//
// 返回值：
// error - 如果查询过程中发生错误，则返回错误信息；否则返回nil。
func GetAllData[T any](ctx *gin.Context, datas *[]T, query DataQuery) error {
	if err := global.DB.Where(query).Find(datas).Error; err != nil {
		ctx.IndentedJSON(http.StatusInternalServerError, gin.H{
			"error":  err.Error(),
			"msg":    "Query All error",
			"query":  query,
			"struct": *datas,
		})
		return err
	}
	return nil
}

func UpdateData[T any](ctx *gin.Context, data *T) error {
	if err := global.DB.Model(data).Updates(data).Error; err != nil {
		ctx.IndentedJSON(http.StatusInternalServerError, gin.H{
			"error":  err.Error(),
			"msg":    "Update error",
			"struct": *data,
		})
		return err
	}
	return nil
}

// DeleteData 是一个泛型函数，用于删除指定类型的数据
// 参数：
//
//	ctx: *gin.Context，Gin框架的上下文对象
//	data: *T，指向要删除的数据的指针，T 是任意类型
//
// 返回值：
//
//	error，如果删除数据时出现错误，则返回错误对象；否则返回 nil
func DeleteData[T any](ctx *gin.Context, data *T) error {
	if err := global.DB.Delete(data).Error; err != nil {
		ctx.IndentedJSON(http.StatusInternalServerError, gin.H{
			"error":  err.Error(),
			"msg":    "Delete error",
			"struct": *data,
		})
		return err
	}
	return nil
}

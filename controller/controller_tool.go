package controller

import (
	"errors"
	"log"
	"net/http"

	"example.com/m/v2/global"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

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
			"msg":    "ShouldBindJSON error(CreateData)",
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

// CreateDataWithoutBind 函数用于在不绑定 JSON 请求体的情况下，将给定的数据对象保存到数据库中
// 该函数会自动迁移数据库表结构，然后尝试创建数据记录
//
// 参数：
//
//	ctx *gin.Context: Gin 框架的上下文对象，用于处理 HTTP 请求和响应
//	data *T: 指向任意类型数据的指针，该数据将被保存到数据库中
//
// 返回值：
//
//	error: 如果函数执行过程中出现错误，将返回一个非空 error 对象；否则返回 nil
func CreateDataWithoutBind[T any](ctx *gin.Context, data *T) error {
	if err := global.DB.AutoMigrate(data); err != nil {
		ctx.IndentedJSON(http.StatusInternalServerError, gin.H{
			"error":  "AutoMigrate error",
			"msg":    err.Error(),
			"struct": *data,
		})
		return err
	}

	if err := global.DB.Create(data).Error; err != nil {
		ctx.IndentedJSON(http.StatusInternalServerError, gin.H{
			"error":  "Create error",
			"msg":    err.Error(),
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
//	query map[string]interface{}: 一个map类型的变量，用于指定查询条件
//
// 返回值：
//
//	error: 如果查询过程中出现错误，则返回错误信息；否则返回nil
//
// 备注：
//
//	如果查询到的记录不存在，函数会以HTTP状态码404返回错误信息和查询条件；
//	如果查询过程中出现其他错误，函数会以HTTP状态码500返回错误信息和查询条件。
func GetData[T any](ctx *gin.Context, data *T, query map[string]interface{}) error {
	// 这里query map[string]interface{}
	// 不能使用type DataQuery map[string]interface{}
	// 会无法识别
	if err := global.DB.Where(query).First(data).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			log.Println()
			log.Printf("Record not found\n")
			log.Printf("query: %v\n", query)
			log.Printf("struct: %v\n", *data)
			log.Printf("%v\n", err.Error())
			log.Println()
			ctx.IndentedJSON(http.StatusNotFound, gin.H{
				"error": "Record not found",
			})
			return err
		} else {
			log.Println()
			log.Printf("Query error\n")
			log.Printf("query: %v\n", query)
			log.Printf("struct: %v\n", *data)
			log.Printf("%v\n", err.Error())
			log.Println()
			ctx.IndentedJSON(http.StatusInternalServerError, gin.H{
				"error": "Query error",
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
// query: map[string]interface{} - 查询条件，用于筛选数据库中的数据。
//
// 返回值：
// error - 如果查询过程中发生错误，则返回错误信息；否则返回nil。
func GetAllData[T any](ctx *gin.Context, datas *[]T, query map[string]interface{}) error {
	if err := global.DB.Where(query).Find(datas).Error; err != nil {
		log.Println()
		log.Printf("Query All error\n")
		log.Printf("query: %v\n", query)
		log.Printf("struct: %v\n", *datas)
		log.Printf("%v\n", err.Error())
		log.Println()
		ctx.IndentedJSON(http.StatusInternalServerError, gin.H{
			"error": "Query All error",
		})
		return err
	}
	return nil
}

func UpdateData[T any](ctx *gin.Context, data *T) error {
	if err := ctx.ShouldBindJSON(data); err != nil {
		log.Println()
		log.Printf("ShouldBindJSON error(UpdateData)\n")
		log.Printf("query: %v\n", *data)
		log.Printf("%v\n", err.Error())
		log.Println()
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{
			"error": "ShouldBindJSON error(UpdateData)",
		})
		return err
	}

	if err := global.DB.Model(data).Updates(data).Error; err != nil {
		log.Println()
		log.Printf("Update error\n")
		log.Printf("query: %v\n", *data)
		log.Printf("%v\n", err.Error())
		log.Println()
		ctx.IndentedJSON(http.StatusInternalServerError, gin.H{
			"error": "Update error",
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
		log.Println()
		log.Printf("Delete error\n")
		log.Printf("query: %v\n", *data)
		log.Printf("%v\n", err.Error())
		log.Println()
		ctx.IndentedJSON(http.StatusInternalServerError, gin.H{
			"error": "Delete error",
		})
		return err
	}
	return nil
}

// BindJSON 是一个泛型函数，用于将传入的 JSON 数据绑定到指定的结构体指针中。
// 参数 ctx 是 gin 框架的上下文对象，data 是要绑定的结构体指针。
// 如果绑定成功，函数返回 nil；如果绑定失败，函数将返回错误信息。
//
// 参数:
//
//	ctx *gin.Context: gin 框架的上下文对象
//	data *T: 要绑定的结构体指针
//
// 返回值:
//
//	error: 如果绑定失败，返回错误信息；否则返回 nil
func BindJSON[T any](ctx *gin.Context, data *T) error {
	if err := ctx.ShouldBindJSON(data); err != nil {
		log.Println()
		log.Printf("ShouldBindJSON error(My BindJSON)\n")
		log.Printf("query: %v\n", *data)
		log.Printf("%v\n", err.Error())
		log.Println()
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{
			"error": "ShouldBindJSON error(My BindJSON)",
		})
		return err
	}
	return nil
}

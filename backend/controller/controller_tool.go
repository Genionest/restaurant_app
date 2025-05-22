package controller

import (
	"errors"
	"log"
	"net/http"
	"time"

	"example.com/m/v2/global"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
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
//	bool: 如果函数执行过程中出现错误，将返回一个非空true；否则返回false
func CreateData[T any](ctx *gin.Context, data *T) bool {
	if err := ctx.ShouldBindJSON(data); err != nil {
		log.Println()
		log.Printf("ShouldBindJSON error(CreateData)\n")
		log.Printf("error: %s\n", err.Error())
		log.Printf("struct: %+v\n", *data)
		log.Println()
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{
			"error": "ShouldBindJSON error(CreateData)",
		})
		return false
	}

	// 后端启动时自动迁移数据库

	if err := global.DB.Create(data).Error; err != nil {
		log.Println()
		log.Printf("Create error(CreateData)\n")
		log.Printf("error: %s\n", err.Error())
		log.Printf("struct: %+v\n", *data)
		log.Println()
		ctx.IndentedJSON(http.StatusInternalServerError, gin.H{
			"error": "Create error",
		})
		return false
	}
	return true
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
//	bool: 如果函数执行过程中出现错误，将返回一个false；否则返回 true
func CreateDataWithoutBind[T any](ctx *gin.Context, data *T) bool {
	if err := global.DB.AutoMigrate(data); err != nil {
		log.Println()
		log.Printf("AutoMigrate error\n")
		log.Printf("error: %s\n", err.Error())
		log.Printf("struct: %+v\n", *data)
		log.Println()
		ctx.IndentedJSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return false
	}

	if err := global.DB.Create(data).Error; err != nil {
		log.Println()
		log.Printf("Create error\n")
		log.Printf("error: %s\n", err.Error())
		log.Printf("struct: %+v\n", *data)
		log.Println()
		ctx.IndentedJSON(http.StatusInternalServerError, gin.H{
			"error": "Create error",
		})
		return false
	}
	return true
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
//	bool: 如果查询过程中出现错误，则返回错误false；否则返回true
//
// 备注：
//
//	如果查询到的记录不存在，函数会以HTTP状态码404返回错误信息和查询条件；
//	如果查询过程中出现其他错误，函数会以HTTP状态码500返回错误信息和查询条件。
func GetData[T any](ctx *gin.Context, data *T, query map[string]interface{}) bool {
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
				"error": "没有发现记录",
			})
			return false
		} else {
			log.Println()
			log.Printf("Query error\n")
			log.Printf("query: %v\n", query)
			log.Printf("struct: %v\n", *data)
			log.Printf("%v\n", err.Error())
			log.Println()
			ctx.IndentedJSON(http.StatusInternalServerError, gin.H{
				"error": "查询错误",
			})
			return false
		}
	}
	return true
}

// GetAllDatas 函数用于根据给定的查询条件从数据库中获取所有数据，并将结果存储到给定的切片中。
//
// 参数：
// ctx: *gin.Context - gin框架的上下文对象，用于处理HTTP请求和响应。
// datas: *[]T - 指向切片类型的指针，用于存储查询结果。T为泛型类型，表示切片中元素的类型。
// query: map[string]interface{} - 查询条件，用于筛选数据库中的数据。
//
// 返回值：
// bool - 如果查询过程中发生错误，则返回错误false；否则返回true。
func GetAllDatas[T any](ctx *gin.Context, datas *[]T, query map[string]interface{}) bool {
	if err := global.DB.Where(query).Find(datas).Error; err != nil {
		log.Println()
		log.Printf("Query All error\n")
		log.Printf("query: %v\n", query)
		log.Printf("struct: %v\n", *datas)
		log.Printf("%v\n", err.Error())
		log.Println()
		ctx.IndentedJSON(http.StatusInternalServerError, gin.H{
			"error": "批量查询错误",
		})
		return false
	}
	return true
}

// GetManyDatas 函数用于根据给定的 SQL 查询语句和参数，从数据库中获取多条数据，并将结果存储到给定的切片中。
// 这是一个泛型函数，支持处理任意类型的数据。
//
// 参数：
// ctx: *gin.Context - gin 框架的上下文对象，用于处理 HTTP 请求和响应。
// datas: *[]T - 指向切片类型的指针，用于存储查询结果。T 为泛型类型，表示切片中元素的类型。
// query: string - SQL 查询语句，用于指定查询条件。
// args ...interface{} - 可变参数，用于填充查询语句中的占位符。
//
// 返回值：
// bool - 如果查询过程中发生错误，则返回false；否则返回 true。
func GetManyDatas[T any](ctx *gin.Context, datas *[]T, query string, args ...interface{}) bool {
	origin := *datas
	if err := global.DB.Where(query, args...).Find(datas).Error; err != nil {
		log.Println()
		log.Printf("Query Many error\n")
		log.Printf("query: %v\n", origin)
		log.Printf("struct: %v\n", *datas)
		log.Printf("%v\n", err.Error())
		log.Println()
		ctx.IndentedJSON(http.StatusInternalServerError, gin.H{
			"error": "局部查询错误",
		})
		return false
	}
	return true
}

func UpdateData[T any](ctx *gin.Context, data *T) bool {
	if err := ctx.ShouldBindJSON(data); err != nil {
		log.Println()
		log.Printf("ShouldBindJSON error(UpdateData)\n")
		log.Printf("query: %v\n", *data)
		log.Printf("%v\n", err.Error())
		log.Println()
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{
			"error": "更新数据时，匹配数据失败",
		})
		return false
	}

	if err := global.DB.Model(data).Updates(data).Error; err != nil {
		log.Println()
		log.Printf("Update error\n")
		log.Printf("query: %v\n", *data)
		log.Printf("%v\n", err.Error())
		log.Println()
		ctx.IndentedJSON(http.StatusInternalServerError, gin.H{
			"error": "更新失败",
		})
		return false
	}
	return true
}

// DeleteData 是一个泛型函数，用于删除指定类型的数据
// 参数：
//
//	ctx: *gin.Context，Gin框架的上下文对象
//	data: *T，指向要删除的数据的指针，T 是任意类型
//
// 返回值：
//
//	bool，如果删除数据时出现错误，则返回true；否则返回 false
func DeleteData[T any](ctx *gin.Context, data *T) bool {
	if err := global.DB.Delete(data).Error; err != nil {
		log.Println()
		log.Printf("Delete error\n")
		log.Printf("query: %v\n", *data)
		log.Printf("%v\n", err.Error())
		log.Println()
		ctx.IndentedJSON(http.StatusInternalServerError, gin.H{
			"error": "删除错误",
		})
		return false
	}
	return true
}

// BindJSON 是一个泛型函数，用于将传入的 JSON 数据绑定到指定的结构体指针中。
// 参数 ctx 是 gin 框架的上下文对象，data 是要绑定的结构体指针。
// 如果绑定成功，函数返回true；如果绑定失败，函数将返回false。
//
// 参数:
//
//	ctx *gin.Context: gin 框架的上下文对象
//	data *T: 要绑定的结构体指针
//
// 返回值:
//
//	bool: 如果绑定失败，返回false；否则返回true
func BindJSON[T any](ctx *gin.Context, data *T) bool {
	if err := ctx.ShouldBindJSON(data); err != nil {
		log.Println()
		log.Printf("ShouldBindJSON error(My BindJSON)\n")
		log.Printf("query: %v\n", *data)
		log.Printf("%v\n", err.Error())
		log.Println()
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{
			"error": "匹配数据失败",
		})
		return false
	}
	return true
}

func EncryptPassword(data *string) (string, error) {
	// 加密
	hash, err := bcrypt.GenerateFromPassword([]byte(*data), bcrypt.DefaultCost)
	return string(hash), err
}

func CheckPassword(data *string, hash *string) bool {
	// 解密
	err := bcrypt.CompareHashAndPassword([]byte(*hash), []byte(*data))
	if err != nil {
		log.Println()
		log.Printf("CheckPassword error\n")
		log.Printf("error: %s\n", err.Error())
		return false
	}
	return true
}

func GenerateJWT(username *string) (string, error) {
	// 生成token
	// claims := jwt.StandardClaims{}
	/*
		Audience  string `json:"aud,omitempty"`
		ExpiresAt int64  `json:"exp,omitempty"`
		Id        string `json:"jti,omitempty"`
		IssuedAt  int64  `json:"iat,omitempty"`
		Issuer    string `json:"iss,omitempty"`
		NotBefore int64  `json:"nbf,omitempty"`
		Subject   string `json:"sub,omitempty"`
	*/
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, jwt.MapClaims{
		"username": username,
		"nbf":      time.Now().Unix(),
		"exp":      time.Now().Add(time.Hour * 2).Unix(),
		"update":   time.Now().Add(time.Minute * 90).Unix(),
	})
	// 加密
	signKey := []byte("secret")
	signedToken, err := token.SignedString(signKey)
	signedToken = "Bearer " + signedToken
	return signedToken, err
}

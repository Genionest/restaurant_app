package controller

import (
	"errors"
	"log"
	"net/http"

	"example.com/m/v2/global"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CheckUsername(ctx *gin.Context, username string) bool {
	user := &User{}
	if err := global.DB.Where("username = ?", username).First(user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			log.Println()
			log.Printf("CheckUsername success\n")
			log.Println()
			return true
		} else {
			log.Println()
			log.Printf("CheckUsername error\n")
			log.Printf("error: %s\n", err.Error())
			log.Println()
			ctx.IndentedJSON(http.StatusInternalServerError, gin.H{
				"error": "用户名审核错误",
			})
			return false
		}
	}
	log.Println()
	log.Printf("CheckUsername error\n")
	log.Printf("error: %s\n", "Username already exists")
	log.Println()
	ctx.IndentedJSON(http.StatusBadRequest, gin.H{
		"error": "用户名已存在",
	})
	return false
}

func UserRegister(ctx *gin.Context) {
	user := &User{}
	// 检查是否存在相同用户名
	if ok := CheckUsername(ctx, user.Username); !ok {
		return
	}
	// 使用ORM，不必担心sql注入
	// 对密码进行加密
	pwd, err := EncryptPassword(&(user.Password))
	if err != nil {
		ctx.IndentedJSON(http.StatusInternalServerError, gin.H{
			"error": "密码解析失败",
		})
		return
	}
	user.Password = pwd

	if ok := CreateData(ctx, user); !ok {
		return
	}
	// jwt token
	// 生成token
	// token, err := GenerateToken(user)
	ctx.IndentedJSON(http.StatusOK, user)
}

func UserLogin(ctx *gin.Context) {
	// user := &User{}
	var input struct {
		Username string
		Password string
	}
	if ok := BindJSON(ctx, &input); !ok {
		return
	}
	// 验证用户名是否存在
	user := &User{}
	if err := global.DB.Where("username =?", input.Username).First(user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			log.Println()
			log.Printf("Username Not Found\n")
			log.Printf("error: %s\n", err.Error())
			log.Println()
			// 不让人知道用户名是否存在，防止暴力破解
			ctx.IndentedJSON(http.StatusUnauthorized, gin.H{
				"error": "用户名或密码不正确",
			})
			return
		} else {
			log.Println()
			log.Printf("Login error\n")
			log.Printf("error: %s\n", err.Error())
			log.Println()
			ctx.IndentedJSON(http.StatusInternalServerError, gin.H{
				"error": "登录错误",
			})
			return
		}
	}
	// 验证密码是否匹配
	if ok := CheckPassword(&(input.Password), &(user.Password)); !ok {
		log.Println()
		log.Printf("error: %s\n", "Password Not Match")
		log.Println()
		// 不让人知道用户名是否存在，防止暴力破解
		ctx.IndentedJSON(http.StatusUnauthorized, gin.H{
			"error": "用户名或密码不正确",
		})
		return
	}
	// jwt token
	// 生成token
	token, err := GenerateJWT(&(user.Username))
	if err != nil {
		log.Println()
		log.Printf("GenerateJWT error\n")
		log.Printf("error: %s\n", err.Error())
		log.Println()
		ctx.IndentedJSON(http.StatusInternalServerError, gin.H{
			"error": "生成token失败",
		})
		return
	}
	ctx.IndentedJSON(http.StatusOK, gin.H{
		"token": token,
	})
}

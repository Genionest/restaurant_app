package controller

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	/* 相当于 gin.New() + Logger,Recovery中间件
	Logger 中间件：自动记录 HTTP 请求的日志（如请求方法、路径、状态码、耗时等）。
	Recovery 中间件：自动捕获处理请求时发生的 panic，防止程序崩溃，并返回 500 错误。
	*/

	// 跨域请求中间件
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"},
		AllowMethods:     []string{"GET", "POST", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		// AllowOriginFunc: func(origin string) bool {
		// 	return origin == "https://github.com"
		// },
		MaxAge: 12 * time.Hour,
	}))

	api := r.Group("/api")
	{
		api.GET("/get_dish/:id", GetDish)
		api.GET("/get_dishes", GetAllDishes)
		api.GET("/get_dishes_by_category/:category", GetDishesByCategory)
	}
	// api.USE(middleware)

	admin := r.Group("/admin")
	{
		admin.POST("/add_dish", AddDish)
		admin.POST("/update_dish", UpdateDish)
		admin.POST("/delete_dish", DeleteDish)
	}

	return r
}

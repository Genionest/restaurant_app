package controller

import (
	"net"
	"net/url"
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

	r.SetTrustedProxies([]string{"127.0.0.1"}) // 信任本地代理
	// r.TrustedPlatform = gin.PlatformGoogleAppEngine // 信任 Google App Engine 平台

	// 日志中间件测试, 要放在前面，不然被cors中间件拦截了
	// r.Use(middleware.RequestLogger())

	// 跨域请求中间件
	r.Use(cors.New(cors.Config{
		// AllowOrigins:     []string{"http://127.0.0.1:5173"},
		AllowOriginFunc: func(origin string) bool {
			// 允许无 Origin 头的请求（如直接 API 调用）
			if origin == "" {
				return true
			}

			// 解析 Origin 的 URL
			parsedOrigin, err := url.Parse(origin)
			if err != nil {
				return false
			}

			// 提取主机名（IP 或域名）
			host := parsedOrigin.Hostname()

			// 验证 IP 是否为允许的地址
			allowedIP := "127.0.0.1"
			if net.ParseIP(host).String() == allowedIP {
				return true
			}

			// 如果需要验证域名，可以在此添加逻辑
			// if host == "example.com" { ... }

			return false
		},
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

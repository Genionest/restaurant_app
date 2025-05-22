package controller

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func MyAllowOriginFunc(origin string) bool {
	// 允许无 Origin 头的请求（如直接 API 调用）
	if origin == "" {
		return true
	}
	if origin == "http://127.0.0.1:5173" {
		return true
	} else if origin == "http://localhost:5173" {
		return true
	}

	log.Println()
	log.Printf("origin don't be allowed: %v\n", origin)
	log.Println()

	// // 解析 Origin 的 URL
	// parsedOrigin, err := url.Parse(origin)
	// if err != nil {
	// 	return false
	// }

	// // 提取主机名（IP 或域名）
	// host := parsedOrigin.Hostname()
	// 验证 IP 是否为允许的地址
	// allowedIP := "127.0.0.1"
	// if net.ParseIP(host).String() == allowedIP {
	// 	return true
	// }

	// 如果需要验证域名，可以在此添加逻辑
	// if host == "example.com" { ... }

	return false
}

func SetMiddlewares(r *gin.Engine) {
	// gin会先匹配路径(router注册的)，再匹配方法(请求method)
	// 判断路径是否正确
	r.NoRoute(func(c *gin.Context) {
		log.Println()
		log.Printf("Path not found: %s\n", c.Request.URL.Path)
		log.Println()
		// 如果是路径不存在，保持默认 404
		c.IndentedJSON(http.StatusNotFound, gin.H{
			"error": "Path Not Found",
			"path":  c.Request.URL.Path,
		})
		// return
	})

	// 需要开启这个选项，r.NoMethod()才会生效
	r.HandleMethodNotAllowed = true
	// 判断方法是否正确
	r.NoMethod(func(c *gin.Context) {
		// 检查状态码
		// 获取相关的请求信息
		requestMethod := c.Request.Method
		requestPath := c.Request.URL.Path
		// 打印请求信息
		log.Println()
		log.Printf("Method '%s' is not allowed, Path: %s\n", requestMethod, requestPath)
		log.Println()
		// 返回状态码405
		c.IndentedJSON(http.StatusMethodNotAllowed, gin.H{
			"error":   "Method Not Allowed",
			"methods": c.Request.Method,
			"path":    c.Request.URL.Path,
		})
	})

	// 日志中间件测试, 要放在前面，不然被cors中间件拦截了
	// r.Use(middleware.RequestLogger())

	// 跨域请求中间件
	r.Use(cors.New(cors.Config{
		// AllowOrigins: []string{"http://127.0.0.1:5173"},
		AllowOriginFunc:  MyAllowOriginFunc,
		AllowMethods:     []string{"GET", "POST", "OPTIONS", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		// AllowOriginFunc: func(origin string) bool {
		// 	return origin == "https://github.com"
		// },
		MaxAge: 12 * time.Hour,
	}))

}

// SetupRouter 用于初始化并配置 Gin 路由引擎，设置中间件和注册 API 路由。
// 返回一个配置好的 *gin.Engine 实例，供后续启动 HTTP 服务使用。
func SetupRouter() *gin.Engine {
	r := gin.Default()
	/* 相当于 gin.New() + Logger,Recovery中间件
	Logger 中间件：自动记录 HTTP 请求的日志（如请求方法、路径、状态码、耗时等）。
	Recovery 中间件：自动捕获处理请求时发生的 panic，防止程序崩溃，并返回 500 错误。
	*/

	// 禁用自动修正路径
	r.RedirectTrailingSlash = false
	r.RedirectFixedPath = false

	r.SetTrustedProxies([]string{"127.0.0.1"}) // 信任本地代理
	// r.TrustedPlatform = gin.PlatformGoogleAppEngine // 信任 Google App Engine 平台

	// middleware需要在router注册之前
	SetMiddlewares(r)

	api := r.Group("/api")
	{
		api.GET("/get_dish/:id", GetDish)
		api.GET("/get_dishes", GetAllDishes)
		api.GET("/get_dishes_by_category/:category", GetDishesByCategory)
		api.GET("/get_hot_dishes", GetHotDishes)
		api.POST("/get_total_price", GetTotalPrice)
		api.POST("/submit_order", SubmitOrder)
	}
	// r.GET("/api/get_total_price", GetTotalPrice)
	// api.USE(middleware)

	admin := r.Group("/admin")
	{
		admin.POST("/add_dish", AddDish)
		admin.PUT("/update_dish", UpdateDish)
		admin.DELETE("/delete_dish", DeleteDish)
	}

	return r
}

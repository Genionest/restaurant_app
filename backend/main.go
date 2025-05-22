package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"example.com/m/v2/config"
	"example.com/m/v2/controller"
	"github.com/gin-gonic/gin"
)

func gracefullyQuit(r *gin.Engine) {
	srv := &http.Server{
		Addr:    fmt.Sprintf("%v:%v", config.APP.Host, config.APP.Port),
		Handler: r,
	}
	go func() {
		// 服务器多线程启动
		// 不会不会优雅的退出
		if err := srv.ListenAndServe(); err != http.ErrServerClosed {
			log.Fatalf("Error, listen:%s\n", err)
		}
	}()
	// 这里是优雅退出的关键
	// 等待中断信号以优雅地关闭服务器（设置 5 秒的超时时间）
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit

	log.Println("Shutting down server...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Error, Server shutdown:", err)
	}
	log.Println("Server exiting")
}

func main() {
	config.InitDB()
	// config.InitRedis()
	config.InitModel()
	config.InitApp()
	r := controller.SetupRouter()

	gracefullyQuit(r)

	// r.Run(fmt.Sprintf("%v:%v", config.APP.Host, config.APP.Port))
}

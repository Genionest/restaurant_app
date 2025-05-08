package main

import (
	"fmt"

	"example.com/m/v2/config"
	"example.com/m/v2/controller"
)

func gracefullyQuit() {

}

func main() {
	config.InitDB()
	// config.InitRedis()
	config.InitApp()
	r := controller.SetupRouter()

	// srv := &http.Server{
	// 	Addr:    fmt.Sprintf("%v:%v", config.APP.Host, config.APP.Port),
	// 	Handler: r,
	// }
	// go func() {
	// 	if err := srv.ListenAndServe(); err != http.ErrServerClosed {
	// 		log.Fatalf("listen:%s\n", err)
	// 	}
	// }()
	// quit := make(chan os.Signal, 1)
	// signal.Notify(quit, os.Interrupt)
	// <-quit

	// log.Println("Shutting down server...")
	// ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	// defer cancel()

	// if err := srv.Shutdown(ctx); err != nil {
	// 	log.Fatal("Server shutdown:", err)
	// }
	// log.Println("Server exiting")

	r.Run(fmt.Sprintf("%v:%v", config.APP.Host, config.APP.Port))
}

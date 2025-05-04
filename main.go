package main

import "example.com/m/v2/config"

func main() {
	config.InitDB()
	config.InitRedis()
	config.InitApp()

}

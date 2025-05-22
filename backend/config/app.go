package config

type App struct {
	Host string
	Port string
}

var APP = &App{}

// InitApp 初始化应用程序配置，从配置文件中加载应用相关的配置信息。
// 该函数会调用 LoadConfig 函数，将配置文件中 "app" 部分的配置信息加载到全局变量 APP 中。
func InitApp() {
	LoadConfig("app", APP)
}

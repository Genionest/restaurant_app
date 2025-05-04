package config

type App struct {
	Host string
	Port string
}

var APP = &App{}

func InitApp() {
	LoadConfig("app", APP)
}

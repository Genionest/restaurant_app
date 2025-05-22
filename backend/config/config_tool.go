package config

import (
	"log"

	"github.com/spf13/viper"
)

func init() {
	viper.SetConfigType("yml")
	viper.AddConfigPath("./yaml")
}

// LoadConfig 是一个泛型函数，用于加载配置文件并将其解码到指定的结构体中
//
// 参数：
//
//	configName string：配置文件的名称（不含文件扩展名）
//	recv *T：指向要解码配置的目标结构体的指针
//
// 说明：
//
//	使用 viper 库加载配置文件，并将配置文件的内容解码到 recv 指向的结构体中。
//	如果配置文件读取失败或解码失败，函数将调用 log.Fatalf 输出错误信息并退出程序。
func LoadConfig[T any](configName string, recv *T) {
	viper.SetConfigName(configName)
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file: %v", err)
	}

	if err := viper.Unmarshal(recv); err != nil {
		log.Fatalf("(%s) Unable to decode into struct, %v", configName, err)
	}
}

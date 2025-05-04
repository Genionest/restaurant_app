package config

import (
	"log"

	"github.com/spf13/viper"
)

func init() {
	viper.SetConfigType("yml")
	viper.AddConfigPath("./yaml")
}

func LoadConfig[T any](configName string, recv *T) {
	viper.SetConfigName(configName)
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file: %v", err)
	}

	if err := viper.Unmarshal(recv); err != nil {
		log.Fatalf("(%s) Unable to decode into struct, %v", configName, err)
	}
}

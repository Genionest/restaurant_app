package config

import (
	"fmt"
	"log"
	"time"

	"example.com/m/v2/global"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type DBConfig struct {
	Host         string
	Port         string
	User         string
	Password     string
	Name         string
	MaxIdelConns int
	MaxOpenConns int
}

var DB_CONFIG = &DBConfig{}

func InitDB() {
	// configDB()
	LoadConfig("db", DB_CONFIG)

	conf := DB_CONFIG
	// fmt.Println(conf)

	str := "%v:%v@tcp(%v:%v)/%v?charset=utf8mb4&parseTime=True&loc=Local"
	dsn := fmt.Sprintf(str,
		conf.User,
		conf.Password,
		conf.Host,
		conf.Port,
		conf.Name)
	// 希望显示sql语句
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info), // 日志配置
	})
	if err != nil {
		log.Fatalf("Failed to initialize database, got error %v", err)
	}

	sqlDB, err := db.DB()
	sqlDB.SetMaxIdleConns(conf.MaxIdelConns)
	sqlDB.SetMaxOpenConns(conf.MaxOpenConns)
	sqlDB.SetConnMaxLifetime(time.Hour / 2)
	if err != nil {
		log.Fatalf("Failed to configure database, got error %v", err)
	}

	// fmt.Println(db)
	global.DB = db
}

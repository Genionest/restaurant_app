package config

import (
	"fmt"
	"log"

	"example.com/m/v2/global"
	"github.com/go-redis/redis"
)

type RedisDBConfig struct {
	Addr     string
	Password string
	DB       int
}

var REDIS_DB_CONFIG = &RedisDBConfig{}

func InitRedis() {
	// configRedisDB()
	LoadConfig("redis", REDIS_DB_CONFIG)

	conf := REDIS_DB_CONFIG
	fmt.Println("REDIS_DB_CONFIG", conf)

	redisClient := redis.NewClient(&redis.Options{
		Addr:     conf.Addr,
		Password: conf.Password, // no password set
		DB:       conf.DB,       // use default DB
	})
	_, err := redisClient.Ping().Result()
	if err != nil {
		log.Fatalf("Failed to connect redis, got error %v", err)
	}

	fmt.Println(redisClient)
	global.REDIS_DB = redisClient
}

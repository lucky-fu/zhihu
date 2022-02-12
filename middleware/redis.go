package middleware

import (
	"context"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	redis "github.com/go-redis/redis/v8"
	"zhihu.com/m/config"
)

var (
	RedisClient *redis.Client
)

// InitRedis ...
func InitRedis() {
	redisConfig, err := config.GetRedisConfig("config/redis")

	if err != nil {
		panic(fmt.Sprintf("init redis failed err : %v", err))
	}

	RedisClient = redis.NewClient(&redis.Options{
		DB:           redisConfig.DB,
		Addr:         fmt.Sprintf("%s:%s", redisConfig.Host, redisConfig.Port),
		DialTimeout:  time.Duration(redisConfig.ConnectTimeout*1000) * time.Millisecond,
		ReadTimeout:  time.Duration(redisConfig.ReadTimeout*1000) * time.Millisecond,
		WriteTimeout: time.Duration(redisConfig.WriteTimeout*1000) * time.Millisecond,
		Password:     redisConfig.Password,
	})
}

// GetRedisClient ...
func GetRedisClient() *redis.Client {
	return RedisClient
}

// Set ...
func Set(ctx *gin.Context, key string, value string, expire time.Duration) (string, error) {
	redisClient := GetRedisClient()

	if redisClient == nil {
		return "", fmt.Errorf("redis nil")
	}

	return redisClient.Set(context.TODO(), key, value, expire).Result()
}

// Get ...
func Get(ctx *gin.Context, key string) (string, error) {
	redisClient := GetRedisClient()

	if redisClient == nil {
		return "", fmt.Errorf("redis nil")
	}

	return redisClient.Get(context.TODO(), key).Result()
}

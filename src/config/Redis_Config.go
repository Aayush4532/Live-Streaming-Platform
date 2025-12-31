package config

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/redis/go-redis/v9"
)

var RDB *redis.Client

func mustEnv(key string) string {
	val := os.Getenv(key)
	if val == "" {
		panic("missing required env: " + key)
	}
	return val
}

func InitRedis() error {
	addr := mustEnv("REDIS_ADDR")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	RDB = redis.NewClient(&redis.Options{
		Addr:     addr,
		Username: os.Getenv("REDIS_USER"),
		Password: os.Getenv("REDIS_PASS"),
		DB:       0,
	})

	if err := RDB.Ping(ctx).Err(); err != nil {
		return fmt.Errorf("redis connection failed: %w", err)
	}

	fmt.Println("Redis connected successfully")
	return nil
}

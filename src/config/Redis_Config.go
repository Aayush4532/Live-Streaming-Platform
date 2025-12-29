package config

import (
	"context"
	"fmt"
	"os"

	"github.com/redis/go-redis/v9"
)

var RDB *redis.Client;

func InitRedis() error {
	ctx := context.Background()

	RDB = redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_ADDR"),
		Username: os.Getenv("REDIS_USER"),
		Password: os.Getenv("REDIS_PASS"),
		DB:       0,
	})

	if err := RDB.Ping(ctx).Err(); err != nil {
		return err
	}
	fmt.Println("Redis Connected Successfully");
	return nil
}

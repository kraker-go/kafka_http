package redis

import (
	"context"

	"github.com/redis/go-redis/v9"
)

func InitRedis() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})
	return client
}

func Ping(client *redis.Client) error {
	return client.Ping(context.Background()).Err()
}

package redis

import (
	"fmt"

	"github.com/go-mservice-bench/lib/config"
	"github.com/go-redis/redis/v8"
)

func New(config config.Config) *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%v:%v", config.RedisHost, config.RedisPort),
	})
}

package db

import (
	"github.com/redis/go-redis/v9"
	"github.com/szymon676/codehund/types"
)

func NewRedisClient(opts *types.RedisConnectionOptions) *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost" + opts.Port,
		Password: opts.Password,
	})
	return client
}

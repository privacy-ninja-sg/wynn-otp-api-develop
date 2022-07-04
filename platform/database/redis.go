package database

import (
	"fmt"
	"wynn-otp-api/pkg/configs"
	"github.com/go-redis/redis/v8"
)

func RedisConnection(db int) *redis.Client {
	host, port, pwd := configs.RedisConfig()

	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", host, port),
		Password: pwd, // no password set
		DB:       db,  // use default DB
	})

	return rdb
}

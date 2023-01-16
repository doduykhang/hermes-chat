package config

import (
	"fmt"

	"github.com/go-redis/redis/v8"
)

func NewRedisClient (conf Redis) *redis.Client {
	return redis.NewClient(&redis.Options{
    		Addr: fmt.Sprintf(
			"%s:%s",
			conf.Host,
			conf.Port,
		),
		Username: conf.Username,
		Password: conf.Password,
	})
}

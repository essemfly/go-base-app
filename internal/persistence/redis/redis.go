package redis

import (
	"errors"
	"essemfly/go_base_app/config"

	"github.com/go-redis/redis/v8"
)

type Redis struct {
	rdb *redis.Client
}

func NewRedisClient(cfg config.Config) (*Redis, error) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     cfg.RedisURL,
		Password: "",
		DB:       0,
	})

	if rdb == nil {
		return nil, errors.New("redis client 생성 실패")
	}

	return &Redis{rdb: rdb}, nil
}
